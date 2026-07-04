# 范围查询

查询「在某区间内的版本」，并用包含策略排除预发布等。

## 🎚 VersionRange

```go
low := versions.NewVersion("1.0.0")
high := versions.NewVersion("2.0.0")

r := versions.NewClosedRange(low, high) // [1.0.0, 2.0.0]
r.Contains(versions.NewVersion("1.5.0")) // true
r.Contains(versions.NewVersion("2.0.0")) // true（闭区间含上界）

open := versions.NewOpenRange(low, high) // (1.0.0, 2.0.0)
open.Contains(versions.NewVersion("2.0.0")) // false（开区间不含上界）

// 自定义两端开闭
r2 := versions.NewVersionRange(low, high, true, false) // [1.0.0, 2.0.0)
r2.Contains(versions.NewVersion("2.0.0")) // false
```

## 🔍 过滤区间内版本

```go
vs := versions.NewVersions("0.9.0", "1.0.0", "1.5.0", "2.0.0", "2.1.0")
r := versions.NewClosedRange(versions.NewVersion("1.0.0"), versions.NewVersion("2.0.0"))
inRange := r.Filter(vs) // [1.0.0 1.5.0 2.0.0]
```

## 🎯 IsBetween 便捷方法

```go
v := versions.NewVersion("1.5.0")
low := versions.NewVersion("1.0.0")
high := versions.NewVersion("2.0.0")
v.IsBetween(low, high) // true（闭区间）
```

## 🚫 排除预发布（ContainsPolicy）

`SortedVersionGroups.QueryRange` 支持用 `ContainsPolicy` 按字符串过滤端点。典型场景：在范围内排除所有预发布版本（后缀含 `-`）。

:::mermaid
flowchart LR
  IN["全部版本<br/>1.0.0 / 1.0.0-rc1 / 1.1.0 / 1.5.0-beta / 2.0.0"]
  IN --> QR["QueryRange<br/>[1.0.0, 2.0.0]"]
  QR --> RANGE["区间内<br/>1.0.0 / 1.0.0-rc1 / 1.1.0 / 1.5.0-beta / 2.0.0"]
  RANGE --> POLICY{"ContainsPolicyNo + '-'"}
  POLICY -->|"含 - 的排除"| KEEP["保留<br/>1.0.0 / 1.1.0 / 2.0.0"]
  POLICY -->|"1.0.0-rc1 / 1.5.0-beta 含 -"| EXCL["排除预发布"]

  style IN fill:#f8fafc,stroke:#475569
  style QR fill:#eff6ff,stroke:#2563eb
  style RANGE fill:#eff6ff,stroke:#2563eb
  style POLICY fill:#fff7ed,stroke:#ea580c
  style KEEP fill:#f0fdf4,stroke:#16a34a,stroke-width:3px
  style EXCL fill:#fef2f2,stroke:#dc2626
:::

详见 [范围与包含策略](/concepts/range-and-policy) 与 [`SortedVersionGroups.QueryRange`](/sdk/api/query-range-sortedversiongroups)。

## 🚀 下一步

- [不可变变更与发布流程](/tutorials/bump-and-release)
- API：[`VersionRange`](/sdk/api/version-range) · [`IsBetween`](/sdk/api/is-between-version)
- CLI：[`range`](/cli/commands/range)
