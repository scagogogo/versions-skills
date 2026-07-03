# versions parse

::: info 命令
```bash
versions parse <version-string>
```
:::

## 📖 简介

解析版本字符串，显示各组成部分

解析版本字符串并显示其结构化组成部分，包括前缀、数字部分、后缀、后缀权重等。

可通过 --delimiters 指定自定义分隔符（默认为 "."），用于解析非标准格式的版本号。

示例:
  versions parse v1.2.3-beta1
  versions parse 2.0.0
  versions parse curl-7_85_0
  versions parse --delimiters "_-" curl-7_85_0

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
