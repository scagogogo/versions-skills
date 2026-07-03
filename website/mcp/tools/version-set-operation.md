# version_set_operation

::: info MCP 工具
```
version_set_operation
```
:::

## 📖 描述

对两个版本集合执行集合运算：差集(difference)、交集(intersection)、并集(union)。

## 📥 参数

| 参数 | 类型 | 必填 | 说明 |
|:--|:--|:--|:--|
| `operation` | `string` | ✅ 是 | 集合运算类型: difference|intersection|union |
| `set_a` | `array` | ✅ 是 | 集合 A 的版本字符串数组 |
| `set_b` | `array` | ✅ 是 | 集合 B 的版本字符串数组 |

## 🔌 调用示例

```json
{
  "tool": "version_set_operation",
  "arguments": {
    "operation": "",
    "set_a": [],
    "set_b": []
  }
}
```

## 📚 相关

- [MCP 总览](/mcp/) · [SDK API](/sdk/) · [CLI 命令](/cli/)

---

::: details 源码
定义于 `internal/mcp/tools.go`，处理逻辑在 `internal/mcp/handlers.go`
:::
