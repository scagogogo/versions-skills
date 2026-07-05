# 范围与包含策略

::: tip 关键
`VersionRange` 描述一个**数值区间**（开/闭边界）；`ContainsPolicy` 控制**端点字符串**的包含/排除。
:::

## 🎚 VersionRange

```go
r := versions.NewClosedRange(versions.NewVersion("1.0.0"), versions.NewVersion("2.0.0"))
r.Contains(versions.NewVersion("1.5.0")) // true（闭区间含两端）
r.Filter(vs) // 过滤出区间内的版本
```

| 构造器 | 边界 | 示例 |
|:--|:--|:--|
| `NewClosedRange(low, high)` | `[low, high]` 含两端 | `[1.0.0, 2.0.0]` |
| `NewOpenRange(low, high)` | `(low, high)` 不含两端 | `(1.0.0, 2.0.0)` |
| `NewVersionRange(low, high, loIn, hiIn)` | 自定义两端开闭 | — |

四种开闭组合示意（`●` 为闭端点，含；`○` 为开端点，不含）：

```mermaid
flowchart LR
  subgraph CC["[1.0.0, 2.0.0]  闭-闭"]
    direction LR
    CC1["● 1.0.0"] --- CC2["1.5.0 ✅"] --- CC3["2.0.0 ●"]
  end
  subgraph OO["(1.0.0, 2.0.0)  开-开"]
    direction LR
    OO1["○ 1.0.0"] --- OO2["1.5.0 ✅"] --- OO3["2.0.0 ○"]
  end
  subgraph CO["[1.0.0, 2.0.0)  闭-开"]
    direction LR
    CO1["● 1.0.0"] --- CO2["1.5.0 ✅"] --- CO3["2.0.0 ○"]
  end
  subgraph OC["(1.0.0, 2.0.0]  开-闭"]
    direction LR
    OC1["○ 1.0.0"] --- OC2["1.5.0 ✅"] --- OC3["2.0.0 ●"]
  end

  style CC2 fill:#f0fdf4,stroke:#16a34a
  style OO2 fill:#f0fdf4,stroke:#16a34a
  style CO2 fill:#f0fdf4,stroke:#16a34a
  style OC2 fill:#f0fdf4,stroke:#16a34a
  style CC1 fill:#eff6ff,stroke:#2563eb
  style CC3 fill:#eff6ff,stroke:#2563eb
  style OC3 fill:#eff6ff,stroke:#2563eb
  style CO1 fill:#eff6ff,stroke:#2563eb
```

## 🎚 ContainsPolicy

`ContainsPolicy` 用于 `SortedVersionGroups.QueryRange`，控制是否包含**匹配某字符串**的端点：

| 枚举 | 值 | 语义 |
|:--|:--|:--|
| `ContainsPolicyNone` | 0 | 不基于字符串过滤 |
| `ContainsPolicyYes` | 1 | 仅保留包含指定字符串的版本 |
| `ContainsPolicyNo` | 2 | 排除包含指定字符串的版本 |

典型用途：在 `[1.0.0, 2.0.0]` 范围内**排除所有预发布版本**（`ContainsPolicyNo` + `"-"`）。

## 📚 延伸

- API：[`VersionRange`](/sdk/api/version-range) · [`ContainsPolicy`](/sdk/api/contains-policy) · [`SortedVersionGroups.QueryRange`](/sdk/api/query-range-sortedversiongroups)
- 概念：[约束表达式](/concepts/constraints) · [分组语义](/concepts/grouping)
- CLI：[`range`](/cli/commands/range)
