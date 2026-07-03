# NewVersionStringParser

::: info 函数 · 根包
```go
func NewVersionStringParser(versionStr string) *VersionStringParser
```
:::

## 📖 说明

NewVersionStringParser 创建一个版本号Parser

该方法创建一个新的版本号解析器实例。每次解析都需要重新创建新的Parser，
因为解析状态保存在Parser对象中。


#### 参数

- `versionStr`：要解析的版本号字符串


#### 返回

- `*VersionStringParser`：新创建的版本号解析器


```go
parser := versions.NewVersionStringParser("v1.2.3-rc1")
version := parser.Parse()
```


---

::: details 源码位置
定义于 [`parser.go`](https://github.com/scagogogo/versions-skills/blob/main/parser.go)
:::
