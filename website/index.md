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
      text: 🚀 一键接入 AI Agent
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
    details: 同时提供 Claude Code Skills（14 个斜杠命令）与 MCP Server（21 个工具）。Codex、Cursor、Windsurf、Cline 等 MCP 客户端一行配置即可接入。
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
    details: 按版本号数字部分分组排序，1.10.0 正确排在 1.2.0 之后。Unicode 树形展示版本层次。
  - icon: 🚀
    title: 零依赖核心
    details: Go 核心库无外部重依赖，纯标准库可编译。CLI/MCP/Skills 三种交付形态共享同一引擎。
---

## 它解决了什么问题

版本号看起来像数字，却是字符串。`"1.10.0"` 在字典序里排在 `"1.2.0"` **前面**；`"1.0.0-rc1"` 和 `"1.0.0"` 谁先谁后、`"^1.2.3"` 到底允许哪些版本——这些在依赖管理、发布流水线、安全补丁筛选里反复出现，手写规则必踩坑。

**versions-skills 把版本号的语义内置成一套可被 AI 直接调用的能力**，让 Agent 像调用函数一样精确处理版本，而不是靠正则和 if-else 猜。

→ 看 [为什么需要它](./why) 了解完整问题面，看 [工作原理](./how-it-works) 了解四层架构。

## 怎么开始

::: tip 三步自动完成
📋 复制提示词 → ⚡ AI 自动配置 → ✅ 冒烟验证
:::

```mermaid
flowchart LR
  COPY["📋 复制提示词"] --> PASTE["💬 粘贴到<br/>Claude Code / Codex"]
  PASTE --> AI["🤖 AI 自动执行<br/>安装 + 配置"]
  AI --> VERIFY["✅ 冒烟测试<br/>解析/比较一个版本"]
  VERIFY --> READY["🎉 回复'已就绪'<br/>你开始下任务"]
  style COPY fill:#eff6ff,stroke:#2563eb
  style AI fill:#fff7ed,stroke:#ea580c
  style READY fill:#f0fdf4,stroke:#16a34a,stroke-width:3px
```

不需要看文档、不需要改配置文件。到 [**一键提示词**](./prompts) 选你的 AI Agent（Claude Code / Codex / Cursor / Windsurf / VS Code / Cline），把对应提示词整段复制粘贴给 AI，它会自己安装、自己配置、自己验证，然后告诉你"已就绪"。

想先手动跑一下？看 [快速开始](./quick-start)。
