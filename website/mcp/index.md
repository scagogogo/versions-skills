# MCP 工具总览

versions-skills 通过 [MCP](https://modelcontextprotocol.io/) 把全部能力暴露给 AI Agent，共 **21 个工具**。MCP server 入口为 `cmd/versions-mcp`，基于 [mcp-go](https://github.com/mark3labs/mcp-go)。

## 🚀 启动 server

```bash
# 从源码构建
go build -o versions-mcp ./cmd/versions-mcp

# 以 stdio 模式启动（默认）
./versions-mcp
```

## 🔧 接入 Claude Code / Claude Desktop

在 MCP 配置中添加：

```json
{
  "mcpServers": {
    "versions": {
      "command": "/path/to/versions-mcp",
      "args": []
    }
  }
}
```

接入后，AI 可调用下方任一工具完成版本号操作，无需用户手敲命令。

## 🧰 工具清单（21 个）

### 🔍 解析与信息

| 工具 | 说明 |
|:--|:--|
| [`version_parse`](/mcp/tools/version-parse) | 解析版本字符串为结构化组件 |
| [`version_validate`](/mcp/tools/version-validate) | 验证版本字符串是否有效 |
| [`version_info`](/mcp/tools/version-info) | 获取版本号完整信息 |
| [`version_core`](/mcp/tools/version-core) | 获取核心版本号（去后缀） |

### ⚖️ 比较

| 工具 | 说明 |
|:--|:--|
| [`version_compare`](/mcp/tools/version-compare) | 比较两个版本号 |

### 📊 排序与极值

| 工具 | 说明 |
|:--|:--|
| [`version_sort`](/mcp/tools/version-sort) | 排序版本号列表 |
| [`version_min`](/mcp/tools/version-min) | 查找最小版本 |
| [`version_max`](/mcp/tools/version-max) | 查找最大版本 |
| [`version_latest_stable`](/mcp/tools/version-latest-stable) | 最新稳定版本 |
| [`version_latest_prerelease`](/mcp/tools/version-latest-prerelease) | 最新预发布版本 |

### 🗃 分组与过滤

| 工具 | 说明 |
|:--|:--|
| [`version_group`](/mcp/tools/version-group) | 按数字部分分组 |
| [`version_filter`](/mcp/tools/version-filter) | 按条件过滤 |
| [`version_unique`](/mcp/tools/version-unique) | 去重 |
| [`version_set_operation`](/mcp/tools/version-set-operation) | 集合运算（差/交/并） |

### 🎯 约束与范围

| 工具 | 说明 |
|:--|:--|
| [`version_constraint_check`](/mcp/tools/version-constraint-check) | 检查是否满足约束 |
| [`version_range_query`](/mcp/tools/version-range-query) | 范围查询 |

### 🛠 变更与构造

| 工具 | 说明 |
|:--|:--|
| [`version_build`](/mcp/tools/version-build) | 从组件构建版本字符串 |
| [`version_bump`](/mcp/tools/version-bump) | 递增版本号 |

### 📁 文件与可视化

| 工具 | 说明 |
|:--|:--|
| [`version_read_file`](/mcp/tools/version-read-file) | 从文件读取版本号 |
| [`version_write_file`](/mcp/tools/version-write-file) | 写入版本号到文件 |
| [`version_visualize`](/mcp/tools/version-visualize) | 文本树状可视化 |

## 📐 参数约定

- 版本号一律以 `string` 传入（如 `"1.2.3"`、`"v1.2.3-rc1"`）
- 版本号列表以 `string[]` 传入
- 布尔 flag 默认 `false`
- 所有工具返回 JSON

## 📚 延伸

- [SDK API](/sdk/) · [CLI 命令](/cli/) · [Skills 斜杠命令](/skills/)
- [AI Agent 接入指南](/ai-agents) · [一键提示词](/prompts)
