# version_min

::: info MCP 工具
```
version_min
```
:::

## 📖 描述

从版本列表中查找最小（最旧）版本。

## 📥 参数

| 参数 | 类型 | 必填 | 说明 |
|:--|:--|:--|:--|
| `versions` | `array` | ✅ 是 | 版本字符串数组 |

## 🔌 调用示例

```json
{
  "tool": "version_min",
  "arguments": {
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
