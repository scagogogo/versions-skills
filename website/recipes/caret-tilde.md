# 用 caret/tilde 判断兼容性

::: tip 问题
判断版本是否满足 `^1.2.0` 或 `~1.2.0`。
:::

## 🛠 SDK (Go)

```go
cs, _ := versions.ParseConstraintSet("^1.2.0")
cs.Match(v)
```

## 🐚 CLI

```bash
versions constraint "^1.2.0" 1.5.0
```

## 🤖 MCP

```json
{
  "tool": "version_constraint_check",
  "arguments": {
    "expression": "^1.2.3",
    "version": "1.5.0",
    "type": "constraint_set"
  }
}
```

## 📚 参考

- [constraints-in-practice](/tutorials/constraints-in-practice)
- [constraint](/cli/commands/constraint)

---

[← 返回配方索引](/recipes/)
