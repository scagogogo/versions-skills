# GetOldest

::: info 方法 · `VersionGroup`
```go
func (x *VersionGroup) GetOldest() *Version
```
:::

## 📖 说明

GetOldest 获取版本组中最旧的版本

返回按排序后最旧的版本，如果组为空则返回 nil。


## 🔗 同类方法

- [`VersionGroup.Add`](/sdk/api/add-versiongroup)
- [`VersionGroup.Contains`](/sdk/api/contains-versiongroup)
- [`VersionGroup.ID`](/sdk/api/id-versiongroup)
- [`VersionGroup.CompareTo`](/sdk/api/compare-to-versiongroup)
- [`VersionGroup.Versions`](/sdk/api/versions-versiongroup)
- [`VersionGroup.SortVersions`](/sdk/api/sort-versions-versiongroup)
- [`VersionGroup.GetLatest`](/sdk/api/get-latest-versiongroup)
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
