# version_read_file

::: info MCP 工具
```
version_read_file
```
:::

## 📖 描述

从文件中读取版本号列表（每行一个版本字符串）。

## 📥 参数

| 参数 | 类型 | 必填 | 说明 |
|:--|:--|:--|:--|
| `filepath` | `string` | ✅ 是 | 文件路径 |

## 🔌 调用示例

```json
{
  "tool": "version_read_file",
  "arguments": {
    "filepath": ""
  }
}
```

## 📚 相关

- [MCP 总览](/mcp/) · [SDK API](/sdk/) · [CLI 命令](/cli/)

---

::: details 源码
定义于 `internal/mcp/tools.go`，处理逻辑在 `internal/mcp/handlers.go`
:::
