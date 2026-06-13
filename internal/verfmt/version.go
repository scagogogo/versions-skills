package verfmt

import (
	"github.com/scagogogo/versions-skills"
)

// FormatVersion 将 Version 转换为 map 用于 JSON/table 输出
func FormatVersion(v *versions.Version) map[string]interface{} {
	if v == nil {
		return nil
	}
	return map[string]interface{}{
		"raw":             v.RawString(),
		"valid":           v.IsValid(),
		"prefix":          v.Prefix.String(),
		"version_numbers": formatIntSlice(v.Segments()),
		"major":           v.Major(),
		"minor":           v.Minor(),
		"patch":           v.Patch(),
		"suffix":          v.Suffix.String(),
		"suffix_weight":   v.SuffixWeight().String(),
		"group_id":        v.BuildGroupID(),
	}
}

// FormatVersionDetailed 将 Version 转换为详细 map（包含所有 Is* 判断）
func FormatVersionDetailed(v *versions.Version) map[string]interface{} {
	if v == nil {
		return nil
	}
	result := FormatVersion(v)

	// 添加详细字段
	result["sub_version"] = v.SubVersion()
	result["metadata"] = v.Metadata
	result["is_prerelease"] = v.IsPrerelease()
	result["is_stable"] = v.IsStable()
	result["is_dev"] = v.IsDev()
	result["is_alpha"] = v.IsAlpha()
	result["is_beta"] = v.IsBeta()
	result["is_rc"] = v.IsRC()
	result["is_snapshot"] = v.IsSnapshot()
	result["is_milestone"] = v.IsMilestone()
	result["is_nightly"] = v.IsNightly()
	result["is_final"] = v.IsFinal()
	result["is_ga"] = v.IsGA()
	result["is_pre"] = v.IsPre()
	result["is_release"] = v.IsRelease()
	result["is_sp"] = v.IsSP()
	result["is_post"] = v.IsPost()
	result["is_zero"] = v.IsZero()
	result["core"] = v.Core().RawString()

	return result
}

// FormatVersionSimple 将 Version 转换为简单 map（仅核心字段）
func FormatVersionSimple(v *versions.Version) map[string]interface{} {
	if v == nil {
		return nil
	}
	return map[string]interface{}{
		"raw":   v.RawString(),
		"valid": v.IsValid(),
	}
}

// FormatVersionStrings 将版本列表转为字符串列表
func FormatVersionStrings(vs []*versions.Version) []string {
	result := make([]string, len(vs))
	for i, v := range vs {
		result[i] = v.RawString()
	}
	return result
}

// formatIntSlice 将 []int 转为 []interface{} 以便 JSON 序列化
func formatIntSlice(nums []int) []interface{} {
	result := make([]interface{}, len(nums))
	for i, n := range nums {
		result[i] = n
	}
	return result
}
