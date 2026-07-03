# versions set-prefix

::: info 命令
```bash
versions set-prefix <version-string> <prefix>
```
:::

## 📖 简介

修改版本号的前缀（不可变修改，返回新版本）

修改版本号的前缀，返回新版本字符串。原版本不变。

示例:
  versions set-prefix 1.2.3 v       # v1.2.3
  versions set-prefix v1.2.3 ""     # 1.2.3

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
