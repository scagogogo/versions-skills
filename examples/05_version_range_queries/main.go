package main

import (
	"fmt"

	"github.com/golang-infrastructure/go-tuple"
	"github.com/scagogogo/versions-skills"
)

func main() {
	// 创建一组测试版本
	testVersions := []*versions.Version{
		versions.NewVersion("1.0.0"),
		versions.NewVersion("1.0.1"),
		versions.NewVersion("1.1.0"),
		versions.NewVersion("1.1.1"),
		versions.NewVersion("1.2.0"),
		versions.NewVersion("2.0.0"),
		versions.NewVersion("2.0.1"),
		versions.NewVersion("2.1.0"),
		versions.NewVersion("3.0.0"),
		versions.NewVersion("3.0.1"),
	}

	// 创建有序版本组
	sortedGroups := versions.NewSortedVersionGroups(testVersions)

	// 示例1：包含边界的范围查询
	fmt.Println("示例1 - 包含边界的范围查询:")

	// 创建查询范围：1.0.0（包含） 到 2.0.0（包含）
	startVersion1 := versions.NewVersion("1.0.0")
	endVersion1 := versions.NewVersion("2.0.0")

	// 创建范围边界条件，使用tuple将版本和包含策略绑定
	startTuple1 := tuple.New2[*versions.Version, versions.ContainsPolicy](startVersion1, versions.ContainsPolicyYes) // 包含起始版本
	endTuple1 := tuple.New2[*versions.Version, versions.ContainsPolicy](endVersion1, versions.ContainsPolicyYes)     // 包含结束版本

	// 执行范围查询
	rangeResult1 := sortedGroups.QueryRange(startTuple1, endTuple1)

	fmt.Printf("查询范围: %s（包含）到 %s（包含）\n", startVersion1.Raw, endVersion1.Raw)
	fmt.Printf("查询结果数量: %d\n", len(rangeResult1))
	fmt.Println("查询结果:")
	for i, version := range versions.SortVersionSlice(rangeResult1) {
		fmt.Printf("  %d. %s\n", i+1, version.Raw)
	}

	/*
		输出示例:
		示例1 - 包含边界的范围查询:
		查询范围: 1.0.0（包含）到 2.0.0（包含）
		查询结果数量: 6
		查询结果:
		  1. 1.0.0
		  2. 1.0.1
		  3. 1.1.0
		  4. 1.1.1
		  5. 1.2.0
		  6. 2.0.0
	*/

	// 示例2：不包含边界的范围查询
	fmt.Println("\n示例2 - 不包含边界的范围查询:")

	// 创建查询范围：1.0.0（不包含） 到 2.0.0（不包含）
	startVersion2 := versions.NewVersion("1.0.0")
	endVersion2 := versions.NewVersion("2.0.0")

	// 创建范围边界条件
	startTuple2 := tuple.New2[*versions.Version, versions.ContainsPolicy](startVersion2, versions.ContainsPolicyNo) // 不包含起始版本
	endTuple2 := tuple.New2[*versions.Version, versions.ContainsPolicy](endVersion2, versions.ContainsPolicyNo)     // 不包含结束版本

	// 执行范围查询
	rangeResult2 := sortedGroups.QueryRange(startTuple2, endTuple2)

	fmt.Printf("查询范围: %s（不包含）到 %s（不包含）\n", startVersion2.Raw, endVersion2.Raw)
	fmt.Printf("查询结果数量: %d\n", len(rangeResult2))
	fmt.Println("查询结果:")
	for i, version := range versions.SortVersionSlice(rangeResult2) {
		fmt.Printf("  %d. %s\n", i+1, version.Raw)
	}

	/*
		输出示例:
		示例2 - 不包含边界的范围查询:
		查询范围: 1.0.0（不包含）到 2.0.0（不包含）
		查询结果数量: 4
		查询结果:
		  1. 1.0.1
		  2. 1.1.0
		  3. 1.1.1
		  4. 1.2.0
	*/

	// 示例3：混合边界的范围查询
	fmt.Println("\n示例3 - 混合边界的范围查询:")

	// 创建查询范围：1.0.0（包含） 到 3.0.0（不包含）
	startVersion3 := versions.NewVersion("1.0.0")
	endVersion3 := versions.NewVersion("3.0.0")

	// 创建范围边界条件
	startTuple3 := tuple.New2[*versions.Version, versions.ContainsPolicy](startVersion3, versions.ContainsPolicyYes) // 包含起始版本
	endTuple3 := tuple.New2[*versions.Version, versions.ContainsPolicy](endVersion3, versions.ContainsPolicyNo)      // 不包含结束版本

	// 执行范围查询
	rangeResult3 := sortedGroups.QueryRange(startTuple3, endTuple3)

	fmt.Printf("查询范围: %s（包含）到 %s（不包含）\n", startVersion3.Raw, endVersion3.Raw)
	fmt.Printf("查询结果数量: %d\n", len(rangeResult3))
	fmt.Println("查询结果:")
	for i, version := range versions.SortVersionSlice(rangeResult3) {
		fmt.Printf("  %d. %s\n", i+1, version.Raw)
	}

	/*
		输出示例:
		示例3 - 混合边界的范围查询:
		查询范围: 1.0.0（包含）到 3.0.0（不包含）
		查询结果数量: 8
		查询结果:
		  1. 1.0.0
		  2. 1.0.1
		  3. 1.1.0
		  4. 1.1.1
		  5. 1.2.0
		  6. 2.0.0
		  7. 2.0.1
		  8. 2.1.0
	*/

	// 示例4：查询特定主版本的所有版本
	fmt.Println("\n示例4 - 查询特定主版本的所有版本:")

	// 获取组ID，查找特定版本组的所有版本
	groupIDs := sortedGroups.GroupIDs()
	if len(groupIDs) >= 2 {
		// 假设我们要查询第2个版本组（通常是2.x.x）的所有版本
		groupID := groupIDs[1] // 数组下标是0开始的，所以 [1] 是第二个组

		fmt.Printf("查询版本组 '%s' 的所有版本:\n", groupID)

		// 仅作演示，我们可以从之前的分组结果中获取
		versionGroups := versions.Group(testVersions)
		groupVersions := versions.SortVersionSlice(versionGroups[groupID].Versions())

		for i, version := range groupVersions {
			fmt.Printf("  %d. %s\n", i+1, version.Raw)
		}
	}

	/*
		输出示例:
		示例4 - 查询特定主版本的所有版本:
		查询版本组 '2' 的所有版本:
		  1. 2.0.0
		  2. 2.0.1
		  3. 2.1.0
	*/
}
