# 找最新稳定版

::: tip 问题
从版本列表中找最新的稳定（非预发布）版本。
:::

## 🛠 SDK (Go)

```go
versions.LatestStable(vs).Raw
```

## 🐚 CLI

```bash
versions latest-stable 1.0.0 1.1.0-rc1 1.2.0
```

## 🤖 MCP

```json
{
  "tool": "version_latest_stable",
  "arguments": {
    "versions": [
      "1.0.0-alpha",
      "1.0.0",
      "2.0.0-beta",
      "2.0.0"
    ]
  }
}
```

## 📚 参考

- [latest-stable](/sdk/api/latest-stable)
- [latest-stable](/cli/commands/latest-stable)

---

[← 返回配方索引](/recipes/)
