# version_build

::: info MCP 工具
```
version_build
```
:::

## 📖 描述

从各组成部分构建版本字符串。

## 📥 参数

| 参数 | 类型 | 必填 | 说明 |
|:--|:--|:--|:--|
| `prefix` | `string` | 否 | 版本前缀（如 'v'） |
| `suffix` | `string` | 否 | 版本后缀（如 '-beta1'） |
| `major` | `number` | 否 | Major 版本号 |
| `minor` | `number` | 否 | Minor 版本号 |
| `patch` | `number` | 否 | Patch 版本号 |

## 🔌 调用示例

```json
{
  "tool": "version_build",
  "arguments": {
    "prefix": "v",
    "suffix": "-beta1",
    "major": 1,
    "minor": 2,
    "patch": 3
  }
}
```

## 📤 返回示例

```json
{
  "built_version": "v1.2.3-beta1"
}
```

## 📚 相关

- [MCP 总览](/mcp/) · [SDK API](/sdk/) · [CLI 命令](/cli/)

---

::: details 源码
定义于 `internal/mcp/tools.go`，处理逻辑在 `internal/mcp/handlers.go`
:::
