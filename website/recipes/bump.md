# 递增版本号

::: tip 问题
把 `1.2.3` bump 到 `1.3.0`。
:::

## 🛠 SDK (Go)

```go
versions.NewVersion("1.2.3").BumpMinor().Raw // 1.3.0
```

## 🐚 CLI

```bash
versions bump 1.2.3 --minor
```

## 🤖 MCP

```json
{
  "tool": "version_bump",
  "arguments": {
    "version_string": "1.2.3",
    "bump_type": "patch"
  }
}
```

## 📚 参考

- [bump-minor-version](/sdk/api/bump-minor-version)
- [bump](/cli/commands/bump)

---

[← 返回配方索引](/recipes/)
