# 版本号结构

::: tip 关键
一个版本号被解析为 **5 个组成部分**：`Raw` / `Prefix` / `VersionNumbers` / `Suffix` / `Metadata`（外加可选的 `PublicTime`）。
:::

## 🧱 五段分解

以 `v1.2.3-beta1+build.7` 为例：

```
v  1.2.3  -beta1  +build.7
│  │      │       └─ Metadata（构建元数据，不参与比较）
│  │      └──────── Suffix（数字后内容，按权重比较）
│  └─────────────── VersionNumbers（数字段 []int）
└────────────────── Prefix（数字前导）
```

| 字段 | 类型 | 值 | 说明 |
|:--|:--|:--|:--|
| `Raw` | `string` | `"v1.2.3-beta1+build.7"` | 原始字符串 |
| `Prefix` | `VersionPrefix` | `"v"` | 数字前导，可为空 |
| `VersionNumbers` | `[]int` | `[1, 2, 3]` | 数字段 |
| `Suffix` | `VersionSuffix` | `"-beta1"` | 数字后内容 |
| `Metadata` | `string` | `"build.7"` | semver 构建元数据 |

```mermaid
flowchart LR
  subgraph RAW["v1.2.3-beta1+build.7"]
    direction LR
    P["v<br/><small>Prefix</small>"] --- VN["1.2.3<br/><small>VersionNumbers</small>"] --- S["-beta1<br/><small>Suffix</small>"] --- M["+build.7<br/><small>Metadata</small>"]
  end

  VN -.->|"比较核心<br/>逐段比较"| CMP["参与比较"]
  S -.->|"数字段相同时<br/>按权重比较"| CMP
  P -.->|"仅展示<br/>不参与"| IGNORE["不参与比较"]
  M -.->|"完全忽略"| IGNORE
  RAW -.->|"PublicTime<br/>可选附加"| PT["发布时间"]

  style P fill:#f8fafc,stroke:#475569
  style VN fill:#eff6ff,stroke:#2563eb,stroke-width:3px
  style S fill:#eff6ff,stroke:#2563eb
  style M fill:#f8fafc,stroke:#94a3b8,stroke-dasharray: 4 3
  style CMP fill:#f0fdf4,stroke:#16a34a
  style IGNORE fill:#fef2f2,stroke:#dc2626
  style PT fill:#fff7ed,stroke:#ea580c
```

## 🔍 各段行为

- **Prefix**：仅显示用，**不参与比较**。`1.2.3` 与 `v1.2.3` 等价。
- **VersionNumbers**：比较的核心。逐段比较，缺失段视为 0。
- **Suffix**：当数字段完全相同时，按[后缀权重](/concepts/suffix-weight)比较。
- **Metadata**：完全忽略，仅作展示。

## 🧭 延伸

- 类型定义见 [Version](/sdk/api/version)
- 解析过程见 [工作原理](/how-it-works)
- 后缀比较见 [后缀权重](/concepts/suffix-weight)
