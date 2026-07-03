# IsPatchChange

::: info 方法 · `VersionDiff`
```go
func (d *VersionDiff) IsPatchChange() bool
```
:::

## 📖 说明

IsPatchChange 判断差异是否仅涉及修订版本号变化（主、次版本号不变）


## 🔗 同类方法

- [`VersionDiff.String`](/sdk/api/string-versiondiff)
- [`VersionDiff.IsUpgrade`](/sdk/api/is-upgrade-versiondiff)
- [`VersionDiff.IsDowngrade`](/sdk/api/is-downgrade-versiondiff)
- [`VersionDiff.IsMajorChange`](/sdk/api/is-major-change-versiondiff)
- [`VersionDiff.IsMinorChange`](/sdk/api/is-minor-change-versiondiff)


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
