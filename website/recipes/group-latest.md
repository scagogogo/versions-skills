# 取分组内最新版本

::: tip 问题
取某个 `major.minor` 分组里的最新版本。
:::

## 🛠 SDK (Go)

```go
g.GetLatest().Raw
```

## 🐚 CLI

```bash
versions group-latest 1.0.0 1.0.1 1.0.2 --group-id 1.0
```

## 🤖 MCP

```json
{
  "tool": "version_group",
  "arguments": {
    "versions": [
      "1.0.0",
      "1.0.0-rc1",
      "1.1.0"
    ]
  }
}
```

## 📚 参考

- [get-latest-versiongroup](/sdk/api/get-latest-versiongroup)
- [group-latest](/cli/commands/group-latest)

---

[← 返回配方索引](/recipes/)
