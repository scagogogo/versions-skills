# Parse

::: info 方法 · `VersionStringParser`
```go
func (x *VersionStringParser) Parse() *Version
```
:::

## 📖 说明

Parse 解析版本号字符串

该方法按照预定义的规则解析版本号字符串，提取前缀、数字部分和后缀，
并构造一个完整的 Version 对象返回。


#### 返回

- `*Version`：解析后的版本对象


```go
parser := versions.NewVersionStringParser("v1.2.3-beta1")
version := parser.Parse()
fmt.Printf("版本号: %s\n", version.Raw)
```


---

::: details 源码位置
定义于 [`parser.go`](https://github.com/scagogogo/versions-skills/blob/main/parser.go)
:::
