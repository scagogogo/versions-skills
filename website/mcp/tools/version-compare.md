# version_compare

::: info MCP 工具
```
version_compare
```
:::

## 📖 描述

比较两个版本号的大小关系。返回 -1(旧于)/0(相等)/1(新于)。

## 📥 参数

| 参数 | 类型 | 必填 | 说明 |
|:--|:--|:--|:--|
| `version1` | `string` | ✅ 是 | 第一个版本字符串 |
| `version2` | `string` | ✅ 是 | 第二个版本字符串 |

## 🔌 调用示例

```json
{
  "tool": "version_compare",
  "arguments": {
    "version1": "1.2.3",
    "version2": "1.2.3"
  }
}
```

## 📚 相关

- [MCP 总览](/mcp/) · [SDK API](/sdk/) · [CLI 命令](/cli/)

---

::: details 源码
定义于 `internal/mcp/tools.go`，处理逻辑在 `internal/mcp/handlers.go`
:::
