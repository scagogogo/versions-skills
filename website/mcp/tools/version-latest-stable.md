# version_latest_stable

::: info MCP 工具
```
version_latest_stable
```
:::

## 📖 描述

从版本列表中查找最新的稳定版本。

## 📥 参数

| 参数 | 类型 | 必填 | 说明 |
|:--|:--|:--|:--|
| `versions` | `array` | ✅ 是 | 版本字符串数组 |

## 🔌 调用示例

```json
{
  "tool": "version_latest_stable",
  "arguments": {
    "versions": [
      "1.0.0-alpha",
      "1.0.0",
      "1.1.0-beta",
      "1.1.0",
      "2.0.0-rc1"
    ]
  }
}
```

## 📤 返回示例

```json
{
  "latest_stable": "1.1.0"
}
```

## 📚 相关

- [MCP 总览](/mcp/) · [SDK API](/sdk/) · [CLI 命令](/cli/)

---

::: details 源码
定义于 `internal/mcp/tools.go`，处理逻辑在 `internal/mcp/handlers.go`
:::
