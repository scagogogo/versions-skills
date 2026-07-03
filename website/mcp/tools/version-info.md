# version_info

::: info MCP 工具
```
version_info
```
:::

## 📖 描述

获取版本号的完整信息，包括所有类型判断（IsPrerelease/IsStable/IsBeta 等）。

## 📥 参数

| 参数 | 类型 | 必填 | 说明 |
|:--|:--|:--|:--|
| `version_string` | `string` | ✅ 是 | 版本字符串 |

## 🔌 调用示例

```json
{
  "tool": "version_info",
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
