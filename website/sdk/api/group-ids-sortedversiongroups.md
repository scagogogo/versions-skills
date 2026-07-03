# GroupIDs

::: info 方法 · `SortedVersionGroups`
```go
func (x *SortedVersionGroups) GroupIDs() []string
```
:::

## 📖 说明

GroupIDs 返回所有版本组的ID列表

该方法返回按顺序排列的版本组ID列表。
返回的ID列表与内部的版本组切片顺序一致，确保保持排序状态。


#### 返回

- `[]string`：含有所有版本组ID的字符串切片


```go
sortedGroups := versions.NewSortedVersionGroups(allVersions)
groupIDs := sortedGroups.GroupIDs()
for _, id := range groupIDs {
    fmt.Printf("版本组: %s\n", id)
}
```


## 🔗 同类方法

- [`SortedVersionGroups.QueryRange`](/sdk/api/query-range-sortedversiongroups)
- [`SortedVersionGroups.Len`](/sdk/api/len-sortedversiongroups)
- [`SortedVersionGroups.Get`](/sdk/api/get-sortedversiongroups)
- [`SortedVersionGroups.At`](/sdk/api/at-sortedversiongroups)
- [`SortedVersionGroups.Contains`](/sdk/api/contains-sortedversiongroups)
- [`SortedVersionGroups.Versions`](/sdk/api/versions-sortedversiongroups)


---

::: details 源码位置
定义于 [`sorted_version_groups.go`](https://github.com/scagogogo/versions-skills/blob/main/sorted_version_groups.go)
:::
