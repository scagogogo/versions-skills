# versions group-contains

::: info 命令
```bash
versions group-contains [version-strings...]
```
:::

## 📖 简介

检查指定分组是否包含某个版本

检查指定分组 ID 是否包含某个版本号。通过 --group-id 和 --version 指定分组和目标版本。

示例:
  versions group-contains --group-id 1.0.0 --version 1.0.0-alpha 1.0.0-alpha 1.0.0 2.0.0

## ⚙️ Flags

| Flag | 类型 | 默认值 | 说明 |
|:--|:--|:--|:--|
| `--version` | `String` | `""` | 要检查的目标版本号 |

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
