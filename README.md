# Versions - Go语言版本号解析计算SDK

<div align="center">

[![Go Tests](https://github.com/scagogogo/versions/actions/workflows/go-test.yml/badge.svg)](https://github.com/scagogogo/versions/actions/workflows/go-test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/versions)](https://goreportcard.com/report/github.com/scagogogo/versions)
[![GoDoc](https://godoc.org/github.com/scagogogo/versions?status.svg)](https://godoc.org/github.com/scagogogo/versions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

<img src="https://user-images.githubusercontent.com/5877/236610549-d20056f0-db64-4ba4-aabd-4f3cf78fb8d5.png" alt="Versions Logo" width="180"/>

</div>

<p align="center">
<b>专业的 Go 语言版本号解析计算SDK - 轻松解析、比较、排序软件版本号</b>
</p>

`versions` 是一个专为 Go 开发者设计的版本号解析计算SDK，专注于处理语义化版本号的解析、比较、排序和查询。它不是一个版本管理系统，而是一个帮助开发者处理版本号字符串的工具库。无论是依赖管理、API兼容性检查还是软件更新逻辑，都能高效完成版本号的计算需求。

---

## 📋 目录

- [✨ 特性](#-特性)
- [📦 安装](#-安装)
- [🚀 快速开始](#-快速开始)
- [📚 详细文档](#-详细文档)
  - [数据类型和常量](#数据类型和常量)
  - [主要函数](#主要函数)
- [🔍 使用示例](#-使用示例)
- [⚠️ 注意事项](#-注意事项)
- [📈 性能](#-性能)
- [📄 许可证](#-许可证)

---

## ✨ 特性

<table>
  <tr>
    <td><b>🔄 全面的版本号支持</b></td>
    <td>支持标准语义化版本格式（如 <code>1.2.3</code>）和多种变体（如 <code>v1.2.3</code>、<code>1.2.3-beta</code> 等）</td>
  </tr>
  <tr>
    <td><b>🧩 灵活的版本号解析</b></td>
    <td>自动识别前缀、版本号和后缀，处理各种版本号格式</td>
  </tr>
  <tr>
    <td><b>📊 版本号比较</b></td>
    <td>基于标准语义化版本规则进行版本号比较，支持前缀和后缀处理</td>
  </tr>
  <tr>
    <td><b>📦 版本号分组和排序</b></td>
    <td>按主版本号、次版本号分组，并提供多种排序方式</td>
  </tr>
  <tr>
    <td><b>🔍 版本号范围查询</b></td>
    <td>支持查询指定版本号范围内的所有版本，带有灵活的包含/排除边界选项</td>
  </tr>
  <tr>
    <td><b>📋 版本号可视化</b></td>
    <td>提供文本方式展示版本号之间的层次关系，直观查看版本号组织结构</td>
  </tr>
  <tr>
    <td><b>📁 文件支持</b></td>
    <td>直接从文件中读取和处理版本号</td>
  </tr>
  <tr>
    <td><b>🚀 无外部依赖</b></td>
    <td>核心功能无需额外依赖，轻量快速</td>
  </tr>
</table>

---

## 📦 安装

使用 `go get` 命令安装:

```bash
go get -u github.com/scagogogo/versions
```

---

## 🚀 快速开始

以下是一个简单的示例，展示如何使用 `versions` 库解析和比较版本号:

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions"
)

func main() {
    // 创建版本对象
    v1 := versions.NewVersion("1.2.3")
    v2 := versions.NewVersion("v1.3.0")
    
    // 比较版本大小
    if v1.CompareTo(v2) < 0 {
        fmt.Printf("%s 小于 %s\n", v1.Raw, v2.Raw)
    }
    
    // 查看版本组成部分
    fmt.Printf("版本号数字: %v\n", v1.VersionNumbers)
    fmt.Printf("前缀: %s\n", v2.Prefix)  // 输出: "v"
    
    // 排序版本号
    versionList := []*versions.Version{
        versions.NewVersion("2.0.0"),
        versions.NewVersion("1.0.0"),
        versions.NewVersion("1.10.0"),
    }
    sortedVersions := versions.SortVersionSlice(versionList)
    for _, v := range sortedVersions {
        fmt.Println(v.Raw)  // 输出: 1.0.0, 1.10.0, 2.0.0
    }
}
```

<details open>
<summary>查看输出结果</summary>

```
1.2.3 小于 v1.3.0
版本号数字: [1 2 3]
前缀: v
1.0.0
1.10.0
2.0.0
```
</details>

---

## 📚 详细文档

### 数据类型和常量

<div align="center">

| 类型 | 描述 |
|:------:|:-----|
| <kbd>Version</kbd> | 表示一个版本号，包含原始字符串、版本号数字、前缀、后缀和发布时间 |
| <kbd>VersionNumbers</kbd> | 整数切片，表示版本号中的数字部分 |
| <kbd>VersionPrefix</kbd> | 字符串，表示版本号数字部分之前的前缀 |
| <kbd>VersionSuffix</kbd> | 字符串，表示版本号数字部分之后的后缀 |
| <kbd>ContainsPolicy</kbd> | 用于控制版本查询时的包含策略（包含、不包含） |
| <kbd>VersionGroup</kbd> | 版本组，包含相同主版本号的一组版本 |
| <kbd>SortedVersionGroups</kbd> | 有序的版本组集合，便于范围查询 |

</div>

### 主要函数

<details open>
<summary><b>版本解析与创建</b></summary>

```go
// 创建版本对象
version := versions.NewVersion("1.2.3")

// 带错误检查的版本创建
version, err := versions.NewVersionE("1.2.3")
if err != nil {
    log.Fatal(err)
}
```
</details>

<details open>
<summary><b>从文件读取版本</b></summary>

```go
// 读取版本号对象
versions, err := versions.ReadVersionsFromFile("path/to/versions.txt")
if err != nil {
    log.Fatal(err)
}

// 读取版本号字符串
versionStrings, err := versions.ReadVersionsStringFromFile("path/to/versions.txt")
if err != nil {
    log.Fatal(err)
}
```
</details>

<details open>
<summary><b>版本分组与排序</b></summary>

```go
// 版本分组
groupedVersions := versions.Group(versionList)

// 字符串版本排序
sortedStrings := versions.SortVersionStringSlice(versionStrings)

// 版本对象排序
sortedVersions := versions.SortVersionSlice(versionList)
```
</details>

<details open>
<summary><b>版本范围查询</b></summary>

```go
// 创建有序版本组
sortedGroups := versions.NewSortedVersionGroups(versionList)

// 定义查询范围和包含策略
startVersion := versions.NewVersion("1.0.0")
endVersion := versions.NewVersion("2.0.0")
startTuple := tuple.New2[*versions.Version, versions.ContainsPolicy](
    startVersion, versions.ContainsPolicyYes) // 包含起始版本
endTuple := tuple.New2[*versions.Version, versions.ContainsPolicy](
    endVersion, versions.ContainsPolicyNo)   // 不包含结束版本

// 执行范围查询
rangeResult := sortedGroups.QueryRange(startTuple, endTuple)
```
</details>

<details open>
<summary><b>版本可视化</b></summary>

```go
// 可视化所有版本（每组显示最多5个版本）
versions.VisualizeVersions(versionList, os.Stdout, 5)

// 可视化版本组层次结构
versions.VisualizeVersionGroups(versionList, os.Stdout)
```

**可视化输出示例:**

```
版本总数: 15
版本组数: 3

┌─ 版本组: 1.0 (3个版本)
├── 1.0.0 (发布时间: 2020-01-01)
├── 1.0.1 (发布时间: 2020-02-01)
└── 1.0.2 (发布时间: 2020-03-01)

┌─ 版本组: 2.0 (4个版本)
├── 2.0.0 (发布时间: 2021-01-01)
├── 2.0.1 (发布时间: 2021-02-01)
├── 2.0.2 (发布时间: 2021-03-01)
└── ...还有1个版本未显示
```
</details>

---

## 🧩 完整API文档

<div align="center">
<h3>核心类型与功能详解</h3>
</div>

### Version 类型

<details open>
<summary><b>结构定义</b></summary>

```go
type Version struct {
    // 原始版本号字符串
    Raw string
    
    // 版本发布时间
    PublicTime time.Time
    
    // 版本号数字部分，例如 1.2.3 中的 [1,2,3]
    VersionNumbers VersionNumbers
    
    // 版本号前缀，例如 v1.2.3 中的 "v"
    Prefix VersionPrefix
    
    // 版本号后缀，例如 1.2.3-beta 中的 "-beta"
    Suffix VersionSuffix
}
```
</details>

<details open>
<summary><b>NewVersion</b> - 创建版本号对象</summary>

```go
func NewVersion(versionString string) *Version
```

**参数:**
- `versionString string`: 版本号字符串，如 "1.2.3", "v1.0.0-beta" 等

**返回值:**
- `*Version`: 解析后的版本对象

**处理逻辑:**
1. 自动识别版本号前缀（如 "v"）
2. 解析版本号数字部分（如 "1.2.3" 中的 [1,2,3]）
3. 提取版本号后缀（如 "-beta", "-rc1" 等）
4. 创建完整的版本对象

**特性:**
- 支持任意数量的版本号段（如 "1.2.3.4.5"）
- 自动处理前导零（如 "1.02.003" 被解析为 [1,2,3]）
- 不会因解析错误而抛出异常，解析失败时会返回空版本号

**示例:**
```go
version := versions.NewVersion("v1.2.3-rc1")
fmt.Printf("前缀: %s, 版本号: %v, 后缀: %s\n", 
    version.Prefix, version.VersionNumbers, version.Suffix)
// 输出: 前缀: v, 版本号: [1 2 3], 后缀: -rc1

// 处理非标准版本号
custom := versions.NewVersion("release-1.0-final")
fmt.Printf("前缀: %s, 版本号: %v, 后缀: %s\n", 
    custom.Prefix, custom.VersionNumbers, custom.Suffix)
// 输出: 前缀: release-, 版本号: [1 0], 后缀: -final
```
</details>

<details open>
<summary><b>NewVersionE</b> - 创建版本号对象（带错误返回）</summary>

```go
func NewVersionE(versionString string) (*Version, error)
```

**参数:**
- `versionString string`: 版本号字符串

**返回值:**
- `*Version`: 解析后的版本对象
- `error`: 解析过程中可能发生的错误，如：
  - 无法识别的版本号格式
  - 版本号中不包含数字部分
  - 版本号数字部分解析失败

**错误处理:**
- 当版本号字符串为空时，返回 `ErrEmptyVersionString` 错误
- 当版本号中没有找到数字时，返回 `ErrNoVersionNumbersFound` 错误
- 当无法识别版本号格式时，返回 `ErrInvalidVersionFormat` 错误

**示例:**
```go
version, err := versions.NewVersionE("v1.2.3-rc1")
if err != nil {
    log.Fatalf("版本号解析失败: %v", err)
}

// 错误处理示例
version, err = versions.NewVersionE("")
if err != nil {
    fmt.Printf("空版本号错误: %v\n", err) 
    // 输出: 空版本号错误: empty version string
}

version, err = versions.NewVersionE("no-numbers")
if err != nil {
    fmt.Printf("无数字版本号错误: %v\n", err)
    // 输出: 无数字版本号错误: no version numbers found
}
```
</details>

<details open>
<summary><b>IsValid</b> - 检查版本号是否有效</summary>

```go
func (v *Version) IsValid() bool
```

**返回值:**
- `bool`: 版本号是否有效，必须至少包含一个版本数字

**验证标准:**
- 版本号对象不能为 nil
- VersionNumbers 字段必须至少包含一个数字
- Raw 字段不能为空字符串

**使用场景:**
- 过滤无效的版本号对象
- 验证用户输入的版本号是否有效
- 在批量处理版本号前进行有效性检查

**示例:**
```go
version := versions.NewVersion("v1.2.3")
if version.IsValid() {
    fmt.Println("版本号有效") // 会执行这行
}

emptyVersion := versions.NewVersion("")
if !emptyVersion.IsValid() {
    fmt.Println("空版本号无效") // 会执行这行
}

invalidVersion := versions.NewVersion("no-numbers")
if !invalidVersion.IsValid() {
    fmt.Println("无数字版本号无效") // 会执行这行
}
```
</details>

<details open>
<summary><b>CompareTo</b> - 比较两个版本号</summary>

```go
func (v *Version) CompareTo(other *Version) int
```

**参数:**
- `other *Version`: 要比较的另一个版本对象

**返回值:**
- `int`: 
  - 小于0: 当前版本小于 other 版本
  - 等于0: 两个版本相等
  - 大于0: 当前版本大于 other 版本

**比较规则:**
1. 首先比较版本号数字部分，按位比较，长度不同时短的版本号在缺失位置视为0
2. 如果版本号数字部分相同，则比较前缀
3. 如果前缀也相同，则比较后缀
4. 如果前缀和后缀都相同，则比较发布时间
5. 如果以上都相同，则认为两个版本相等

**兼容性说明:**
- 能正确处理数字部分不同长度的版本号比较（如 "1.0" 和 "1.0.0"）
- 前缀比较区分大小写（如 "v1.0" 和 "V1.0" 被视为不同）
- 后缀比较遵循预发布版本规则（如 "-alpha" < "-beta" < "-rc" < 正式版）

**示例:**
```go
v1 := versions.NewVersion("1.2.3")
v2 := versions.NewVersion("1.3.0")
result := v1.CompareTo(v2)
if result < 0 {
    fmt.Printf("%s 小于 %s\n", v1.Raw, v2.Raw) // 会执行这行
}

// 比较相同数字部分但有不同前后缀的版本
v3 := versions.NewVersion("v1.0.0")
v4 := versions.NewVersion("1.0.0-beta")
if v4.CompareTo(v3) < 0 {
    fmt.Println("预发布版本小于正式版本") // 会执行这行
}

// 比较发布时间
v5 := versions.NewVersion("1.0.0")
v6 := versions.NewVersion("1.0.0")
v5.PublicTime = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
v6.PublicTime = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
if v5.CompareTo(v6) < 0 {
    fmt.Println("早期发布的版本小于晚期发布的版本") // 会执行这行
}
```
</details>

<details open>
<summary><b>String</b> - 获取版本号字符串表示</summary>

```go
func (v *Version) String() string
```

**返回值:**
- `string`: 版本号的字符串表示，通常等同于原始版本号

**行为说明:**
- 如果 Raw 字段不为空，则直接返回 Raw 字段
- 如果 Raw 字段为空，则根据版本号各组成部分拼接成字符串返回
- 返回的字符串格式为：`前缀 + 版本号数字（以点分隔） + 后缀`

**使用场景:**
- 在日志或输出中显示版本号
- 将版本对象转换回字符串用于存储或传输
- 在比较后选择特定版本时获取其字符串表示

**示例:**
```go
version := versions.NewVersion("v1.2.3")
fmt.Println(version.String()) // 输出: v1.2.3

// 使用String()方法在日志中显示版本信息
log.Printf("当前使用的版本: %s", version)

// 通过比较后选择最大版本并显示
v1 := versions.NewVersion("1.0.0")
v2 := versions.NewVersion("2.0.0")
var latestVersion *versions.Version
if v1.CompareTo(v2) > 0 {
    latestVersion = v1
} else {
    latestVersion = v2
}
fmt.Printf("最新版本: %s\n", latestVersion) // 输出: 最新版本: 2.0.0
```
</details>

### VersionNumbers 类型

<details open>
<summary><b>结构定义与方法</b></summary>

```go
// VersionNumbers 是整数切片，表示版本号的数字部分
type VersionNumbers []int

// 获取主版本号
func (v VersionNumbers) MajorVersion() int

// 获取次版本号
func (v VersionNumbers) MinorVersion() int

// 获取修订版本号
func (v VersionNumbers) PatchVersion() int

// 比较两个版本号数字部分
func (v VersionNumbers) CompareTo(other VersionNumbers) int
```

**详细方法说明:**

**`MajorVersion()`**: 获取主版本号（第一个数字）
- **返回值:** `int` - 版本号的第一个数字
- **行为:** 如果版本号为空，返回0；否则返回第一个数字
- **用途:** 用于识别不兼容API变更的主版本

**`MinorVersion()`**: 获取次版本号（第二个数字）
- **返回值:** `int` - 版本号的第二个数字
- **行为:** 如果版本号不足两段，返回0；否则返回第二个数字
- **用途:** 用于识别向后兼容的功能性变更

**`PatchVersion()`**: 获取修订版本号（第三个数字）
- **返回值:** `int` - 版本号的第三个数字
- **行为:** 如果版本号不足三段，返回0；否则返回第三个数字
- **用途:** 用于识别向后兼容的问题修复

**`CompareTo(other VersionNumbers)`**: 比较两个版本号数字部分
- **参数:** `other VersionNumbers` - 要比较的另一个版本号数字
- **返回值:** `int` - 小于0表示当前版本小，等于0表示相等，大于0表示当前版本大
- **比较规则:**
  1. 从左到右逐位比较，高位优先
  2. 对于长度不同的版本号，缺失部分视为0（如1.0和1.0.0被视为相等）
  3. 数字大的版本号大（如1.2.0大于1.1.9）

**高级用法:**
- 可以处理任意长度的版本号（不限于语义化版本的三段式）
- 支持零值安全比较（空的VersionNumbers被视为[0]）
- 可直接访问底层切片获取特定位置的版本号数字

**示例:**
```go
version := versions.NewVersion("1.2.3")
major := version.VersionNumbers.MajorVersion() // 返回 1
minor := version.VersionNumbers.MinorVersion() // 返回 2
patch := version.VersionNumbers.PatchVersion() // 返回 3

// 比较版本号数字部分
v1Numbers := versions.NewVersion("1.2.0").VersionNumbers
v2Numbers := versions.NewVersion("1.2.3").VersionNumbers
if v1Numbers.CompareTo(v2Numbers) < 0 {
    fmt.Println("1.2.0 的数字部分小于 1.2.3") // 会执行这行
}

// 处理超过三段的版本号
longVersion := versions.NewVersion("1.2.3.4.5")
fmt.Printf("第5段版本号: %d\n", longVersion.VersionNumbers[4]) // 输出: 第5段版本号: 5

// 根据版本号数字部分生成分支名
version := versions.NewVersion("2.5.1")
branchName := fmt.Sprintf("release/%d.%d.x", 
    version.VersionNumbers.MajorVersion(),
    version.VersionNumbers.MinorVersion())
fmt.Println(branchName) // 输出: release/2.5.x
```
</details>

### VersionPrefix 类型

<details open>
<summary><b>结构定义与方法</b></summary>

```go
// VersionPrefix 是字符串，表示版本号前缀
type VersionPrefix string

// 检查前缀是否为空
func (v VersionPrefix) IsEmpty() bool

// 比较两个前缀
func (v VersionPrefix) CompareTo(other VersionPrefix) int
```

**详细方法说明:**

**`IsEmpty()`**: 检查前缀是否为空
- **返回值:** `bool` - 前缀是否为空字符串
- **行为:** 当前缀为空字符串时返回true，否则返回false
- **用途:** 判断版本号是否有前缀，通常用于决定是否需要特殊处理

**`CompareTo(other VersionPrefix)`**: 比较两个前缀
- **参数:** `other VersionPrefix` - 要比较的另一个前缀
- **返回值:** `int` - 小于0表示当前前缀字典序小，等于0表示相等，大于0表示当前前缀字典序大
- **比较规则:**
  1. 使用字符串字典序比较
  2. 区分大小写（如"v"和"V"被视为不同）
  3. 空前缀小于任何非空前缀

**常见前缀类型:**
- `v` - 常见的版本号前缀，如 "v1.2.3"
- `version-` - 一些项目使用的冗长前缀
- `release-` - 表示发布版本的前缀
- `Ver` - 带大写字母的变体

**使用场景:**
- 识别版本号的类型或来源
- 保持与特定工具或生态系统的兼容性
- 在显示时维持一致的格式

**示例:**
```go
version := versions.NewVersion("v1.2.3")
if !version.Prefix.IsEmpty() {
    fmt.Printf("版本前缀: %s\n", version.Prefix) // 输出: 版本前缀: v
}

// 比较前缀
v1 := versions.NewVersion("v1.0.0")
v2 := versions.NewVersion("release-1.0.0")
if v1.Prefix.CompareTo(v2.Prefix) < 0 {
    fmt.Println("v 前缀字典序小于 release- 前缀") // 不会执行，因为 "v" 字典序大于 "release-"
} else {
    fmt.Println("release- 前缀字典序小于 v 前缀") // 会执行这行
}

// 去除前缀获取纯版本号
version := versions.NewVersion("v2.0.1")
pureVersionStr := strings.TrimPrefix(version.Raw, string(version.Prefix))
fmt.Println(pureVersionStr) // 输出: 2.0.1
```
</details>

### VersionSuffix 类型

<details open>
<summary><b>结构定义与方法</b></summary>

```go
// VersionSuffix 是字符串，表示版本号后缀
type VersionSuffix string

// 检查后缀是否为空
func (v VersionSuffix) IsEmpty() bool

// 比较两个后缀
func (v VersionSuffix) CompareTo(other VersionSuffix) int
```

**详细方法说明:**

**`IsEmpty()`**: 检查后缀是否为空
- **返回值:** `bool` - 后缀是否为空字符串
- **行为:** 当后缀为空字符串时返回true，否则返回false
- **用途:** 判断是否为预发布版本，通常正式版本没有后缀

**`CompareTo(other VersionSuffix)`**: 比较两个后缀
- **参数:** `other VersionSuffix` - 要比较的另一个后缀
- **返回值:** `int` - 小于0表示当前后缀优先级低，等于0表示相等，大于0表示当前后缀优先级高
- **比较规则:**
  1. 空后缀大于任何非空后缀（正式版大于预发布版）
  2. 基于预发布版本通用优先级规则
  3. 对于相同类型的后缀，进一步按字符串比较

**常见后缀类型和优先级（从低到高）:**
1. `-dev`, `-alpha`, `-a` - 开发预览版/测试版
2. `-beta`, `-b` - 测试版
3. `-milestone`, `-m` - 里程碑版本
4. `-rc`, `-pre` - 发布候选版
5. 无后缀 - 正式发布版

**使用场景:**
- 标识预发布版本
- 确定版本稳定性和发布阶段
- 控制客户端更新策略（可能排除某些后缀类型）

**示例:**
```go
version := versions.NewVersion("1.2.3-beta")
if !version.Suffix.IsEmpty() {
    fmt.Printf("版本后缀: %s\n", version.Suffix) // 输出: 版本后缀: -beta
}

// 比较后缀优先级
v1 := versions.NewVersion("1.0.0-alpha")
v2 := versions.NewVersion("1.0.0-beta")
v3 := versions.NewVersion("1.0.0-rc1")
v4 := versions.NewVersion("1.0.0")

// 优先级: alpha < beta < rc < 正式版
if v1.Suffix.CompareTo(v2.Suffix) < 0 {
    fmt.Println("alpha 后缀优先级低于 beta 后缀") // 会执行这行
}
if v2.Suffix.CompareTo(v3.Suffix) < 0 {
    fmt.Println("beta 后缀优先级低于 rc 后缀") // 会执行这行
}
if v3.Suffix.CompareTo(v4.Suffix) < 0 {
    fmt.Println("rc 后缀优先级低于无后缀(正式版)") // 会执行这行
}

// 确定版本是否为预发布版本
if !version.Suffix.IsEmpty() {
    fmt.Println("这是预发布版本，不推荐用于生产环境")
}
```
</details>

### ContainsPolicy 类型

<details open>
<summary><b>定义与常量</b></summary>

```go
// ContainsPolicy 用于控制版本范围查询时是否包含边界版本
type ContainsPolicy int

const (
    // 未指定包含策略
    ContainsPolicyNone ContainsPolicy = iota
    
    // 包含边界版本
    ContainsPolicyYes
    
    // 不包含边界版本
    ContainsPolicyNo
)
```

**详细说明:**

**`ContainsPolicy`** 是一个枚举类型，用于在版本范围查询中指定是否包含边界版本。

**常量值:**

**`ContainsPolicyNone (0)`**: 
- **含义:** 未指定包含策略，通常使用默认策略
- **默认行为:** 在大多数上下文中等同于 ContainsPolicyYes
- **使用场景:** 当不关心边界包含性或希望使用系统默认行为时

**`ContainsPolicyYes (1)`**: 
- **含义:** 包含边界版本
- **符号表示:** 对应数学符号中的闭区间 `[`, `]`
- **使用场景:** 当查询需要包含指定的起始或结束版本时

**`ContainsPolicyNo (2)`**: 
- **含义:** 不包含边界版本
- **符号表示:** 对应数学符号中的开区间 `(`, `)`
- **使用场景:** 当查询需要排除指定的起始或结束版本时

**行为举例:**
- **[1.0.0, 2.0.0]**: 包含1.0.0和2.0.0及其间的所有版本
- **(1.0.0, 2.0.0)**: 不包含1.0.0和2.0.0，只包含其间的版本
- **[1.0.0, 2.0.0)**: 包含1.0.0但不包含2.0.0
- **(1.0.0, 2.0.0]**: 不包含1.0.0但包含2.0.0

**在实际代码中的应用:**
- 与版本范围查询函数配合使用
- 用于构建版本兼容性条件
- 实现依赖规范中定义的版本范围

**示例:**
```go
// 创建版本
v1 := versions.NewVersion("1.0.0")
v2 := versions.NewVersion("2.0.0")

// 创建版本范围查询条件
startWithInclude := tuple.New2[*versions.Version, versions.ContainsPolicy](
    v1, versions.ContainsPolicyYes) // 包含起始版本，相当于 [1.0.0
endWithExclude := tuple.New2[*versions.Version, versions.ContainsPolicy](
    v2, versions.ContainsPolicyNo)  // 不包含结束版本，相当于 2.0.0)

// 使用版本组执行范围查询（模拟 [1.0.0, 2.0.0) 范围）
result := sortedGroups.QueryRange(startWithInclude, endWithExclude)

// 在日志中显示查询范围
fmt.Printf("查询版本范围: %s%s, %s%s\n",
    inclusionSymbol(startWithInclude.V2), startWithInclude.V1.Raw,
    endWithExclude.V1.Raw, exclusionSymbol(endWithExclude.V2))
// 输出: 查询版本范围: [1.0.0, 2.0.0)

// 辅助函数示例
func inclusionSymbol(policy versions.ContainsPolicy) string {
    if policy == versions.ContainsPolicyNo {
        return "("
    }
    return "["
}

func exclusionSymbol(policy versions.ContainsPolicy) string {
    if policy == versions.ContainsPolicyNo {
        return ")"
    }
    return "]"
}
```
</details>

### VersionGroup 类型

<details open>
<summary><b>结构定义与方法</b></summary>

```go
// VersionGroup 表示具有相同主版本号的一组版本
type VersionGroup struct {
    // ...内部字段
}

// 创建新的版本组
func NewVersionGroup(id string) *VersionGroup

// 添加版本到组中
func (g *VersionGroup) Add(version *Version)

// 检查组是否包含某个版本
func (g *VersionGroup) Contains(version *Version) bool

// 获取组ID
func (g *VersionGroup) ID() string

// 获取组中的所有版本
func (g *VersionGroup) Versions() []*Version

// 获取组中版本的数量
func (g *VersionGroup) Count() int

// 按版本号排序组内的版本
func (g *VersionGroup) SortVersions() []*Version

// 查询范围内的版本
func (g *VersionGroup) QueryRangeVersions(start, end *Version) []*Version
```

**详细方法说明:**

**`NewVersionGroup(id string)`**: 创建新的版本组
- **参数:** `id string` - 组的标识符，通常是主版本号
- **返回值:** `*VersionGroup` - 新创建的版本组对象
- **行为:** 初始化一个空的版本组，带有指定的ID
- **用途:** 手动创建版本组，用于后续添加版本

**`Add(version *Version)`**: 添加版本到组中
- **参数:** `version *Version` - 要添加的版本对象
- **行为:** 将版本添加到组中，如果版本已存在则不重复添加
- **注意:** 不会验证版本的主版本号是否与组ID匹配
- **用途:** 向版本组添加新版本

**`Contains(version *Version) bool`**: 检查组是否包含某个版本
- **参数:** `version *Version` - 要检查的版本对象
- **返回值:** `bool` - 组中是否包含该版本
- **比较方式:** 使用版本的 Raw 字符串进行比较
- **用途:** 确定特定版本是否已在组中

**`ID() string`**: 获取组ID
- **返回值:** `string` - 组的标识符
- **行为:** 返回创建组时指定的ID
- **用途:** 识别版本组，通常等同于主版本号

**`Versions() []*Version`**: 获取组中的所有版本
- **返回值:** `[]*Version` - 组中所有版本的切片
- **行为:** 返回版本组内部存储的所有版本对象
- **注意:** 返回的切片不保证有特定顺序
- **用途:** 获取组内所有版本进行遍历或处理

**`Count() int`**: 获取组中版本的数量
- **返回值:** `int` - 组中版本的数量
- **行为:** 返回组中包含的版本对象数量
- **用途:** 快速获取组大小，无需遍历版本

**`SortVersions() []*Version`**: 按版本号排序组内的版本
- **返回值:** `[]*Version` - 排序后的版本对象切片
- **排序规则:** 使用 Version.CompareTo 方法比较，按版本号从小到大排序
- **用途:** 获取组内按序排列的版本列表

**`QueryRangeVersions(start, end *Version) []*Version`**: 查询范围内的版本
- **参数:** 
  - `start *Version` - 范围起始版本
  - `end *Version` - 范围结束版本
- **返回值:** `[]*Version` - 符合范围条件的版本对象切片
- **行为:** 返回组内版本号在 start 和 end 之间的所有版本（包含 start 和 end）
- **注意:** 返回的结果已按版本号排序
- **用途:** 获取特定版本范围内的所有版本

**版本组常见使用场景:**
- 按主版本号对版本进行分组管理
- 在UI中展示按主版本组织的版本列表
- 对特定主版本内的版本执行范围查询
- 获取特定主版本的最新版本或稳定版本

**示例:**
```go
// 创建版本组
group := versions.NewVersionGroup("1")

// 添加版本
group.Add(versions.NewVersion("1.0.0"))
group.Add(versions.NewVersion("1.1.0"))
group.Add(versions.NewVersion("1.2.0"))

// 获取组内所有版本
allVersions := group.Versions()
fmt.Printf("版本组 %s 包含 %d 个版本\n", group.ID(), group.Count())
// 输出: 版本组 1 包含 3 个版本

// 排序组内版本
sortedVersions := group.SortVersions()
fmt.Println("排序后的版本:")
for i, v := range sortedVersions {
    fmt.Printf("%d. %s\n", i+1, v.Raw)
}
// 输出:
// 排序后的版本:
// 1. 1.0.0
// 2. 1.1.0
// 3. 1.2.0

// 检查版本是否在组中
v := versions.NewVersion("1.1.0")
if group.Contains(v) {
    fmt.Printf("版本组 %s 包含版本 %s\n", group.ID(), v.Raw)
    // 输出: 版本组 1 包含版本 1.1.0
}

// 范围查询
start := versions.NewVersion("1.0.5")
end := versions.NewVersion("1.1.5")
rangeVersions := group.QueryRangeVersions(start, end)
fmt.Printf("范围 %s 到 %s 内的版本数: %d\n", start.Raw, end.Raw, len(rangeVersions))
// 输出: 范围 1.0.5 到 1.1.5 内的版本数: 1 (只有1.1.0在这个范围内)

// 获取组内最新版本
latestVersion := group.SortVersions()[group.Count()-1]
fmt.Printf("版本组 %s 中的最新版本: %s\n", group.ID(), latestVersion.Raw)
// 输出: 版本组 1 中的最新版本: 1.2.0
```
</details>

### SortedVersionGroups 类型

<details open>
<summary><b>结构定义与方法</b></summary>

```go
// SortedVersionGroups 表示一组有序的版本组
type SortedVersionGroups struct {
    // ...内部字段
}

// 创建新的有序版本组
func NewSortedVersionGroups(versions []*Version) *SortedVersionGroups

// 获取所有版本组ID
func (s *SortedVersionGroups) GroupIDs() []string

// 查询指定范围内的版本
func (s *SortedVersionGroups) QueryRange(
    start *tuple.Tuple2[*Version, ContainsPolicy],
    end *tuple.Tuple2[*Version, ContainsPolicy],
) []*Version
```

**详细方法说明:**

**`NewSortedVersionGroups(versions []*Version)`**: 创建新的有序版本组
- **参数:** `versions []*Version` - 版本对象切片
- **返回值:** `*SortedVersionGroups` - 有序版本组对象
- **行为:** 
  1. 将版本按主版本号分组创建 VersionGroup
  2. 对各组内版本进行排序
  3. 将版本组按组ID排序
- **用途:** 构建用于高效查询和遍历的有序版本结构

**`GroupIDs() []string`**: 获取所有版本组ID
- **返回值:** `[]string` - 所有版本组ID的切片
- **行为:** 返回已排序的版本组ID列表（通常是主版本号）
- **顺序:** 按字符串顺序排序，数字ID会考虑数值大小
- **用途:** 
  - 获取所有可用的主版本号
  - 在UI中创建版本选择器
  - 按主版本顺序遍历版本

**`QueryRange(...) []*Version`**: 查询指定范围内的版本
- **参数:** 
  - `start *tuple.Tuple2[*Version, ContainsPolicy]` - 起始版本及其包含策略
  - `end *tuple.Tuple2[*Version, ContainsPolicy]` - 结束版本及其包含策略
- **返回值:** `[]*Version` - 符合范围条件的版本对象切片
- **行为:** 
  1. 确定范围覆盖的版本组
  2. 在每个相关组内执行范围查询
  3. 合并结果并按版本顺序排序
- **特点:**
  - 支持跨版本组的范围查询
  - 考虑包含策略确定边界处理
  - 结果按版本顺序排序
- **用途:**
  - 实现版本选择器中的范围过滤
  - 查找指定范围内的兼容版本
  - 获取特定范围内的版本更新

**使用场景:**
- 在前端UI中展示分层的版本选择器
- 实现符合语义化版本规范的版本范围查询
- 处理大量版本时提供高效的版本过滤和分组显示

**性能特点:**
- 初始化时完成分组和排序，查询操作高效
- 范围查询利用排序特性，具有对数级时间复杂度
- 适合处理大规模版本集合，并支持频繁的查询操作

**示例:**
```go
// 创建测试版本列表
versionList := []*versions.Version{
    versions.NewVersion("1.0.0"),
    versions.NewVersion("1.1.0"),
    versions.NewVersion("2.0.0"),
    versions.NewVersion("2.1.0"),
    versions.NewVersion("3.0.0"),
}

// 创建有序版本组
sortedGroups := versions.NewSortedVersionGroups(versionList)

// 获取所有组ID
groupIDs := sortedGroups.GroupIDs()
fmt.Printf("共有 %d 个版本组: %v\n", len(groupIDs), groupIDs)
// 输出: 共有 3 个版本组: [1 2 3]

// 执行范围查询：获取1.0.0（包含）到2.0.0（不包含）之间的所有版本
startVersion := versions.NewVersion("1.0.0")
endVersion := versions.NewVersion("2.0.0")

// 创建包含策略元组
startTuple := tuple.New2[*versions.Version, versions.ContainsPolicy](
    startVersion, versions.ContainsPolicyYes) // 包含起始版本
endTuple := tuple.New2[*versions.Version, versions.ContainsPolicy](
    endVersion, versions.ContainsPolicyNo)    // 不包含结束版本
    
// 执行查询
result := sortedGroups.QueryRange(startTuple, endTuple)
fmt.Printf("查询结果包含 %d 个版本\n", len(result))
// 输出: 查询结果包含 2 个版本 (1.0.0, 1.1.0)

// 打印查询结果
for i, v := range result {
    fmt.Printf("%d. %s\n", i+1, v.Raw)
}
// 输出:
// 1. 1.0.0
// 2. 1.1.0

// 另一个范围查询示例：获取所有大于等于2.0.0的版本
laterVersionQuery := sortedGroups.QueryRange(
    tuple.New2[*versions.Version, versions.ContainsPolicy](
        versions.NewVersion("2.0.0"), 
        versions.ContainsPolicyYes),
    tuple.New2[*versions.Version, versions.ContainsPolicy](
        nil, versions.ContainsPolicyNone), // nil表示不设上限
)
fmt.Printf("2.0.0及以后的版本共有 %d 个\n", len(laterVersionQuery))
// 输出: 2.0.0及以后的版本共有 3 个 (2.0.0, 2.1.0, 3.0.0)
```
</details>

### 文件操作函数

<details open>
<summary><b>从文件读取版本号</b></summary>

```go
// 读取文件中的版本号字符串
func ReadVersionsStringFromFile(filepath string) ([]string, error)

// 读取并解析文件中的版本号
func ReadVersionsFromFile(filepath string) ([]*Version, error)
```

**详细函数说明:**

**`ReadVersionsStringFromFile(filepath string)`**: 读取文件中的版本号字符串
- **参数:** `filepath string` - 文件路径
- **返回值:** 
  - `[]string` - 版本号字符串切片
  - `error` - 读取过程中可能发生的错误
- **行为:**
  1. 打开并读取指定文件
  2. 按行分割文件内容
  3. 去除每行首尾空白字符
  4. 过滤空行和注释行（以#开头）
  5. 返回有效的版本号字符串行
- **错误处理:**
  - 文件不存在：返回 os.ErrNotExist
  - 权限问题：返回相应的文件系统错误
  - 读取失败：返回 io.EOF 或其他I/O错误
- **用途:** 读取版本列表文件，获取原始版本号字符串

**`ReadVersionsFromFile(filepath string)`**: 读取并解析文件中的版本号
- **参数:** `filepath string` - 文件路径
- **返回值:** 
  - `[]*Version` - 版本对象切片
  - `error` - 读取或解析过程中可能发生的错误
- **行为:**
  1. 调用 ReadVersionsStringFromFile 获取版本号字符串
  2. 对每个字符串调用 NewVersion 创建版本对象
  3. 返回创建的版本对象切片
- **错误处理:**
  - 继承 ReadVersionsStringFromFile 的所有错误
  - 不会因版本解析失败而返回错误（使用 NewVersion 而非 NewVersionE）
- **用途:** 一次性读取并解析版本列表文件

**文件格式要求:**
- 每行一个版本号
- 支持空行（会被忽略）
- 支持注释行（以#开头，会被忽略）
- 行首尾的空白字符会被自动移除

**性能考虑:**
- 对于大文件，考虑使用缓冲读取
- 如需处理无效版本，应在读取后手动过滤
- 读取后可考虑缓存结果，避免频繁IO操作

**示例:**
```go
// 从文件读取版本号字符串
versionStrings, err := versions.ReadVersionsStringFromFile("versions.txt")
if err != nil {
    log.Fatalf("读取版本号失败: %v", err)
}
fmt.Printf("共读取 %d 个版本号字符串\n", len(versionStrings))

// 从文件读取并解析版本号
versionObjects, err := versions.ReadVersionsFromFile("versions.txt")
if err != nil {
    log.Fatalf("解析版本号失败: %v", err)
}
fmt.Printf("共解析 %d 个版本号对象\n", len(versionObjects))

// 版本号文件示例 (versions.txt):
// # 稳定版本
// 1.0.0
// 1.0.1
// 1.0.2
// 
// # 预发布版本
// 1.1.0-alpha
// 1.1.0-beta
// 1.1.0-rc1

// 高级应用：读取并分组版本
versionObjects, err := versions.ReadVersionsFromFile("versions.txt")
if err != nil {
    log.Fatalf("读取版本号失败: %v", err)
}
sortedGroups := versions.NewSortedVersionGroups(versionObjects)
fmt.Printf("共读取 %d 个版本，分为 %d 个版本组\n", 
    len(versionObjects), len(sortedGroups.GroupIDs()))

// 过滤无效版本
validVersions := make([]*versions.Version, 0)
for _, v := range versionObjects {
    if v.IsValid() {
        validVersions = append(validVersions, v)
    } else {
        fmt.Printf("忽略无效版本: %s\n", v.Raw)
    }
}
fmt.Printf("有效版本数: %d\n", len(validVersions))
```
</details>

### 排序函数

<details open>
<summary><b>版本排序函数</b></summary>

```go
// 对版本字符串切片进行排序
func SortVersionStringSlice(versionStringSlice []string) []string

// 对版本对象切片进行排序
func SortVersionSlice(versions []*Version) []*Version
```

**详细函数说明:**

**`SortVersionStringSlice(versionStringSlice []string)`**: 对版本字符串切片进行排序
- **参数:** `versionStringSlice []string` - 要排序的版本号字符串切片
- **返回值:** `[]string` - 排序后的版本号字符串切片
- **行为:**
  1. 将字符串转换为版本对象
  2. 对版本对象进行排序
  3. 将排序后的版本对象转回字符串
- **排序规则:** 基于 Version.CompareTo 方法比较
- **处理无效版本:** 
  - 无效或无法解析的版本会出现在结果中
  - 无效版本之间的顺序取决于其原始字符串
- **用途:** 直接对版本号字符串进行排序，无需手动创建版本对象

**`SortVersionSlice(versions []*Version)`**: 对版本对象切片进行排序
- **参数:** `versions []*Version` - 要排序的版本对象切片
- **返回值:** `[]*Version` - 排序后的版本对象切片
- **行为:** 对输入的版本对象切片进行原地排序
- **排序规则:** 基于 Version.CompareTo 方法比较
- **处理无效版本:**
  - 无效版本会被保留在结果中
  - 排序时将无效版本视为最小值
- **用途:** 对已创建的版本对象集合进行排序

**排序算法特点:**
- 稳定排序：相等的版本保持原始顺序
- 时间复杂度：O(n log n)，其中n为版本数量
- 空间复杂度：SortVersionStringSlice需要O(n)额外空间，SortVersionSlice为O(1)

**注意事项:**
- 两个函数都不会修改原始输入切片，而是返回新的排序结果
- 排序时会考虑版本的所有组成部分（前缀、数字、后缀）
- 大多数情况下，排序结果符合语义化版本规范的预期顺序

**示例:**
```go
// 排序版本号字符串
unsortedStrings := []string{
    "2.0.0", 
    "1.0.0", 
    "1.10.0", 
    "1.2.0",
    "v1.5.0",
    "1.5.0-beta",
}
sortedStrings := versions.SortVersionStringSlice(unsortedStrings)
fmt.Println("排序后的版本号字符串:")
for i, v := range sortedStrings {
    fmt.Printf("%d. %s\n", i+1, v)
}
// 输出:
// 排序后的版本号字符串:
// 1. 1.0.0
// 2. 1.2.0
// 3. 1.5.0-beta
// 4. v1.5.0
// 5. 1.10.0
// 6. 2.0.0

// 排序版本对象
unsortedVersions := []*versions.Version{
    versions.NewVersion("2.0.0"),
    versions.NewVersion("1.0.0"),
    versions.NewVersion("1.10.0"),
    versions.NewVersion("1.2.0-alpha"),
}
sortedVersions := versions.SortVersionSlice(unsortedVersions)
fmt.Println("排序后的版本对象:")
for i, v := range sortedVersions {
    fmt.Printf("%d. %s\n", i+1, v.Raw)
}
// 输出:
// 排序后的版本对象:
// 1. 1.0.0
// 2. 1.2.0-alpha
// 3. 1.10.0
// 4. 2.0.0

// 自定义排序：按发布时间倒序
// 注：需配合自定义排序函数使用
type ByReleaseTimeDesc []*versions.Version
func (v ByReleaseTimeDesc) Len() int           { return len(v) }
func (v ByReleaseTimeDesc) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v ByReleaseTimeDesc) Less(i, j int) bool { return v[j].PublicTime.Before(v[i].PublicTime) }

// 设置发布时间并排序
for i, v := range unsortedVersions {
    // 模拟不同的发布时间，每个版本间隔1个月
    v.PublicTime = time.Now().AddDate(0, -i, 0)
}
sort.Sort(ByReleaseTimeDesc(unsortedVersions))
fmt.Println("按发布时间倒序排序:")
for i, v := range unsortedVersions {
    fmt.Printf("%d. %s (发布于: %s)\n", i+1, v.Raw, v.PublicTime.Format("2006-01-02"))
}
```
</details>

### 分组函数

<details open>
<summary><b>版本分组函数</b></summary>

```go
// 将版本列表按主版本号分组
func Group(versions []*Version) map[string]*VersionGroup
```

**详细函数说明:**

**`Group(versions []*Version)`**: 将版本列表按主版本号分组
- **参数:** `versions []*Version` - 要分组的版本对象列表
- **返回值:** `map[string]*VersionGroup` - 以组ID为键的版本组映射
- **行为:**
  1. 遍历版本对象列表
  2. 对每个版本提取其主版本号（VersionNumbers[0]）作为组ID
  3. 按组ID创建或更新 VersionGroup
  4. 将版本添加到对应的组中
- **分组标准:**
  - 默认按主版本号（第一个数字）分组
  - 无效版本或无主版本号的版本使用"0"作为组ID
- **返回格式:**
  - 键：字符串形式的组ID，通常是主版本号（如"1", "2"）
  - 值：包含对应主版本号所有版本的 VersionGroup 对象
- **用途:**
  - 按主版本号组织版本
  - 创建层次化版本选择界面
  - 辅助版本号可视化

**常见分组用例:**
- 按语义化版本的主版本号分组，快速区分不兼容版本
- UI中显示分级版本选择器，方便用户选择合适版本
- 对大量版本进行结构化管理，提高版本浏览和选择效率

**性能特点:**
- 时间复杂度：O(n)，其中n为版本数量
- 空间复杂度：O(n)，需要存储所有版本的组织结构
- 适合处理任意规模的版本集合，性能随版本数量线性增长

**示例:**
```go
// 创建版本列表
versionList := []*versions.Version{
    versions.NewVersion("1.0.0"),
    versions.NewVersion("1.1.0"),
    versions.NewVersion("1.2.0"),
    versions.NewVersion("2.0.0"),
    versions.NewVersion("2.1.0"),
    versions.NewVersion("3.0.0-beta"),
}

// 按主版本号分组
groupMap := versions.Group(versionList)

// 打印分组结果
fmt.Printf("共有 %d 个版本组\n", len(groupMap))
// 输出: 共有 3 个版本组

// 遍历分组结果
for groupID, group := range groupMap {
    fmt.Printf("\n版本组 %s 包含 %d 个版本:\n", groupID, group.Count())
    // 对组内版本排序
    sortedVersions := group.SortVersions()
    for i, v := range sortedVersions {
        fmt.Printf("  %d. %s\n", i+1, v.Raw)
    }
}
// 输出:
// 版本组 1 包含 3 个版本:
//   1. 1.0.0
//   2. 1.1.0
//   3. 1.2.0
//
// 版本组 2 包含 2 个版本:
//   1. 2.0.0
//   2. 2.1.0
//
// 版本组 3 包含 1 个版本:
//   1. 3.0.0-beta

// 获取特定组内的最新版本
group1 := groupMap["1"]
if group1 != nil && group1.Count() > 0 {
    sortedGroup1 := group1.SortVersions()
    latestVersion := sortedGroup1[len(sortedGroup1)-1]
    fmt.Printf("\n版本组 1 中的最新版本: %s\n", latestVersion.Raw)
    // 输出: 版本组 1 中的最新版本: 1.2.0
}

// 高级应用：创建版本选择器数据
type VersionOption struct {
    Label    string
    Value    string
    Children []VersionOption
}

versionSelector := make([]VersionOption, 0, len(groupMap))
for groupID, group := range groupMap {
    children := make([]VersionOption, 0, group.Count())
    for _, v := range group.SortVersions() {
        children = append(children, VersionOption{
            Label: v.Raw,
            Value: v.Raw,
        })
    }
    versionSelector = append(versionSelector, VersionOption{
        Label:    fmt.Sprintf("版本 %s.x", groupID),
        Value:    groupID,
        Children: children,
    })
}
```
</details>

### 可视化函数

<details open>
<summary><b>版本可视化函数</b></summary>

```go
// 以文本树形式可视化版本结构
func VisualizeVersions(versions []*Version, w io.Writer, maxItemsPerGroup int)

// 以文本树形式可视化版本组层次结构
func VisualizeVersionGroups(versions []*Version, w io.Writer)
```

**详细函数说明:**

**`VisualizeVersions(versions []*Version, w io.Writer, maxItemsPerGroup int)`**: 以文本树形式可视化版本结构
- **参数:** 
  - `versions []*Version` - 要可视化的版本对象列表
  - `w io.Writer` - 输出写入目标
  - `maxItemsPerGroup int` - 每组最多显示的版本数量，0表示不限制
- **行为:**
  1. 按主版本号对版本进行分组
  2. 为每个版本组创建树形结构
  3. 对组内版本排序并输出有限数量
  4. 如超过限制，显示剩余版本数
- **输出格式:**
  - 显示总版本数和版本组数
  - 每个版本组以树形结构展示
  - 可选显示版本的发布时间
- **用途:**
  - 在终端或日志中展示版本组织结构
  - 提供版本分布的直观视图
  - 方便调试和检查版本管理逻辑

**`VisualizeVersionGroups(versions []*Version, w io.Writer)`**: 以文本树形式可视化版本组层次结构
- **参数:** 
  - `versions []*Version` - 要可视化的版本对象列表
  - `w io.Writer` - 输出写入目标
- **行为:**
  1. 创建版本组层次结构
  2. 以树形结构输出版本组信息
  3. 对每个组显示组ID和版本数量
- **输出格式:**
  - 不显示具体版本，只展示组级别信息
  - 每组显示ID和包含的版本数量
- **用途:**
  - 获取版本分组的概览
  - 查看不同主版本的分布情况
  - 为大型版本集合提供简化视图

**视觉效果增强:**
- 使用Unicode字符创建树形结构（如 "├──", "└──"）
- 缩进和连接线提高可读性
- 可选颜色高亮（输出到终端时）

**应用场景:**
- 版本库管理和维护
- 发布日志和版本跟踪
- 在CLI工具中展示版本信息
- 调试版本比较和排序逻辑

**示例:**
```go
// 准备版本列表
versionList := []*versions.Version{
    versions.NewVersion("1.0.0"),
    versions.NewVersion("1.0.1"),
    versions.NewVersion("1.1.0"),
    versions.NewVersion("2.0.0"),
    versions.NewVersion("2.0.1"),
    versions.NewVersion("2.1.0"),
    versions.NewVersion("3.0.0-alpha"),
    versions.NewVersion("3.0.0-beta"),
}

// 设置发布时间（示例）
currentTime := time.Now()
for i, v := range versionList {
    // 模拟不同发布时间，每个版本间隔1个月
    v.PublicTime = currentTime.AddDate(0, -i, 0)
}

// 可视化完整版本结构，每组显示最多2个版本
fmt.Println("完整版本可视化 (每组最多2个版本):")
versions.VisualizeVersions(versionList, os.Stdout, 2)
// 输出示例:
// 版本总数: 8
// 版本组数: 3
//
// ┌─ 版本组: 1 (3个版本)
// ├── 1.0.0 (发布时间: 2023-05-15)
// ├── 1.0.1 (发布时间: 2023-04-15)
// └── ...还有1个版本未显示
//
// ┌─ 版本组: 2 (3个版本)
// ├── 2.0.0 (发布时间: 2023-02-15)
// ├── 2.0.1 (发布时间: 2023-01-15)
// └── ...还有1个版本未显示
//
// ┌─ 版本组: 3 (2个版本)
// ├── 3.0.0-alpha (发布时间: 2022-12-15)
// └── 3.0.0-beta (发布时间: 2022-11-15)

// 仅可视化版本组结构
fmt.Println("\n版本组结构可视化:")
versions.VisualizeVersionGroups(versionList, os.Stdout)
// 输出示例:
// 版本总数: 8
// 版本组数: 3
//
// ┌─ 版本组: 1 (3个版本)
// ├─ 版本组: 2 (3个版本)
// └─ 版本组: 3 (2个版本)

// 可视化特定版本范围
fmt.Println("\n筛选后的版本可视化:")
// 获取1.x和2.x系列的所有版本
filteredVersions := make([]*versions.Version, 0)
for _, v := range versionList {
    major := v.VersionNumbers.MajorVersion()
    if major == 1 || major == 2 {
        filteredVersions = append(filteredVersions, v)
    }
}
versions.VisualizeVersions(filteredVersions, os.Stdout, 0) // 0表示显示所有版本
```
</details>

---

## 🔍 使用示例

<div align="center">

| 示例 | 描述 |
|:------:|:-----|
| [📚 基本版本解析](./examples/01_basic_version_parsing/main.go) | 如何解析和比较不同格式的版本号 |
| [📂 从文件读取版本](./examples/02_reading_versions_from_file/main.go) | 如何从文件中读取版本信息 |
| [🔢 版本排序](./examples/03_version_sorting/main.go) | 如何对版本号进行排序 |
| [📦 版本分组](./examples/04_version_grouping/main.go) | 如何对版本进行分组管理 |
| [🔍 版本范围查询](./examples/05_version_range_queries/main.go) | 如何查询特定版本范围 |
| [📊 版本可视化](./examples/06_version_visualization/main.go) | 如何可视化版本结构 |

</div>

<div align="center">
<img src="https://user-images.githubusercontent.com/5877/236610730-b2f3c58f-b70b-4621-9f1a-ae99928dae99.png" alt="版本树示例" width="600"/>
<br>
<i>版本树可视化示例</i>
</div>

---

## ⚠️ 注意事项

- ⚠️ 确保传入的文件路径正确，且文件格式符合要求
- ⚠️ 版本号的解析和比较依赖于正确的格式，非标准格式可能会导致解析错误
- ⚠️ 时间比较依赖于 `PublicTime` 被正确设置
- ⚠️ 对于非标准格式的版本号可能需要额外的处理

---

## 📈 性能

- 版本解析: `O(n)`，其中 n 是版本号字符串的长度
- 版本比较: `O(m)`，其中 m 是版本号中数字部分的最大长度
- 版本排序: `O(n log n)`，其中 n 是版本列表的长度
- 范围查询: `O(log n)`，基于有序版本组的二分查找

---

## 🤖 Claude Code Skills

本项目同时是一个 **Claude Code Skills 仓库**，提供版本号处理的专业技能。安装后，您可以在 Claude Code 中通过斜杠命令调用这些技能。

### 安装方式

```bash
# 添加 marketplace
claude marketplace add versions https://github.com/scagogogo/versions

# 安装插件
claude plugin install versions@versions
```

### 可用 Skills

| Skill | 命令 | 用途 |
|:------|:-----|:-----|
| 版本号解析 | `/version-parsing` | 解析、验证版本字符串 |
| 版本号比较 | `/version-comparison` | 比较两个版本号大小 |
| 版本号排序 | `/version-sorting` | 对版本号列表排序 |
| 版本号分组 | `/version-grouping` | 按主版本号分组管理 |
| 版本范围查询 | `/version-range-query` | 查询指定范围内的版本 |
| 版本号可视化 | `/version-visualization` | 以树形结构展示版本层次 |
| 版本号文件操作 | `/version-file-operations` | 从文件读取版本号列表 |

### 使用示例

在 Claude Code 中输入：

```
/version-parsing v1.2.3-beta1
```

Claude 将基于此 Skill 提供专业的版本号解析指导，包括正确的 Go 代码示例和 API 用法。

---

## 📄 许可证

<div align="center">
  
本项目采用 [MIT 许可证](./LICENSE) - 详见 LICENSE 文件

<b>Copyright © 2023-2025 scagogogo</b>

</div>