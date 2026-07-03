# QueryRange

::: info 方法 · `SortedVersionGroups`
```go
func (x *SortedVersionGroups) QueryRange(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version
```
:::

## 📖 说明

QueryRange 在有序版本组中查询指定范围内的版本

该方法根据给定的起始和结束版本范围，返回所有符合条件的版本对象数组。
方法利用预排序的版本组结构，快速定位并收集符合范围条件的版本。


#### 参数

- `start`：包含起始版本和包含策略的元组
- `end`：包含结束版本和包含策略的元组


#### 返回

- `[]*Version`：符合查询范围条件的版本对象数组


```go
// 创建已排序的版本组
sortedGroups := versions.NewSortedVersionGroups(allVersions)

// 定义查询范围
startVer := versions.NewVersion("1.0.0")
endVer := versions.NewVersion("2.0.0")
startTuple := tuple.NewTuple2(startVer, versions.ContainsPolicyYes) // 包含1.0.0
endTuple := tuple.NewTuple2(endVer, versions.ContainsPolicyNo)      // 不包含2.0.0

// 执行范围查询
rangeResult := sortedGroups.QueryRange(startTuple, endTuple)
fmt.Printf("在范围内的版本数: %d\n", len(rangeResult))
```


## 🔗 同类方法

- [`SortedVersionGroups.GroupIDs`](/sdk/api/group-ids-sortedversiongroups)
- [`SortedVersionGroups.Len`](/sdk/api/len-sortedversiongroups)
- [`SortedVersionGroups.Get`](/sdk/api/get-sortedversiongroups)
- [`SortedVersionGroups.At`](/sdk/api/at-sortedversiongroups)
- [`SortedVersionGroups.Contains`](/sdk/api/contains-sortedversiongroups)
- [`SortedVersionGroups.Versions`](/sdk/api/versions-sortedversiongroups)


---

::: details 源码位置
定义于 [`sorted_version_groups.go`](https://github.com/scagogogo/versions-skills/blob/main/sorted_version_groups.go)
:::
