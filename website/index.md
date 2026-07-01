---
# https://vitepress.dev/reference/default-theme-home-page
layout: home

hero:
  name: Versions-Skills
  text: 面向 AI 原生的版本号工具集
  tagline: 让 AI Agent 直接解析、比较、排序、约束检查版本号 —— Skills + MCP + Go SDK + CLI，开箱即用。
  image:
    src: /favicon.svg
    alt: Versions-Skills
  actions:
    - theme: brand
      text: 🚀 一键接入 Claude Code
      link: /prompts
    - theme: alt
      text: 为什么需要它
      link: /why
    - theme: alt
      text: 工作原理
      link: /how-it-works

features:
  - icon: 🤖
    title: AI 原生接入
    details: 同时提供 Claude Code Skills（13 个斜杠命令）与 MCP Server（21 个工具）。Codex、Cursor、Windsurf、Cline 等 MCP 客户端一行配置即可接入。
  - icon: 🧩
    title: 宽口径解析
    details: 标准 semver、带前缀（v1.2.3）、预发布（1.2.3-beta1）、Maven/Scala/PEP 440 风格全支持。还能从任意字符串里 Coerce 出版本号。
  - icon: 📊
    title: 语义化比较
    details: 不是字典序。按"数字段 → 后缀权重 → 发布时间 → 原始串"四级优先级比较，dev < alpha < beta < rc < 正式版，贴合真实发布阶梯。
  - icon: 📋
    title: npm 风格约束
    details: 完整支持 >=、^、~、1.x、|| 组合约束表达式。约束语法三层（Union OR / Set AND / Single），语义与 npm/pip 一致。
  - icon: 🌳
    title: 分组与可视化
    details: 按主/次版本号分组排序，1.10.0 正确排在 1.2.0 之后。Unicode 树形展示版本层次。
  - icon: 🚀
    title: 零依赖核心
    details: Go 核心库无外部重依赖，纯标准库可编译。CLI/MCP/Skills 三种交付形态共享同一引擎。
---

<div style="margin-top: 48px;">

## 🤖 一键接入：把提示词丢给 AI，等它装好

**AI-First 接入** —— 你不需要看文档、不需要改配置文件。把下面提示词复制给你的 AI Agent，它会自己安装、自己配置、自己验证，然后告诉你"已就绪"。

<div style="display:flex;gap:8px;flex-wrap:wrap;margin:12px 0 4px;">
  <span style="background:#eff6ff;color:#1e3a8a;border:1px solid #bfdbfe;padding:3px 10px;border-radius:4px;font-size:13px;">📋 整段复制</span>
  <span style="background:#f0fdf4;color:#14532d;border:1px solid #bbf7d0;padding:3px 10px;border-radius:4px;font-size:13px;">⚡ AI 自动配置</span>
  <span style="background:#fff7ed;color:#7c2d12;border:1px solid #fed7aa;padding:3px 10px;border-radius:4px;font-size:13px;">✅ 自动验证</span>
</div>

</div>

:::mermaid
flowchart LR
  COPY["📋 复制提示词"] --> PASTE["💬 粘贴到<br/>Claude Code / Codex"]
  PASTE --> AI["🤖 AI 自动执行<br/>安装 + 配置"]
  AI --> VERIFY["✅ 冒烟测试<br/>解析/比较一个版本"]
  VERIFY --> READY["🎉 回复'已就绪'<br/>你开始下任务"]
  style COPY fill:#eff6ff,stroke:#2563eb
  style AI fill:#fff7ed,stroke:#ea580c
  style READY fill:#f0fdf4,stroke:#16a34a,stroke-width:3px
:::

::: code-group

```markdown title="Claude Code"
我需要处理版本号，请帮我接入 versions-skills 工具。

按以下步骤执行，每步完成后简要汇报：

1. 安装 Claude Code 的 versions 插件（提供 13 个版本号斜杠命令）：
   运行：claude plugin marketplace add https://github.com/scagogogo/versions-skills
   运行：claude plugin install versions

2. 注册 versions 的 MCP Server（提供 21 个 version_* 工具供你直接调用）：
   运行：claude mcp add versions -- versions-mcp --transport stdio
   如果 versions-mcp 二进制不存在，先安装：
   go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest
   若 Go 未安装，从 https://github.com/scagogogo/versions-skills/releases/latest 下载对应平台的 versions-mcp 放到 PATH（如 /usr/local/bin）

3. 验证：
   - 确认插件已加载（输入 / 应能看到 /version-parsing、/version-sorting 等命令）
   - 用 version_parse 工具解析 "v1.2.3-beta1" 做冒烟测试，把结构化结果告诉我

完成后回复"已就绪"，接下来我会给你版本号任务。
```

```markdown title="Codex"
我需要处理版本号，请帮我接入 versions-skills 的 MCP Server。

按以下步骤执行：

1. 安装二进制：
   go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest
   若 Go 未安装，从 https://github.com/scagogogo/versions-skills/releases/latest 下载对应平台的 versions-mcp 放到 PATH（如 /usr/local/bin）

2. 注册到 Codex（写入 ~/.codex/config.toml）：
   运行：codex mcp add versions -- versions-mcp --transport stdio

3. 验证：用 version_sort 工具对 ["3.0.0","1.0.0","1.10.0","1.2.0-beta"] 排序，
   把结果告诉我（期望 1.2.0-beta 排在 1.10.0 之前、1.10.0 排在 1.2.0 之后）。

完成后回复"已就绪"，接下来我会给你版本号任务。
```

```markdown title="Cursor / Windsurf / VS Code / Cline"
我需要处理版本号，请帮我接入 versions-skills 的 MCP Server。

按以下步骤执行：

1. 安装二进制：
   go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest
   若 Go 未安装，从 https://github.com/scagogogo/versions-skills/releases/latest 下载对应平台的 versions-mcp 放到 PATH

2. 把以下配置写入我所用编辑器的 MCP 配置文件（请先确认我用的哪个编辑器，再选对应文件）：
   - Cursor → 项目根 .cursor/mcp.json
   - Windsurf → 项目根 .windsurf/mcp.json
   - VS Code Copilot → 项目根 .vscode/mcp.json（注意字段名是 servers）
   - Cline → ~/.cline/cline_mcp_settings.json

   配置内容（VS Code Copilot 用 "servers"，其余用 "mcpServers"）：
   {
     "mcpServers": {
       "versions": {
         "command": "versions-mcp",
         "args": ["--transport", "stdio"]
       }
     }
   }

3. 验证：用 version_compare 工具比较 "1.0.0" 和 "2.0.0"，把结果告诉我。

完成后回复"已就绪"，接下来我会给你版本号任务。
```

:::

> 💡 **怎么用**：点上方标签选你的 AI Agent → 点代码块右上角复制按钮 → 粘贴到 Agent 输入框回车。AI 会按提示词里的步骤自动装好并验证，全程无需你手动改任何配置文件。更多提示词与任务模板见 [一键提示词](./prompts)。

<div style="margin-top: 48px;">

## 它解决了什么问题

版本号看起来像数字，却是字符串。`"1.10.0"` 在字典序里排在 `"1.2.0"` **前面**；`"1.0.0-rc1"` 和 `"1.0.0"` 谁先谁后、`"^1.2.3"` 到底允许哪些版本——这些在依赖管理、发布流水线、安全补丁筛选里反复出现，手写规则必踩坑。

**versions-skills 把版本号的语义内置成一套可被 AI 直接调用的能力**，让 Agent 像调用函数一样精确处理版本，而不是靠正则和 if-else 猜。

→ 看 [为什么需要它](./why) 了解完整问题面，看 [工作原理](./how-it-works) 了解四层架构。

</div>
