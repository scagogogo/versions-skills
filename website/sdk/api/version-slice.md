# VersionSlice

::: info 类型 · 根包
```go
type VersionSlice []*Version
```
:::

## 📖 说明

VersionSlice 是 []*Version 的有序集合类型，实现了 sort.Interface

VersionSlice 提供了 Go 标准库 sort.Sort() 的直接支持，
允许使用 sort.Sort(slice) 而不是 sort.Slice() 配合闭包。


```go
slice := versions.VersionSlice(versions.NewVersions("2.0.0", "1.0.0", "1.5.0"))
sort.Sort(slice)
// slice 现在按版本号排序
```


---

::: details 源码位置
定义于 [`version_slice.go`](https://github.com/scagogogo/versions-skills/blob/main/version_slice.go)
:::
