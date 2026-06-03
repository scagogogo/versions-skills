package versions

import (
	"fmt"
	"testing"

	"github.com/golang-infrastructure/go-shuffle"
	"github.com/golang-infrastructure/go-tuple"
	"github.com/stretchr/testify/assert"
)

// TestSortedVersionGroups_QueryRange 测试有序版本组的范围查询功能
//
// 这个测试函数验证 SortedVersionGroups 的 QueryRange 方法是否能正确地
// 查询指定范围内的版本。测试使用实际的版本数据集进行范围查询测试。
//
// 测试步骤:
// 1. 从测试数据文件中读取版本号列表（使用 fastjson 的版本列表）
// 2. 随机打乱版本号顺序，以确保排序逻辑正常工作
// 3. 创建 SortedVersionGroups 实例
// 4. 定义查询范围（从 "0" 到 "1.2.83"）和包含策略
// 5. 调用 QueryRange 方法进行范围查询
// 6. 输出查询结果的版本号
//
// 注意:
// 测试文件中还包含了其他版本数据集的查询代码，这些部分已被注释掉，
// 但可以取消注释以测试其他数据集上的范围查询功能。可用的测试数据集包括:
// - org.jboss_jboss-ejb-client.txt
// - de.tum.in.ase_artemis-java-test-sandbox.txt
func TestSortedVersionGroups_QueryRange(t *testing.T) {

	// 1. 基本范围查询测试（使用fastjson数据集）
	versions, err := ReadVersionsFromFile("./test_data/fast_json_versions.txt")
	assert.Nil(t, err)

	// 故意打乱顺序，以防止影响到测试结果
	shuffle.Shuffle(versions)

	groups := NewSortedVersionGroups(versions)

	start := tuple.New2[*Version, ContainsPolicy](NewVersion("0"), ContainsPolicyYes)
	end := tuple.New2[*Version, ContainsPolicy](NewVersion("1.2.83"), ContainsPolicyNo)
	queryRange := groups.QueryRange(start, end)
	for _, v := range queryRange {
		fmt.Println(v.Raw)
	}

	// 2. 测试构造版本数据集，增加覆盖度
	customTestVersions := []*Version{
		NewVersion("1.0.0"),
		NewVersion("1.1.0"),
		NewVersion("1.2.0"),
		NewVersion("2.0.0"),
		NewVersion("2.1.0"),
	}
	customGroups := NewSortedVersionGroups(customTestVersions)

	// 2.1 测试版本组不存在的情况（如查询3.0.0之前的版本，但实际没有3.0.0）
	nonExistStart := tuple.New2[*Version, ContainsPolicy](NewVersion("3.0.0"), ContainsPolicyYes)
	nonExistEnd := tuple.New2[*Version, ContainsPolicy](NewVersion("4.0.0"), ContainsPolicyYes)
	nonExistResult := customGroups.QueryRange(nonExistStart, nonExistEnd)
	assert.Empty(t, nonExistResult) // 结果应该为空

	// 2.2 测试版本"0"特殊情况（表示从最小版本开始）
	zeroStart := tuple.New2[*Version, ContainsPolicy](NewVersion("0"), ContainsPolicyYes)
	midEnd := tuple.New2[*Version, ContainsPolicy](NewVersion("1.2.0"), ContainsPolicyYes)
	zeroResult := customGroups.QueryRange(zeroStart, midEnd)
	// 检查结果中包含的版本号
	var resultVersions []string
	for _, v := range zeroResult {
		resultVersions = append(resultVersions, v.Raw)
	}
	t.Logf("zeroResult包含的版本: %v", resultVersions)
	// 实际结果是否包含预期版本
	assert.Contains(t, resultVersions, "1.0.0")
	assert.Contains(t, resultVersions, "1.1.0")
	assert.Contains(t, resultVersions, "1.2.0")

	// 2.3 测试不同的包含策略组合
	// 2.3.1 起始版本包含，结束版本不包含
	includeStartExcludeEnd := customGroups.QueryRange(
		tuple.New2[*Version, ContainsPolicy](NewVersion("1.0.0"), ContainsPolicyYes),
		tuple.New2[*Version, ContainsPolicy](NewVersion("2.0.0"), ContainsPolicyNo),
	)
	// 检查结果中包含的版本号
	var includeStartResults []string
	for _, v := range includeStartExcludeEnd {
		includeStartResults = append(includeStartResults, v.Raw)
	}
	t.Logf("includeStartExcludeEnd包含的版本: %v", includeStartResults)
	// 验证包含预期版本，不包含意外版本
	assert.Contains(t, includeStartResults, "1.0.0")
	assert.Contains(t, includeStartResults, "1.1.0")
	assert.Contains(t, includeStartResults, "1.2.0")
	assert.NotContains(t, includeStartResults, "2.0.0")

	// 2.3.2 起始版本不包含，结束版本包含
	excludeStartIncludeEnd := customGroups.QueryRange(
		tuple.New2[*Version, ContainsPolicy](NewVersion("1.0.0"), ContainsPolicyNo),
		tuple.New2[*Version, ContainsPolicy](NewVersion("2.0.0"), ContainsPolicyYes),
	)
	// 检查结果中包含的版本号
	var excludeStartResults []string
	for _, v := range excludeStartIncludeEnd {
		excludeStartResults = append(excludeStartResults, v.Raw)
	}
	t.Logf("excludeStartIncludeEnd包含的版本: %v", excludeStartResults)
	// 验证包含预期版本，不包含意外版本
	assert.NotContains(t, excludeStartResults, "1.0.0")
	assert.Contains(t, excludeStartResults, "1.1.0")
	assert.Contains(t, excludeStartResults, "1.2.0")
	assert.Contains(t, excludeStartResults, "2.0.0")

	// 以下是其他数据集的测试代码，目前已注释掉
	// 可以取消注释以测试不同数据集上的范围查询功能

	//versions, err := ReadVersionsFromFile("./test_data/org.jboss_jboss-ejb-client.txt")
	//assert.Nil(t, err)
	//groups := NewSortedVersionGroups(versions)
	//
	//start := tuple.New2[*Version, ContainsPolicy](NewVersion("0"), ContainsPolicyYes)
	//end := tuple.New2[*Version, ContainsPolicy](NewVersion("4.0.39"), ContainsPolicyNo)
	//queryRange := groups.QueryRange(start, end)
	//for _, v := range queryRange {
	//	fmt.Println(v.Raw)
	//}

	//versions, err := ReadVersionsFromFile("./test_data/de.tum.in.ase_artemis-java-test-sandbox.txt")
	//assert.Nil(t, err)
	//groups := NewSortedVersionGroups(versions)
	//
	//start := tuple.New2[*Version, ContainsPolicy](NewVersion("0"), ContainsPolicyYes)
	//end := tuple.New2[*Version, ContainsPolicy](NewVersion("1.8.0"), ContainsPolicyNo)
	//queryRange := groups.QueryRange(start, end)
	//for _, v := range queryRange {
	//	fmt.Println(v.Raw)
	//}
}

func TestSortedVersionGroups_Len(t *testing.T) {
	versions := NewVersions("1.0.0", "1.1.0", "2.0.0")
	svg := NewSortedVersionGroups(versions)
	if svg.Len() != 3 {
		t.Errorf("Len() = %d, want 3", svg.Len())
	}
}

func TestSortedVersionGroups_Get(t *testing.T) {
	versions := NewVersions("1.0.0", "1.1.0", "2.0.0")
	svg := NewSortedVersionGroups(versions)
	g := svg.Get("1.0.0")
	if g == nil {
		t.Fatal("Get(\"1.0.0\") returned nil")
	}
	if g.ID() != "1.0.0" {
		t.Errorf("Get().ID() = %q, want %q", g.ID(), "1.0.0")
	}
	if svg.Get("99.99") != nil {
		t.Error("Get(\"99.99\") should return nil for non-existent group")
	}
}

func TestSortedVersionGroups_At(t *testing.T) {
	versions := NewVersions("2.0.0", "1.0.0")
	svg := NewSortedVersionGroups(versions)
	first := svg.At(0)
	if first == nil {
		t.Fatal("At(0) returned nil")
	}
	if first.ID() != "1.0.0" {
		t.Errorf("At(0).ID() = %q, want %q", first.ID(), "1.0.0")
	}
	if svg.At(-1) != nil {
		t.Error("At(-1) should return nil")
	}
	if svg.At(99) != nil {
		t.Error("At(99) should return nil")
	}
}

func TestSortedVersionGroups_Contains(t *testing.T) {
	versions := NewVersions("1.0.0", "2.0.0")
	svg := NewSortedVersionGroups(versions)
	if !svg.Contains("1.0.0") {
		t.Error("Should contain group 1.0.0")
	}
	if svg.Contains("99.99") {
		t.Error("Should not contain group 99.99")
	}
}
