package main

import (
	"fmt"

	"github.com/scagogogo/versions-skills"
)

func main() {
	// 创建一组包含不同版本号的版本对象
	versionObjs := []*versions.Version{
		versions.NewVersion("1.0.0"),
		versions.NewVersion("1.0.1"),
		versions.NewVersion("1.1.0"),
		versions.NewVersion("v1.1.1"),
		versions.NewVersion("2.0.0"),
		versions.NewVersion("2.0.1"),
		versions.NewVersion("v2.1.0"),
		versions.NewVersion("v2.1.1-beta"),
		versions.NewVersion("3.0.0"),
		versions.NewVersion("v3.0.0-rc1"),
	}

	// 示例1：基本版本分组
	fmt.Println("示例1 - 基本版本分组:")
	versionGroups := versions.Group(versionObjs)

	fmt.Printf("版本总数: %d\n", len(versionObjs))
	fmt.Printf("版本组数: %d\n", len(versionGroups))

	for groupID, group := range versionGroups {
		fmt.Printf("\n版本组 '%s': (包含 %d 个版本)\n", groupID, len(group.Versions()))
		for i, version := range group.Versions() {
			fmt.Printf("  %d. %s\n", i+1, version.Raw)
		}
	}

	// 示例1输出:
	// 示例1 - 基本版本分组:
	// 版本总数: 10
	// 版本组数: 3
	//
	// 版本组 '1': (包含 4 个版本)
	//   1. 1.0.0
	//   2. 1.0.1
	//   3. 1.1.0
	//   4. v1.1.1
	//
	// 版本组 '2': (包含 4 个版本)
	//   1. 2.0.0
	//   2. 2.0.1
	//   3. v2.1.0
	//   4. v2.1.1-beta
	//
	// 版本组 '3': (包含 2 个版本)
	//   1. 3.0.0
	//   2. v3.0.0-rc1

	// 示例2：创建有序版本组
	fmt.Println("\n\n示例2 - 创建有序版本组:")
	sortedGroups := versions.NewSortedVersionGroups(versionObjs)

	// 获取所有组的ID
	groupIDs := sortedGroups.GroupIDs()

	fmt.Printf("版本总数: %d\n", len(versionObjs))
	fmt.Printf("版本组数: %d\n", len(groupIDs))

	// 遍历每个组ID，输出组信息
	for _, groupID := range groupIDs {
		group := versionGroups[groupID]
		fmt.Printf("\n版本组 '%s': (包含 %d 个版本)\n", groupID, len(group.Versions()))

		// 获取并排序每个组内的版本
		groupVersions := versions.SortVersionSlice(group.Versions())
		for i, version := range groupVersions {
			fmt.Printf("  %d. %s\n", i+1, version.Raw)
		}
	}

	// 示例2输出:
	// 示例2 - 创建有序版本组:
	// 版本总数: 10
	// 版本组数: 3
	//
	// 版本组 '1': (包含 4 个版本)
	//   1. 1.0.0
	//   2. 1.0.1
	//   3. 1.1.0
	//   4. v1.1.1
	//
	// 版本组 '2': (包含 4 个版本)
	//   1. 2.0.0
	//   2. 2.0.1
	//   3. v2.1.0
	//   4. v2.1.1-beta
	//
	// 版本组 '3': (包含 2 个版本)
	//   1. 3.0.0
	//   2. v3.0.0-rc1

	// 示例3：版本组元数据
	fmt.Println("\n\n示例3 - 版本组元数据:")
	for _, groupID := range groupIDs {
		group := versionGroups[groupID]
		// 选择任意一个版本获取其前缀和后缀作为示例
		sampleVersion := group.Versions()[0]
		fmt.Printf("版本组 '%s': 示例版本='%s', 组ID='%s'\n",
			groupID, sampleVersion.Raw, group.ID())
	}

	// 示例3输出:
	// 示例3 - 版本组元数据:
	// 版本组 '1': 示例版本='1.0.0', 组ID='1'
	// 版本组 '2': 示例版本='2.0.0', 组ID='2'
	// 版本组 '3': 示例版本='3.0.0', 组ID='3'
}

/*
输出示例:
示例1 - 基本版本分组:
版本总数: 10
版本组数: 3

版本组 '1': (包含 4 个版本)
  1. 1.0.0
  2. 1.0.1
  3. 1.1.0
  4. v1.1.1

版本组 '2': (包含 4 个版本)
  1. 2.0.0
  2. 2.0.1
  3. v2.1.0
  4. v2.1.1-beta

版本组 '3': (包含 2 个版本)
  1. 3.0.0
  2. v3.0.0-rc1


示例2 - 创建有序版本组:
版本总数: 10
版本组数: 3

版本组 '1': (包含 4 个版本)
  1. 1.0.0
  2. 1.0.1
  3. 1.1.0
  4. v1.1.1

版本组 '2': (包含 4 个版本)
  1. 2.0.0
  2. 2.0.1
  3. v2.1.0
  4. v2.1.1-beta

版本组 '3': (包含 2 个版本)
  1. 3.0.0
  2. v3.0.0-rc1


示例3 - 版本组元数据:
版本组 '1': 示例版本='1.0.0', 组ID='1'
版本组 '2': 示例版本='2.0.0', 组ID='2'
版本组 '3': 示例版本='3.0.0', 组ID='3'
*/
