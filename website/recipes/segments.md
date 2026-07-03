# 获取版本号各数字段

::: tip 问题
取出 `1.2.3` 的数字段 `[1,2,3]`。
:::

## 🛠 SDK (Go)

```go
v := versions.NewVersion("1.2.3")
v.Segments() // [1 2 3]
```

## 🐚 CLI

```bash
versions segments 1.2.3
```

## 🤖 MCP

```
version_parse(version_string="1.2.3") → version_numbers
```

## 📚 参考

- [segments-version](/sdk/api/segments-version)
- [segments](/cli/commands/segments)

---

[← 返回配方索引](/recipes/)
