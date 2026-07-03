# 集合交集/并集/差集

::: tip 问题
对两个版本集合做集合运算。
:::

## 🛠 SDK (Go)

```go
versions.Union(a, b)
versions.Intersection(a, b)
versions.Difference(a, b)
```

## 🐚 CLI

```bash
# CLI 无直接命令
```

## 🤖 MCP

```
version_set_operation(operation="union", set_a=[...], set_b=[...])
```

## 📚 参考

- [union](/sdk/api/union)
- [intersection](/sdk/api/intersection)
- [difference](/sdk/api/difference)

---

[← 返回配方索引](/recipes/)
