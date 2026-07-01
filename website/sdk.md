# Go SDK API

核心库包名 `versions`，导入路径 `github.com/scagogogo/versions-skills`。零外部重依赖。

## 核心类型

| 类型 | 说明 |
|:--|:--|
| `Version` | 版本对象，含 Raw / PublicTime / VersionNumbers / Prefix / Suffix / Metadata |
| `VersionNumbers` | `[]int`，数字段，实现 `CompareTo` |
| `VersionPrefix` | `string`，数字前导（如 `"v"`） |
| `VersionSuffix` | `string`，数字后内容（如 `"-beta1"`），按权重可比 |
| `VersionRange` | 版本区间，支持开/闭边界 |
| `VersionDiff` | 两版本差值（major/minor/patch） |
| `VersionGroup` | 同数字前缀的版本组 |
| `SortedVersionGroups` | 有序索引，高效范围查询 |
| `Constraint` / `ConstraintSet` / `ConstraintUnion` | 单约束 / AND 集 / OR 联合 |
| `VersionBuilder` | 流式构造 Version |
| `VersionSlice` | `[]*Version`，实现 `sort.Interface` |
| `SuffixWeight` | 后缀语义权重枚举 |

## 关键函数

| 类别 | 函数 |
|:--|:--|
| **解析** | `NewVersion`、`NewVersionE`、`MustParse`、`NewVersions`、`Coerce`、`CoerceE` |
| **比较** | `CompareTo`、`IsNewerThan`、`IsOlderThan`、`Equals`、`IsBetween`、`Diff` |
| **排序** | `SortVersionSlice`、`SortVersionStringSlice`、`VersionSlice.Sort()` |
| **分组** | `Group`、`GroupByMajor`、`GroupByMinor`、`NewSortedVersionGroups` |
| **过滤** | `Filter`、`FilterByConstraint`、`FilterByStable`、`FilterByMajor`、`Unique` |
| **约束** | `ParseConstraint`、`ParseConstraintSet`、`ParseConstraintUnion`、`NegateConstraint` |
| **范围** | `NewClosedRange`、`NewOpenRange`、`VersionRange.Contains`、`VersionRange.Filter` |
| **检查** | `IsPrerelease`、`IsStable`、`IsSemver`、`ValidateSemver`、`PreReleaseType` |
| **变更** | `BumpMajor`、`BumpMinor`、`BumpPatch`、`WithPrefix`、`WithSuffix`、`WithMajor`、`Increment` |
| **集合** | `Min`、`Max`、`LatestStable`、`ContainsVersion`、`IndexOf`、`Difference`、`Intersection`、`Union`、`Partition` |
| **文件** | `ReadVersionsFromFile`、`WriteVersionsToFile`、`ReadVersionsFromReader` |
| **可视化** | `VisualizeVersions`、`VisualizeVersionGroups` |
| **序列化** | `MarshalJSON`、`UnmarshalJSON`、`MarshalText`、`UnmarshalText`、`Scan`、`Value` |

## 常用示例

### 解析与检查

```go
v := versions.NewVersion("v1.2.3-beta1")
v.Major()          // 1
v.IsValid()        // true
v.IsPrerelease()   // true
v.PreReleaseType() // "beta"
v.IsSemver()       // false（不够 3 段？实际 1.2.3-beta1 是 3 段，true）
v.SuffixWeight()   // 200 (beta)
```

### 比较

```go
a := versions.NewVersion("1.0.0")
b := versions.NewVersion("1.0.0-rc1")
a.IsNewerThan(b)        // true
a.Diff(b).IsUpgrade()   // true
a.Diff(b).IsPatchChange() // true（同 major/minor，patch 都 0... 见 Diff 语义）
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

→ 算法语义见 [算法详解](./algorithms)。CLI 等价命令见 [CLI](./cli)。
