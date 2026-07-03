# 判断版本是否为预发布

::: tip 问题
判断 `1.2.3-rc1` 这类版本是否是预发布版。
:::

## 🛠 SDK (Go)

```go
v := versions.NewVersion("1.2.3-rc1")
v.IsPrerelease() // true
```

## 🐚 CLI

```bash
versions check 1.2.3-rc1 --prerelease
```

## 🤖 MCP

```
version_info(version_string="1.2.3-rc1") → is_prerelease
```

## 📚 参考

- [is-prerelease-version](/sdk/api/is-prerelease-version)
- [check](/cli/commands/check)

---

[← 返回配方索引](/recipes/)
