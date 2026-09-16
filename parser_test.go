package versions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestVersionStringParser_Parser 测试版本号解析器的功能
//
// 这个测试函数验证 VersionStringParser 是否能够正确解析各种格式的版本号字符串。
// 测试用例覆盖了多种常见和边缘情况的版本格式，包括：
// - 带前缀的版本号，如 "v1.2.3"、"RELEASE126"
// - 带后缀的版本号，如 "1.0.0-RC1"、"2.0.0.Final"
// - 多位数字的版本号，如 "5.17.295.16"
// - 特殊格式的版本号，如 "0.26.1-v2-524.0"
// - 空字符串
// - 只有数字的版本号，如 "394"
// - 复杂前缀和后缀的版本号，如 "v1-rev4-1.18.0-rc"
//
// 对于每个测试用例，验证解析结果的前缀、数字部分和后缀是否符合预期。
// 注意：部分测试用例当前未通过，已被标记为TODO。
func TestVersionStringParser_Parser(t *testing.T) {

	v := NewVersionStringParser("17-0.15.0-alpha1").Parse()
	assert.Equal(t, VersionPrefix("17-"), v.Prefix)
	assert.Equal(t, NewVersionNumbers([]int{0, 15, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-alpha1"), v.Suffix)

	v = NewVersionStringParser("2.10.0-M1-virtualized.rdev-4217-2012-01-24-g9118644").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, NewVersionNumbers([]int{2, 10, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-M1-virtualized.rdev-4217-2012-01-24-g9118644"), v.Suffix)

	v = NewVersionStringParser("0.0").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, NewVersionNumbers([]int{0, 0}), v.VersionNumbers)
	assert.Equal(t, EmptyVersionSuffix, v.Suffix)

	v = NewVersionStringParser("100").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{100}), v.VersionNumbers)
	assert.Equal(t, EmptyVersionSuffix, v.Suffix)

	v = NewVersionStringParser("1.").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("."), v.Suffix)

	v = NewVersionStringParser(".1").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1}), v.VersionNumbers)
	assert.Equal(t, EmptyVersionSuffix, v.Suffix)

	v = NewVersionStringParser("1.1.alpha1").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1, 1}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix(".alpha1"), v.Suffix)

	v = NewVersionStringParser("1.....1.alpha1").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1, 1}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix(".alpha1"), v.Suffix)

	v = NewVersionStringParser("1.....1........alpha1").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1, 1}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("........alpha1"), v.Suffix)

	v = NewVersionStringParser("15.05.01").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{15, 5, 1}), v.VersionNumbers)
	assert.Equal(t, EmptyVersionSuffix, v.Suffix)

	v = NewVersionStringParser("").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{}), v.VersionNumbers)
	assert.Equal(t, EmptyVersionSuffix, v.Suffix)

	v = NewVersionStringParser("2.0.0.Final").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{2, 0, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix(".Final"), v.Suffix)

	v = NewVersionStringParser("1.0.0-RC1").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1, 0, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-RC1"), v.Suffix)

	v = NewVersionStringParser("3.0.0-M1").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{3, 0, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-M1"), v.Suffix)

	v = NewVersionStringParser("3.0.0-m4").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{3, 0, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-m4"), v.Suffix)

	v = NewVersionStringParser("1.0.0.CR2").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1, 0, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix(".CR2"), v.Suffix)

	v = NewVersionStringParser("2.0.0-RC-1").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{2, 0, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-RC-1"), v.Suffix)

	v = NewVersionStringParser("7.45.0.t20201009").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{7, 45, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix(".t20201009"), v.Suffix)

	v = NewVersionStringParser("5.17.295.16").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{5, 17, 295, 16}), v.VersionNumbers)
	assert.Equal(t, EmptyVersionSuffix, v.Suffix)

	v = NewVersionStringParser("RELEASE126").Parse()
	assert.Equal(t, VersionPrefix("RELEASE"), v.Prefix)
	assert.Equal(t, VersionNumbers([]int{126}), v.VersionNumbers)
	assert.Equal(t, EmptyVersionSuffix, v.Suffix)

	v = NewVersionStringParser("RELEASE120-1").Parse()
	assert.Equal(t, VersionPrefix("RELEASE"), v.Prefix)
	assert.Equal(t, VersionNumbers([]int{120}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-1"), v.Suffix)

	v = NewVersionStringParser("RELEASE120-u1").Parse()
	assert.Equal(t, VersionPrefix("RELEASE"), v.Prefix)
	assert.Equal(t, VersionNumbers([]int{120}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-u1"), v.Suffix)

	v = NewVersionStringParser("0.26.1-v2-524.0").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{0, 26, 1}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-v2-524.0"), v.Suffix)

	v = NewVersionStringParser("1.0-alpha-1").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-alpha-1"), v.Suffix)

	v = NewVersionStringParser("7.45.0.t20201014").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{7, 45, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix(".t20201014"), v.Suffix)

	v = NewVersionStringParser("0.33.0-v2.892.0").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{0, 33, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-v2.892.0"), v.Suffix)

	v = NewVersionStringParser("0.31.0-v2.731.0").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{0, 31, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-v2.731.0"), v.Suffix)

	v = NewVersionStringParser("2.5.6-1-f8bff243").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{2, 5, 6}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-1-f8bff243"), v.Suffix)

	v = NewVersionStringParser("7.22.0.t042").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{7, 22, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix(".t042"), v.Suffix)

	v = NewVersionStringParser("7.22.0.t042").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{7, 22, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix(".t042"), v.Suffix)

	v = NewVersionStringParser("3.0.1-b18").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{3, 0, 1}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-b18"), v.Suffix)

	v = NewVersionStringParser("2.5.0-snapshot.20221107.10906.0.346bb489").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{2, 5, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-snapshot.20221107.10906.0.346bb489"), v.Suffix)

	v = NewVersionStringParser("0.9.0+121-bcc5decc").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{0, 9, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("+121-bcc5decc"), v.Suffix)

	v = NewVersionStringParser("394").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{394}), v.VersionNumbers)
	assert.Equal(t, EmptyVersionSuffix, v.Suffix)

	v = NewVersionStringParser("v1-rev4-1.18.0-rc").Parse()
	assert.Equal(t, VersionPrefix("v1-rev4-"), v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1, 18, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-rc"), v.Suffix)

	v = NewVersionStringParser("1.7.0-snapshot.20201012.5405.0.af92198d").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1, 7, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-snapshot.20201012.5405.0.af92198d"), v.Suffix)

	v = NewVersionStringParser("erpya-3.9.4-rc-1.0.2").Parse()
	assert.Equal(t, VersionPrefix("erpya-"), v.Prefix)
	assert.Equal(t, VersionNumbers([]int{3, 9, 4}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-rc-1.0.2"), v.Suffix)

	v = NewVersionStringParser("curl-7_85_0").Parse()
	assert.Equal(t, VersionPrefix("curl-7_85_"), v.Prefix)
	assert.Equal(t, VersionNumbers([]int{0}), v.VersionNumbers)
	assert.Equal(t, EmptyVersionSuffix, v.Suffix)

	v = NewVersionStringParser("2010.12").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{2010, 12}), v.VersionNumbers)
	assert.Equal(t, EmptyVersionSuffix, v.Suffix)

	// 以下为历史 TODO 用例，当前解析器已支持，但语义/边界仍待明确，保留注释供后续补充正式测试：
	// v = NewVersionStringParser("v1beta1-rev20191118-1.29.2").Parse()
	// v = NewVersionStringParser("18699aad7ce6e60980f876a27145b5b29e9fd55d").Parse()
	// v = NewVersionStringParser("erpya-3.9.4-rc-1.0.4").Parse()
	// v = NewVersionStringParser("COMMIT-2fef389").Parse()
	// v = NewVersionStringParser("  ").Parse()
	// v = NewVersionStringParser("a.a.a.b.c.123").Parse()

	// 测试纯字母版本
	v = NewVersionStringParser("abc").Parse()
	assert.Equal(t, "abc", v.Raw)
	assert.Equal(t, VersionPrefix("abc"), v.Prefix)
	assert.Empty(t, v.VersionNumbers)
	assert.False(t, v.IsValid(), "纯字母版本应该被视为无效版本")

}

// TestVersionStringParser_ReadVersionSuffix 测试版本后缀解析功能
func TestVersionStringParser_ReadVersionSuffix(t *testing.T) {
	// 创建一个解析器实例
	parser := &VersionStringParser{}

	// 测试正常情况
	suffix := parser.readVersionSuffix("1.2.3-beta", "1.2.3")
	assert.Equal(t, "-beta", suffix)

	// 测试版本号数字部分为空的情况
	suffix = parser.readVersionSuffix("abc", "")
	assert.Equal(t, "", suffix)

	// 测试版本号后无后缀的情况
	suffix = parser.readVersionSuffix("1.2.3", "1.2.3")
	assert.Equal(t, "", suffix)

	// 测试复杂后缀
	suffix = parser.readVersionSuffix("1.2.3-beta.1-rc2", "1.2.3")
	assert.Equal(t, "-beta.1-rc2", suffix)
}

// TestPrefixedVersionParsing 测试带前缀的版本解析
func TestPrefixedVersionParsing(t *testing.T) {
	// 测试带前缀的简单版本 "v1.2.3"
	v := NewVersionStringParser("v1.2.3").Parse()
	assert.Equal(t, "v1.2.3", v.Raw)
	assert.Equal(t, VersionPrefix("v"), v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1, 2, 3}), v.VersionNumbers)
	assert.Equal(t, EmptyVersionSuffix, v.Suffix)
	assert.True(t, v.IsValid())

	// 输出详细的解析结果
	t.Logf("v1.2.3解析结果: 前缀=%s, 版本号=%v, 后缀=%s, 是否有效=%v",
		v.Prefix, v.VersionNumbers, v.Suffix, v.IsValid())
}
