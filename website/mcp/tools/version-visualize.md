# version_visualize

::: info MCP 工具
```
version_visualize
```
:::

## 📖 描述

生成版本层级的文本可视化树状图。

## 📥 参数

| 参数 | 类型 | 必填 | 说明 |
|:--|:--|:--|:--|
| `max_items_per_group` | `number` | 否 | 每组最多显示的版本数（0=不限，默认 0） |
| `groups_only` | `boolean` | 否 | 仅显示分组摘要 |
| `versions` | `array` | ✅ 是 | 版本字符串数组 |

## 🔌 调用示例

```json
{
  "tool": "version_visualize",
  "arguments": {
    "max_items_per_group": 3,
    "groups_only": false,
    "versions": [
      "1.0.0",
      "1.0.1",
      "1.1.0",
      "2.0.0-rc1",
      "2.0.0"
    ]
  }
}
```

## 📤 返回示例

```json
{
  "visualization": "1
├─0
│ ├─0
│ └─1
└─1
2
├─0-rc1
└─0"
}
```

## 📚 相关

- [MCP 总览](/mcp/) · [SDK API](/sdk/) · [CLI 命令](/cli/)

---

::: details 源码
定义于 `internal/mcp/tools.go`，处理逻辑在 `internal/mcp/handlers.go`
:::
