# LatestPrerelease

::: info 方法 · `VersionGroup`
```go
func (x *VersionGroup) LatestPrerelease() *Version
```
:::

## 📖 说明

LatestPrerelease 获取版本组中最新预发布版本

返回按排序后最新的预发布版本，如果组中无预发布版本则返回 nil。


---

::: details 源码位置
定义于 [`version_group.go`](https://github.com/scagogogo/versions-skills/blob/main/version_group.go)
:::
