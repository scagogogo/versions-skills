# SortVersionGroupSlice

::: info 函数 · 根包
```go
func SortVersionGroupSlice(groupSlice []*VersionGroup)
```
:::

## 📖 说明

SortVersionGroupSlice 对版本组切片进行排序

该函数直接对传入的版本组切片进行原地排序，排序依据是版本组之间的比较规则。


#### 参数

- `groupSlice`：待排序的版本组切片，排序操作直接修改该切片


```go
groups := []*VersionGroup{group1, group2, group3}
SortVersionGroupSlice(groups)
// 现在 groups 已按版本组规则排序
```


---

::: details 源码位置
定义于 [`sort.go`](https://github.com/scagogogo/versions-skills/blob/main/sort.go)
:::
