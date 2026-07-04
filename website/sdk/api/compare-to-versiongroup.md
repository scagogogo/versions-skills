# CompareTo

::: info 方法 · `VersionGroup`
```go
func (x *VersionGroup) CompareTo(target *VersionGroup) int
```
:::

## 📖 说明

CompareTo 比较两个版本组的大小

该方法通过比较版本组的数字部分来确定两个版本组的先后顺序。


#### 参数

- `target`：要比较的目标版本组


#### 返回

- `int`：如果当前版本组小于目标版本组，返回负数；如果相等，返回0；如果大于，返回正数


```go
group1 := versions.NewVersionGroup(vs.NewVersionNumbers([]int{1, 2}))
group2 := versions.NewVersionGroup(vs.NewVersionNumbers([]int{1, 3}))

if group1.CompareTo(group2) < 0 {
    fmt.Println("group1 比 group2 旧")
}
```


::: details 同类方法（点击展开）

**第 1 组**

- [`VersionGroup.Add`](/sdk/api/add-versiongroup)
- [`VersionGroup.Contains`](/sdk/api/contains-versiongroup)
- [`VersionGroup.ID`](/sdk/api/id-versiongroup)
- [`VersionGroup.Versions`](/sdk/api/versions-versiongroup)
- [`VersionGroup.SortVersions`](/sdk/api/sort-versions-versiongroup)
- [`VersionGroup.GetLatest`](/sdk/api/get-latest-versiongroup)
- [`VersionGroup.GetOldest`](/sdk/api/get-oldest-versiongroup)
- [`VersionGroup.Count`](/sdk/api/count-versiongroup)

**第 2 组**

- [`VersionGroup.StableVersions`](/sdk/api/stable-versions-versiongroup)
- [`VersionGroup.PrereleaseVersions`](/sdk/api/prerelease-versions-versiongroup)
- [`VersionGroup.LatestStable`](/sdk/api/latest-stable-versiongroup)
- [`VersionGroup.Remove`](/sdk/api/remove-versiongroup)
- [`VersionGroup.LatestPrerelease`](/sdk/api/latest-prerelease-versiongroup)
- [`VersionGroup.String`](/sdk/api/string-versiongroup)
- [`VersionGroup.Filter`](/sdk/api/filter-versiongroup)
- [`VersionGroup.QueryRangeVersions`](/sdk/api/query-range-versions-versiongroup)

:::


::: details 源码位置
定义于 [`version_group.go`](https://github.com/scagogogo/versions-skills/blob/main/version_group.go)
:::
