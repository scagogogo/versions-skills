# IsUpgrade

::: info 方法 · `VersionDiff`
```go
func (d *VersionDiff) IsUpgrade() bool
```
:::

## 📖 说明

IsUpgrade 判断差异是否为升级（主、次或修订版本号至少有一项增加）


## 🔗 同类方法

- [`VersionDiff.String`](/sdk/api/string-versiondiff)
- [`VersionDiff.IsDowngrade`](/sdk/api/is-downgrade-versiondiff)
- [`VersionDiff.IsMajorChange`](/sdk/api/is-major-change-versiondiff)
- [`VersionDiff.IsMinorChange`](/sdk/api/is-minor-change-versiondiff)
- [`VersionDiff.IsPatchChange`](/sdk/api/is-patch-change-versiondiff)


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
