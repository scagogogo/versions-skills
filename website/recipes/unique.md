# 版本列表去重

::: tip 问题
去除重复的版本。
:::

## 🛠 SDK (Go)

```go
versions.Unique(vs)
```

## 🐚 CLI

```bash
# CLI 无直接命令；sort 可配合
```

## 🤖 MCP

```json
{
  "tool": "version_unique",
  "arguments": {
    "versions": [
      "1.0.0",
      "1.0.0",
      "2.0.0"
    ]
  }
}
```

## 📚 参考

- [unique](/sdk/api/unique)

---

[← 返回配方索引](/recipes/)
