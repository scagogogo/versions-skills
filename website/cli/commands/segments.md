# versions segments

::: info 命令
```bash
versions segments <version-string>
```
:::

## 📖 简介

获取版本号的数字段列表

获取版本号的数字段（Segments），返回 int 数组。

示例:
  versions segments 1.2.3       # [1, 2, 3]
  versions segments v1.2.3.4   # [1, 2, 3, 4]

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
