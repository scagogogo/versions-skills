# 从任意文本提取版本号

::: tip 问题
从 `app-1.2.3-linux-amd64` 这类字符串中提取版本号。
:::

## 🛠 SDK (Go)

```go
v := versions.Coerce("app-1.2.3-linux-amd64")
v.Raw // 1.2.3
```

## 🐚 CLI

```bash
# CLI 暂无直接 coerce；用 parse 配合 shell 截取
```

## 🤖 MCP

```
version_parse(version_string="app-1.2.3-linux-amd64")
```

## 📚 参考

- [coerce](/sdk/api/coerce)
- [parse-and-check](/tutorials/parse-and-check)

---

[← 返回配方索引](/recipes/)
