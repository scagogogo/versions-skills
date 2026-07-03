# version_validate

::: info MCP 工具
```
version_validate
```
:::

## 📖 描述

验证版本字符串是否有效。有效版本须包含数字部分且数字非负。

## 📥 参数

| 参数 | 类型 | 必填 | 说明 |
|:--|:--|:--|:--|
| `version_string` | `string` | ✅ 是 | 要验证的版本字符串 |

## 🔌 调用示例

```json
{
  "tool": "version_validate",
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
