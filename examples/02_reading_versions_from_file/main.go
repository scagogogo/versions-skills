package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/scagogogo/versions-skills"
)

func main() {
	// 获取版本文件的绝对路径
	versionFilePath := filepath.Join("examples", "02_reading_versions_from_file", "versions.txt")

	// 示例1：读取文件中的版本号字符串
	fmt.Println("示例1 - 读取文件中的版本号字符串:")
	versionStrings, err := versions.ReadVersionsStringFromFile(versionFilePath)
	if err != nil {
		log.Fatalf("读取版本号字符串失败: %v", err)
	}

	fmt.Printf("从文件读取到 %d 个版本号字符串:\n", len(versionStrings))
	for i, versionStr := range versionStrings {
		fmt.Printf("  %d. %s\n", i+1, versionStr)
	}

	/*
		输出示例:
		示例1 - 读取文件中的版本号字符串:
		从文件读取到 8 个版本号字符串:
		  1. 1.0.0
		  2. 1.0.1
		  3. v1.1.0
		  4. v1.1.1-beta
		  5. 2.0.0
		  6. 2.0.1-rc1
		  7. v2.1.0
		  8. 3.0.0
	*/

	// 示例2：读取并解析文件中的版本号
	fmt.Println("\n示例2 - 读取并解析文件中的版本号:")
	versionObjects, err := versions.ReadVersionsFromFile(versionFilePath)
	if err != nil {
		log.Fatalf("读取版本号对象失败: %v", err)
	}

	fmt.Printf("从文件读取并解析到 %d 个版本号对象:\n", len(versionObjects))
	for i, version := range versionObjects {
		fmt.Printf("  %d. 原始字符串: %s, 版本号数字: %v, 前缀: '%s', 后缀: '%s'\n",
			i+1, version.Raw, version.VersionNumbers, version.Prefix, version.Suffix)
	}

	/*
		输出示例:
		示例2 - 读取并解析文件中的版本号:
		从文件读取并解析到 8 个版本号对象:
		  1. 原始字符串: 1.0.0, 版本号数字: [1 0 0], 前缀: '', 后缀: ''
		  2. 原始字符串: 1.0.1, 版本号数字: [1 0 1], 前缀: '', 后缀: ''
		  3. 原始字符串: v1.1.0, 版本号数字: [1 1 0], 前缀: 'v', 后缀: ''
		  4. 原始字符串: v1.1.1-beta, 版本号数字: [1 1 1], 前缀: 'v', 后缀: '-beta'
		  5. 原始字符串: 2.0.0, 版本号数字: [2 0 0], 前缀: '', 后缀: ''
		  6. 原始字符串: 2.0.1-rc1, 版本号数字: [2 0 1], 前缀: '', 后缀: '-rc1'
		  7. 原始字符串: v2.1.0, 版本号数字: [2 1 0], 前缀: 'v', 后缀: ''
		  8. 原始字符串: 3.0.0, 版本号数字: [3 0 0], 前缀: '', 后缀: ''
	*/

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

	/*
		输出示例:
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
}
