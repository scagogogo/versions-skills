# Skills 斜杠命令

Claude Code 专属：14 个 Skills 作为领域知识注入，输入 `/` 即可触发。每个 Skill 是仓库 `skills/*/SKILL.md` 里的一个文档，告诉 Claude **何时用、怎么用**，再由 Claude 调用底层 CLI/MCP/SDK 执行。

## 安装

```bash
claude plugin marketplace add https://github.com/scagogogo/versions-skills
claude plugin install versions
```

## 命令清单

| 命令 | 作用 |
|:--|:--|
| `/version-parsing` | 解析、验证、提取版本号组件 |
| `/version-comparison` | 比较版本号，检查排序关系 |
| `/version-sorting` | 升序/降序排序版本号列表 |
| `/version-grouping` | 按主/次版本号分组 |
| `/version-constraints` | 解析和检查约束表达式 |
| `/version-range-query` | 查询范围内的版本号 |
| `/version-visualization` | 树形版本层次结构展示 |
| `/version-file-operations` | 读写版本号列表文件 |
| `/version-check` | 布尔类型检查（IsBeta、IsStable 等） |
| `/version-mutation` | 版本号 Bump，不可变修改 |
| `/version-properties` | 访问版本号段落、后缀权重、前缀 |
| `/cli-operations` | 完整 CLI 命令参考 |
| `/mcp-operations` | MCP 服务器设置与工具参考 |
| `/installation` | 安装与接入引导 |

## 工作原理

Skills **不执行逻辑**，它注入的是判断力。当你在 Claude Code 里输入 `/version-sorting`：

1. Claude 加载 `version-sorting` Skill 的全部知识（API 参考、代码示例、决策树）。
2. Claude 据此理解：排序要按数值序而非字典序，预发布版按权重排。
3. Claude 调用底层工具（CLI 或 MCP）执行，拿到确定性结果。

这避免了 AI 靠"直觉"猜版本关系——比如 `1.10.0 < 1.2.0` 这种字符串直觉错误。

## 与 MCP 配合

::: tip 推荐
同时安装 Skills 插件和 MCP Server（见 [AI Agent 接入指南](./ai-agents)）。Skills 管"怎么想"，MCP 管"怎么做"。
:::

→ 不想手动装？用 [一键提示词](./prompts) 让 AI 自己配置。
