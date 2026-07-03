# version_write_file

::: info MCP 工具
```
version_write_file
```
:::

## 📖 描述

将版本号列表排序后写入文件（每行一个版本字符串）。

## 📥 参数

| 参数 | 类型 | 必填 | 说明 |
|:--|:--|:--|:--|
| `filepath` | `string` | ✅ 是 | 文件路径 |
| `versions` | `array` | ✅ 是 | 版本字符串数组 |

## 🔌 调用示例

```json
{
  "tool": "version_write_file",
  "arguments": {
    "filepath": "",
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
