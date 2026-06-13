package main

import (
	"fmt"
	"time"

	"github.com/scagogogo/versions-skills"
)

func main() {
	// 示例1：解析简单版本号
	version1 := versions.NewVersion("1.2.3")
	fmt.Println("示例1 - 解析简单版本号 1.2.3:")
	fmt.Printf("  原始字符串: %s\n", version1.Raw)
	fmt.Printf("  版本号数字: %v\n", version1.VersionNumbers)
	fmt.Printf("  前缀: %s\n", version1.Prefix)
	fmt.Printf("  后缀: %s\n", version1.Suffix)

	/*
		输出:
		示例1 - 解析简单版本号 1.2.3:
		  原始字符串: 1.2.3
		  版本号数字: [1 2 3]
		  前缀:
		  后缀:
	*/

	// 示例2：解析带前缀的版本号
	version2 := versions.NewVersion("v2.0.0")
	fmt.Println("\n示例2 - 解析带前缀的版本号 v2.0.0:")
	fmt.Printf("  原始字符串: %s\n", version2.Raw)
	fmt.Printf("  版本号数字: %v\n", version2.VersionNumbers)
	fmt.Printf("  前缀: %s\n", version2.Prefix)
	fmt.Printf("  后缀: %s\n", version2.Suffix)

	/*
		输出:
		示例2 - 解析带前缀的版本号 v2.0.0:
		  原始字符串: v2.0.0
		  版本号数字: [2 0 0]
		  前缀: v
		  后缀:
	*/

	// 示例3：解析带后缀的版本号
	version3 := versions.NewVersion("1.0.0-beta")
	fmt.Println("\n示例3 - 解析带后缀的版本号 1.0.0-beta:")
	fmt.Printf("  原始字符串: %s\n", version3.Raw)
	fmt.Printf("  版本号数字: %v\n", version3.VersionNumbers)
	fmt.Printf("  前缀: %s\n", version3.Prefix)
	fmt.Printf("  后缀: %s\n", version3.Suffix)

	/*
		输出:
		示例3 - 解析带后缀的版本号 1.0.0-beta:
		  原始字符串: 1.0.0-beta
		  版本号数字: [1 0 0]
		  前缀:
		  后缀: -beta
	*/

	// 示例4：带发布时间的版本号
	version4 := versions.NewVersion("3.1.4")
	// 设置发布时间
	version4.PublicTime = time.Date(2023, 4, 15, 0, 0, 0, 0, time.UTC)
	fmt.Println("\n示例4 - 带发布时间的版本号:")
	fmt.Printf("  原始字符串: %s\n", version4.Raw)
	fmt.Printf("  版本号数字: %v\n", version4.VersionNumbers)
	fmt.Printf("  发布时间: %s\n", version4.PublicTime.Format("2006-01-02"))

	/*
		输出:
		示例4 - 带发布时间的版本号:
		  原始字符串: 3.1.4
		  版本号数字: [3 1 4]
		  发布时间: 2023-04-15
	*/

	// 示例5：比较两个版本号
	version5a := versions.NewVersion("1.2.3")
	version5b := versions.NewVersion("1.3.0")
	fmt.Println("\n示例5 - 比较两个版本号:")
	fmt.Printf("  %s < %s: %t\n", version5a.Raw, version5b.Raw, version5a.CompareTo(version5b) < 0)
	fmt.Printf("  %s > %s: %t\n", version5a.Raw, version5b.Raw, version5a.CompareTo(version5b) > 0)

	/*
		输出:
		示例5 - 比较两个版本号:
		  1.2.3 < 1.3.0: true
		  1.2.3 > 1.3.0: false
	*/
}

/*
输出示例:
示例1 - 解析简单版本号 1.2.3:
  原始字符串: 1.2.3
  版本号数字: [1 2 3]
  前缀:
  后缀:

示例2 - 解析带前缀的版本号 v2.0.0:
  原始字符串: v2.0.0
  版本号数字: [2 0 0]
  前缀: v
  后缀:

示例3 - 解析带后缀的版本号 1.0.0-beta:
  原始字符串: 1.0.0-beta
  版本号数字: [1 0 0]
  前缀:
  后缀: -beta

示例4 - 带发布时间的版本号:
  原始字符串: 3.1.4
  版本号数字: [3 1 4]
  发布时间: 2023-04-15

示例5 - 比较两个版本号:
  1.2.3 < 1.3.0: true
  1.2.3 > 1.3.0: false
*/
