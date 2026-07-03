# 比较两个版本新旧

::: tip 问题
判断 `1.2.3` 与 `1.2.4` 谁更新。
:::

## 🛠 SDK (Go)

```go
a := versions.NewVersion("1.2.3")
b := versions.NewVersion("1.2.4")
a.IsNewerThan(b) // false
```

## 🐚 CLI

```bash
versions compare 1.2.3 1.2.4
```

## 🤖 MCP

```
version_compare(version1="1.2.3", version2="1.2.4")
```

## 📚 参考

- [is-newer-than-version](/sdk/api/is-newer-than-version)
- [compare](/cli/commands/compare)

---

[← 返回配方索引](/recipes/)
