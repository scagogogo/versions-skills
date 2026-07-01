# MCP 工具

`versions-mcp` 把核心库的 21 个能力暴露为 MCP 工具，任何 MCP 兼容客户端可调。安装与配置见 [AI Agent 接入指南](./ai-agents)。

## 工具清单

| 工具 | 作用 | 对应 SDK / CLI |
|:--|:--|:--|
| `version_parse` | 解析版本字符串，返回各组成部分 | `NewVersion` / `parse` |
| `version_validate` | 校验版本有效性 | `IsValid` / `validate` |
| `version_info` | 返回版本完整信息 | `info` |
| `version_compare` | 比较两个版本 | `CompareTo` / `compare` |
| `version_sort` | 排序版本列表 | `SortVersionSlice` / `sort` |
| `version_filter` | 按条件过滤 | `Filter` / `filter` |
| `version_group` | 分组 | `Group` / `group` |
| `version_range_query` | 范围查询 | `SortedVersionGroups.QueryRange` / `range` |
| `version_constraint_check` | 检查是否满足约束表达式 | `Matches` / `satisfies` |
| `version_min` | 最小版本 | `Min` / `min` |
| `version_max` | 最大版本 | `Max` / `max` |
| `version_latest_stable` | 最新稳定版 | `LatestStable` / `latest-stable` |
| `version_latest_prerelease` | 最新预发布版 | `LatestPrerelease` / `latest-prerelease` |
| `version_unique` | 去重 | `Unique` |
| `version_set_operation` | 集合运算（交/并/差） | `Intersection`/`Union`/`Difference` |
| `version_build` | 构造版本 | `VersionBuilder` / `build` |
| `version_bump` | 递增版本号 | `BumpMajor/Minor/Patch` / `bump` |
| `version_core` | 去后缀核心版本 | `Core` / `core` |
| `version_read_file` | 从文件读取版本列表 | `ReadVersionsFromFile` / `read` |
| `version_write_file` | 写入文件 | `WriteVersionsToFile` / `write` |
| `version_visualize` | 树形可视化 | `VisualizeVersions` / `visualize` |

## 传输方式

- **stdio**（默认）：本地单机使用，进程间通信。
- **SSE**：`versions-mcp --transport sse --port 8080`，团队共享部署，端点 `http://localhost:8080/sse`。

## 调用示例（概念）

MCP 工具由 AI Agent 自行调用，你通常不需要手写 JSON-RPC。但概念上一次 `version_sort` 调用形如：

```json
{
  "method": "tools/call",
  "params": {
    "name": "version_sort",
    "arguments": {
      "versions": ["2.0.0", "1.0.0", "1.10.0", "1.2.0-beta"],
      "order": "asc"
    }
  }
}
```

返回结构化 JSON：

```json
{
  "sorted": ["1.0.0", "1.2.0-beta", "1.10.0", "2.0.0"]
}
```

::: tip 让 AI 自动调
接入后直接用自然语言下指令即可，AI 会自动选对工具。任务模板见 [一键提示词](./prompts)。
:::

→ 所有工具的底层语义见 [算法详解](./algorithms)。
