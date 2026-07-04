# 修改版本号某段

::: tip 问题
把 `1.2.3` 的 Patch 改成 5。
:::

## 🛠 SDK (Go)

```go
versions.NewVersion("1.2.3").WithPatch(5).Raw // 1.2.5
```

## 🐚 CLI

```bash
versions set-patch 1.2.3 5
```

## 🤖 MCP

```json
{
  "tool": "version_build",
  "arguments": {
    "prefix": "v",
    "major": 1,
    "minor": 2,
    "patch": 5
  }
}
```

## 📚 参考

- [with-patch-version](/sdk/api/with-patch-version)
- [set-patch](/cli/commands/set-patch)

---

[← 返回配方索引](/recipes/)
