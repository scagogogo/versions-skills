# Filter

::: info 方法 · `VersionSlice`
```go
func (s VersionSlice) Filter(predicate func(*Version) bool) VersionSlice
```
:::

## 📖 说明

Filter 根据谓词函数过滤版本切片

返回所有满足谓词条件的版本。


## 🔗 同类方法

- [`VersionSlice.Len`](/sdk/api/len-versionslice)
- [`VersionSlice.Less`](/sdk/api/less-versionslice)
- [`VersionSlice.Swap`](/sdk/api/swap-versionslice)
- [`VersionSlice.Min`](/sdk/api/min-versionslice)
- [`VersionSlice.Max`](/sdk/api/max-versionslice)
- [`VersionSlice.Contains`](/sdk/api/contains-versionslice)
- [`VersionSlice.IndexOf`](/sdk/api/index-of-versionslice)
- [`VersionSlice.Unique`](/sdk/api/unique-versionslice)
- [`VersionSlice.Sort`](/sdk/api/sort-versionslice)
- [`VersionSlice.Sorted`](/sdk/api/sorted-versionslice)


---

::: details 源码位置
定义于 [`version_slice.go`](https://github.com/scagogogo/versions-skills/blob/main/version_slice.go)
:::
