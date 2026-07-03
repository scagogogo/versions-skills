# CLI 命令总览

`versions` CLI 基于 [cobra](https://github.com/spf13/cobra) 构建，提供 44 个子命令，覆盖版本号的**解析、比较、排序、分组、过滤、约束、范围查询、不可变变更、文件 IO 与可视化**。

## 📦 安装

::: tip 前置条件
需要 Go ≥ 1.21。
:::

```bash
# 从源码构建
go install github.com/scagogogo/versions-skills/cmd/versions@latest

# 或在仓库内
go build -o versions ./cmd/versions
```

## 🗂 命令分组

### 🔍 解析与信息

| 命令 | 说明 |
|:--|:--|
| [`parse`](/cli/commands/parse) | 解析版本字符串，显示各组成部分 |
| [`info`](/cli/commands/info) | 显示版本号的完整信息 |
| [`validate`](/cli/commands/validate) | 验证版本字符串是否有效 |
| [`segments`](/cli/commands/segments) | 获取数字段列表 |
| [`sub-version`](/cli/commands/sub-version) | 获取后缀中的子版本号 |
| [`suffix-weight`](/cli/commands/suffix-weight) | 获取后缀语义权重 |
| [`pure-prefix`](/cli/commands/pure-prefix) | 获取纯净前缀 |
| [`group-id`](/cli/commands/group-id) | 获取分组 ID |
| [`core`](/cli/commands/core) | 获取核心版本号（去后缀） |
| [`clone`](/cli/commands/clone) | 克隆版本号 |

### ⚖️ 比较与检查

| 命令 | 说明 |
|:--|:--|
| [`compare`](/cli/commands/compare) | 比较两个版本号 |
| [`check`](/cli/commands/check) | 检查版本号属性（布尔，exit code 0=真/1=假） |

### 📊 排序与极值

| 命令 | 说明 |
|:--|:--|
| [`sort`](/cli/commands/sort) | 对版本号排序 |
| [`sort-strings`](/cli/commands/sort-strings) | 对原始字符串排序 |
| [`min`](/cli/commands/min) | 查找最小版本 |
| [`max`](/cli/commands/max) | 查找最大版本 |
| [`latest-stable`](/cli/commands/latest-stable) | 最新稳定版本 |
| [`latest-prerelease`](/cli/commands/latest-prerelease) | 最新预发布版本 |

### 🗃 分组与过滤

| 命令 | 说明 |
|:--|:--|
| [`group`](/cli/commands/group) | 按数字部分分组 |
| [`group-ids`](/cli/commands/group-ids) | 列出分组 ID |
| [`group-latest`](/cli/commands/group-latest) | 分组最新版本 |
| [`group-oldest`](/cli/commands/group-oldest) | 分组最旧版本 |
| [`group-stable`](/cli/commands/group-stable) | 分组稳定版本 |
| [`group-prerelease`](/cli/commands/group-prerelease) | 分组预发布版本 |
| [`group-latest-stable`](/cli/commands/group-latest-stable) | 分组最新稳定版本 |
| [`group-latest-prerelease`](/cli/commands/group-latest-prerelease) | 分组最新预发布版本 |
| [`group-contains`](/cli/commands/group-contains) | 检查分组是否包含某版本 |
| [`filter`](/cli/commands/filter) | 按条件过滤 |
| [`partition`](/cli/commands/partition) | 分为两组 |
| [`count`](/cli/commands/count) | 统计满足条件的数量 |

### 🎯 约束与范围

| 命令 | 说明 |
|:--|:--|
| [`constraint`](/cli/commands/constraint) | 检查是否满足约束表达式 |
| [`satisfies`](/cli/commands/satisfies) | 以版本为中心检查约束 |
| [`range`](/cli/commands/range) | 查询范围内的版本 |

### 🛠 变更与构造

| 命令 | 说明 |
|:--|:--|
| [`build`](/cli/commands/build) | 构建版本号字符串 |
| [`bump`](/cli/commands/bump) | 递增版本号 |
| [`set-prefix`](/cli/commands/set-prefix) | 修改前缀 |
| [`set-suffix`](/cli/commands/set-suffix) | 修改后缀 |
| [`set-major`](/cli/commands/set-major) | 修改 Major |
| [`set-minor`](/cli/commands/set-minor) | 修改 Minor |
| [`set-patch`](/cli/commands/set-patch) | 修改 Patch |
| [`set-numbers`](/cli/commands/set-numbers) | 修改数字部分 |

### 📁 文件与可视化

| 命令 | 说明 |
|:--|:--|
| [`read`](/cli/commands/read) | 从文件读取版本号 |
| [`write`](/cli/commands/write) | 写入版本号到文件（自动排序） |
| [`read-strings`](/cli/commands/read-strings) | 读取原始字符串（不解析） |
| [`visualize`](/cli/commands/visualize) | 可视化层级树 |

## 🧰 通用 flag

多数接受 `[version-strings...]` 的命令支持：

| Flag | 说明 |
|:--|:--|
| `--from-file <path>` | 从文件读取版本号列表（每行一个） |
| `--help` / `-h` | 查看命令帮助 |

## 📚 延伸

- [Go SDK API](/sdk/) — 每个 CLI 命令背后都有对应的 SDK 函数
- [MCP 工具](/mcp/) — 同样能力暴露给 AI Agent
- [一键提示词](/prompts) — 让 AI 直接调用本 CLI
