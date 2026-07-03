# versions sub-version

::: info 命令
```bash
versions sub-version <version-string>
```
:::

## 📖 简介

获取后缀中的子版本号

获取版本号后缀中的子版本号数字部分。

例如 1.2.3-beta2 中的子版本号为 2。

示例:
  versions sub-version 1.2.3-beta2  # 2
  versions sub-version 1.2.3        # 0

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
