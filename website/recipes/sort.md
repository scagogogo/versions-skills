# 按语义排序版本列表

::: tip 问题
把 `1.10.0, 1.2.0, 1.1.0` 按语义（非字典序）排序。
:::

## 🛠 SDK (Go)

```go
sorted := versions.SortVersionSlice(versions.NewVersions("1.10.0","1.2.0","1.1.0"))
```

## 🐚 CLI

```bash
versions sort 1.10.0 1.2.0 1.1.0
```

## 🤖 MCP

```json
{
  "tool": "version_sort",
  "arguments": {
    "versions": [
      "1.10.0",
      "1.2.0",
      "1.1.0"
    ]
  }
}
```

## 📚 参考

- [sort-version-slice](/sdk/api/sort-version-slice)
- [sort](/cli/commands/sort)

---

[← 返回配方索引](/recipes/)
