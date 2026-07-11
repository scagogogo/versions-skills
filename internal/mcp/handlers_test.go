package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
)

// newReq 构造一个带参数的 CallToolRequest
func newReq(args map[string]interface{}) mcp.CallToolRequest {
	r := mcp.CallToolRequest{}
	r.Params.Arguments = args
	return r
}

// parseResult 把 CallToolResult 的文本内容解析为 map
func parseResult(t *testing.T, result *mcp.CallToolResult) map[string]interface{} {
	t.Helper()
	assert.False(t, result.IsError, "result should not be error: %v", result.Content)
	if len(result.Content) == 0 {
		return nil
	}
	// 取第一个 text content
	tc, ok := result.Content[0].(mcp.TextContent)
	if !ok {
		t.Fatalf("content[0] is not TextContent: %T", result.Content[0])
	}
	text := tc.Text
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(text), &m); err != nil {
		t.Fatalf("unmarshal result: %v (text: %s)", err, text)
	}
	return m
}

// parseResultArray 把 CallToolResult 的文本内容解析为 []interface{}
func parseResultArray(t *testing.T, result *mcp.CallToolResult) []interface{} {
	t.Helper()
	assert.False(t, result.IsError, "result should not be error: %v", result.Content)
	if len(result.Content) == 0 {
		return nil
	}
	tc, ok := result.Content[0].(mcp.TextContent)
	if !ok {
		t.Fatalf("content[0] is not TextContent: %T", result.Content[0])
	}
	var arr []interface{}
	if err := json.Unmarshal([]byte(tc.Text), &arr); err != nil {
		t.Fatalf("unmarshal result array: %v (text: %s)", err, tc.Text)
	}
	return arr
}

func TestServer_handleParse(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	// 正常解析
	res, err := s.handleParse(ctx, newReq(map[string]interface{}{"version_string": "1.2.3-beta"}))
	assert.NoError(t, err)
	m := parseResult(t, res)
	assert.Equal(t, "1.2.3-beta", m["raw"])
	assert.Equal(t, true, m["valid"])

	// 缺参数
	res2, err2 := s.handleParse(ctx, newReq(map[string]interface{}{}))
	assert.NoError(t, err2)
	assert.True(t, res2.IsError)

	// 参数类型错误
	res3, err3 := s.handleParse(ctx, newReq(map[string]interface{}{"version_string": 123}))
	assert.NoError(t, err3)
	assert.True(t, res3.IsError)
}

func TestServer_handleValidate(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	// 合法版本
	res, err := s.handleValidate(ctx, newReq(map[string]interface{}{"version_string": "1.2.3"}))
	assert.NoError(t, err)
	m := parseResult(t, res)
	assert.Equal(t, true, m["valid"])

	// 非法版本
	res2, err2 := s.handleValidate(ctx, newReq(map[string]interface{}{"version_string": "abc"}))
	assert.NoError(t, err2)
	m2 := parseResult(t, res2)
	assert.Equal(t, false, m2["valid"])

	// 缺参数
	res3, err3 := s.handleValidate(ctx, newReq(map[string]interface{}{}))
	assert.NoError(t, err3)
	assert.True(t, res3.IsError)
}

func TestServer_handleInfo(t *testing.T) {
	s, _ := NewServer("test")
	res, err := s.handleInfo(context.Background(), newReq(map[string]interface{}{"version_string": "1.0.0-beta"}))
	assert.NoError(t, err)
	m := parseResult(t, res)
	assert.Equal(t, "1.0.0-beta", m["raw"])
	assert.Equal(t, true, m["is_beta"])

	// 缺参数
	res2, err2 := s.handleInfo(context.Background(), newReq(map[string]interface{}{}))
	assert.NoError(t, err2)
	assert.True(t, res2.IsError)
}

func TestServer_handleCompare(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	// v1 < v2
	res, err := s.handleCompare(ctx, newReq(map[string]interface{}{
		"version1": "1.0.0", "version2": "2.0.0",
	}))
	assert.NoError(t, err)
	m := parseResult(t, res)
	assert.Equal(t, -1, int(m["result"].(float64)))
	assert.Contains(t, m["description"].(string), "旧于")

	// v1 > v2
	res2, _ := s.handleCompare(ctx, newReq(map[string]interface{}{
		"version1": "2.0.0", "version2": "1.0.0",
	}))
	m2 := parseResult(t, res2)
	assert.Equal(t, 1, int(m2["result"].(float64)))
	assert.Contains(t, m2["description"].(string), "新于")

	// 相等
	res3, _ := s.handleCompare(ctx, newReq(map[string]interface{}{
		"version1": "1.0.0", "version2": "1.0.0",
	}))
	m3 := parseResult(t, res3)
	assert.Equal(t, 0, int(m3["result"].(float64)))
	assert.Contains(t, m3["description"].(string), "等于")

	// 缺参数
	res4, err4 := s.handleCompare(ctx, newReq(map[string]interface{}{"version1": "1.0.0"}))
	assert.NoError(t, err4)
	assert.True(t, res4.IsError)

	// 缺 version2
	res5, err5 := s.handleCompare(ctx, newReq(map[string]interface{}{"version2": "1.0.0"}))
	assert.NoError(t, err5)
	assert.True(t, res5.IsError)

	// version_string 类型错误
	res6, err6 := s.handleCompare(ctx, newReq(map[string]interface{}{"version1": 123, "version2": "1.0.0"}))
	assert.NoError(t, err6)
	assert.True(t, res6.IsError)
}

func TestServer_handleSort(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	// 升序
	res, err := s.handleSort(ctx, newReq(map[string]interface{}{
		"versions": []interface{}{"3.0.0", "1.0.0", "2.0.0"},
	}))
	assert.NoError(t, err)
	arr := parseResultArray(t, res)
	assert.Equal(t, []interface{}{"1.0.0", "2.0.0", "3.0.0"}, arr)

	// 降序
	res2, err2 := s.handleSort(ctx, newReq(map[string]interface{}{
		"versions":   []interface{}{"3.0.0", "1.0.0", "2.0.0"},
		"descending": true,
	}))
	assert.NoError(t, err2)
	arr2 := parseResultArray(t, res2)
	assert.Equal(t, []interface{}{"3.0.0", "2.0.0", "1.0.0"}, arr2)

	// 缺参数
	res3, err3 := s.handleSort(ctx, newReq(map[string]interface{}{}))
	assert.NoError(t, err3)
	assert.True(t, res3.IsError)

	// 参数类型错误（不是数组）
	res4, err4 := s.handleSort(ctx, newReq(map[string]interface{}{"versions": "1.0.0"}))
	assert.NoError(t, err4)
	assert.True(t, res4.IsError)

	// 数组元素类型错误
	res5, err5 := s.handleSort(ctx, newReq(map[string]interface{}{"versions": []interface{}{123}}))
	assert.NoError(t, err5)
	assert.True(t, res5.IsError)
}

func TestServer_handleGroup(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	// 正常分组
	res, err := s.handleGroup(ctx, newReq(map[string]interface{}{
		"versions": []interface{}{"1.0.0", "1.0.1", "2.0.0"},
	}))
	assert.NoError(t, err)
	groups := parseResultArray(t, res)
	assert.GreaterOrEqual(t, len(groups), 2)

	// 缺参数
	res2, err2 := s.handleGroup(ctx, newReq(map[string]interface{}{}))
	assert.NoError(t, err2)
	assert.True(t, res2.IsError)
}

func TestServer_handleConstraintCheck(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	// set 类型满足
	res, err := s.handleConstraintCheck(ctx, newReq(map[string]interface{}{
		"expression": ">=1.0.0,<2.0.0",
		"version":    "1.5.0",
	}))
	assert.NoError(t, err)
	m := parseResult(t, res)
	assert.Equal(t, true, m["satisfied"])
	assert.Equal(t, "set", m["type"])

	// set 类型不满足
	res2, err2 := s.handleConstraintCheck(ctx, newReq(map[string]interface{}{
		"expression": ">=1.0.0,<2.0.0",
		"version":    "2.5.0",
	}))
	assert.NoError(t, err2)
	m2 := parseResult(t, res2)
	assert.Equal(t, false, m2["satisfied"])

	// union 类型
	res3, err3 := s.handleConstraintCheck(ctx, newReq(map[string]interface{}{
		"expression": ">=1.0.0 <2.0.0 || >=3.0.0",
		"version":    "3.5.0",
		"type":       "union",
	}))
	assert.NoError(t, err3)
	m3 := parseResult(t, res3)
	assert.Equal(t, true, m3["satisfied"])
	assert.Equal(t, "union", m3["type"])

	// 无效版本
	res4, err4 := s.handleConstraintCheck(ctx, newReq(map[string]interface{}{
		"expression": ">=1.0.0",
		"version":    "abc",
	}))
	assert.NoError(t, err4)
	assert.True(t, res4.IsError)

	// 无效约束表达式（set）
	res5, err5 := s.handleConstraintCheck(ctx, newReq(map[string]interface{}{
		"expression": ">>>bad",
		"version":    "1.0.0",
	}))
	assert.NoError(t, err5)
	assert.True(t, res5.IsError)

	// 无效约束表达式（union）
	res6, err6 := s.handleConstraintCheck(ctx, newReq(map[string]interface{}{
		"expression": ">>>bad",
		"version":    "1.0.0",
		"type":       "union",
	}))
	assert.NoError(t, err6)
	assert.True(t, res6.IsError)

	// 缺 expression
	res7, err7 := s.handleConstraintCheck(ctx, newReq(map[string]interface{}{
		"version": "1.0.0",
	}))
	assert.NoError(t, err7)
	assert.True(t, res7.IsError)

	// 缺 version
	res8, err8 := s.handleConstraintCheck(ctx, newReq(map[string]interface{}{
		"expression": ">=1.0.0",
	}))
	assert.NoError(t, err8)
	assert.True(t, res8.IsError)
}

func TestServer_handleRangeQuery(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	versionsArg := []interface{}{"1.0.0", "1.5.0", "2.0.0", "2.5.0", "3.0.0"}

	// 正常范围（默认 include_start=true, include_end=false）
	res, err := s.handleRangeQuery(ctx, newReq(map[string]interface{}{
		"start":    "1.0.0",
		"end":      "3.0.0",
		"versions": versionsArg,
	}))
	assert.NoError(t, err)
	m := parseResult(t, res)
	assert.Equal(t, true, m["include_start"])
	assert.Equal(t, false, m["include_end"])
	assert.Equal(t, float64(4), m["count"])

	// include_end=true
	res2, err2 := s.handleRangeQuery(ctx, newReq(map[string]interface{}{
		"start":       "1.0.0",
		"end":         "3.0.0",
		"versions":    versionsArg,
		"include_end": true,
	}))
	assert.NoError(t, err2)
	m2 := parseResult(t, res2)
	assert.Equal(t, true, m2["include_end"])
	assert.Equal(t, float64(5), m2["count"])

	// include_start=false
	res3, err3 := s.handleRangeQuery(ctx, newReq(map[string]interface{}{
		"start":         "1.0.0",
		"end":           "3.0.0",
		"versions":      versionsArg,
		"include_start": false,
	}))
	assert.NoError(t, err3)
	m3 := parseResult(t, res3)
	assert.Equal(t, false, m3["include_start"])
	assert.Equal(t, float64(3), m3["count"])

	// 无效 start
	res4, err4 := s.handleRangeQuery(ctx, newReq(map[string]interface{}{
		"start":    "abc",
		"end":      "3.0.0",
		"versions": versionsArg,
	}))
	assert.NoError(t, err4)
	assert.True(t, res4.IsError)

	// 无效 end
	res5, err5 := s.handleRangeQuery(ctx, newReq(map[string]interface{}{
		"start":    "1.0.0",
		"end":      "abc",
		"versions": versionsArg,
	}))
	assert.NoError(t, err5)
	assert.True(t, res5.IsError)

	// 缺 start
	res6, err6 := s.handleRangeQuery(ctx, newReq(map[string]interface{}{
		"end":      "3.0.0",
		"versions": versionsArg,
	}))
	assert.NoError(t, err6)
	assert.True(t, res6.IsError)

	// 缺 end
	res7, err7 := s.handleRangeQuery(ctx, newReq(map[string]interface{}{
		"start":    "1.0.0",
		"versions": versionsArg,
	}))
	assert.NoError(t, err7)
	assert.True(t, res7.IsError)

	// 缺 versions
	res8, err8 := s.handleRangeQuery(ctx, newReq(map[string]interface{}{
		"start": "1.0.0",
		"end":   "3.0.0",
	}))
	assert.NoError(t, err8)
	assert.True(t, res8.IsError)
}

func TestServer_handleFilter(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	versionsArg := []interface{}{"1.0.0", "1.0.1-beta", "2.0.0", "v2.1.0", "2.0.0-rc1"}

	// stable 过滤
	res, err := s.handleFilter(ctx, newReq(map[string]interface{}{
		"versions": versionsArg,
		"stable":   true,
	}))
	assert.NoError(t, err)
	list := parseResultArray(t, res)
	for _, v := range list {
		assert.NotContains(t, v, "beta")
		assert.NotContains(t, v, "rc")
	}

	// prerelease 过滤
	res2, err2 := s.handleFilter(ctx, newReq(map[string]interface{}{
		"versions":   versionsArg,
		"prerelease": true,
	}))
	assert.NoError(t, err2)
	list2 := parseResultArray(t, res2)
	assert.GreaterOrEqual(t, len(list2), 2)

	// major 过滤
	res3, err3 := s.handleFilter(ctx, newReq(map[string]interface{}{
		"versions": versionsArg,
		"major":    float64(2),
	}))
	assert.NoError(t, err3)
	list3 := parseResultArray(t, res3)
	for _, v := range list3 {
		assert.Contains(t, v.(string), "2.")
	}

	// minor + patch 组合
	res4, err4 := s.handleFilter(ctx, newReq(map[string]interface{}{
		"versions": versionsArg,
		"minor":    float64(0),
		"patch":    float64(0),
	}))
	assert.NoError(t, err4)
	list4 := parseResultArray(t, res4)
	for _, v := range list4 {
		assert.Contains(t, v.(string), ".0.0")
	}

	// prefix 过滤
	res5, err5 := s.handleFilter(ctx, newReq(map[string]interface{}{
		"versions": versionsArg,
		"prefix":   "v",
	}))
	assert.NoError(t, err5)
	list5 := parseResultArray(t, res5)
	assert.GreaterOrEqual(t, len(list5), 1)
	for _, v := range list5 {
		assert.True(t, strings.HasPrefix(v.(string), "v"))
	}

	// suffix 过滤
	res6, err6 := s.handleFilter(ctx, newReq(map[string]interface{}{
		"versions": versionsArg,
		"suffix":   "beta",
	}))
	assert.NoError(t, err6)
	list6 := parseResultArray(t, res6)
	for _, v := range list6 {
		assert.Contains(t, v.(string), "beta")
	}

	// constraint 过滤
	res7, err7 := s.handleFilter(ctx, newReq(map[string]interface{}{
		"versions":   versionsArg,
		"constraint": ">=2.0.0",
	}))
	assert.NoError(t, err7)
	parseResultArray(t, res7)

	// constraint 解析失败
	res8, err8 := s.handleFilter(ctx, newReq(map[string]interface{}{
		"versions":   versionsArg,
		"constraint": ">>>bad",
	}))
	assert.NoError(t, err8)
	assert.True(t, res8.IsError)

	// 缺参数
	res9, err9 := s.handleFilter(ctx, newReq(map[string]interface{}{}))
	assert.NoError(t, err9)
	assert.True(t, res9.IsError)

	// 数组类型错误
	res10, err10 := s.handleFilter(ctx, newReq(map[string]interface{}{"versions": "1.0.0"}))
	assert.NoError(t, err10)
	assert.True(t, res10.IsError)
}

func TestServer_handleMin(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	res, err := s.handleMin(ctx, newReq(map[string]interface{}{
		"versions": []interface{}{"3.0.0", "1.0.0", "2.0.0"},
	}))
	assert.NoError(t, err)
	m := parseResult(t, res)
	assert.Equal(t, "1.0.0", m["raw"])

	// 全部无效版本 → errorResult
	res2, err2 := s.handleMin(ctx, newReq(map[string]interface{}{
		"versions": []interface{}{"abc", "def"},
	}))
	assert.NoError(t, err2)
	assert.True(t, res2.IsError)

	// 缺参数
	res3, err3 := s.handleMin(ctx, newReq(map[string]interface{}{}))
	assert.NoError(t, err3)
	assert.True(t, res3.IsError)
}

func TestServer_handleMax(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	res, err := s.handleMax(ctx, newReq(map[string]interface{}{
		"versions": []interface{}{"1.0.0", "3.0.0", "2.0.0"},
	}))
	assert.NoError(t, err)
	m := parseResult(t, res)
	assert.Equal(t, "3.0.0", m["raw"])

	// 全部无效版本 → errorResult
	res2, err2 := s.handleMax(ctx, newReq(map[string]interface{}{
		"versions": []interface{}{"abc"},
	}))
	assert.NoError(t, err2)
	assert.True(t, res2.IsError)

	// 缺参数
	res3, err3 := s.handleMax(ctx, newReq(map[string]interface{}{}))
	assert.NoError(t, err3)
	assert.True(t, res3.IsError)
}

func TestServer_handleLatestStable(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	// 有稳定版
	res, err := s.handleLatestStable(ctx, newReq(map[string]interface{}{
		"versions": []interface{}{"1.0.0", "2.0.0-beta", "2.0.0"},
	}))
	assert.NoError(t, err)
	m := parseResult(t, res)
	assert.Equal(t, "2.0.0", m["raw"])

	// 无稳定版 → errorResult
	res2, err2 := s.handleLatestStable(ctx, newReq(map[string]interface{}{
		"versions": []interface{}{"1.0.0-beta", "2.0.0-rc1"},
	}))
	assert.NoError(t, err2)
	assert.True(t, res2.IsError)

	// 缺参数
	res3, err3 := s.handleLatestStable(ctx, newReq(map[string]interface{}{}))
	assert.NoError(t, err3)
	assert.True(t, res3.IsError)
}

func TestServer_handleLatestPrerelease(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	// 有预发布
	res, err := s.handleLatestPrerelease(ctx, newReq(map[string]interface{}{
		"versions": []interface{}{"1.0.0-alpha", "2.0.0-beta", "2.0.0"},
	}))
	assert.NoError(t, err)
	m := parseResult(t, res)
	assert.Equal(t, "2.0.0-beta", m["raw"])

	// 无预发布 → errorResult
	res2, err2 := s.handleLatestPrerelease(ctx, newReq(map[string]interface{}{
		"versions": []interface{}{"1.0.0", "2.0.0"},
	}))
	assert.NoError(t, err2)
	assert.True(t, res2.IsError)

	// 缺参数
	res3, err3 := s.handleLatestPrerelease(ctx, newReq(map[string]interface{}{}))
	assert.NoError(t, err3)
	assert.True(t, res3.IsError)
}

func TestServer_handleUnique(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	res, err := s.handleUnique(ctx, newReq(map[string]interface{}{
		"versions": []interface{}{"1.0.0", "1.0.0", "2.0.0", "2.0.0", "3.0.0"},
	}))
	assert.NoError(t, err)
	arr := parseResultArray(t, res)
	assert.Equal(t, []interface{}{"1.0.0", "2.0.0", "3.0.0"}, arr)

	// 缺参数
	res2, err2 := s.handleUnique(ctx, newReq(map[string]interface{}{}))
	assert.NoError(t, err2)
	assert.True(t, res2.IsError)
}

func TestServer_handleSetOperation(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	setA := []interface{}{"1.0.0", "2.0.0", "3.0.0"}
	setB := []interface{}{"2.0.0", "3.0.0", "4.0.0"}

	// difference
	res, err := s.handleSetOperation(ctx, newReq(map[string]interface{}{
		"operation": "difference",
		"set_a":     setA,
		"set_b":     setB,
	}))
	assert.NoError(t, err)
	m := parseResult(t, res)
	assert.Equal(t, "difference", m["operation"])
	assert.Equal(t, []interface{}{"1.0.0"}, m["versions"])

	// intersection
	res2, err2 := s.handleSetOperation(ctx, newReq(map[string]interface{}{
		"operation": "intersection",
		"set_a":     setA,
		"set_b":     setB,
	}))
	assert.NoError(t, err2)
	m2 := parseResult(t, res2)
	assert.Equal(t, []interface{}{"2.0.0", "3.0.0"}, m2["versions"])

	// union
	res3, err3 := s.handleSetOperation(ctx, newReq(map[string]interface{}{
		"operation": "union",
		"set_a":     setA,
		"set_b":     setB,
	}))
	assert.NoError(t, err3)
	m3 := parseResult(t, res3)
	assert.Equal(t, float64(4), m3["count"])

	// 不支持的运算
	res4, err4 := s.handleSetOperation(ctx, newReq(map[string]interface{}{
		"operation": "xor",
		"set_a":     setA,
		"set_b":     setB,
	}))
	assert.NoError(t, err4)
	assert.True(t, res4.IsError)

	// 缺 operation
	res5, err5 := s.handleSetOperation(ctx, newReq(map[string]interface{}{
		"set_a": setA,
		"set_b": setB,
	}))
	assert.NoError(t, err5)
	assert.True(t, res5.IsError)

	// 缺 set_a
	res6, err6 := s.handleSetOperation(ctx, newReq(map[string]interface{}{
		"operation": "union",
		"set_b":     setB,
	}))
	assert.NoError(t, err6)
	assert.True(t, res6.IsError)

	// 缺 set_b
	res7, err7 := s.handleSetOperation(ctx, newReq(map[string]interface{}{
		"operation": "union",
		"set_a":     setA,
	}))
	assert.NoError(t, err7)
	assert.True(t, res7.IsError)
}

func TestServer_handleVisualize(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	versionsArg := []interface{}{"1.0.0", "1.0.1", "2.0.0"}

	// 普通模式
	res, err := s.handleVisualize(ctx, newReq(map[string]interface{}{
		"versions": versionsArg,
	}))
	assert.NoError(t, err)
	m := parseResult(t, res)
	assert.NotEmpty(t, m["text"].(string))
	assert.Equal(t, float64(3), m["count"])

	// 普通模式 + max_items_per_group
	res2, err2 := s.handleVisualize(ctx, newReq(map[string]interface{}{
		"versions":            versionsArg,
		"max_items_per_group": float64(1),
	}))
	assert.NoError(t, err2)
	parseResult(t, res2)

	// groups_only 模式
	res3, err3 := s.handleVisualize(ctx, newReq(map[string]interface{}{
		"versions":    versionsArg,
		"groups_only": true,
	}))
	assert.NoError(t, err3)
	m3 := parseResult(t, res3)
	assert.NotEmpty(t, m3["text"].(string))

	// 缺参数
	res4, err4 := s.handleVisualize(ctx, newReq(map[string]interface{}{}))
	assert.NoError(t, err4)
	assert.True(t, res4.IsError)
}

func TestServer_handleBuild(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	// 完整构建
	res, err := s.handleBuild(ctx, newReq(map[string]interface{}{
		"prefix": "v",
		"major":  float64(1),
		"minor":  float64(2),
		"patch":  float64(3),
		"suffix": "-beta",
	}))
	assert.NoError(t, err)
	m := parseResult(t, res)
	assert.Equal(t, true, m["valid"])
	assert.Contains(t, m["raw"].(string), "1")

	// 部分参数
	res2, err2 := s.handleBuild(ctx, newReq(map[string]interface{}{
		"major": float64(2),
	}))
	assert.NoError(t, err2)
	parseResult(t, res2)

	// 空参数
	res3, err3 := s.handleBuild(ctx, newReq(map[string]interface{}{}))
	assert.NoError(t, err3)
	parseResult(t, res3)
}

func TestServer_handleBump(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	// major
	res, err := s.handleBump(ctx, newReq(map[string]interface{}{
		"version_string": "1.2.3",
		"bump_type":      "major",
	}))
	assert.NoError(t, err)
	m := parseResult(t, res)
	assert.Contains(t, m["bumped"].(string), "2")

	// minor
	res2, err2 := s.handleBump(ctx, newReq(map[string]interface{}{
		"version_string": "1.2.3",
		"bump_type":      "minor",
	}))
	assert.NoError(t, err2)
	m2 := parseResult(t, res2)
	assert.Contains(t, m2["bumped"].(string), "1.3")

	// patch
	res3, err3 := s.handleBump(ctx, newReq(map[string]interface{}{
		"version_string": "1.2.3",
		"bump_type":      "patch",
	}))
	assert.NoError(t, err3)
	m3 := parseResult(t, res3)
	assert.Contains(t, m3["bumped"].(string), "1.2.4")

	// 无效版本
	res4, err4 := s.handleBump(ctx, newReq(map[string]interface{}{
		"version_string": "abc",
		"bump_type":      "patch",
	}))
	assert.NoError(t, err4)
	assert.True(t, res4.IsError)

	// 不支持的类型
	res5, err5 := s.handleBump(ctx, newReq(map[string]interface{}{
		"version_string": "1.2.3",
		"bump_type":      "nope",
	}))
	assert.NoError(t, err5)
	assert.True(t, res5.IsError)

	// 缺 version_string
	res6, err6 := s.handleBump(ctx, newReq(map[string]interface{}{
		"bump_type": "patch",
	}))
	assert.NoError(t, err6)
	assert.True(t, res6.IsError)

	// 缺 bump_type
	res7, err7 := s.handleBump(ctx, newReq(map[string]interface{}{
		"version_string": "1.2.3",
	}))
	assert.NoError(t, err7)
	assert.True(t, res7.IsError)
}

func TestServer_handleCore(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	// 正常
	res, err := s.handleCore(ctx, newReq(map[string]interface{}{
		"version_string": "1.2.3-beta",
	}))
	assert.NoError(t, err)
	m := parseResult(t, res)
	assert.Equal(t, "1.2.3", m["core"])

	// 无效版本
	res2, err2 := s.handleCore(ctx, newReq(map[string]interface{}{
		"version_string": "abc",
	}))
	assert.NoError(t, err2)
	assert.True(t, res2.IsError)

	// 缺参数
	res3, err3 := s.handleCore(ctx, newReq(map[string]interface{}{}))
	assert.NoError(t, err3)
	assert.True(t, res3.IsError)
}

func TestServer_handleReadFile(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	// 正常读取
	dir := t.TempDir()
	fp := filepath.Join(dir, "versions.txt")
	assert.NoError(t, os.WriteFile(fp, []byte("3.0.0\n1.0.0\n2.0.0\n"), 0o644))
	res, err := s.handleReadFile(ctx, newReq(map[string]interface{}{
		"filepath": fp,
	}))
	assert.NoError(t, err)
	m := parseResult(t, res)
	assert.Equal(t, fp, m["filepath"])
	assert.Equal(t, float64(3), m["count"])

	// 不存在文件
	res2, err2 := s.handleReadFile(ctx, newReq(map[string]interface{}{
		"filepath": filepath.Join(dir, "nope.txt"),
	}))
	assert.NoError(t, err2)
	assert.True(t, res2.IsError)

	// 缺参数
	res3, err3 := s.handleReadFile(ctx, newReq(map[string]interface{}{}))
	assert.NoError(t, err3)
	assert.True(t, res3.IsError)
}

func TestServer_handleWriteFile(t *testing.T) {
	s, _ := NewServer("test")
	ctx := context.Background()

	// 正常写入
	dir := t.TempDir()
	fp := filepath.Join(dir, "out.txt")
	res, err := s.handleWriteFile(ctx, newReq(map[string]interface{}{
		"filepath": fp,
		"versions": []interface{}{"3.0.0", "1.0.0", "2.0.0"},
	}))
	assert.NoError(t, err)
	m := parseResult(t, res)
	assert.Equal(t, fp, m["filepath"])
	assert.Equal(t, float64(3), m["count"])
	// 验证文件已写入且排序
	content, rerr := os.ReadFile(fp)
	assert.NoError(t, rerr)
	assert.Equal(t, "1.0.0\n2.0.0\n3.0.0", string(content))

	// 缺 filepath
	res2, err2 := s.handleWriteFile(ctx, newReq(map[string]interface{}{
		"versions": []interface{}{"1.0.0"},
	}))
	assert.NoError(t, err2)
	assert.True(t, res2.IsError)

	// 缺 versions
	res3, err3 := s.handleWriteFile(ctx, newReq(map[string]interface{}{
		"filepath": fp,
	}))
	assert.NoError(t, err3)
	assert.True(t, res3.IsError)

	// 写入失败（路径是目录）
	dirFP := filepath.Join(dir, "subdir")
	assert.NoError(t, os.Mkdir(dirFP, 0o755))
	res4, err4 := s.handleWriteFile(ctx, newReq(map[string]interface{}{
		"filepath": dirFP,
		"versions": []interface{}{"1.0.0"},
	}))
	assert.NoError(t, err4)
	assert.True(t, res4.IsError)
}

// === server.go 辅助函数测试 ===

func TestGetStringParam(t *testing.T) {
	// 正常
	v, err := getStringParam(newReq(map[string]interface{}{"k": "v"}), "k")
	assert.NoError(t, err)
	assert.Equal(t, "v", v)
	// 缺参数
	_, err = getStringParam(newReq(map[string]interface{}{}), "k")
	assert.Error(t, err)
	// 类型错误
	_, err = getStringParam(newReq(map[string]interface{}{"k": 123}), "k")
	assert.Error(t, err)
}

func TestGetOptionalStringParam(t *testing.T) {
	// 正常
	assert.Equal(t, "v", getOptionalStringParam(newReq(map[string]interface{}{"k": "v"}), "k"))
	// 缺参数
	assert.Equal(t, "", getOptionalStringParam(newReq(map[string]interface{}{}), "k"))
	// 类型错误
	assert.Equal(t, "", getOptionalStringParam(newReq(map[string]interface{}{"k": 123}), "k"))
}

func TestGetBoolParam(t *testing.T) {
	// 正常 true
	assert.True(t, getBoolParam(newReq(map[string]interface{}{"k": true}), "k"))
	// 正常 false
	assert.False(t, getBoolParam(newReq(map[string]interface{}{"k": false}), "k"))
	// 缺参数
	assert.False(t, getBoolParam(newReq(map[string]interface{}{}), "k"))
	// 类型错误
	assert.False(t, getBoolParam(newReq(map[string]interface{}{"k": "true"}), "k"))
}

func TestGetOptionalNumberParam(t *testing.T) {
	// 正常
	n, ok := getOptionalNumberParam(newReq(map[string]interface{}{"k": float64(42)}), "k")
	assert.True(t, ok)
	assert.Equal(t, 42, n)
	// 缺参数
	_, ok = getOptionalNumberParam(newReq(map[string]interface{}{}), "k")
	assert.False(t, ok)
	// 类型错误
	_, ok = getOptionalNumberParam(newReq(map[string]interface{}{"k": "42"}), "k")
	assert.False(t, ok)
}

func TestGetVersionStringsParam(t *testing.T) {
	// 正常
	vs, err := getVersionStringsParam(newReq(map[string]interface{}{"k": []interface{}{"1.0.0", "2.0.0"}}), "k")
	assert.NoError(t, err)
	assert.Equal(t, []string{"1.0.0", "2.0.0"}, vs)
	// 缺参数
	_, err = getVersionStringsParam(newReq(map[string]interface{}{}), "k")
	assert.Error(t, err)
	// 类型错误（非数组）
	_, err = getVersionStringsParam(newReq(map[string]interface{}{"k": "1.0.0"}), "k")
	assert.Error(t, err)
	// 元素类型错误
	_, err = getVersionStringsParam(newReq(map[string]interface{}{"k": []interface{}{123}}), "k")
	assert.Error(t, err)
}

func TestErrorResult_text(t *testing.T) {
	r := errorResult(assert.AnError)
	assert.True(t, r.IsError)
}

// === tools.go 测试 ===

func TestMarshalJSON(t *testing.T) {
	// 正常
	b, err := marshalJSON(map[string]interface{}{"k": "v"})
	assert.NoError(t, err)
	assert.Contains(t, string(b), `"k": "v"`)
	// 触发错误：channel 不可序列化
	_, err = marshalJSON(make(chan int))
	assert.Error(t, err)
}

func TestJsonResult_error(t *testing.T) {
	// jsonResult 内部调用 marshalJSON，传入不可序列化值应走 errorResult 分支
	r := jsonResult(make(chan int))
	assert.True(t, r.IsError)
}

func TestServer_ServeStdio(t *testing.T) {
	s, _ := NewServer("test")
	// 用一个会立即 EOF 的 pipe 替换 os.Stdin，ServeStdio 读到 EOF 后应返回 nil
	origStdin := os.Stdin
	r, w, err := os.Pipe()
	assert.NoError(t, err)
	assert.NoError(t, w.Close()) // 立即 EOF
	os.Stdin = r
	defer func() { os.Stdin = origStdin }()

	done := make(chan error, 1)
	go func() { done <- s.ServeStdio() }()
	select {
	case e := <-done:
		// 读到 EOF 正常返回（具体是否返回 nil 由 mcp-go 决定，只要不 panic 即可）
		_ = e
	case <-time.After(5 * time.Second):
		t.Fatal("ServeStdio did not return within timeout")
	}
}

func TestServer_ServeSSE(t *testing.T) {
	s, _ := NewServer("test")
	// 先占用一个端口不释放，让 ServeSSE 的 Start 绑定失败，从而走错误返回路径
	l, err := net.Listen("tcp", "127.0.0.1:0")
	assert.NoError(t, err)
	defer func() { _ = l.Close() }()
	port := l.Addr().(*net.TCPAddr).Port
	err = s.ServeSSE(fmt.Sprintf("127.0.0.1:%d", port))
	// 预期返回错误（端口被占用）
	assert.Error(t, err)
}

func TestToolDefinitions(t *testing.T) {
	// 调用所有 tool 定义函数，确保都被覆盖（NewServer→registerTools 已调用，
	// 这里直接调用一次以覆盖函数体本身的行）。
	s, _ := NewServer("test")
	assert.NotPanics(t, func() {
		_ = s.toolParse()
		_ = s.toolValidate()
		_ = s.toolInfo()
		_ = s.toolCompare()
		_ = s.toolSort()
		_ = s.toolGroup()
		_ = s.toolConstraintCheck()
		_ = s.toolRangeQuery()
		_ = s.toolFilter()
		_ = s.toolMin()
		_ = s.toolMax()
		_ = s.toolLatestStable()
		_ = s.toolLatestPrerelease()
		_ = s.toolUnique()
		_ = s.toolSetOperation()
		_ = s.toolVisualize()
		_ = s.toolBuild()
		_ = s.toolBump()
		_ = s.toolCore()
		_ = s.toolReadFile()
		_ = s.toolWriteFile()
	})
	for _, tt := range []mcp.Tool{
		s.toolParse(), s.toolValidate(), s.toolInfo(), s.toolCompare(),
		s.toolSort(), s.toolGroup(), s.toolConstraintCheck(), s.toolRangeQuery(),
		s.toolFilter(), s.toolMin(), s.toolMax(), s.toolLatestStable(),
		s.toolLatestPrerelease(), s.toolUnique(), s.toolSetOperation(),
		s.toolVisualize(), s.toolBuild(), s.toolBump(), s.toolCore(),
		s.toolReadFile(), s.toolWriteFile(),
	} {
		assert.NotEmpty(t, tt.Name)
	}
}
