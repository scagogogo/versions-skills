package versions

import (
	"testing"

	"github.com/golang-infrastructure/go-tuple"
	"github.com/stretchr/testify/assert"
)

// TestVersionGroup_QueryRangeVersions 测试版本组范围查询功能
//
// 这个测试函数验证 VersionGroup 的 QueryRangeVersions 方法是否能正确地
// 查询指定范围内的版本。测试使用特定的版本范围边界条件，检查版本组的范围查询行为。
//
// 测试步骤:
// 1. 创建包含指定版本的版本组
// 2. 定义查询范围（起始版本和结束版本）及包含策略
// 3. 调用 QueryRangeVersions 方法进行范围查询
// 4. 检查查询结果是否符合预期
//
// 注意:
// 当前有一个已知问题（TODO项），需要修复范围查询的bug。
// 该bug与特定版本号范围的查询相关，如 1.2.69 到 1.2.70 之间的版本查询。
func TestVersionGroup_QueryRangeVersions(t *testing.T) {

	// TODO 2023-5-16 19:12:11 单元测试，范围查询有个bug
	// https://repo1.maven.org/maven2/com/alibaba/fastjson/
	// 1.2.69/
	versions := NewVersions("1.2.6", "")
	group := NewVersionGroupFromVersions(versions)

	start := tuple.New2[*Version, ContainsPolicy](NewVersion("1.2.69"), ContainsPolicyYes)
	end := tuple.New2[*Version, ContainsPolicy](NewVersion("1.2.70"), ContainsPolicyYes)
	rangeVersions := group.QueryRangeVersions(start, end)
	t.Log(rangeVersions)

}

// TestVersionGroup_Contains 测试版本组的包含检查方法
func TestVersionGroup_Contains(t *testing.T) {
	// 创建包含多个版本的版本组
	versions := NewVersions("1.0.0", "1.1.0", "1.2.0")
	group := NewVersionGroupFromVersions(versions)

	// 测试组中已有的版本
	v1 := NewVersion("1.0.0")
	assert.True(t, group.Contains(v1))

	v2 := NewVersion("1.1.0")
	assert.True(t, group.Contains(v2))

	v3 := NewVersion("1.2.0")
	assert.True(t, group.Contains(v3))

	// 测试组中不存在的版本
	v4 := NewVersion("1.3.0")
	assert.False(t, group.Contains(v4))

	v5 := NewVersion("0.9.0")
	assert.False(t, group.Contains(v5))

	// 测试与已有版本相同但实例不同的版本
	v6 := NewVersion("1.0.0") // 与v1相同版本号但是不同实例
	assert.True(t, group.Contains(v6))
}

func TestVersionGroup_PrereleaseVersions(t *testing.T) {
	versions := []*Version{
		NewVersion("1.0.0"),
		NewVersion("1.1.0-alpha"),
		NewVersion("1.2.0-beta"),
		NewVersion("2.0.0"),
	}
	group := NewVersionGroupFromVersions(versions)
	pre := group.PrereleaseVersions()
	assert.Equal(t, 2, len(pre))

	// 全稳定版的组
	stableGroup := NewVersionGroupFromVersions([]*Version{
		NewVersion("1.0.0"),
		NewVersion("2.0.0"),
	})
	assert.Equal(t, 0, len(stableGroup.PrereleaseVersions()))
}
