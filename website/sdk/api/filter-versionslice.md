# Filter

::: info 方法 · `VersionSlice`
```go
func (s VersionSlice) Filter(predicate func(*Version) bool) VersionSlice
```
:::

## 📖 说明

Filter 根据谓词函数过滤版本切片

返回所有满足谓词条件的版本。


---

::: details 源码位置
定义于 [`version_slice.go`](https://github.com/scagogogo/versions-skills/blob/main/version_slice.go)
:::
