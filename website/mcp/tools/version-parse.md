# version_parse

::: info MCP 工具
```
version_parse
```
:::

## 📖 描述

解析版本字符串为结构化组件（前缀、数字、后缀等）。支持标准 semver 和多种变体如 'v1.2.3-beta1'。

## 📥 参数

| 参数 | 类型 | 必填 | 说明 |
|:--|:--|:--|:--|
| `version_string` | `string` | ✅ 是 | 要解析的版本字符串，如 '1.2.3', 'v1.2.3-beta1' |

## 🔌 调用示例

```json
{
  "tool": "version_parse",
  "arguments": {
    "version_string": "1.2.3"
  }
}
```

## 📚 相关

- [MCP 总览](/mcp/) · [SDK API](/sdk/) · [CLI 命令](/cli/)

---

::: details 源码
定义于 `internal/mcp/tools.go`，处理逻辑在 `internal/mcp/handlers.go`
:::
