# versions validate

::: info 命令
```bash
versions validate <version-string>
```
:::

## 📖 简介

严格验证版本字符串是否有效

严格验证版本字符串是否有效。有效版本返回 exit code 0，无效返回 exit code 1。

验证规则:
  - 版本必须包含数字部分
  - 数字部分不能包含负数

示例:
  versions validate 1.2.3         # 有效 (exit 0)
  versions validate not-a-version # 无效 (exit 1)

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
