# versions core

::: info 命令
```bash
versions core <version-string>
```
:::

## 📖 简介

获取核心版本号（去除后缀）

获取版本号的核心部分，即去除预发布后缀后的版本。

例如: v1.2.3-beta1 → v1.2.3

示例:
  versions core v1.2.3-beta1
  versions core 2.0.0-rc1

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
