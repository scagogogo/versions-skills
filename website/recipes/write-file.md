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

```
version_write_file(filepath="sorted.txt", versions=[...])
```

## 📚 参考

- [write-versions-to-file](/sdk/api/write-versions-to-file)
- [write](/cli/commands/write)

---

[← 返回配方索引](/recipes/)
