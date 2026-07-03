# Get

::: info 方法 · `SortedVersionGroups`
```go
func (x *SortedVersionGroups) Get(groupID string) *VersionGroup
```
:::

## 📖 说明

Get 根据组 ID 获取版本组

如果组 ID 不存在则返回 nil。


#### 参数

- `groupID`：版本组 ID，如 "1.2"


#### 返回

- `*VersionGroup`：对应的版本组，不存在则返回 nil


---

::: details 源码位置
定义于 [`sorted_version_groups.go`](https://github.com/scagogogo/versions-skills/blob/main/sorted_version_groups.go)
:::
