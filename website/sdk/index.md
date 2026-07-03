# SDK API 索引

::: tip 包信息
- **导入路径**：`github.com/scagogogo/versions-skills`
- **包名**：`versions`
- **零重依赖**：核心仅依赖 `go-tuple` / `go-shuffle` / `go-compare-anything`（golang-infrastructure 自研库）
- **Go 版本**：≥ 1.21
:::

完整的 SDK API 由若干**类型**与**函数**组成，按能力域分组如下。每个条目都对应一篇专属文档页。

## 🧱 核心类型

| 类型 | 文档 | 说明 |
|:--|:--|:--|
| `Version` | [sdk/api/version.md](/sdk/api/version) | 版本对象，封装 Raw / PublicTime / VersionNumbers / Prefix / Suffix / Metadata |
| `VersionNumbers` | [sdk/api/version-numbers.md](/sdk/api/version-numbers) | `[]int` 数字段，实现 `CompareTo` |
| `VersionPrefix` | [sdk/api/version-prefix.md](/sdk/api/version-prefix) | 数字前导（如 `"v"`） |
| `VersionSuffix` | [sdk/api/version-suffix.md](/sdk/api/version-suffix) | 数字后内容（如 `"-beta1"`），按权重可比 |
| `SuffixWeight` | [sdk/api/suffix-weight.md](/sdk/api/suffix-weight) | 后缀语义权重枚举 |
| `VersionBuilder` | [sdk/api/version-builder.md](/sdk/api/version-builder) | 流式构造 Version |
| `VersionSlice` | [sdk/api/version-slice.md](/sdk/api/version-slice) | `[]*Version`，实现 `sort.Interface` |
| `VersionRange` | [sdk/api/version-range.md](/sdk/api/version-range) | 版本区间，支持开/闭边界 |
| `VersionDiff` | [sdk/api/version-diff.md](/sdk/api/version-diff) | 两版本差值（major/minor/patch） |
| `VersionGroup` | [sdk/api/version-group.md](/sdk/api/version-group) | 同数字前缀的版本组 |
| `SortedVersionGroups` | [sdk/api/sorted-version-groups.md](/sdk/api/sorted-version-groups) | 有序索引，高效范围查询 |
| `ContainsPolicy` | [sdk/api/contains-policy.md](/sdk/api/contains-policy) | 范围查询的包含策略 |

## 🔗 约束类型

| 类型 | 文档 | 说明 |
|:--|:--|:--|
| `Constraint` | [sdk/api/constraint.md](/sdk/api/constraint) | 单约束条件，如 `>=1.0.0` |
| `ConstraintOperator` | [sdk/api/constraint-operator.md](/sdk/api/constraint-operator) | 运算符枚举（`= != > >= < <= ^ ~ x`） |
| `ConstraintSet` | [sdk/api/constraint-set.md](/sdk/api/constraint-set) | AND 约束集（逗号分隔） |
| `ConstraintUnion` | [sdk/api/constraint-union.md](/sdk/api/constraint-union) | OR 约束联合（`\|\|` 分隔） |

## 🛠 能力域函数

| 能力域 | 文档 | 代表函数 |
|:--|:--|:--|
| 解析构造 | [sdk/parsing.md](/sdk/parsing) | `NewVersion` / `NewVersionE` / `MustParse` / `NewVersions` / `Coerce` / `CoerceE` |
| 版本属性 | [sdk/properties.md](/sdk/properties) | `Major` / `Minor` / `Patch` / `Segments` / `SubVersion` / `PreReleaseType` |
| 类型判断 | [sdk/predicates.md](/sdk/predicates) | `IsStable` / `IsPrerelease` / `IsAlpha` / `IsBeta` / `IsRC` / `IsSemver` |
| 比较 | [sdk/comparison.md](/sdk/comparison) | `CompareTo` / `IsNewerThan` / `IsOlderThan` / `Equals` / `IsBetween` / `Diff` |
| 排序 | [sdk/sorting.md](/sdk/sorting) | `SortVersionSlice` / `SortVersionStringSlice` / `VersionSlice.Sort` |
| 分组 | [sdk/grouping.md](/sdk/grouping) | `Group` / `GroupByMajor` / `GroupByMinor` / `NewSortedVersionGroups` |
| 过滤 | [sdk/filtering.md](/sdk/filtering) | `Filter` / `FilterByConstraint` / `FilterByStable` / `FilterByMajor` |
| 集合运算 | [sdk/set-operations.md](/sdk/set-operations) | `Min` / `Max` / `Difference` / `Intersection` / `Union` / `Partition` |
| 约束 | [sdk/constraints.md](/sdk/constraints) | `ParseConstraint` / `ParseConstraintSet` / `NegateConstraint` |
| 范围查询 | [sdk/range-query.md](/sdk/range-query) | `NewClosedRange` / `NewOpenRange` / `VersionRange.Contains` |
| 不可变变更 | [sdk/mutation.md](/sdk/mutation) | `BumpMajor` / `WithPrefix` / `WithSuffix` / `WithMajor` / `Increment` |
| 文件 IO | [sdk/file-io.md](/sdk/file-io) | `ReadVersionsFromFile` / `WriteVersionsToFile` / `ReadVersionsFromReader` |
| 可视化 | [sdk/visualization.md](/sdk/visualization) | `VisualizeVersions` / `VisualizeVersionGroups` |
| 序列化 | [sdk/serialization.md](/sdk/serialization) | `MarshalJSON` / `UnmarshalJSON` / `MarshalText` / `Scan` / `Value` |
## 常用示例

### 解析与检查

```go
v := versions.NewVersion("v1.2.3-beta1")
v.Major()          // 1
v.IsValid()        // true
v.IsPrerelease()   // true
v.PreReleaseType() // "beta"
v.IsSemver()       // true（semver 正则允许 v 前缀与预发布段）
v.SuffixWeight()   // 200 (beta)
```

### 比较

```go
a := versions.NewVersion("1.0.0")
b := versions.NewVersion("1.0.0-rc1")
a.IsNewerThan(b)        // true（数字段相同，稳定版 > 预发布版，由后缀权重决定）

// 注意：Diff 只比数字段差值，不看后缀
a.Diff(b).IsUpgrade()   // false（major/minor/patch 差值都为 0，不算升级）
a.Diff(b).IsPatchChange() // false（同理，patch 差值为 0）

d := versions.NewVersion("1.2.3").Diff(versions.NewVersion("1.2.5"))
d.IsPatchChange()       // true（仅 patch +2）
d.IsUpgrade()           // true
```

### 排序与分组

```go
list := versions.NewVersions("2.0.0", "1.0.0", "1.10.0", "1.5.0-beta")
sorted := versions.SortVersionSlice(list) // [1.0.0, 1.5.0-beta, 1.10.0, 2.0.0]

byMajor := versions.GroupByMajor(list) // map[int][]*Version{1: [...], 2: [...]}
```

### 约束

```go
v := versions.NewVersion("1.5.0")
ok, _ := v.Matches(">=1.0.0,<2.0.0 || >=3.0.0") // true

c, _ := versions.ParseConstraint("^1.2.3")
v.Satisfies(c) // 用 Version 侧方法
```

### 范围查询

```go
low := versions.NewVersion("1.0.0")
high := versions.NewVersion("2.0.0")
r := versions.NewClosedRange(low, high)
r.Contains(versions.NewVersion("1.5.0")) // true
r.Contains(versions.NewVersion("2.1.0")) // false

// 大列表反复查询
sg := versions.NewSortedVersionGroups(allVersions)
// 见 algorithms.md 第 7 节
```

### 不可变变更

```go
v := versions.NewVersion("1.2.3")
v2 := v.BumpPatch()    // 1.2.4（v 不变）
v3 := v.WithSuffix("-rc1") // 1.2.3-rc1
```

### 从任意字符串提取

```go
v := versions.Coerce("program-1.2.3-linux-amd64")
fmt.Println(v.Raw) // 1.2.3
```

## 设计要点

- **不可变**：所有 `With*` / `Bump*` 返回新对象，原对象不变。
- **Never nil from NewVersion**：`NewVersion` 永远返回非 nil 的 `Version`，用 `IsValid()` 判有效性。
- **SQL 支持**：`Version` 实现 `sql.Scanner` / `driver.Valuer`，可直接作为数据库字段。
- **零依赖核心**：核心库只依赖 `go-tuple`、`go-shuffle`、`go-compare-anything`（golang-infrastructure）。

