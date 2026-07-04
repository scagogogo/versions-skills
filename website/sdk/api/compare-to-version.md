# CompareTo

::: info 方法 · `Version`
```go
func (x *Version) CompareTo(target *Version) int
```
:::

## 📖 说明

CompareTo 比较两个版本号

该方法按以下顺序比较两个版本号：
1. 首先比较版本号数字部分（VersionNumbers 数组，如 [1,2,3]，逐段比较，非仅主版本号）
2. 其次比较后缀（空后缀=正式版 > 有后缀=预发布版）
3. 然后比较发布时间
4. 最后比较原始版本号字符串

::: warning 注释与实现
源码 doc 注释把第 2、3 步写成「发布时间→后缀」、第 1 步写成「主版本号数字部分」，但实现实际是「后缀→发布时间」，且第 1 步比较的是整个 VersionNumbers 数组而非仅主版本号。以本文档为准，详见 [比较优先级](/concepts/compare-priority)。
:::


#### 参数

- `target`：要比较的目标版本对象


#### 返回

- `int`：如果当前版本小于目标版本，返回-1；如果相等，返回0；如果大于，返回1


```go
v1 := versions.NewVersion("1.0.0")
v2 := versions.NewVersion("1.1.0")

switch v1.CompareTo(v2) {
case -1:
    fmt.Println("v1 < v2")
case 0:
    fmt.Println("v1 = v2")
case 1:
    fmt.Println("v1 > v2")
}
```


::: details 同类方法（点击展开）

**第 1 组**

- [`Version.IsValid`](/sdk/api/is-valid-version)
- [`Version.BuildGroupID`](/sdk/api/build-group-id-version)
- [`Version.String`](/sdk/api/string-version)
- [`Version.IsPrerelease`](/sdk/api/is-prerelease-version)
- [`Version.IsStable`](/sdk/api/is-stable-version)
- [`Version.IsDev`](/sdk/api/is-dev-version)
- [`Version.IsAlpha`](/sdk/api/is-alpha-version)
- [`Version.IsBeta`](/sdk/api/is-beta-version)

**第 2 组**

- [`Version.IsRC`](/sdk/api/is-rc-version)
- [`Version.IsSnapshot`](/sdk/api/is-snapshot-version)
- [`Version.IsMilestone`](/sdk/api/is-milestone-version)
- [`Version.IsNightly`](/sdk/api/is-nightly-version)
- [`Version.IsFinal`](/sdk/api/is-final-version)
- [`Version.IsGA`](/sdk/api/is-ga-version)
- [`Version.IsPre`](/sdk/api/is-pre-version)
- [`Version.IsRelease`](/sdk/api/is-release-version)

**第 3 组**

- [`Version.IsSP`](/sdk/api/is-sp-version)
- [`Version.IsPost`](/sdk/api/is-post-version)
- [`Version.Satisfies`](/sdk/api/satisfies-version)
- [`Version.Matches`](/sdk/api/matches-version)
- [`Version.IsNewerThan`](/sdk/api/is-newer-than-version)
- [`Version.IsOlderThan`](/sdk/api/is-older-than-version)
- [`Version.Equals`](/sdk/api/equals-version)
- [`Version.IsBetween`](/sdk/api/is-between-version)

**第 4 组**

- [`Version.Major`](/sdk/api/major-version)

:::


::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
