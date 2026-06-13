package main

import (
	"fmt"
	"os"
	"time"

	"github.com/scagogogo/versions-skills"
)

func main() {
	// 创建一组测试版本，包含不同的主版本和子版本
	testVersions := []*versions.Version{
		// 1.x 系列版本
		versions.NewVersion("1.0.0"),
		versions.NewVersion("1.0.1"),
		versions.NewVersion("1.1.0"),
		versions.NewVersion("1.1.1"),
		versions.NewVersion("1.2.0"),

		// 2.x 系列版本
		versions.NewVersion("2.0.0"),
		versions.NewVersion("2.0.1"),
		versions.NewVersion("v2.1.0-beta"),
		versions.NewVersion("v2.1.0"),
		versions.NewVersion("v2.1.1"),

		// 3.x 系列版本
		versions.NewVersion("3.0.0-rc1"),
		versions.NewVersion("3.0.0-rc2"),
		versions.NewVersion("3.0.0"),
		versions.NewVersion("v3.0.1"),
		versions.NewVersion("v3.1.0"),
	}

	// 添加一些发布时间信息
	testVersions[0].PublicTime = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)  // 1.0.0
	testVersions[1].PublicTime = time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)  // 1.0.1
	testVersions[5].PublicTime = time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)  // 2.0.0
	testVersions[12].PublicTime = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC) // 3.0.0

	// 示例1：可视化所有版本
	fmt.Println("示例1 - 可视化所有版本:")
	fmt.Println("------------------------------")
	versions.VisualizeVersions(testVersions, os.Stdout, 0) // 0表示不限制每个组显示的版本数

	// 示例2：限制每个组显示的版本数量
	fmt.Println("\n\n示例2 - 限制每个组显示的版本数量（每组最多显示2个）:")
	fmt.Println("------------------------------")
	versions.VisualizeVersions(testVersions, os.Stdout, 2) // 每组最多显示2个版本

	// 示例3：可视化版本组层次结构
	fmt.Println("\n\n示例3 - 可视化版本组层次结构:")
	fmt.Println("------------------------------")
	versions.VisualizeVersionGroups(testVersions, os.Stdout)
}

/*
输出示例:
示例1 - 可视化所有版本:
------------------------------
版本总数: 15
版本组数: 3

┌─ 版本组: 1 (5个版本)
├── 1.0.0 (发布时间: 2021-01-01)
├── 1.0.1 (发布时间: 2021-02-01)
├── 1.1.0
├── 1.1.1
└── 1.2.0

┌─ 版本组: 2 (5个版本)
├── 2.0.0 (发布时间: 2022-01-01)
├── 2.0.1
├── v2.1.0-beta
├── v2.1.0
└── v2.1.1

┌─ 版本组: 3 (5个版本)
├── 3.0.0-rc1
├── 3.0.0-rc2
├── 3.0.0 (发布时间: 2023-01-01)
├── v3.0.1
└── v3.1.0


示例2 - 限制每个组显示的版本数量（每组最多显示2个）:
------------------------------
版本总数: 15
版本组数: 3

┌─ 版本组: 1 (5个版本)
├── 1.0.0 (发布时间: 2021-01-01)
├── 1.0.1 (发布时间: 2021-02-01)
└── ...还有3个版本未显示

┌─ 版本组: 2 (5个版本)
├── 2.0.0 (发布时间: 2022-01-01)
├── 2.0.1
└── ...还有3个版本未显示

┌─ 版本组: 3 (5个版本)
├── 3.0.0-rc1
├── 3.0.0-rc2
└── ...还有3个版本未显示


示例3 - 可视化版本组层次结构:
------------------------------
版本总数: 15
版本组数: 3

├─ 1 (5个版本)
│ ├─ 1.0 (2个版本)
│ ├─ 1.1 (2个版本)
│ └─ 1.2 (1个版本)
├─ 2 (5个版本)
│ ├─ 2.0 (2个版本)
│ └─ 2.1 (3个版本)
└─ 3 (5个版本)
  ├─ 3.0 (4个版本)
  └─ 3.1 (1个版本)
*/
