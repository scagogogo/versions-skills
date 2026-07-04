# IndexOf

::: info 方法 · `VersionSlice`
```go
func (s VersionSlice) IndexOf(target *Version) int
```
:::

## 📖 说明

IndexOf 查找版本在切片中的位置

如果未找到则返回 -1。


::: details 同类方法（点击展开）

**第 1 组**

- [`VersionSlice.Len`](/sdk/api/len-versionslice)
- [`VersionSlice.Less`](/sdk/api/less-versionslice)
- [`VersionSlice.Swap`](/sdk/api/swap-versionslice)
- [`VersionSlice.Min`](/sdk/api/min-versionslice)
- [`VersionSlice.Max`](/sdk/api/max-versionslice)
- [`VersionSlice.Filter`](/sdk/api/filter-versionslice)
- [`VersionSlice.Contains`](/sdk/api/contains-versionslice)
- [`VersionSlice.Unique`](/sdk/api/unique-versionslice)

**第 2 组**

- [`VersionSlice.Sort`](/sdk/api/sort-versionslice)
- [`VersionSlice.Sorted`](/sdk/api/sorted-versionslice)

:::


::: details 源码位置
定义于 [`version_slice.go`](https://github.com/scagogogo/versions-skills/blob/main/version_slice.go)
:::
