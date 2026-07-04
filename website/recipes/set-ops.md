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

```json
{
  "tool": "version_set_operation",
  "arguments": {
    "operation": "intersection",
    "set_a": [
      "1.0.0",
      "2.0.0"
    ],
    "set_b": [
      "1.0.0",
      "3.0.0"
    ]
  }
}
```

## 📚 参考

- [union](/sdk/api/union)
- [intersection](/sdk/api/intersection)
- [difference](/sdk/api/difference)

---

[← 返回配方索引](/recipes/)
