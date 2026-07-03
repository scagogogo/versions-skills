# versions compare

::: info 命令
```bash
versions compare <version1> <version2>
```
:::

## 📖 简介

比较两个版本号

比较两个版本号的大小关系。

返回结果:
  -1 表示 version1 旧于 version2
   0 表示两个版本相等
   1 表示 version1 新于 version2

示例:
  versions compare 1.2.3 2.0.0     # 1.2.3 旧于 2.0.0
  versions compare v1.0 v1.0.0     # 相等
  versions compare 2.0 1.0         # 2.0 新于 1.0

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
