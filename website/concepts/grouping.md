# 分组语义

::: tip 关键
分组按版本号的**数字部分**归组。`BuildGroupID` 把数字段用点号**全部拼接**作为分组 ID。
:::

## 🏷 分组 ID

`BuildGroupID` 把所有数字段点号连接：

| 版本 | VersionNumbers | 分组 ID |
|:--|:--|:--|
| `v1.2.3` | `[1,2,3]` | `"1.2.3"` |
| `1.2.3-beta1` | `[1,2,3]` | `"1.2.3"` |
| `1.0.0` | `[1,0,0]` | `"1.0.0"` |
| `2.0` | `[2,0]` | `"2.0"` |

`v1.2.3` 与 `1.2.3-beta1` 数字段相同，归入**同一分组**——分组只看数字段，后缀不影响。

## 🧱 两种分组 API

### 简单分组 — `Group`

```go
groups := versions.Group(vs) // map[string]*VersionGroup
```

返回 `map[分组ID]*VersionGroup`，无序。

### 有序索引 — `SortedVersionGroups`

```go
svg := versions.NewSortedVersionGroups(vs)
svg.GroupIDs()  // 排序后的分组 ID 列表，如 ["1.0.0","1.1.0","2.0.0"]
svg.Get("1.0.0")  // 按ID取组
svg.QueryRange(start, end) // 高效范围查询
```

构建一次后，范围查询走有序索引，比每次重排快得多。

## 📚 延伸

- API：[`Group`](/sdk/api/group) · [`VersionGroup`](/sdk/api/version-group) · [`SortedVersionGroups`](/sdk/api/sorted-version-groups)
- 工具：[`GroupByMajor`](/sdk/api/group-by-major) · [`GroupByMinor`](/sdk/api/group-by-minor)
- 概念：[版本号结构](/concepts/version-anatomy)
