# PreReleaseType

::: info 方法 · `Version`
```go
func (x *Version) PreReleaseType() string
```
:::

## 📖 说明

PreReleaseType 返回预发布版本的类型标识

返回后缀的语义权重类型名称字符串，如 "alpha"、"beta"、"rc" 等。
如果不是预发布版本则返回空字符串。


#### 返回

- `string`：预发布类型名称


```go
v := versions.NewVersion("1.0.0-beta2")
fmt.Println(v.PreReleaseType()) // "beta"

s := versions.NewVersion("1.0.0")
fmt.Println(s.PreReleaseType()) // ""
```


::: details 同类方法（点击展开）

**第 1 组**

- [`Version.IsValid`](/sdk/api/is-valid-version)
- [`Version.BuildGroupID`](/sdk/api/build-group-id-version)
- [`Version.CompareTo`](/sdk/api/compare-to-version)
- [`Version.String`](/sdk/api/string-version)
- [`Version.IsPrerelease`](/sdk/api/is-prerelease-version)
- [`Version.IsStable`](/sdk/api/is-stable-version)
- [`Version.IsDev`](/sdk/api/is-dev-version)
- [`Version.IsAlpha`](/sdk/api/is-alpha-version)

**第 2 组**

- [`Version.IsBeta`](/sdk/api/is-beta-version)
- [`Version.IsRC`](/sdk/api/is-rc-version)
- [`Version.IsSnapshot`](/sdk/api/is-snapshot-version)
- [`Version.IsMilestone`](/sdk/api/is-milestone-version)
- [`Version.IsNightly`](/sdk/api/is-nightly-version)
- [`Version.IsFinal`](/sdk/api/is-final-version)
- [`Version.IsGA`](/sdk/api/is-ga-version)
- [`Version.IsPre`](/sdk/api/is-pre-version)

**第 3 组**

- [`Version.IsRelease`](/sdk/api/is-release-version)
- [`Version.IsSP`](/sdk/api/is-sp-version)
- [`Version.IsPost`](/sdk/api/is-post-version)
- [`Version.Satisfies`](/sdk/api/satisfies-version)
- [`Version.Matches`](/sdk/api/matches-version)
- [`Version.IsNewerThan`](/sdk/api/is-newer-than-version)
- [`Version.IsOlderThan`](/sdk/api/is-older-than-version)
- [`Version.Equals`](/sdk/api/equals-version)

**第 4 组**

- [`Version.IsBetween`](/sdk/api/is-between-version)

:::


::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
