# 把版本列表分两组

::: tip 问题
按谓词把版本分为满足/不满足两组。
:::

## 🛠 SDK (Go)

```go
yes, no := versions.Partition(vs, func(v *versions.Version) bool { return v.IsStable() })
```

## 🐚 CLI

```bash
versions partition 1.0.0 1.1.0-rc1 --stable
```

## 🤖 MCP

```json
{
  "tool": "version_filter",
  "arguments": {
    "stable": true,
    "versions": [
      "1.0.0",
      "1.0.0-beta",
      "2.0.0"
    ]
  }
}
```

## 📚 参考

- [partition](/sdk/api/partition)
- [partition](/cli/commands/partition)

---

[← 返回配方索引](/recipes/)
