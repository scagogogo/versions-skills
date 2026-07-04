# IsSemver

::: info 方法 · `Version`
```go
func (x *Version) IsSemver() bool
```
:::

## 📖 说明

IsSemver 判断版本字符串是否符合 SemVer 2.0.0 规范

严格的 semver 格式要求：主版本号.次版本号.修订版本号，
可选的预发布标识（以 - 分隔）和构建元数据（以 + 分隔）。
不允许前导零（如 01.02.03）。


#### 返回

- `bool`：如果符合 semver 规范则返回 true


```go
v := versions.NewVersion("1.2.3")
fmt.Println(v.IsSemver()) // true

v2 := versions.NewVersion("1.2.3-alpha.1+build.123")
fmt.Println(v2.IsSemver()) // true

v3 := versions.NewVersion("1.2")
fmt.Println(v3.IsSemver()) // false（不够3段）
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
