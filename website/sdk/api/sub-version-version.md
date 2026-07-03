# SubVersion

::: info 方法 · `Version`
```go
func (x *Version) SubVersion() int
```
:::

## 📖 说明

SubVersion 返回后缀中的子版本号

例如 "-beta2" 返回 2，"-rc1" 返回 1。如果后缀中没有数字则返回 0。
该方法将内部使用的 extractSubVersion 函数暴露为公开 API。


#### 返回

- `int`：后缀中的子版本号数字


```go
v := versions.NewVersion("1.0.0-beta2")
fmt.Println(v.SubVersion()) // 输出: 2
```


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
