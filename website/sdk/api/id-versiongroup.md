# ID

::: info 方法 · `VersionGroup`
```go
func (x *VersionGroup) ID() string
```
:::

## 📖 说明

ID 返回组的ID

该方法返回版本组的唯一标识符，由其数字部分生成。


#### 返回

- `string`：版本组的ID，例如 "1.2"


```go
group := versions.NewVersionGroup(vs.NewVersionNumbers([]int{1, 2}))
groupID := group.ID() // 返回 "1.2"
```


::: details 同类方法（点击展开）

**第 1 组**

- [`VersionGroup.Add`](/sdk/api/add-versiongroup)
- [`VersionGroup.Contains`](/sdk/api/contains-versiongroup)
- [`VersionGroup.CompareTo`](/sdk/api/compare-to-versiongroup)
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
