# 读取版本文件

::: tip 问题
从文件读取版本号列表。
:::

## 🛠 SDK (Go)

```go
vs, _ := versions.ReadVersionsFromFile("releases.txt")
```

## 🐚 CLI

```bash
versions read releases.txt
```

## 🤖 MCP

```json
{
  "tool": "version_read_file",
  "arguments": {
    "filepath": "/path/to/releases.txt"
  }
}
```

## 📚 参考

- [read-versions-from-file](/sdk/api/read-versions-from-file)
- [read](/cli/commands/read)

---

[← 返回配方索引](/recipes/)
