# 排序后写回文件

::: tip 问题
把版本列表排序后写入文件。
:::

## 🛠 SDK (Go)

```go
versions.WriteVersionsToFile(vs, "sorted.txt")
```

## 🐚 CLI

```bash
versions write sorted.txt 1.2.0 1.10.0 1.0.0
```

## 🤖 MCP

```json
{
  "tool": "version_write_file",
  "arguments": {
    "filepath": "/path/to/output.txt",
    "versions": [
      "1.0.0",
      "2.0.0"
    ]
  }
}
```

## 📚 参考

- [write-versions-to-file](/sdk/api/write-versions-to-file)
- [write](/cli/commands/write)

---

[← 返回配方索引](/recipes/)
