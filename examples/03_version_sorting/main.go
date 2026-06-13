package main

import (
	"fmt"
	"time"

	"github.com/scagogogo/versions-skills"
)

func main() {
	// 示例1：排序版本号字符串
	fmt.Println("示例1 - 排序版本号字符串:")
	versionStrings := []string{
		"2.0.0", "1.0.0", "1.10.0", "1.2.0", "1.1.0",
		"v3.0.0", "v2.0.0-beta", "1.0.0-alpha",
	}

	fmt.Println("排序前:")
	for i, v := range versionStrings {
		fmt.Printf("  %d. %s\n", i+1, v)
	}

	// 按语义版本顺序排序
	sortedVersionStrings := versions.SortVersionStringSlice(versionStrings)

	fmt.Println("\n排序后:")
	for i, v := range sortedVersionStrings {
		fmt.Printf("  %d. %s\n", i+1, v)
	}

	/*
		输出示例:
		示例1 - 排序版本号字符串:
		排序前:
		  1. 2.0.0
		  2. 1.0.0
		  3. 1.10.0
		  4. 1.2.0
		  5. 1.1.0
		  6. v3.0.0
		  7. v2.0.0-beta
		  8. 1.0.0-alpha

		排序后:
		  1. 1.0.0
		  2. 1.0.0-alpha
		  3. 1.1.0
		  4. 1.2.0
		  5. 1.10.0
		  6. 2.0.0
		  7. v2.0.0-beta
		  8. v3.0.0
	*/

	// 示例2：排序版本对象
	fmt.Println("\n示例2 - 排序版本对象:")
	versionObjs := []*versions.Version{
		versions.NewVersion("2.0.0"),
		versions.NewVersion("1.0.0"),
		versions.NewVersion("1.10.0"),
		versions.NewVersion("1.2.0"),
		versions.NewVersion("1.1.0"),
	}

	// 为部分版本添加发布时间
	versionObjs[1].PublicTime = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC) // 1.0.0
	versionObjs[4].PublicTime = time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC) // 1.1.0
	versionObjs[3].PublicTime = time.Date(2020, 2, 1, 0, 0, 0, 0, time.UTC) // 1.2.0

	fmt.Println("排序前:")
	for i, v := range versionObjs {
		timeStr := ""
		if !v.PublicTime.IsZero() {
			timeStr = fmt.Sprintf(", 发布时间: %s", v.PublicTime.Format("2006-01-02"))
		}
		fmt.Printf("  %d. %s%s\n", i+1, v.Raw, timeStr)
	}

	// 排序版本对象切片
	sortedVersionObjs := versions.SortVersionSlice(versionObjs)

	fmt.Println("\n排序后:")
	for i, v := range sortedVersionObjs {
		timeStr := ""
		if !v.PublicTime.IsZero() {
			timeStr = fmt.Sprintf(", 发布时间: %s", v.PublicTime.Format("2006-01-02"))
		}
		fmt.Printf("  %d. %s%s\n", i+1, v.Raw, timeStr)
	}

	/*
		输出示例:
		示例2 - 排序版本对象:
		排序前:
		  1. 2.0.0
		  2. 1.0.0, 发布时间: 2020-01-01
		  3. 1.10.0
		  4. 1.2.0, 发布时间: 2020-02-01
		  5. 1.1.0, 发布时间: 2020-03-01

		排序后:
		  1. 1.0.0, 发布时间: 2020-01-01
		  2. 1.1.0, 发布时间: 2020-03-01
		  3. 1.2.0, 发布时间: 2020-02-01
		  4. 1.10.0
		  5. 2.0.0
	*/

	// 示例3：排序带前缀和后缀的版本
	fmt.Println("\n示例3 - 排序带前缀和后缀的版本:")
	mixedVersions := []*versions.Version{
		versions.NewVersion("1.0.0"),
		versions.NewVersion("v1.0.0"),
		versions.NewVersion("1.0.0-alpha"),
		versions.NewVersion("1.0.0-beta"),
		versions.NewVersion("v1.0.0-rc1"),
		versions.NewVersion("v1.0.0-rc2"),
	}

	fmt.Println("排序前:")
	for i, v := range mixedVersions {
		fmt.Printf("  %d. %s (前缀: '%s', 后缀: '%s')\n", i+1, v.Raw, v.Prefix, v.Suffix)
	}

	// 排序混合版本
	sortedMixedVersions := versions.SortVersionSlice(mixedVersions)

	fmt.Println("\n排序后:")
	for i, v := range sortedMixedVersions {
		fmt.Printf("  %d. %s (前缀: '%s', 后缀: '%s')\n", i+1, v.Raw, v.Prefix, v.Suffix)
	}

	/*
		输出示例:
		示例3 - 排序带前缀和后缀的版本:
		排序前:
		  1. 1.0.0 (前缀: '', 后缀: '')
		  2. v1.0.0 (前缀: 'v', 后缀: '')
		  3. 1.0.0-alpha (前缀: '', 后缀: '-alpha')
		  4. 1.0.0-beta (前缀: '', 后缀: '-beta')
		  5. v1.0.0-rc1 (前缀: 'v', 后缀: '-rc1')
		  6. v1.0.0-rc2 (前缀: 'v', 后缀: '-rc2')

		排序后:
		  1. 1.0.0 (前缀: '', 后缀: '')
		  2. 1.0.0-alpha (前缀: '', 后缀: '-alpha')
		  3. 1.0.0-beta (前缀: '', 后缀: '-beta')
		  4. v1.0.0 (前缀: 'v', 后缀: '')
		  5. v1.0.0-rc1 (前缀: 'v', 后缀: '-rc1')
		  6. v1.0.0-rc2 (前缀: 'v', 后缀: '-rc2')
	*/
}
