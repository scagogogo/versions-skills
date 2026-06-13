package main

import (
	"fmt"
	"time"

	"github.com/scagogogo/versions-skills"
)

func main() {
	fmt.Println("=== 新版本示例程序 ===")
	fmt.Println("当前时间:", time.Now().Format("2006-01-02 15:04:05"))

	// 硬编码版本字符串
	versionStrings := []string{
		"1.0.0",
		"1.0.1",
		"v1.1.0",
		"v1.1.1-beta",
		"2.0.0",
		"2.0.1-rc1",
		"v2.1.0",
		"3.0.0",
	}

	// 示例1：使用版本号字符串
	fmt.Println("\n示例1 - 直接使用硬编码版本号字符串:")
	fmt.Printf("版本号字符串数量: %d\n", len(versionStrings))
	for i, versionStr := range versionStrings {
		fmt.Printf("  %d. %s\n", i+1, versionStr)
	}

	// 示例2：解析版本号字符串为版本对象
	fmt.Println("\n示例2 - 解析版本号字符串为版本对象:")
	versionObjects := make([]*versions.Version, 0, len(versionStrings))
	for _, versionStr := range versionStrings {
		versionObjects = append(versionObjects, versions.NewVersion(versionStr))
	}

	fmt.Printf("解析到 %d 个版本号对象:\n", len(versionObjects))
	for i, version := range versionObjects {
		fmt.Printf("  %d. 原始字符串: %s, 版本号数字: %v, 前缀: '%s', 后缀: '%s'\n",
			i+1, version.Raw, version.VersionNumbers, version.Prefix, version.Suffix)
	}

	// 示例3：过滤有效的版本号
	fmt.Println("\n示例3 - 过滤有效的版本号:")
	validVersions := make([]*versions.Version, 0)
	for _, version := range versionObjects {
		if version.IsValid() {
			validVersions = append(validVersions, version)
		}
	}

	fmt.Printf("有效的版本号数量: %d\n", len(validVersions))
	for i, version := range validVersions {
		fmt.Printf("  %d. %s\n", i+1, version.Raw)
	}
}

/*
输出示例:
=== 新版本示例程序 ===
当前时间: 2025-04-22 15:02:56

示例1 - 直接使用硬编码版本号字符串:
版本号字符串数量: 8
  1. 1.0.0
  2. 1.0.1
  3. v1.1.0
  4. v1.1.1-beta
  5. 2.0.0
  6. 2.0.1-rc1
  7. v2.1.0
  8. 3.0.0

示例2 - 解析版本号字符串为版本对象:
解析到 8 个版本号对象:
  1. 原始字符串: 1.0.0, 版本号数字: [1 0 0], 前缀: '', 后缀: ''
  2. 原始字符串: 1.0.1, 版本号数字: [1 0 1], 前缀: '', 后缀: ''
  3. 原始字符串: v1.1.0, 版本号数字: [1 1 0], 前缀: 'v', 后缀: ''
  4. 原始字符串: v1.1.1-beta, 版本号数字: [1 1 1], 前缀: 'v', 后缀: '-beta'
  5. 原始字符串: 2.0.0, 版本号数字: [2 0 0], 前缀: '', 后缀: ''
  6. 原始字符串: 2.0.1-rc1, 版本号数字: [2 0 1], 前缀: '', 后缀: '-rc1'
  7. 原始字符串: v2.1.0, 版本号数字: [2 1 0], 前缀: 'v', 后缀: ''
  8. 原始字符串: 3.0.0, 版本号数字: [3 0 0], 前缀: '', 后缀: ''

示例3 - 过滤有效的版本号:
有效的版本号数量: 8
  1. 1.0.0
  2. 1.0.1
  3. v1.1.0
  4. v1.1.1-beta
  5. 2.0.0
  6. 2.0.1-rc1
  7. v2.1.0
  8. 3.0.0
*/
