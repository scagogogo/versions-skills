# version_sort

::: info MCP 工具
```
version_sort
```
:::

## 📖 描述

对版本号列表进行排序。默认升序，可设置降序。

## 📥 参数

| 参数 | 类型 | 必填 | 说明 |
|:--|:--|:--|:--|
| `descending` | `boolean` | 否 | 是否降序排列（默认 false） |
| `versions` | `array` | ✅ 是 | 版本字符串数组 |

## 🔌 调用示例

```json
{
  "tool": "version_sort",
  "arguments": {
    "descending": false,
    "versions": [
      "2.0.0",
      "1.0.0",
      "1.10.0",
      "1.2.0-beta"
    ]
  }
}
```

## 📤 返回示例

```json
{
  "sorted_versions": [
    "1.2.0-beta",
    "1.0.0",
    "1.10.0",
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
