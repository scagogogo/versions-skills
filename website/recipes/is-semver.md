# 判断版本是否符合 SemVer

::: tip 问题
判断版本字符串是否严格符合 SemVer 2.0.0。
:::

## 🛠 SDK (Go)

```go
versions.NewVersion("1.2.3").IsSemver() // true
```

## 🐚 CLI

```bash
versions check 1.2.3 --semver
```

## 🤖 MCP

```json
{
  "tool": "version_validate",
  "arguments": {
    "version_string": "1.2.3"
  }
}
```

## 📚 参考

- [is-semver-version](/sdk/api/is-semver-version)
- [semver](/concepts/semver)

---

[← 返回配方索引](/recipes/)
