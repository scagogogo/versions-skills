# IsMinorChange

::: info 方法 · `VersionDiff`
```go
func (d *VersionDiff) IsMinorChange() bool
```
:::

## 📖 说明

IsMinorChange 判断差异是否仅涉及次版本号变化（主版本号不变）


## 🔗 同类方法

- [`VersionDiff.String`](/sdk/api/string-versiondiff)
- [`VersionDiff.IsUpgrade`](/sdk/api/is-upgrade-versiondiff)
- [`VersionDiff.IsDowngrade`](/sdk/api/is-downgrade-versiondiff)
- [`VersionDiff.IsMajorChange`](/sdk/api/is-major-change-versiondiff)
- [`VersionDiff.IsPatchChange`](/sdk/api/is-patch-change-versiondiff)


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
