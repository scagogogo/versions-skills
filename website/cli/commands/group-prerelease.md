# versions group-prerelease

::: info 命令
```bash
versions group-prerelease [version-strings...]
```
:::

## 📖 简介

获取指定分组的预发布版本列表

获取指定分组 ID 的所有预发布版本。通过 --group-id 指定分组 ID。

示例:
  versions group-prerelease --group-id 1.0.0 1.0.0-alpha 1.0.0 1.0.0-beta 2.0.0

## ⚙️ Flags

| Flag | 类型 | 默认值 | 说明 |
|:--|:--|:--|:--|
| `--from-file` | `String` | `""` | 从文件读取版本号列表 |
| `--group-id` | `String` | `""` | 分组 ID（版本号数字部分，如 '1.0.0'） |
| `--from-file` | `String` | `""` | 从文件读取版本号列表 |
| `--group-id` | `String` | `""` | 分组 ID（版本号数字部分，如 '1.0.0'） |

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
