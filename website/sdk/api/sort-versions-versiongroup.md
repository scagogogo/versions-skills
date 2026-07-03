# SortVersions

::: info 方法 · `VersionGroup`
```go
func (x *VersionGroup) SortVersions() []*Version
```
:::

## 📖 说明

SortVersions 对组下的所有版本进行排序返回

该方法返回版本组中所有版本的有序数组，排序遵循版本号的自然排序规则。


#### 返回

- `[]*Version`：排序后的版本数组


```go
group := versions.NewVersionGroupFromVersions(versions.NewVersions("1.2.2", "1.2.0", "1.2.1"))
sortedVersions := group.SortVersions()
// 结果顺序: ["1.2.0", "1.2.1", "1.2.2"]
```


## 🔗 同类方法

- [`VersionGroup.Add`](/sdk/api/add-versiongroup)
- [`VersionGroup.Contains`](/sdk/api/contains-versiongroup)
- [`VersionGroup.ID`](/sdk/api/id-versiongroup)
- [`VersionGroup.CompareTo`](/sdk/api/compare-to-versiongroup)
- [`VersionGroup.Versions`](/sdk/api/versions-versiongroup)
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
