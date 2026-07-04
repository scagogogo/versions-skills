# 排序与极值

给版本号列表排序，找出最新/最旧、最新稳定/预发布版本。

:::mermaid
flowchart LR
  INPUT["无序列表<br/>1.10.0 / 1.2.0 / 1.1.0<br/>1.0.0 / 1.2.0-beta"]
  INPUT --> SORT["SortVersionSlice<br/>（数字段→后缀→时间→Raw）"]
  SORT --> ORDERED["有序列表<br/>1.0.0 → 1.1.0 → 1.2.0-beta → 1.2.0 → 1.10.0"]

  ORDERED --> MIN["Min 最旧<br/>1.0.0"]
  ORDERED --> MAX["Max 最新<br/>1.10.0"]
  ORDERED --> LS["LatestStable<br/>1.10.0"]
  ORDERED --> LP["LatestPrerelease<br/>1.2.0-beta"]
  INPUT -.->|"Filter"| FILTER["过滤<br/>稳定版/预发布/按 major"]

  style INPUT fill:#f8fafc,stroke:#475569
  style SORT fill:#eff6ff,stroke:#2563eb
  style ORDERED fill:#eff6ff,stroke:#2563eb,stroke-width:3px
  style MIN fill:#f0fdf4,stroke:#16a34a
  style MAX fill:#f0fdf4,stroke:#16a34a
  style LS fill:#f0fdf4,stroke:#16a34a
  style LP fill:#f0fdf4,stroke:#16a34a
  style FILTER fill:#fff7ed,stroke:#ea580c,stroke-dasharray:4 3
:::

## 📊 排序

```go
vs := versions.NewVersions("1.10.0", "1.2.0", "1.1.0", "1.0.0")

asc := versions.SortVersionSlice(vs)                          // 升序
desc := versions.SortVersionSlice(vs)                         // 原地已排
// 降序可用 VersionSlice
slice := versions.VersionSlice(asc)
slice.Sort()
// 反转
for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
    slice[i], slice[j] = slice[j], slice[i]
}
```

字符串版（保留原始书写）：

```go
strs := versions.SortVersionStringSlice([]string{"v1.10.0", "1.2.0"})
// ["1.2.0" "v1.10.0"]
```

## 🏆 极值

```go
vs := versions.NewVersions("1.0.0", "1.0.0-rc1", "1.1.0-beta1", "2.0.0")

versions.Max(vs).Raw               // 2.0.0（最新，含预发布）
versions.Min(vs).Raw               // 1.0.0-rc1（最旧）
versions.LatestStable(vs).Raw      // 2.0.0（最新稳定）
versions.LatestPrerelease(vs).Raw  // 1.1.0-beta1（最新预发布）
```

## 🎚 过滤

```go
stable := versions.FilterByStable(vs)       // 仅稳定版
pre := versions.FilterByPrerelease(vs)      // 仅预发布
maj2 := versions.FilterByMajor(vs, 2)       // 仅 Major=2
```

## 🚀 下一步

- [分组与聚合](/tutorials/grouping)
- API：[`SortVersionSlice`](/sdk/api/sort-version-slice) · [`Max`](/sdk/api/max) · [`LatestStable`](/sdk/api/latest-stable) · [`FilterByStable`](/sdk/api/filter-by-stable)
- CLI：[`sort`](/cli/commands/sort) · [`max`](/cli/commands/max) · [`latest-stable`](/cli/commands/latest-stable)
