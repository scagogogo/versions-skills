# 按主版本号分组

::: tip 问题
把版本按 Major 号分组。
:::

## 🛠 SDK (Go)

```go
groups := versions.GroupByMajor(vs) // map[int][]*Version
```

## 🐚 CLI

```bash
versions group 1.0.0 1.1.0 2.0.0
```

## 🤖 MCP

```json
{
  "tool": "version_group",
  "arguments": {
    "versions": [
      "1.0.0",
      "1.1.0",
      "2.0.0"
    ]
  }
}
```

## 📚 参考

- [group-by-major](/sdk/api/group-by-major)
- [group](/cli/commands/group)

---

[← 返回配方索引](/recipes/)
