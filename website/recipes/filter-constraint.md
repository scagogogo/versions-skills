# 过滤满足约束的版本

::: tip 问题
从列表中筛出满足约束的版本。
:::

## 🛠 SDK (Go)

```go
cs, _ := versions.ParseConstraintSet(">=1.0.0,<2.0.0")
versions.FilterByConstraintSet(vs, cs)
```

## 🐚 CLI

```bash
versions filter 1.0.0 1.5.0 2.0.0 --constraint ">=1.0.0,<2.0.0"
```

## 🤖 MCP

```json
{
  "tool": "version_filter",
  "arguments": {
    "constraint": ">=1.0.0,<2.0.0",
    "versions": [
      "0.5.0",
      "1.0.0",
      "1.5.0",
      "2.0.0"
    ]
  }
}
```

## 📚 参考

- [filter-by-constraint-set](/sdk/api/filter-by-constraint-set)
- [filter](/cli/commands/filter)

---

[← 返回配方索引](/recipes/)
