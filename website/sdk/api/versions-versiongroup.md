# Versions

::: info 方法 · `VersionGroup`
```go
func (x *VersionGroup) Versions() []*Version
```
:::

## 📖 说明

Versions 返回组下的所有版本

该方法返回版本组中包含的所有版本的数组，结果不保证有序。


#### 返回

- `[]*Version`：版本组中所有版本的数组


```go
group := versions.NewVersionGroupFromVersions(versions.NewVersions("1.2.0", "1.2.1"))
allVersions := group.Versions()
fmt.Printf("组中包含 %d 个版本\n", len(allVersions))
```


## 🔗 同类方法

- [`VersionGroup.Add`](/sdk/api/add-versiongroup)
- [`VersionGroup.Contains`](/sdk/api/contains-versiongroup)
- [`VersionGroup.ID`](/sdk/api/id-versiongroup)
- [`VersionGroup.CompareTo`](/sdk/api/compare-to-versiongroup)
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
