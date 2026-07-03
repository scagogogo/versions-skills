# Matches

::: info 方法 · `Version`
```go
func (x *Version) Matches(expr string) (bool, error)
```
:::

## 📖 说明

Matches 判断版本是否满足约束表达式字符串

该方法是 ParseConstraintSet + Match 的便捷组合，
适用于需要从字符串快速判断版本是否满足约束的场景。


#### 参数

- `expr`：约束表达式字符串，如 ">=1.0.0"


#### 返回

- `bool`：如果版本满足约束则返回 true
- `error`：如果约束表达式解析失败则返回错误


```go
v := versions.NewVersion("1.5.0")
ok, err := v.Matches(">=1.0.0,<2.0.0")
if ok {
    fmt.Println("版本在范围内")
}
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
- [`Version.IsNewerThan`](/sdk/api/is-newer-than-version)
- [`Version.IsOlderThan`](/sdk/api/is-older-than-version)
- [`Version.Equals`](/sdk/api/equals-version)
- [`Version.IsBetween`](/sdk/api/is-between-version)
- [`Version.Major`](/sdk/api/major-version)


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
