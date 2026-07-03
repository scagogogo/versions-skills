# versions constraint

::: info 命令
```bash
versions constraint <expression> <version>
```
:::

## 📖 简介

检查版本号是否满足约束表达式

检查版本号是否满足给定的约束表达式。

支持三种约束类型:
  - single: 单个 Constraint，如 "&gt;=1.0.0"
  - set (默认): ConstraintSet，逗号分隔，AND 逻辑
  - union: ConstraintUnion，|| 分隔，OR 逻辑

支持的运算符:
  =, !=, &gt;, &gt;=, &lt;, &lt;=, ^ (caret), ~ (tilde), x/X/* (wildcard)

示例:
  versions constraint "&gt;=1.0.0" 1.5.0
  versions constraint "&gt;=1.0.0" 1.5.0 --type single
  versions constraint "&gt;=1.0.0,&lt;2.0.0" 1.5.0
  versions constraint "&gt;=1.0.0 || &gt;=3.0.0" 3.5.0 --type union

## ⚙️ Flags

| Flag | 类型 | 默认值 | 说明 |
|:--|:--|:--|:--|
| `--type` | `String` | `"set"` | 约束类型: single|set|union |

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
