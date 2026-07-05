# 零依赖设计

::: tip 关键
根包 `versions` **零外部重依赖**——仅依赖 golang-infrastructure 自研的 3 个小库。
:::

## 📦 依赖清单

| 依赖 | 来源 | 用途 |
|:--|:--|:--|
| `go-tuple` | golang-infrastructure | 二元组（范围查询的端点对） |
| `go-shuffle` | golang-infrastructure | 切片乱序（测试辅助） |
| `go-compare-anything` | golang-infrastructure | 通用 `Comparable` 接口 |

CLI 额外依赖 `cobra`，MCP 额外依赖 `mcp-go`——但**核心库本身**保持轻量。

```mermaid
flowchart TB
  subgraph CORE["根包 versions（核心库）"]
    direction LR
    P["解析/比较/排序<br/>分组/约束/范围"]
  end

  CORE --> D1["go-tuple<br/>golang-infrastructure"]
  CORE --> D2["go-shuffle<br/>golang-infrastructure"]
  CORE --> D3["go-compare-anything<br/>golang-infrastructure"]

  subgraph EXT["扩展入口（独立二进制）"]
    direction LR
    CLI["CLI"]
    MCP["MCP"]
    SDK["SDK 用户代码"]
  end

  CORE --> EXT
  CLI -.->|"额外"| COBRA["cobra"]
  MCP -.->|"额外"| MCPGO["mcp-go"]

  D1 ~~~ D2 ~~~ D3

  style CORE fill:#eff6ff,stroke:#2563eb,stroke-width:3px
  style D1 fill:#f0fdf4,stroke:#16a34a
  style D2 fill:#f0fdf4,stroke:#16a34a
  style D3 fill:#f0fdf4,stroke:#16a34a
  style CLI fill:#fff7ed,stroke:#ea580c
  style MCP fill:#fff7ed,stroke:#ea580c
  style SDK fill:#fff7ed,stroke:#ea580c
  style COBRA fill:#f8fafc,stroke:#94a3b8,stroke-dasharray:4 3
  style MCPGO fill:#f8fafc,stroke:#94a3b8,stroke-dasharray:4 3
```

## ✅ 为什么重要

- **可移植**：无传递依赖地狱，`go get` 即用
- **安全**：依赖面小，供应链风险低
- **可维护**：所有依赖来自同一作者，风格统一、升级可控
- **可嵌入**：适合嵌入到对依赖敏感的工具链 / agent 框架

## 🧪 验证

```bash
go list -m all          # 查看实际依赖
go mod why golang-infrastructure/go-tuple
```

## 📚 延伸

- 概念：[三层接入](/concepts/three-layers)
- 入门：[快速开始](/quick-start)
