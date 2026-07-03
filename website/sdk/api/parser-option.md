# ParserOption

::: info 类型 · 根包
```go
type ParserOption struct {
	// Delimiters 版本号数字部分的分隔符集合
	// 默认值: "." (仅点号)
	// 常见扩展: ".-_" (支持 RPM/Debian 的连字符和 Python 的下划线)
	Delimiters string
}
```
:::

## 📖 说明

ParserOption 配置版本号解析器的行为

ParserOption 允许调用者自定义解析器支持的数字分隔符，
以适应不同语言生态系统的版本号格式差异。


```go
// 支持 underscore 分隔（Python/RPM 生态）
v := versions.NewVersionWithOption("1_2_3",
    versions.ParserOption{Delimiters: ".-_"})
```


---

::: details 源码位置
定义于 [`parser_option.go`](https://github.com/scagogogo/versions-skills/blob/main/parser_option.go)
:::
