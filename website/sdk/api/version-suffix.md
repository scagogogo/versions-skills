# VersionSuffix

::: info 类型 · 根包
```go
type VersionSuffix string
```
:::

## 📖 说明

VersionSuffix 表示版本号的后缀，版本号中数字后面的部分

VersionSuffix 是一个字符串类型，用于表示和操作版本号的后缀部分。
在版本号格式中，后缀是位于数字部分之后的字符串，如 "1.2.3-beta1" 中的 "-beta1"。
后缀通常用于表示预发布版本、构建元数据或其他版本标识信息。

后缀可以为空，表示版本号仅包含数字部分，如 "1.2.3"。


```go
// 检查后缀是否为空
suffix := versions.VersionSuffix("-beta1")
if !suffix.IsEmpty() {
    fmt.Printf("版本后缀: %s\n", suffix)
}

// 比较后缀的优先级
suffix1 := versions.VersionSuffix("-alpha")
suffix2 := versions.VersionSuffix("-beta")
if suffix1.CompareTo(suffix2) < 0 {
    fmt.Println("alpha 后缀的优先级低于 beta 后缀")
}
```


---

::: details 源码位置
定义于 [`version_suffix.go`](https://github.com/scagogogo/versions-skills/blob/main/version_suffix.go)
:::
