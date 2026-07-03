# version_range_query

::: info MCP 工具
```
version_range_query
```
:::

## 📖 描述

查询指定范围内的版本号列表。

## 📥 参数

| 参数 | 类型 | 必填 | 说明 |
|:--|:--|:--|:--|
| `start` | `string` | ✅ 是 | 范围起始版本 |
| `end` | `string` | ✅ 是 | 范围结束版本 |
| `include_start` | `boolean` | 否 | 是否包含起始版本（默认 true） |
| `include_end` | `boolean` | 否 | 是否包含结束版本（默认 false） |
| `versions` | `array` | ✅ 是 | 待查询的版本字符串数组 |

## 🔌 调用示例

```json
{
  "tool": "version_range_query",
  "arguments": {
    "start": "1.2.3",
    "end": "1.2.3",
    "include_start": false,
    "include_end": false,
    "versions": []
  }
}
```

## 📚 相关

- [MCP 总览](/mcp/) · [SDK API](/sdk/) · [CLI 命令](/cli/)

---

::: details 源码
定义于 `internal/mcp/tools.go`，处理逻辑在 `internal/mcp/handlers.go`
:::
