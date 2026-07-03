# 分组语义

::: tip 关键
分组按版本号的**数字部分前缀**归组。默认取前两段（`major.minor`）作为分组 ID。
:::

## 🏷 分组 ID

`BuildGroupID` 取数字段前两段，点号连接：

| 版本 | VersionNumbers | 分组 ID |
|:--|:--|:--|
| `v1.2.3` | `[1,2,3]` | `"1.2"` |
| `1.2.3-beta1` | `[1,2,3]` | `"1.2"` |
| `1.0.0` | `[1,0,0]` | `"1.0"` |
| `2.0` | `[2,0]` | `"2.0"` |

`v1.2.3` 与 `1.2.3-beta1` 归入**同一分组** `1.2`——因为分组只看数字前缀。

## 🧱 两种分组 API

### 简单分组 — `Group`

```go
groups := versions.Group(vs) // map[string]*VersionGroup
```

返回 `map[分组ID]*VersionGroup`，无序。

### 有序索引 — `SortedVersionGroups`

```go
svg := versions.NewSortedVersionGroups(vs)
svg.GroupIDs()  // 排序后的分组 ID 列表
svg.Get("1.2")  // 按ID取组
svg.QueryRange(start, end) // 高效范围查询
```

构建一次后，范围查询走有序索引，比每次重排快得多。

## 📚 延伸

- API：[`Group`](/sdk/api/group) · [`VersionGroup`](/sdk/api/version-group) · [`SortedVersionGroups`](/sdk/api/sorted-version-groups)
- 工具：[`GroupByMajor`](/sdk/api/group-by-major) · [`GroupByMinor`](/sdk/api/group-by-minor)
- 概念：[版本号结构](/concepts/version-anatomy)
