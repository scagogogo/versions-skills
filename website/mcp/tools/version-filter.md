# version_filter

::: info MCP 工具
```
version_filter
```
:::

## 📖 描述

按条件过滤版本号列表。多个条件为 AND 逻辑。

## 📥 参数

| 参数 | 类型 | 必填 | 说明 |
|:--|:--|:--|:--|
| `prefix` | `string` | 否 | 按前缀过滤 |
| `suffix` | `string` | 否 | 按后缀过滤 |
| `constraint` | `string` | 否 | 按约束表达式过滤（ConstraintSet 语法） |
| `major` | `number` | 否 | 按 Major 版本号过滤 |
| `minor` | `number` | 否 | 按 Minor 版本号过滤 |
| `patch` | `number` | 否 | 按 Patch 版本号过滤 |
| `stable` | `boolean` | 否 | 仅保留稳定版本 |
| `prerelease` | `boolean` | 否 | 仅保留预发布版本 |
| `versions` | `array` | ✅ 是 | 版本字符串数组 |

## 🔌 调用示例

```json
{
  "tool": "version_filter",
  "arguments": {
    "stable": true,
    "versions": [
      "1.0.0",
      "1.0.0-rc1",
      "1.1.0",
      "1.1.0-beta",
      "2.0.0"
    ]
  }
}
```

## 📤 返回示例

```json
{
  "filtered_versions": [
    "1.0.0",
    "1.1.0",
    "2.0.0"
  ]
}
```

## 📚 相关

- [MCP 总览](/mcp/) · [SDK API](/sdk/) · [CLI 命令](/cli/)

---

::: details 源码
定义于 `internal/mcp/tools.go`，处理逻辑在 `internal/mcp/handlers.go`
:::
