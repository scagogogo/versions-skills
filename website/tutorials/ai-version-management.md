# 让 Claude 管理版本

通过 MCP 让 AI Agent 自动完成版本号解析、比较、排序等操作。

:::mermaid
flowchart LR
  USER["💬 你下达任务<br/>「检查这些版本<br/>哪个最新稳定」"]
  USER --> AGENT["🤖 AI Agent<br/>Claude/Codex/Cursor"]
  AGENT -->|"调用"| MCP["MCP Server<br/>versions-mcp<br/>21 个工具"]
  MCP --> CORE["核心库 versions"]
  CORE --> RESULT["结构化结果<br/>返回给 Agent"]
  RESULT --> AGENT
  AGENT -->|"自然语言回复"| USER

  style USER fill:#eff6ff,stroke:#2563eb
  style AGENT fill:#fff7ed,stroke:#ea580c
  style MCP fill:#eff6ff,stroke:#2563eb,stroke-width:3px
  style CORE fill:#f0fdf4,stroke:#16a34a
  style RESULT fill:#f0fdf4,stroke:#16a34a
:::

## 🔧 接入 MCP

1. 构建并启动 server：

```bash
go build -o versions-mcp ./cmd/versions-mcp
```

2. 在 Claude Code / Claude Desktop 的 MCP 配置中加入：

```json
{
  "mcpServers": {
    "versions": { "command": "/path/to/versions-mcp" }
  }
}
```

接入后，Claude 即可调用 [21 个版本工具](/mcp/)。

## 💬 实际对话

接入后，你可以直接对 Claude 说：

> 「我有这些版本：1.10.0、1.2.0、1.1.0。帮我排序并找出最新稳定版。」

Claude 会自动调用 `version_sort` 与 `version_latest_stable` 工具，返回结果，无需你手敲命令。

## 📝 一键提示词

把 [一键提示词](/prompts) 里的内容粘给 Claude，它会知道如何组合调用版本工具完成复杂任务（如「找出满足 `^1.2.0` 且非预发布的最新版本」）。

## 🆚 MCP vs CLI

| 维度 | MCP | CLI（+ 提示词） |
|:--|:--|:--|
| 安装 | 配置一次 | 装好 CLI 即可 |
| 调用 | Claude 自动选工具 | Claude 生成 shell 命令 |
| 适合 | 长期集成、频繁使用 | 临时任务、不想跑 server |
| 结构化输出 | JSON 直接消费 | 文本解析 |

两种方式详见 [AI Agent 接入指南](/ai-agents)。

## 🚀 下一步

- [CI/CD 中的版本判断](/tutorials/ci-cd)
- [AI Agent 接入指南](/ai-agents) · [一键提示词](/prompts)
- [MCP 工具总览](/mcp/)
