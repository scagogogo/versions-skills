# versions satisfies

::: info 命令
```bash
versions satisfies <version-string> <constraint-expression>
```
:::

## 📖 简介

检查版本号是否满足约束表达式（版本为中心）

检查版本号是否满足给定的约束表达式。

与 constraint 命令相反：constraint 是约束为中心，satisfies 是版本为中心。
satisfies 使用 Version.Matches() 方法，自动解析表达式。

示例:
  versions satisfies 1.5.0 "&gt;=1.0.0,&lt;2.0.0"
  versions satisfies 2.0.0 "^1.0.0"

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
