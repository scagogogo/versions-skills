# VersionPrefix

::: info 类型 · 根包
```go
type VersionPrefix string
```
:::

## 📖 说明

VersionPrefix 表示版本中数字部分之前的部分

VersionPrefix 是一个字符串类型，用于表示和操作版本号的前缀部分。
在版本号格式中，前缀是位于数字部分之前的字符串，如 "v1.2.3" 中的 "v"。
前缀可以为空，表示版本号直接以数字部分开始，如 "1.2.3"。


```go
// 检查前缀是否为空
prefix := versions.VersionPrefix("v")
if !prefix.IsEmpty() {
fmt.Printf("版本前缀: %s\n", prefix)
}

// 在解析版本号时处理前缀
version := versions.NewVersion("v1.2.3")
fmt.Printf("版本 %s 的前缀是: %s\n", version.Raw, version.Prefix)
```


---

::: details 源码位置
定义于 [`version_prefix.go`](https://github.com/scagogogo/versions-skills/blob/main/version_prefix.go)
:::
