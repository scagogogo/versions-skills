# VersionStringParser

::: info 类型 · 根包
```go
type VersionStringParser struct {

	// versionStr 被解析的字符串原始样子
	versionStr string
	// versionRunes 转为字符序列方便处理
	versionRunes []rune
	// i 上面的字符序列，当前解析到哪个下标了
	i int

	// v 解析结果
	v *Version

	// option 解析器选项
	option ParserOption
}
```
:::

## 📖 说明

VersionStringParser 把版本从字符串形式解析为struct

VersionStringParser 负责将版本号字符串解析为结构化的 Version 对象。
它实现了版本号字符串的词法分析，将字符串划分为前缀、数字部分和后缀三个组成部分。

解析过程：
1. 首先识别和提取版本号的前缀部分（如 "v"）
2. 然后识别和提取版本号的数字部分（如 "1.2.3"）
3. 最后识别和提取版本号的后缀部分（如 "-beta1"）


```go
// 创建一个版本解析器
parser := versions.NewVersionStringParser("v1.2.3-beta1")

// 解析版本字符串
version := parser.Parse()

// 使用解析结果
fmt.Printf("前缀: %s\n", version.Prefix)
fmt.Printf("数字部分: %v\n", version.VersionNumbers)
fmt.Printf("后缀: %s\n", version.Suffix)
```


---

::: details 源码位置
定义于 [`parser.go`](https://github.com/scagogogo/versions-skills/blob/main/parser.go)
:::
