# IsDowngrade

::: info 方法 · `VersionDiff`
```go
func (d *VersionDiff) IsDowngrade() bool
```
:::

## 📖 说明

IsDowngrade 判断差异是否为降级（主、次或修订版本号至少有一项减少）


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
