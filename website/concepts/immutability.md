# 不可变性

::: tip 关键
所有 `With*` / `Bump*` 方法都返回**新的 Version 对象**，原对象永不修改。
:::

## 🔄 不可变变更

```go
v := versions.NewVersion("1.2.3")

// 不修改 v，返回新对象
v2 := v.WithMajor(2)         // v2 = "2.2.3"，v 仍是 "1.2.3"
v3 := v.BumpMinor()          // v3 = "1.3.0"（Patch 清零，后缀清除）
```

:::mermaid
flowchart LR
  V["v = 1.2.3<br/>（原对象）"]

  V -->|"WithMajor(2)<br/>不改 v"| V2["v2 = 2.2.3<br/>（新对象）"]
  V -->|"BumpMinor()<br/>不改 v"| V3["v3 = 1.3.0<br/>（新对象，Patch 清零+去后缀）"]
  V -->|"Clone()"| V4["v4 = 1.2.3<br/>（深拷贝）"]

  V -.->|"原对象永不修改"| SAFE["✅ 并发安全<br/>可共享、可复用"]

  style V fill:#eff6ff,stroke:#2563eb,stroke-width:3px
  style V2 fill:#f0fdf4,stroke:#16a34a
  style V3 fill:#f0fdf4,stroke:#16a34a
  style V4 fill:#f0fdf4,stroke:#16a34a
  style SAFE fill:#fff7ed,stroke:#ea580c,stroke-dasharray:4 3
:::

| 方法 | 作用 |
|:--|:--|
| `WithPrefix` / `WithSuffix` / `WithMajor` / `WithMinor` / `WithPatch` / `WithNumbers` / `WithPublicTime` / `WithMetadata` | 替换指定字段 |
| `BumpMajor` / `BumpMinor` / `BumpPatch` | 递增指定段并清零下级段 + 清后缀 |
| `Increment(segment)` | 递增指定段 |
| `Clone()` | 深拷贝 |
| `Core()` | 返回去后缀的核心版本 |

## 🧪 Bump 语义

| 输入 | 操作 | 输出 |
|:--|:--|:--|
| `1.2.3-beta1` | `BumpMajor()` | `2.0.0` |
| `1.2.3-beta1` | `BumpMinor()` | `1.3.0` |
| `1.2.3-beta1` | `BumpPatch()` | `1.2.4` |

::: warning 注意
`Bump*` 总是**清除后缀**——预发布版本 bump 后直接晋升为正式版。
:::

## ✅ 为什么不可变

- **并发安全**：多 goroutine 可共享同一 `*Version`，无需锁
- **可预测**：方法链不会产生隐式副作用
- **可复用**：原对象可作为基准派生多个变体

## 📚 延伸

- API：[变更方法索引](/sdk/mutation) · [`BumpMajor`](/sdk/api/bump-major-version) · [`WithPrefix`](/sdk/api/with-prefix-version) · [`Clone`](/sdk/api/clone-version)
- CLI：[`bump`](/cli/commands/bump) · [`set-*`](/cli/commands/set-prefix)
- 概念：[版本号结构](/concepts/version-anatomy)
