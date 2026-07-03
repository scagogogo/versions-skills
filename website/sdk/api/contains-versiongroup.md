# Contains

::: info 方法 · `VersionGroup`
```go
func (x *VersionGroup) Contains(v *Version) bool
```
:::

## 📖 说明

Contains 判断本版本组中是否包含给定的版本

该方法检查指定的版本是否已存在于版本组中。


#### 参数

- `v`：要检查的版本对象


#### 返回

- `bool`：如果版本存在于组中则返回 true，否则返回 false


```go
group := versions.NewVersionGroup(versions.NewVersionNumbers([]int{1, 2}))
version := versions.NewVersion("1.2.3")
if !group.Contains(version) {
    group.Add(version)
}
```


## 🔗 同类方法

- [`VersionGroup.Add`](/sdk/api/add-versiongroup)
- [`VersionGroup.ID`](/sdk/api/id-versiongroup)
- [`VersionGroup.CompareTo`](/sdk/api/compare-to-versiongroup)
- [`VersionGroup.Versions`](/sdk/api/versions-versiongroup)
- [`VersionGroup.SortVersions`](/sdk/api/sort-versions-versiongroup)
- [`VersionGroup.GetLatest`](/sdk/api/get-latest-versiongroup)
- [`VersionGroup.GetOldest`](/sdk/api/get-oldest-versiongroup)
- [`VersionGroup.Count`](/sdk/api/count-versiongroup)
- [`VersionGroup.StableVersions`](/sdk/api/stable-versions-versiongroup)
- [`VersionGroup.PrereleaseVersions`](/sdk/api/prerelease-versions-versiongroup)
- [`VersionGroup.LatestStable`](/sdk/api/latest-stable-versiongroup)
- [`VersionGroup.Remove`](/sdk/api/remove-versiongroup)
- [`VersionGroup.LatestPrerelease`](/sdk/api/latest-prerelease-versiongroup)
- [`VersionGroup.String`](/sdk/api/string-versiongroup)
- [`VersionGroup.Filter`](/sdk/api/filter-versiongroup)
- [`VersionGroup.QueryRangeVersions`](/sdk/api/query-range-versions-versiongroup)


---

::: details 源码位置
定义于 [`version_group.go`](https://github.com/scagogogo/versions-skills/blob/main/version_group.go)
:::
