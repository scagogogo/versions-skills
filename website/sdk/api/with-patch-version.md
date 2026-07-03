# WithPatch

::: info 方法 · `Version`
```go
func (x *Version) WithPatch(patch int) *Version
```
:::

## 📖 说明

WithPatch 返回一个修改修订版本号的新版本对象

原版本对象不变，返回一个新对象，其修订版本号被替换为指定值。
前缀和后缀保持不变。


#### 参数

- `patch`：新的修订版本号


#### 返回

- `*Version`：修改修订版本号后的新版本对象


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
定义于 [`version_clone.go`](https://github.com/scagogogo/versions-skills/blob/main/version_clone.go)
:::
