# SortedVersionGroups

::: info 类型 · 根包
```go
type SortedVersionGroups struct {

	// groupIDToIndexMap 用于根据组的ID快速定位到这个组在排好序的切片中的位置
	// 键为版本组ID，值为该组在groupSlice中的索引位置
	groupIDToIndexMap map[string]int

	// groupSlice 排好序的版本组切片
	// 按照版本组的大小顺序排列
	groupSlice []*VersionGroup
}
```
:::

## 📖 说明

SortedVersionGroups 表示已排序的版本组集合

SortedVersionGroups 封装了已排序的版本组切片和索引映射，
用于高效地进行版本组查询和范围检索。它通过预先排序和建立索引，
优化了版本组的查找性能。

结构特点:
1. 保持版本组的有序性，便于范围查询
2. 维护组ID到数组索引的映射，支持快速定位
3. 支持基于版本范围的高效查询


```go
// 创建已排序的版本组
allVersions := versions.NewVersions("1.0.0", "1.1.0", "2.0.0", "2.1.0")
sortedGroups := versions.NewSortedVersionGroups(allVersions)

// 执行范围查询
startVer := versions.NewVersion("1.0.0")
endVer := versions.NewVersion("2.0.0")
startTuple := tuple.NewTuple2(startVer, versions.ContainsPolicyYes)
endTuple := tuple.NewTuple2(endVer, versions.ContainsPolicyYes)

rangeResult := sortedGroups.QueryRange(startTuple, endTuple)
```


---

::: details 源码位置
定义于 [`sorted_version_groups.go`](https://github.com/scagogogo/versions-skills/blob/main/sorted_version_groups.go)
:::
