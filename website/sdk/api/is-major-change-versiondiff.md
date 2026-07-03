# IsMajorChange

::: info 方法 · `VersionDiff`
```go
func (d *VersionDiff) IsMajorChange() bool
```
:::

## 📖 说明

IsMajorChange 判断差异是否涉及主版本号变化


## 🔗 同类方法

- [`VersionDiff.String`](/sdk/api/string-versiondiff)
- [`VersionDiff.IsUpgrade`](/sdk/api/is-upgrade-versiondiff)
- [`VersionDiff.IsDowngrade`](/sdk/api/is-downgrade-versiondiff)
- [`VersionDiff.IsMinorChange`](/sdk/api/is-minor-change-versiondiff)
- [`VersionDiff.IsPatchChange`](/sdk/api/is-patch-change-versiondiff)


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
