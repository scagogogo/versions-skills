# versions set-suffix

::: info 命令
```bash
versions set-suffix <version-string> <suffix>
```
:::

## 📖 简介

修改版本号的后缀（不可变修改，返回新版本）

修改版本号的后缀，返回新版本字符串。原版本不变。

示例:
  versions set-suffix 1.2.3 -beta1   # 1.2.3-beta1
  versions set-suffix 1.2.3-beta1 "" # 1.2.3

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
