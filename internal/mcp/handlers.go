package mcp

import (
	"bytes"
	"context"
	"fmt"

	"github.com/golang-infrastructure/go-tuple"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/scagogogo/versions-skills"
	"github.com/scagogogo/versions-skills/internal/verfmt"
)

// === 工具处理函数 ===

func (s *Server) handleParse(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	versionStr, err := getStringParam(request, "version_string")
	if err != nil {
		return errorResult(err), nil
	}
	v := versions.NewVersion(versionStr)
	return jsonResult(verfmt.FormatVersion(v)), nil
}

func (s *Server) handleValidate(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	versionStr, err := getStringParam(request, "version_string")
	if err != nil {
		return errorResult(err), nil
	}
	v := versions.NewVersion(versionStr)
	isValid := v.IsValid()

	data := map[string]interface{}{
		"raw":   versionStr,
		"valid": isValid,
	}

	if isValid {
		if validateErr := v.Validate(); validateErr != nil {
			data["valid"] = false
			data["error"] = validateErr.Error()
		}
	}

	return jsonResult(data), nil
}

func (s *Server) handleInfo(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	versionStr, err := getStringParam(request, "version_string")
	if err != nil {
		return errorResult(err), nil
	}
	v := versions.NewVersion(versionStr)
	return jsonResult(verfmt.FormatVersionDetailed(v)), nil
}

func (s *Server) handleCompare(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	v1Str, err := getStringParam(request, "version1")
	if err != nil {
		return errorResult(err), nil
	}
	v2Str, err := getStringParam(request, "version2")
	if err != nil {
		return errorResult(err), nil
	}

	v1 := versions.NewVersion(v1Str)
	v2 := versions.NewVersion(v2Str)
	result := v1.CompareTo(v2)

	var description string
	switch {
	case result < 0:
		description = fmt.Sprintf("%s 旧于 %s", v1.RawString(), v2.RawString())
	case result > 0:
		description = fmt.Sprintf("%s 新于 %s", v1.RawString(), v2.RawString())
	default:
		description = fmt.Sprintf("%s 等于 %s", v1.RawString(), v2.RawString())
	}

	data := map[string]interface{}{
		"v1":          v1.RawString(),
		"v2":          v2.RawString(),
		"result":      result,
		"description": description,
	}

	return jsonResult(data), nil
}

func (s *Server) handleSort(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	versionStrs, err := getVersionStringsParam(request, "versions")
	if err != nil {
		return errorResult(err), nil
	}

	vs := parseVersionStrings(versionStrs)
	sorted := versions.SortVersionSlice(vs)

	descending := getBoolParam(request, "descending")
	if descending {
		for i, j := 0, len(sorted)-1; i < j; i, j = i+1, j-1 {
			sorted[i], sorted[j] = sorted[j], sorted[i]
		}
	}

	return jsonResult(verfmt.FormatVersionStrings(sorted)), nil
}

func (s *Server) handleGroup(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	versionStrs, err := getVersionStringsParam(request, "versions")
	if err != nil {
		return errorResult(err), nil
	}

	vs := parseVersionStrings(versionStrs)
	groupMap := versions.Group(vs)
	data := verfmt.FormatVersionGroupMap(groupMap)
	return jsonResult(data), nil
}

func (s *Server) handleConstraintCheck(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	expr, err := getStringParam(request, "expression")
	if err != nil {
		return errorResult(err), nil
	}
	versionStr, err := getStringParam(request, "version")
	if err != nil {
		return errorResult(err), nil
	}

	v := versions.NewVersion(versionStr)
	if !v.IsValid() {
		return errorResult(fmt.Errorf("无效的版本号: %s", versionStr)), nil
	}

	constraintType := getOptionalStringParam(request, "type")
	if constraintType == "" {
		constraintType = "set"
	}

	var satisfied bool

	switch constraintType {
	case "union":
		cu, parseErr := versions.ParseConstraintUnion(expr)
		if parseErr != nil {
			return errorResult(fmt.Errorf("解析约束表达式失败: %w", parseErr)), nil
		}
		satisfied = cu.Satisfies(v)
	default:
		cs, parseErr := versions.ParseConstraintSet(expr)
		if parseErr != nil {
			return errorResult(fmt.Errorf("解析约束表达式失败: %w", parseErr)), nil
		}
		satisfied = cs.Satisfies(v)
	}

	data := map[string]interface{}{
		"expression": expr,
		"version":    v.RawString(),
		"satisfied":  satisfied,
		"type":       constraintType,
	}

	return jsonResult(data), nil
}

func (s *Server) handleRangeQuery(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	startStr, err := getStringParam(request, "start")
	if err != nil {
		return errorResult(err), nil
	}
	endStr, err := getStringParam(request, "end")
	if err != nil {
		return errorResult(err), nil
	}
	versionStrs, err := getVersionStringsParam(request, "versions")
	if err != nil {
		return errorResult(err), nil
	}

	startV := versions.NewVersion(startStr)
	endV := versions.NewVersion(endStr)

	if !startV.IsValid() {
		return errorResult(fmt.Errorf("无效的起始版本: %s", startStr)), nil
	}
	if !endV.IsValid() {
		return errorResult(fmt.Errorf("无效的结束版本: %s", endStr)), nil
	}

	vs := parseVersionStrings(versionStrs)

	includeStart := true
	includeEnd := false
	if val, ok := getArgs(request)["include_start"]; ok {
		if b, ok := val.(bool); ok {
			includeStart = b
		}
	}
	if val, ok := getArgs(request)["include_end"]; ok {
		if b, ok := val.(bool); ok {
			includeEnd = b
		}
	}

	startPolicy := versions.ContainsPolicyNo
	if includeStart {
		startPolicy = versions.ContainsPolicyYes
	}
	endPolicy := versions.ContainsPolicyNo
	if includeEnd {
		endPolicy = versions.ContainsPolicyYes
	}

	svg := versions.NewSortedVersionGroups(vs)
	result := svg.QueryRange(
		tuple.New2(startV, startPolicy),
		tuple.New2(endV, endPolicy),
	)

	data := map[string]interface{}{
		"start":         startV.RawString(),
		"end":           endV.RawString(),
		"include_start": includeStart,
		"include_end":   includeEnd,
		"versions":      verfmt.FormatVersionStrings(result),
		"count":         len(result),
	}

	return jsonResult(data), nil
}

func (s *Server) handleFilter(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	versionStrs, err := getVersionStringsParam(request, "versions")
	if err != nil {
		return errorResult(err), nil
	}

	vs := parseVersionStrings(versionStrs)
	result := vs

	if getBoolParam(request, "stable") {
		result = versions.FilterByStable(result)
	}
	if getBoolParam(request, "prerelease") {
		result = versions.FilterByPrerelease(result)
	}
	if n, ok := getOptionalNumberParam(request, "major"); ok {
		result = versions.FilterByMajor(result, n)
	}
	if n, ok := getOptionalNumberParam(request, "minor"); ok {
		result = versions.FilterByMinor(result, n)
	}
	if n, ok := getOptionalNumberParam(request, "patch"); ok {
		result = versions.FilterByPatch(result, n)
	}
	if prefix := getOptionalStringParam(request, "prefix"); prefix != "" {
		result = versions.FilterByPrefix(result, prefix)
	}
	if suffix := getOptionalStringParam(request, "suffix"); suffix != "" {
		result = versions.FilterBySuffix(result, suffix)
	}
	if constraint := getOptionalStringParam(request, "constraint"); constraint != "" {
		cs, parseErr := versions.ParseConstraintSet(constraint)
		if parseErr != nil {
			return errorResult(fmt.Errorf("解析约束表达式失败: %w", parseErr)), nil
		}
		result = versions.FilterByConstraintSet(result, cs)
	}

	return jsonResult(verfmt.FormatVersionStrings(result)), nil
}

func (s *Server) handleMin(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	versionStrs, err := getVersionStringsParam(request, "versions")
	if err != nil {
		return errorResult(err), nil
	}
	vs := parseVersionStringsValid(versionStrs)
	if len(vs) == 0 {
		return errorResult(fmt.Errorf("没有有效版本")), nil
	}
	result := versions.Min(vs)
	return jsonResult(verfmt.FormatVersion(result)), nil
}

func (s *Server) handleMax(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	versionStrs, err := getVersionStringsParam(request, "versions")
	if err != nil {
		return errorResult(err), nil
	}
	vs := parseVersionStringsValid(versionStrs)
	if len(vs) == 0 {
		return errorResult(fmt.Errorf("没有有效版本")), nil
	}
	result := versions.Max(vs)
	return jsonResult(verfmt.FormatVersion(result)), nil
}

func (s *Server) handleLatestStable(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	versionStrs, err := getVersionStringsParam(request, "versions")
	if err != nil {
		return errorResult(err), nil
	}
	vs := parseVersionStringsValid(versionStrs)
	result := versions.LatestStable(vs)
	if result == nil {
		return errorResult(fmt.Errorf("未找到稳定版本")), nil
	}
	return jsonResult(verfmt.FormatVersion(result)), nil
}

func (s *Server) handleLatestPrerelease(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	versionStrs, err := getVersionStringsParam(request, "versions")
	if err != nil {
		return errorResult(err), nil
	}
	vs := parseVersionStringsValid(versionStrs)
	result := versions.LatestPrerelease(vs)
	if result == nil {
		return errorResult(fmt.Errorf("未找到预发布版本")), nil
	}
	return jsonResult(verfmt.FormatVersion(result)), nil
}

func (s *Server) handleUnique(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	versionStrs, err := getVersionStringsParam(request, "versions")
	if err != nil {
		return errorResult(err), nil
	}
	vs := parseVersionStrings(versionStrs)
	result := versions.Unique(vs)
	return jsonResult(verfmt.FormatVersionStrings(result)), nil
}

func (s *Server) handleSetOperation(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	operation, err := getStringParam(request, "operation")
	if err != nil {
		return errorResult(err), nil
	}
	setAStrs, err := getVersionStringsParam(request, "set_a")
	if err != nil {
		return errorResult(err), nil
	}
	setBStrs, err := getVersionStringsParam(request, "set_b")
	if err != nil {
		return errorResult(err), nil
	}

	a := parseVersionStrings(setAStrs)
	b := parseVersionStrings(setBStrs)

	var result []*versions.Version

	switch operation {
	case "difference":
		result = versions.Difference(a, b)
	case "intersection":
		result = versions.Intersection(a, b)
	case "union":
		result = versions.Union(a, b)
	default:
		return errorResult(fmt.Errorf("不支持的运算: %s，请使用 difference/intersection/union", operation)), nil
	}

	data := map[string]interface{}{
		"operation": operation,
		"versions":  verfmt.FormatVersionStrings(result),
		"count":     len(result),
	}

	return jsonResult(data), nil
}

func (s *Server) handleVisualize(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	versionStrs, err := getVersionStringsParam(request, "versions")
	if err != nil {
		return errorResult(err), nil
	}

	vs := parseVersionStrings(versionStrs)
	groupsOnly := getBoolParam(request, "groups_only")

	maxItems := 0
	if n, ok := getOptionalNumberParam(request, "max_items_per_group"); ok {
		maxItems = n
	}

	var buf bytes.Buffer

	if groupsOnly {
		versions.VisualizeVersionGroups(vs, &buf)
	} else {
		versions.VisualizeVersions(vs, &buf, maxItems)
	}

	data := map[string]interface{}{
		"text":  buf.String(),
		"count": len(vs),
	}

	return jsonResult(data), nil
}

func (s *Server) handleBuild(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	builder := versions.NewVersionBuilder()

	if prefix := getOptionalStringParam(request, "prefix"); prefix != "" {
		builder.Prefix(prefix)
	}
	if n, ok := getOptionalNumberParam(request, "major"); ok {
		builder.Major(n)
	}
	if n, ok := getOptionalNumberParam(request, "minor"); ok {
		builder.Minor(n)
	}
	if n, ok := getOptionalNumberParam(request, "patch"); ok {
		builder.Patch(n)
	}
	if suffix := getOptionalStringParam(request, "suffix"); suffix != "" {
		builder.Suffix(suffix)
	}

	v := builder.Build()

	data := map[string]interface{}{
		"raw":   v.RawString(),
		"valid": v.IsValid(),
	}

	return jsonResult(data), nil
}

func (s *Server) handleBump(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	versionStr, err := getStringParam(request, "version_string")
	if err != nil {
		return errorResult(err), nil
	}
	bumpType, err := getStringParam(request, "bump_type")
	if err != nil {
		return errorResult(err), nil
	}

	v := versions.NewVersion(versionStr)
	if !v.IsValid() {
		return errorResult(fmt.Errorf("无效的版本号: %s", versionStr)), nil
	}

	var result *versions.Version

	switch bumpType {
	case "major":
		result = v.BumpMajor()
	case "minor":
		result = v.BumpMinor()
	case "patch":
		result = v.BumpPatch()
	default:
		return errorResult(fmt.Errorf("不支持的递增类型: %s，请使用 major/minor/patch", bumpType)), nil
	}

	data := map[string]interface{}{
		"original": v.RawString(),
		"bumped":   result.RawString(),
		"type":     bumpType,
	}

	return jsonResult(data), nil
}

func (s *Server) handleCore(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	versionStr, err := getStringParam(request, "version_string")
	if err != nil {
		return errorResult(err), nil
	}

	v := versions.NewVersion(versionStr)
	if !v.IsValid() {
		return errorResult(fmt.Errorf("无效的版本号: %s", versionStr)), nil
	}

	coreV := v.Core()

	data := map[string]interface{}{
		"original": v.RawString(),
		"core":     coreV.RawString(),
	}

	return jsonResult(data), nil
}

func (s *Server) handleReadFile(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	filepath, err := getStringParam(request, "filepath")
	if err != nil {
		return errorResult(err), nil
	}

	vs, readErr := versions.ReadVersionsFromFile(filepath)
	if readErr != nil {
		return errorResult(fmt.Errorf("读取文件失败: %w", readErr)), nil
	}

	data := map[string]interface{}{
		"filepath": filepath,
		"count":    len(vs),
		"versions": verfmt.FormatVersionStrings(vs),
	}

	return jsonResult(data), nil
}

func (s *Server) handleWriteFile(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	filepath, err := getStringParam(request, "filepath")
	if err != nil {
		return errorResult(err), nil
	}
	versionStrs, err := getVersionStringsParam(request, "versions")
	if err != nil {
		return errorResult(err), nil
	}

	vs := parseVersionStrings(versionStrs)
	sorted := versions.SortVersionSlice(vs)

	if writeErr := versions.WriteVersionsToFile(sorted, filepath); writeErr != nil {
		return errorResult(fmt.Errorf("写入文件失败: %w", writeErr)), nil
	}

	data := map[string]interface{}{
		"filepath": filepath,
		"count":    len(sorted),
	}

	return jsonResult(data), nil
}

// === 辅助函数 ===

// parseVersionStrings 将字符串数组解析为版本列表
func parseVersionStrings(strs []string) []*versions.Version {
	result := make([]*versions.Version, 0, len(strs))
	for _, s := range strs {
		v := versions.NewVersion(s)
		result = append(result, v)
	}
	return result
}

// parseVersionStringsValid 仅保留有效版本
func parseVersionStringsValid(strs []string) []*versions.Version {
	result := make([]*versions.Version, 0, len(strs))
	for _, s := range strs {
		v := versions.NewVersion(s)
		if v.IsValid() {
			result = append(result, v)
		}
	}
	return result
}
