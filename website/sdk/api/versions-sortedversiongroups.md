# Versions

::: info 方法 · `SortedVersionGroups`
```go
func (x *SortedVersionGroups) Versions() []*Version
```
:::

## 📖 说明

Versions 返回所有版本组中的所有版本

版本按组排序，组内按版本排序。


#### 返回

- `[]*Version`：所有版本的有序列表


---

::: details 源码位置
定义于 [`sorted_version_groups.go`](https://github.com/scagogogo/versions-skills/blob/main/sorted_version_groups.go)
:::
