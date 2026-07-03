# versions group-id

::: info 命令
```bash
versions group-id <version-string>
```
:::

## 📖 简介

获取版本号的分组 ID

获取版本号的分组 ID（BuildGroupID），即版本号数字部分以 . 连接。

例如 v1.2.3-beta1 的分组 ID 为 "1.2.3"。

示例:
  versions group-id v1.2.3-beta1  # 1.2.3
  versions group-id 2.0.0         # 2.0.0

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
