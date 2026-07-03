# versions group-ids

::: info 命令
```bash
versions group-ids [version-strings...]
```
:::

## 📖 简介

列出所有版本分组的 ID

列出所有版本分组的 ID（即 VersionNumbers 以 . 连接）。

示例:
  versions group-ids 1.0.0 1.0.0-alpha 2.0.0 2.0.0-rc1

## ⚙️ Flags

| Flag | 类型 | 默认值 | 说明 |
|:--|:--|:--|:--|
| `--from-file` | `String` | `""` | 从文件读取版本号列表 |
| `--group-id` | `String` | `""` | 分组 ID（版本号数字部分，如 '1.0.0'） |

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
