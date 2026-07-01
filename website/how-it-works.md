# 工作原理

versions-skills 是一个**四层架构**的工具集：核心库在底，能力向上通过三种接口暴露给 AI Agent 和人类。

## 四层架构

:::mermaid
flowchart TB
  subgraph L1["🤖 AI Agent / IDE 层"]
    direction LR
    CC["Claude Code"]
    CX["Codex"]
    CUR["Cursor"]
    WS["Windsurf · Cline · VS Code"]
  end
  subgraph L2["接口层"]
    direction LR
    SK["Skills Plugin<br/>14 SKILL.md → 斜杠命令<br/>领域知识注入"]
    MCP["MCP Server<br/>21 version_* 工具<br/>结构化 JSON 响应"]
    CLI["CLI Binary<br/>shell / CI/CD"]
    SDK["Go SDK<br/>程序内嵌"]
  end
  subgraph L3["核心库（Go · 零依赖）"]
    CORE["parser · compare · sort · constraint · range · group"]
  end
  CC --> SK
  CC --> MCP
  CX --> MCP
  CUR --> MCP
  WS --> MCP
  SK --> CLI
  SK --> SDK
  MCP --> CORE
  CLI --> CORE
  SDK --> CORE
  style L1 fill:#eff6ff,stroke:#2563eb
  style L2 fill:#f0fdf4,stroke:#16a34a
  style L3 fill:#f8fafc,stroke:#475569
  style CORE fill:#1e293b,color:#fff,stroke:#0f172a
:::

**核心库**（仓库根目录的 `*.go`）是所有能力的真正实现——解析、比较、排序、分组、约束、范围查询都在这里。上层三种接口（CLI、Go SDK、MCP）只是对同一核心的薄封装，所以**无论从哪条路径进来，结果语义完全一致**。

## AI Agent 的两条路径

### 路径 1：Skills Plugin（Claude Code 专属）

Claude Code 的 Skills 机制会把 `skills/*/SKILL.md` 这 14 个文件作为**领域知识**注入。每个 SKILL.md 里写的是"什么时候用、怎么用、API 参考、代码示例、决策树"。

当你在 Claude Code 里输入 `/version-sorting`，发生的事是：

1. Claude Code 加载 `version-sorting` 这个 Skill 的全部知识。
2. Claude 据此理解：你要排序版本号，正确方式是按数值序而非字典序，预发布版要按权重排。
3. Claude 调用底层工具（CLI 或 MCP）执行，拿到结果。

**关键点**：Skill 本身不执行逻辑，它给 AI 的是**判断力**——什么时候用哪个工具、怎么解读结果。执行仍走 CLI/MCP/SDK。这避免了 AI 靠"直觉"猜版本关系。

### 路径 2：MCP Server（任何 MCP 客户端）

MCP Server（`cmd/versions-mcp`）把核心库的 21 个能力暴露成标准 MCP 工具，任何 MCP 兼容客户端都能直接调用：

- Claude Code、Codex、Cursor、Windsurf、Cline、VS Code Copilot ……
- AI 不需要"懂"版本号规则，它只需决定"调用 `version_sort` 工具"，工具返回的就是确定性的正确答案。

### 两条路径配合使用

:::mermaid
flowchart LR
  Q{"你的 AI Agent 是?"} -->|"Claude Code"| BOTH["装 Skills + MCP<br/>（知识 + 执行）"]
  Q -->|"Codex / Cursor / 其它"| MCPONLY["只装 MCP<br/>（确定性执行足够）"]
  BOTH --> BEST["✅ 最佳体验"]
  MCPONLY --> GOOD["✅ 已够用"]
  style BOTH fill:#eff6ff,stroke:#2563eb
  style BEST fill:#f0fdf4,stroke:#16a34a,stroke-width:3px
  style GOOD fill:#f0fdf4,stroke:#16a34a
:::

::: tip 推荐做法
在 Claude Code 上**同时**装 Skills 插件和 MCP Server：
- **Skills** 给 Claude 版本号领域知识（何时用约束、后缀如何排序）。
- **MCP** 给 Claude 快速、结构化的执行通道。

Skills 管"怎么想"，MCP 管"怎么做"。其它 Agent（Codex 等）只有 MCP 一条路径，但已经足够——工具的确定性结果就是 AI 需要的全部。
:::

## 数据流

一个版本号字符串从输入到输出，经过这条管道：

:::mermaid
flowchart LR
  IN["原始字符串<br/>'v1.2.3-beta1'"] --> IF["接口层<br/>CLI / MCP / SDK / Skill"]
  IF --> P["解析 parser.go<br/>Prefix / Numbers / Suffix / Metadata"]
  P --> PROC["处理<br/>compare · sort · constraint · range"]
  PROC --> OUT["输出<br/>结构化 JSON / 文本"]
  style IN fill:#eff6ff,stroke:#2563eb
  style P fill:#fff7ed,stroke:#ea580c
  style PROC fill:#fff7ed,stroke:#ea580c
  style OUT fill:#f0fdf4,stroke:#16a34a
:::

1. **输入**：字符串从 CLI 参数、MCP 调用参数、SDK 调用、或 Skill 引导进入。
2. **解析**（`parser.go`）：拆成 Prefix / VersionNumbers / Suffix / Metadata 四字段。
3. **处理**（`version.go` / `sort.go` / `constraint.go` / `version_range.go`）：比较、排序、约束检查、范围查询。
4. **输出**：结构化 JSON（MCP/SDK）或人类可读文本（CLI）。

## 为什么核心库零依赖

核心库只依赖三个 `golang-infrastructure` 的小工具包（`go-tuple`、`go-shuffle`、`go-compare-anything`），实质上等价于标准库。这意味着：

- **可复现**：构建行为不随上游版本漂移。
- **可嵌入**：作为库引入任何 Go 项目不会带来依赖冲突。
- **快**：解析 O(n)、比较 O(m)、排序 O(n log n)、范围查询经有序索引 O(组数)。

→ 想看每一步的精确规则，进 [算法详解](./algorithms)。想直接上手，进 [一键提示词](./prompts)。
