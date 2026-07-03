# versions latest-stable

::: info 命令
```bash
versions latest-stable [version-strings...]
```
:::

## 📖 简介

查找最新稳定版本

从版本号列表中查找最新的稳定版本（不含预发布标识）。

示例:
  versions latest-stable 1.0.0-alpha 1.0.0 2.0.0-beta 2.0.0

## ⚙️ Flags

| Flag | 类型 | 默认值 | 说明 |
|:--|:--|:--|:--|
| `--from-file` | `String` | `""` | 从文件读取版本号列表 |

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
