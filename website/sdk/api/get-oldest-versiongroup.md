# GetOldest

::: info 方法 · `VersionGroup`
```go
func (x *VersionGroup) GetOldest() *Version
```
:::

## 📖 说明

GetOldest 获取版本组中最旧的版本

返回按排序后最旧的版本，如果组为空则返回 nil。


---

::: details 源码位置
定义于 [`version_group.go`](https://github.com/scagogogo/versions-skills/blob/main/version_group.go)
:::
