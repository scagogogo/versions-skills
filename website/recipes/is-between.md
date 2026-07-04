# 判断版本是否在某区间

::: tip 问题
判断 `1.5.0` 是否在 `[1.0.0, 2.0.0]` 内。
:::

## 🛠 SDK (Go)

```go
v := versions.NewVersion("1.5.0")
v.IsBetween(versions.NewVersion("1.0.0"), versions.NewVersion("2.0.0")) // true
```

## 🐚 CLI

```bash
# CLI 无直接命令；用 range 配合
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
      "1.5.0",
      "2.5.0"
    ]
  }
}
```

## 📚 参考

- [is-between-version](/sdk/api/is-between-version)
- [range-query](/tutorials/range-query)

---

[← 返回配方索引](/recipes/)
