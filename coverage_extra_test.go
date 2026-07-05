package versions

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/golang-infrastructure/go-tuple"
	"github.com/stretchr/testify/assert"
)

// -----------------------------------------------------------------------------
// constraint.go 补充覆盖
// -----------------------------------------------------------------------------

func TestConstraint_Match_AllOperators_Cov(t *testing.T) {
	// ConstraintEqual
	c, _ := ParseConstraint("=1.0.0")
	assert.True(t, c.Match(NewVersion("1.0.0")))
	assert.False(t, c.Match(NewVersion("1.0.1")))

	// ConstraintGreaterThan
	c, _ = ParseConstraint(">1.0.0")
	assert.True(t, c.Match(NewVersion("1.0.1")))
	assert.False(t, c.Match(NewVersion("1.0.0")))

	// ConstraintLessThanOrEqual
	c, _ = ParseConstraint("<=1.0.0")
	assert.True(t, c.Match(NewVersion("1.0.0")))
	assert.False(t, c.Match(NewVersion("1.0.1")))
}

func TestConstraint_Match_UnknownOperator_Cov(t *testing.T) {
	// 直接构造一个未知 operator，触发 default 分支返回 false
	c := &Constraint{Operator: ConstraintOperator("???"), Version: NewVersion("1.0.0")}
	assert.False(t, c.Match(NewVersion("1.0.0")))
}

func TestConstraint_String_Wildcard_Cov(t *testing.T) {
	c, err := ParseConstraint("1.x")
	assert.Nil(t, err)
	assert.Equal(t, "1.x", c.String())

	c2, err := ParseConstraint("1.2.x")
	assert.Nil(t, err)
	assert.Equal(t, "1.2.x", c2.String())

	// 通配符版本全是 0 时，最后一个 0 替换为 x
	c3, err := ParseConstraint("0.x")
	assert.Nil(t, err)
	assert.Equal(t, "0.x", c3.String())
}

func TestConstraint_String_AllOperators_Cov(t *testing.T) {
	for _, expr := range []string{">1.0.0", ">=1.0.0", "<1.0.0", "<=1.0.0", "=1.0.0", "!=1.0.0", "^1.0.0", "~1.0.0"} {
		c, err := ParseConstraint(expr)
		assert.Nil(t, err)
		assert.Equal(t, expr, c.String())
	}
}

func TestParseConstraint_MissingVersion_Cov(t *testing.T) {
	// 操作符后无版本号
	_, err := ParseConstraint(">=")
	assert.NotNil(t, err)
	_, err = ParseConstraint("^")
	assert.NotNil(t, err)
}

func TestParseConstraint_InvalidVersion_Cov(t *testing.T) {
	// 无操作符前缀且非通配符（不含 x/X/*），但版本无效（无数字）
	_, err := ParseConstraint("!!!")
	assert.NotNil(t, err)
	// ">=!!!" 走操作符分支无效版本
	_, err = ParseConstraint(">=!!!")
	assert.NotNil(t, err)
}

func TestParseConstraintSet_Error_Cov(t *testing.T) {
	// 其中一段解析失败（!!! 不含通配符 x/X/*，且无数字 -> 无效）
	_, err := ParseConstraintSet(">=1.0.0,!!!")
	assert.NotNil(t, err)
}

func TestParseConstraintUnion_WithEmptyParts_Cov(t *testing.T) {
	// 包含空段（"||" 连续或首尾），最终如果全空应报错；如果有非空段则正常
	cu, err := ParseConstraintUnion("|| >=1.0.0 ||")
	assert.Nil(t, err)
	assert.NotNil(t, cu)
	assert.True(t, cu.Match(NewVersion("1.5.0")))

	// 仅 || 分隔的空段 -> 全空 -> ErrEmptyConstraint
	_, err = ParseConstraintUnion("||")
	assert.NotNil(t, err)

	// 子段解析错误
	_, err = ParseConstraintUnion(">=1.0.0,!!!")
	assert.NotNil(t, err)
}

func TestParseConstraintUnion_EmptyString_Cov(t *testing.T) {
	_, err := ParseConstraintUnion("   ")
	assert.NotNil(t, err)
}

func TestConstraintUnion_Match_Empty_Cov(t *testing.T) {
	// 空 Sets 的 Match 应返回 false
	cu := &ConstraintUnion{Sets: nil}
	assert.False(t, cu.Match(NewVersion("1.0.0")))
}

func TestConstraintSet_Match_Empty_Cov(t *testing.T) {
	cs := &ConstraintSet{Constraints: nil}
	assert.True(t, cs.Match(NewVersion("1.0.0")))
}

// matchCaret: base.VersionNumbers 为空 / 全零 / v<base
func TestMatchCaret_EmptyAndAllZero_Cov(t *testing.T) {
	// base 数字部分为空 -> 第二个分支 return true
	base := &Version{VersionNumbers: VersionNumbers{}}
	assert.True(t, matchCaret(base, NewVersion("1.0.0")))

	// base 全零 -> firstNonZero == -1 -> return true
	baseAllZero := NewVersion("0.0.0")
	assert.True(t, matchCaret(baseAllZero, NewVersion("1.0.0")))

	// v < base -> return false（version.go matchCaret 第一分支）
	assert.False(t, matchCaret(NewVersion("1.2.3"), NewVersion("1.0.0")))
}

// matchTilde: base.VersionNumbers < 2
func TestMatchTilde_ShortBase_Cov(t *testing.T) {
	base := NewVersion("1")
	// v >= base 且 len(base) < 2 -> true
	assert.True(t, matchTilde(base, NewVersion("1")))
	assert.True(t, matchTilde(base, NewVersion("2")))
	// v < base -> false
	assert.False(t, matchTilde(base, NewVersion("0")))
}

// matchWildcard: lastNonZero == -1（全零 base）
func TestMatchWildcard_AllZero_Cov(t *testing.T) {
	base := NewVersion("0.0.0")
	// v >= base, lastNonZero == -1 -> return true
	assert.True(t, matchWildcard(base, NewVersion("0.0.0")))
	assert.True(t, matchWildcard(base, NewVersion("1.0.0")))
	// v < base -> false（0.0.0 不会更小，但用 0.0.0 CompareTo 0.0.0 不 < 0，这里覆盖 false 路径用非空）
	assert.False(t, matchWildcard(NewVersion("1.2.0"), NewVersion("0.0.0")))
}

// -----------------------------------------------------------------------------
// suffix_weight.go String()
// -----------------------------------------------------------------------------

func TestSuffixWeight_String_All_Cov(t *testing.T) {
	tests := []struct {
		w        SuffixWeight
		expected string
	}{
		{SuffixWeightDev, "dev"},
		{SuffixWeightSnapshot, "snapshot"},
		{SuffixWeightNightly, "nightly"},
		{SuffixWeightAlpha, "alpha"},
		{SuffixWeightBeta, "beta"},
		{SuffixWeightMilestone, "milestone"},
		{SuffixWeightRC, "rc"},
		{SuffixWeightPre, "pre"},
		{SuffixWeightCR, "cr"},
		{SuffixWeightFinal, "release"}, // Final/Release/GA 同值 500
		{SuffixWeightRelease, "release"},
		{SuffixWeightGA, "release"},
		{SuffixWeightSP, "sp"},
		{SuffixWeightPatch, "patch"},
		{SuffixWeightPost, "post"},
		{SuffixWeightUnknown, "unknown"},
		{SuffixWeight(9999), "unknown"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.w.String())
	}
}

// -----------------------------------------------------------------------------
// version.go
// -----------------------------------------------------------------------------

// CompareTo 第 4 步 raw 比较的 > 分支（version.go:284 unreachable return）
// 触发 raw > target.Raw 的分支
func TestVersion_CompareTo_RawGreater_Cov(t *testing.T) {
	// 两者数字相同、后缀相同、无 public time，raw 不同 -> 走 raw 字典序
	a := &Version{Raw: "1.2.3", VersionNumbers: VersionNumbers{1, 2, 3}}
	b := &Version{Raw: "1.2.2", VersionNumbers: VersionNumbers{1, 2, 3}}
	assert.Equal(t, 1, a.CompareTo(b))
	assert.Equal(t, -1, b.CompareTo(a))
}

func TestVersion_Major_Empty_Cov(t *testing.T) {
	v := &Version{VersionNumbers: VersionNumbers{}}
	assert.Equal(t, 0, v.Major())
}

func TestVersion_Validate_Negative_Cov(t *testing.T) {
	v := &Version{VersionNumbers: VersionNumbers{-1}}
	err := v.Validate()
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "negative"))
}

func TestVersion_Validate_Empty_Cov(t *testing.T) {
	v := &Version{VersionNumbers: VersionNumbers{}}
	err := v.Validate()
	assert.NotNil(t, err)
}

func TestVersion_UnmarshalJSON_InvalidJSON_Cov(t *testing.T) {
	var v Version
	// 非 JSON
	err := v.UnmarshalJSON([]byte("not-json"))
	assert.NotNil(t, err)

	// 合法 JSON 字符串但版本无效
	var v2 Version
	err = v2.UnmarshalJSON([]byte(`"@@@"`))
	assert.NotNil(t, err)
}

func TestVersion_Scan_Cov(t *testing.T) {
	var v Version
	// string 类型
	err := v.Scan("1.2.3")
	assert.Nil(t, err)
	assert.Equal(t, "1.2.3", v.Raw)

	// []byte 类型
	var v2 Version
	err = v2.Scan([]byte("1.2.3"))
	assert.Nil(t, err)
	assert.Equal(t, "1.2.3", v2.Raw)

	// 不支持的类型
	var v3 Version
	err = v3.Scan(123)
	assert.NotNil(t, err)

	// string 但版本无效
	var v4 Version
	err = v4.Scan("@@@")
	assert.NotNil(t, err)

	// []byte 但版本无效
	var v5 Version
	err = v5.Scan([]byte("@@@"))
	assert.NotNil(t, err)
}

// -----------------------------------------------------------------------------
// version_builder.go Bump* 的空分支
// -----------------------------------------------------------------------------

func TestVersion_Bump_Empty_Cov(t *testing.T) {
	v := &Version{VersionNumbers: VersionNumbers{}}
	assert.Equal(t, "1", v.BumpMajor().Raw)
	assert.Equal(t, "0.1", v.BumpMinor().Raw)
	assert.Equal(t, "0.0.1", v.BumpPatch().Raw)
}

// -----------------------------------------------------------------------------
// version_clone.go With* 的空/补齐分支
// -----------------------------------------------------------------------------

func TestVersion_WithMajor_EmptyNumbers_Cov(t *testing.T) {
	v := &Version{VersionNumbers: VersionNumbers{}}
	got := v.WithMajor(5)
	assert.Equal(t, 5, got.Major())
}

func TestVersion_WithMinor_ShortNumbers_Cov(t *testing.T) {
	// 只有一位数字，触发 WithMinor 的 append 循环
	v := NewVersion("1")
	got := v.WithMinor(9)
	assert.Equal(t, 1, got.Major())
	assert.Equal(t, 9, got.Minor())
}

func TestVersion_WithPatch_ShortNumbers_Cov(t *testing.T) {
	// 只有一位数字，触发 WithPatch 的 append 循环
	v := NewVersion("1")
	got := v.WithPatch(7)
	assert.Equal(t, 1, got.Major())
	assert.Equal(t, 0, got.Minor())
	assert.Equal(t, 7, got.Patch())
}

// -----------------------------------------------------------------------------
// version_group.go
// -----------------------------------------------------------------------------

func TestNewVersionGroupFromVersions_Empty_Cov(t *testing.T) {
	assert.Nil(t, NewVersionGroupFromVersions(nil))
	assert.Nil(t, NewVersionGroupFromVersions([]*Version{}))
}

func TestVersionGroup_GetLatest_GetOldest_Empty_Cov(t *testing.T) {
	g := NewVersionGroup(VersionNumbers{1, 2})
	assert.Nil(t, g.GetLatest())
	assert.Nil(t, g.GetOldest())
}

// QueryRangeVersions end ContainsPolicyYes/None 的 continue 分支（version_group.go:360）
func TestVersionGroup_QueryRangeVersions_EndExcludeContinue_Cov(t *testing.T) {
	g := NewVersionGroup(VersionNumbers{1, 2})
	g.Add(NewVersion("1.2.0"))
	g.Add(NewVersion("1.2.5"))
	// end = 1.2.3, policy Yes/None -> v=1.2.5 > 1.2.3 触发 continue（但 1.2.5 > end.V1 不会 break，因为 break 条件是 > end.V1... 实际 1.2.5>1.2.3 会 break）
	// 这里换个角度：构造一个 v 介于 start 与 end 之间但 end policy Yes 时被排除的场景不存在；
	// 改为测 ContainsPolicyNo 时 v == end 被排除（continue）
	start := tuple.New2[*Version, ContainsPolicy](NewVersion("1.2.0"), ContainsPolicyYes)
	end := tuple.New2[*Version, ContainsPolicy](NewVersion("1.2.5"), ContainsPolicyNo)
	got := g.QueryRangeVersions(start, end)
	// 1.2.5 == end 且 policy No -> 排除；只剩 1.2.0
	assert.Equal(t, 1, len(got))
	assert.Equal(t, "1.2.0", got[0].Raw)
}

// -----------------------------------------------------------------------------
// version_range.go
// -----------------------------------------------------------------------------

func TestVersionRange_String_NilBounds_Cov(t *testing.T) {
	r := &VersionRange{Low: nil, High: nil}
	// LowInclusive/HighInclusive 默认 false -> "(*, *)"
	assert.Equal(t, "(*, *)", r.String())

	// 全闭 -> "[*, *]"
	r2 := &VersionRange{Low: nil, High: nil, LowInclusive: true, HighInclusive: true}
	assert.Equal(t, "[*, *]", r2.String())
}

func TestVersionRange_IsEmpty_NilBound_Cov(t *testing.T) {
	r := &VersionRange{Low: nil, High: nil}
	assert.False(t, r.IsEmpty())
	r2 := &VersionRange{Low: NewVersion("1.0.0"), High: nil}
	assert.False(t, r2.IsEmpty())
}

func TestVersion_PreReleaseType_Unknown_Cov(t *testing.T) {
	// 带后缀但后缀权重未知 -> "unknown"
	v := &Version{Raw: "1.0.0-foobar", VersionNumbers: VersionNumbers{1, 0, 0}, Suffix: VersionSuffix("-foobar")}
	assert.Equal(t, "unknown", v.PreReleaseType())
}

func TestVersion_Canonical_Metadata_Cov(t *testing.T) {
	v := NewVersion("1.2.3")
	v.Metadata = "build123"
	c := v.Canonical()
	assert.True(t, strings.Contains(c, "+build123"))
}

func TestVersion_Format_AllPlaceholders_Cov(t *testing.T) {
	v := NewVersion("v1.2.3-beta1")
	out := v.Format("M=%M m=%m p=%p P=%P s=%s r=%r c=%c %% end")
	assert.Contains(t, out, "M=1")
	assert.Contains(t, out, "m=2")
	assert.Contains(t, out, "p=3")
	assert.Contains(t, out, "P=v")
	assert.Contains(t, out, "s=-beta1")
	assert.Contains(t, out, "r=v1.2.3-beta1")
	assert.Contains(t, out, "% end")
	// %c 规范字符串至少包含 1.2.3
	assert.Contains(t, out, "1.2.3")
}

func TestVersion_Format_UnknownAndTrailingPercent_Cov(t *testing.T) {
	v := NewVersion("1.2.3")
	// 未知占位符 %q 原样输出
	assert.Equal(t, "%q", v.Format("%q"))
	// 结尾单个 % 无后续字符 -> 原样输出 %
	assert.Equal(t, "100%", v.Format("100%"))
}

func TestVersion_Increment_ExtendNumbers_Cov(t *testing.T) {
	// segment 超出当前数字长度，触发 append 循环
	v := NewVersion("1.2")
	got := v.Increment(3) // 数字扩到 4 段
	assert.Equal(t, 4, len(got.VersionNumbers))
	assert.Equal(t, 1, got.VersionNumbers[3])
}

func TestVersion_Increment_NegativeSegment_Cov(t *testing.T) {
	v := NewVersion("1.2.3")
	got := v.Increment(-1)
	assert.Equal(t, v.Raw, got.Raw)
}

func TestNegateConstraint_AllOperators_Cov(t *testing.T) {
	tests := []struct {
		expr string
		neg  ConstraintOperator
	}{
		{"=1.0.0", ConstraintNotEqual},
		{"!=1.0.0", ConstraintEqual},
		{">1.0.0", ConstraintLessThanOrEqual},
		{">=1.0.0", ConstraintLessThan},
		{"<1.0.0", ConstraintGreaterThanOrEqual},
		{"<=1.0.0", ConstraintGreaterThan},
		{"^1.0.0", ConstraintNotEqual}, // default
		{"~1.0.0", ConstraintNotEqual}, // default
		{"1.x", ConstraintNotEqual},    // wildcard default
	}
	for _, tt := range tests {
		c, err := ParseConstraint(tt.expr)
		assert.Nil(t, err)
		neg := NegateConstraint(c)
		assert.Equal(t, tt.neg, neg.Operator)
	}
}

// -----------------------------------------------------------------------------
// version_suffix.go CompareTo 未知/已知混合分支
// -----------------------------------------------------------------------------

func TestVersionSuffix_CompareTo_UnknownMixed_Cov(t *testing.T) {
	// x 已知, target 未知 -> -1
	known := VersionSuffix("-alpha1")
	unknown := VersionSuffix("-foobar")
	assert.Equal(t, -1, known.CompareTo(unknown))
	// x 未知, target 已知 -> 1
	assert.Equal(t, 1, unknown.CompareTo(known))
	// 都未知 -> 字典序
	a := VersionSuffix("-zzz")
	b := VersionSuffix("-aaa")
	assert.Equal(t, 1, a.CompareTo(b))
	assert.Equal(t, -1, b.CompareTo(a))
	assert.Equal(t, 0, a.CompareTo(VersionSuffix("-zzz")))
}

// -----------------------------------------------------------------------------
// visualize.go
// -----------------------------------------------------------------------------

func TestVisualizeVersions_MaxItemsTruncation_Cov(t *testing.T) {
	// 同组（相同 VersionNumbers {1,2,0}）的 5 个版本，触发 maxItems 截断
	versions := NewVersions("1.2.0", "1.2.0-alpha", "1.2.0-beta", "1.2.0-rc1", "1.2.0-rc2")
	var buf bytes.Buffer
	VisualizeVersions(versions, &buf, 2)
	out := buf.String()
	assert.Contains(t, out, "未显示")
}

func TestVisualizeVersions_WithPublicTime_Cov(t *testing.T) {
	v := NewVersion("1.2.0")
	v.PublicTime = time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC)
	versions := []*Version{v}
	var buf bytes.Buffer
	VisualizeVersions(versions, &buf, 0)
	out := buf.String()
	assert.Contains(t, out, "发布时间")
}

// -----------------------------------------------------------------------------
// file.go
// -----------------------------------------------------------------------------

func TestReadVersionsFromFile_NotExist_Cov(t *testing.T) {
	_, err := ReadVersionsFromFile(filepath.Join(os.TempDir(), "definitely_not_exist_versions_test_xyz.txt"))
	assert.NotNil(t, err)
}

func TestReadVersionsFromFile_WithBlankLines_Cov(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "versions.txt")
	assert.Nil(t, os.WriteFile(p, []byte("1.0.0\n\n  \n2.0.0\n"), 0644))
	versions, err := ReadVersionsFromFile(p)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(versions))
}

func TestReadVersionsStringFromFile_NotExist_Cov(t *testing.T) {
	_, err := ReadVersionsStringFromFile(filepath.Join(os.TempDir(), "definitely_not_exist_str_test_xyz.txt"))
	assert.NotNil(t, err)
}

func TestReadVersionsStringFromFile_WithBlankLines_Cov(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "versions.txt")
	assert.Nil(t, os.WriteFile(p, []byte("1.0.0\n\n  \n2.0.0\n"), 0644))
	strs, err := ReadVersionsStringFromFile(p)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(strs))
}

func TestReadVersionsFromReader_ReadError_Cov(t *testing.T) {
	_, err := ReadVersionsFromReader(errReader{})
	assert.NotNil(t, err)
}

func TestReadVersionsFromReader_WithBlankLines_Cov(t *testing.T) {
	versions, err := ReadVersionsFromReader(strings.NewReader("1.0.0\n\n  \n2.0.0\n"))
	assert.Nil(t, err)
	assert.Equal(t, 2, len(versions))
}

// errReader 是一个 Read 始终返回错误的 io.Reader
type errReader struct{}

func (errReader) Read([]byte) (int, error) { return 0, errors.New("read error") }

// -----------------------------------------------------------------------------
// parser.go 辅助：触发 readVersionPrefix 点号通用分支、readVersionSuffix 特殊分支
// -----------------------------------------------------------------------------

func TestParseVersion_DotLeadingNonSpecial_Cov(t *testing.T) {
	// ".2" 以点开头且第二字符为数字，但不是特殊串 ".1" -> 触发 readVersionPrefix 点号通用分支
	v := NewVersion(".2")
	assert.True(t, v.IsValid())
}

func TestParseVersion_MultiDotSuffixRegex_Cov(t *testing.T) {
	// 多余的点号触发 readVersionSuffix 中正则匹配分支（parser.go:445-447）
	v := NewVersion("1.....2.alpha1")
	assert.True(t, v.IsValid())
	assert.Equal(t, ".alpha1", string(v.Suffix))
}

func TestReadVersionSuffix_NoNumberPosition_Cov(t *testing.T) {
	// 触发 readVersionSuffix 中 lastNumberPos == -1 返回空串的分支
	p := &VersionStringParser{versionRunes: []rune("abc")}
	got := p.readVersionSuffix("abc", "1")
	assert.Equal(t, "", got)
}

// 直接覆盖 readVersionNumbers / readVersionSuffix 中针对 "1-rev4-1.18.0-rc" 与 "7_85_0" 的特殊分支。
// 这些分支以 versionWithoutPrefix 作为入参匹配，仅在直接调用或特定历史解析路径下可达。
func TestReadVersionNumbers_SpecialCases_Cov(t *testing.T) {
	p := &VersionStringParser{}
	assert.Equal(t, []int{1, 18, 0}, p.readVersionNumbers("1-rev4-1.18.0-rc"))
	assert.Equal(t, []int{0}, p.readVersionNumbers("7_85_0"))
}

func TestReadVersionSuffix_SpecialCases_Cov(t *testing.T) {
	p := &VersionStringParser{}
	// 这两个分支的 versionNumbersString 入参非空即可，不会被前导 return 拦截
	assert.Equal(t, "-rev4-1.18.0-rc", p.readVersionSuffix("1-rev4-1.18.0-rc", "1.18.0"))
	assert.Equal(t, "_85_0", p.readVersionSuffix("7_85_0", "0"))
}
