# version_constraint_check

::: info MCP 工具
```
version_constraint_check
```
:::

## 📖 描述

检查版本号是否满足约束表达式。支持运算符: =, !=, >, >=, <, <=, ^(caret), ~(tilde), x/X/*(通配符)。

## 📥 参数

| 参数 | 类型 | 必填 | 说明 |
|:--|:--|:--|:--|
| `expression` | `string` | ✅ 是 | 约束表达式，如 '>=1.0.0,<2.0.0' |
| `version` | `string` | ✅ 是 | 要检查的版本字符串 |
| `type` | `string` | 否 | 约束类型: set(逗号分隔AND,默认) 或 union(||分隔OR) |

## 🔌 调用示例

```json
{
  "tool": "version_constraint_check",
  "arguments": {
    "expression": "",
    "version": "1.2.3",
    "type": ""
  }
}
```

## 📚 相关

- [MCP 总览](/mcp/) · [SDK API](/sdk/) · [CLI 命令](/cli/)

---

::: details 源码
定义于 `internal/mcp/tools.go`，处理逻辑在 `internal/mcp/handlers.go`
:::
