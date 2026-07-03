# Patch

::: info 方法 · `Version`
```go
func (x *Version) Patch() int
```
:::

## 📖 说明

Patch 返回修订版本号

如果版本号数字部分少于3位则返回 0。


## 🔗 同类方法

- [`Version.IsValid`](/sdk/api/is-valid-version)
- [`Version.BuildGroupID`](/sdk/api/build-group-id-version)
- [`Version.CompareTo`](/sdk/api/compare-to-version)
- [`Version.String`](/sdk/api/string-version)
- [`Version.IsPrerelease`](/sdk/api/is-prerelease-version)
- [`Version.IsStable`](/sdk/api/is-stable-version)
- [`Version.IsDev`](/sdk/api/is-dev-version)
- [`Version.IsAlpha`](/sdk/api/is-alpha-version)
- [`Version.IsBeta`](/sdk/api/is-beta-version)
- [`Version.IsRC`](/sdk/api/is-rc-version)
- [`Version.IsSnapshot`](/sdk/api/is-snapshot-version)
- [`Version.IsMilestone`](/sdk/api/is-milestone-version)
- [`Version.IsNightly`](/sdk/api/is-nightly-version)
- [`Version.IsFinal`](/sdk/api/is-final-version)
- [`Version.IsGA`](/sdk/api/is-ga-version)
- [`Version.IsPre`](/sdk/api/is-pre-version)
- [`Version.IsRelease`](/sdk/api/is-release-version)
- [`Version.IsSP`](/sdk/api/is-sp-version)
- [`Version.IsPost`](/sdk/api/is-post-version)
- [`Version.Satisfies`](/sdk/api/satisfies-version)
- [`Version.Matches`](/sdk/api/matches-version)
- [`Version.IsNewerThan`](/sdk/api/is-newer-than-version)
- [`Version.IsOlderThan`](/sdk/api/is-older-than-version)
- [`Version.Equals`](/sdk/api/equals-version)
- [`Version.IsBetween`](/sdk/api/is-between-version)


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
