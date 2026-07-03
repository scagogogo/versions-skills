# LatestStable

::: info 函数 · 根包
```go
func LatestStable(versions []*Version) *Version
```
:::

## 📖 说明

LatestStable 从版本列表中找到最新的稳定版本

稳定版本是指不带后缀的版本。如果不存在稳定版本则返回 nil。


---

::: details 源码位置
定义于 [`version_utils.go`](https://github.com/scagogogo/versions-skills/blob/main/version_utils.go)
:::
