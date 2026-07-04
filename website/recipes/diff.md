# 计算两版本差异

::: tip 问题
计算从 `1.2.3` 到 `2.0.0` 的版本号变化。
:::

## 🛠 SDK (Go)

```go
d := versions.NewVersion("1.2.3").Diff(versions.NewVersion("2.0.0"))
d.IsMajorChange() // true
```

## 🐚 CLI

```bash
# CLI 无直接命令；用 info 拼接
```

## 🤖 MCP

```json
{
  "tool": "version_compare",
  "arguments": {
    "version1": "1.2.3",
    "version2": "1.3.0"
  }
}
```

## 📚 参考

- [diff-version](/sdk/api/diff-version)
- [version-diff](/sdk/api/version-diff)

---

[← 返回配方索引](/recipes/)
