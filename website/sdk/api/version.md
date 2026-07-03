# Version

::: info 类型 · 根包
```go
type Version struct {

	// Raw 原始的版本号字符串
	Raw string `json:"raw"`

	// PublicTime 此版本的发布时间
	PublicTime time.Time `json:"public_time"`

	// VersionNumbers 版本号中的数字部分
	// 例如对于版本号 "v1.2.3-beta1"，VersionNumbers 为 [1,2,3]
	VersionNumbers VersionNumbers `json:"version_numbers"`

	// Prefix 版本号数字部分之前的前缀
	// 例如对于版本号 "v1.2.3"，Prefix 为 "v"
	Prefix VersionPrefix `json:"prefix"`

	// Suffix 版本号数字部分之后的后缀
	// 例如对于版本号 "1.2.3-beta1"，Suffix 为 "-beta1"
	Suffix VersionSuffix `json:"suffix"`

	// Metadata semver 构建元数据
	//
	// 在 semver 规范中，构建元数据是版本号中 + 号后面的部分，如 "1.0.0+build123" 中的 "build123"。
	// 根据 semver 规范，构建元数据不参与版本比较。
	Metadata string `json:"metadata,omitempty"`
}
```
:::

## 📖 说明

Version 用于表示一个版本号

Version 结构体封装了版本号的各个组成部分，包括原始字符串、发布时间、数字部分、
前缀和后缀。它支持版本号的解析、比较和分组等操作，实现了 Comparable 接口以便
进行版本排序。

一个典型的版本号格式可能为：v1.2.3-beta1，其中：
- "v" 是前缀
- "1.2.3" 是数字部分
- "-beta1" 是后缀


```go
// 创建一个版本对象
version := versions.NewVersion("v1.2.3-rc1")

// 检查版本是否有效
if version.IsValid() {
fmt.Printf("版本号有效: %s\n", version.Raw)
fmt.Printf("版本号数字部分: %v\n", version.VersionNumbers)
}

// 比较两个版本
v1 := versions.NewVersion("1.2.3")
v2 := versions.NewVersion("1.3.0")
if v1.CompareTo(v2) < 0 {
fmt.Println("v1 比 v2 旧")
}
```


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
