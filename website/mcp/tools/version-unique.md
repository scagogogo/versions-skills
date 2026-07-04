# version_unique

::: info MCP 工具
```
version_unique
```
:::

## 📖 描述

去除版本列表中的重复项。

## 📥 参数

| 参数 | 类型 | 必填 | 说明 |
|:--|:--|:--|:--|
| `versions` | `array` | ✅ 是 | 版本字符串数组 |

## 🔌 调用示例

```json
{
  "tool": "version_unique",
  "arguments": {
    "versions": [
      "1.0.0",
      "1.0.0",
      "1.1.0",
      "1.1.0",
      "2.0.0"
    ]
  }
}
```

## 📤 返回示例

```json
{
  "unique_versions": [
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
