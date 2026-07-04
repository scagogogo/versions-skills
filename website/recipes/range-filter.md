# 查询区间内版本

::: tip 问题
筛出 `[1.0.0, 2.0.0]` 内的版本。
:::

## 🛠 SDK (Go)

```go
r := versions.NewClosedRange(versions.NewVersion("1.0.0"), versions.NewVersion("2.0.0"))
r.Filter(vs)
```

## 🐚 CLI

```bash
versions range 1.0.0 2.0.0 0.9.0 1.5.0 2.1.0
```

## 🤖 MCP

```json
{
  "tool": "version_range_query",
  "arguments": {
    "start": "1.0.0",
    "end": "2.0.0",
    "include_start": true,
    "include_end": true,
    "versions": [
      "0.9.0",
      "1.0.0",
      "1.5.0",
      "2.0.0",
      "2.1.0"
    ]
  }
}
```

## 📚 参考

- [version-range](/sdk/api/version-range)
- [range](/cli/commands/range)

---

[← 返回配方索引](/recipes/)
