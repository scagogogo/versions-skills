# version_bump

::: info MCP 工具
```
version_bump
```
:::

## 📖 描述

递增版本号的指定部分并清除后缀。major: 1.2.3→2.0.0; minor: 1.2.3→1.3.0; patch: 1.2.3→1.2.4。

## 📥 参数

| 参数 | 类型 | 必填 | 说明 |
|:--|:--|:--|:--|
| `version_string` | `string` | ✅ 是 | 原始版本字符串 |
| `bump_type` | `string` | ✅ 是 | 递增类型: major|minor|patch |

## 🔌 调用示例

```json
{
  "tool": "version_bump",
  "arguments": {
    "version_string": "1.2.3",
    "bump_type": "patch"
  }
}
```

## 📤 返回示例

```json
{
  "bumped_version": "1.2.4"
}
```

## 📚 相关

- [MCP 总览](/mcp/) · [SDK API](/sdk/) · [CLI 命令](/cli/)

---

::: details 源码
定义于 `internal/mcp/tools.go`，处理逻辑在 `internal/mcp/handlers.go`
:::
