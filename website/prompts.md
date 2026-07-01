# 一键提示词

> 把下面任一提示词**整段复制**，粘贴到对应 AI Agent 的输入框，回车。AI 会自动完成安装、配置、验证——你只需描述需求。

## Claude Code

### 完整接入（推荐）

一次性同时安装 Skills 插件与 MCP Server，接入后即可用 13 个斜杠命令 + 21 个 MCP 工具。

```markdown
请帮我接入 versions-skills 工具，用于版本号的解析、比较、排序、约束检查。

按以下步骤执行，每步完成后简要汇报：

1. 安装 Claude Code 的 versions 插件（提供 13 个版本号斜杠命令）：
   - 运行：claude plugin marketplace add https://github.com/scagogogo/versions-skills
   - 运行：claude plugin install versions

2. 注册 versions 的 MCP Server（提供 21 个 version_* 工具供你直接调用）：
   - 运行：claude mcp add versions -- versions-mcp --transport stdio
   - 如果 versions-mcp 二进制不存在，先安装：
     go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest
   - 若 Go 未安装，从 https://github.com/scagogogo/versions-skills/releases/latest
     下载对应平台的 versions-mcp 二进制，放到 PATH 中（如 /usr/local/bin）

3. 验证：
   - 确认插件已加载（输入 / 应能看到 /version-parsing、/version-sorting 等命令）
   - 用 version_parse 工具解析 "v1.2.3-beta1" 做冒烟测试，把结构化结果告诉我

完成后回复"已就绪"，接下来我会给你版本号任务。
```

### 只要 MCP（不装插件）

如果你只想要工具调用、不需要斜杠命令：

```markdown
请帮我接入 versions-skills 的 MCP Server：

1. 运行：claude mcp add versions -- versions-mcp --transport stdio
   - 若 versions-mcp 不存在，先安装：go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest
   - 若 Go 未安装，从 https://github.com/scagogogo/versions-skills/releases/latest 下载 versions-mcp 放到 PATH

2. 用 version_compare 工具比较 "1.0.0" 和 "2.0.0" 做冒烟测试，告诉我结果。
完成后回复"已就绪"。
```

## Codex（OpenAI CLI）

Codex 通过 MCP 接入，配置写入 `~/.codex/config.toml`。

```markdown
请帮我接入 versions-skills 的 MCP Server，用于版本号处理。

按以下步骤执行：

1. 安装二进制：go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest
   - 若 Go 未安装，从 https://github.com/scagogogo/versions-skills/releases/latest
     下载对应平台的 versions-mcp 放到 PATH（如 /usr/local/bin）

2. 注册到 Codex（写入 ~/.codex/config.toml）：
   - 运行：codex mcp add versions -- versions-mcp --transport stdio

3. 验证：用 version_sort 工具对 ["3.0.0","1.0.0","1.10.0","1.2.0-beta"] 排序，
   把结果告诉我（期望 1.2.0-beta 排在 1.10.0 之前、1.10.0 排在 1.2.0 之后）。

完成后回复"已就绪"，接下来我会给你版本号任务。
```

## 其它 MCP 客户端

Cursor / Windsurf / VS Code Copilot / Cline 的提示词结构相同，只是配置文件位置不同。把下方对应你客户端的 JSON 写进指定文件，然后让 AI 调用 `version_*` 工具即可。

::: details 配置文件位置
| 客户端 | 配置文件 |
|:--|:--|
| Cursor | `.cursor/mcp.json` |
| Windsurf | `.windsurf/mcp.json` |
| VS Code Copilot | `.vscode/mcp.json`（字段名是 `servers`） |
| Cline | `~/.cline/cline_mcp_settings.json` |

各文件内容见 [AI Agent 接入指南](./ai-agents)。
:::

## 团队共享部署（SSE）

把 MCP Server 部署为网络服务，团队成员共用一个实例：

```markdown
请帮我部署 versions-skills 的 MCP Server 为 SSE 模式：

1. 安装：go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest
2. 启动：versions-mcp --transport sse --port 8080
3. 告诉我 SSE 端点地址（默认 http://localhost:8080/sse），以便我分发给团队成员配置。
```

客户端侧把 MCP 配置指向该 SSE 端点（用客户端的 HTTP/SSE 传输选项，而非 command）。

---

## 接入后的任务模板

接入完成后，下面这些提示词可以直接用，AI 会自动调用合适的工具：

::: code-group

```markdown title="排序一批版本号"
我有这些版本号，请用 versions-skills 升序排序，并标注哪些是预发布版：
2.0.0, 1.0.0, 1.10.0, 1.2.0-beta, 1.2.0, 1.0.0-rc1, 0.9.0+121-bcc5decc
```

```markdown title="检查版本是否满足约束"
判断版本 1.5.0 是否满足约束 ">=1.0.0,<2.0.0 || >=3.0.0"，
并用版本号比较的原理解释为什么。
```

```markdown title="找最新稳定版"
从这份 versions.txt 里读取所有版本，找出最新的稳定版（不带后缀的），
再找出最新的预发布版，告诉我两者差距（major/minor/patch diff）。
```

```markdown title="范围查询"
在 [1.0.0, 2.0.0) 范围内（左闭右开）有哪些版本？
请用范围查询工具，并说明 2.0.0 为何被排除。
```

:::

→ 提示词背后的能力详见 [AI Agent 接入指南](./ai-agents) 与 [MCP 工具](./mcp)。
