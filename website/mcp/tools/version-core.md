# version_core

::: info MCP 工具
```
version_core
```
:::

## 📖 描述

获取版本号的核心部分（去除预发布后缀）。如 v1.2.3-beta1 → v1.2.3。

## 📥 参数

| 参数 | 类型 | 必填 | 说明 |
|:--|:--|:--|:--|
| `version_string` | `string` | ✅ 是 | 版本字符串 |

## 🔌 调用示例

```json
{
  "tool": "version_core",
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
