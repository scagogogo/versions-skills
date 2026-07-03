# versions count

::: info 命令
```bash
versions count [version-strings...]
```
:::

## 📖 简介

统计满足条件的版本号数量

统计版本号列表中满足指定条件的版本数量。

多个条件之间为 AND 逻辑。

示例:
  versions count --stable 1.0.0 1.0.0-beta 2.0.0
  versions count --prerelease 1.0.0-alpha 1.0.0 2.0.0-beta
  versions count --major 1 1.0.0 2.0.0 1.5.0
  versions count --stable --major 2 1.0.0 2.0.0 2.5.0-beta 3.0.0
  cat versions.txt | versions count --stable

## ⚙️ Flags

| Flag | 类型 | 默认值 | 说明 |
|:--|:--|:--|:--|
| `--stable` | `Bool` | `false` | 统计稳定版本数量 |
| `--prerelease` | `Bool` | `false` | 统计预发布版本数量 |
| `--major` | `String` | `""` | 按 Major 版本号统计 |
| `--minor` | `String` | `""` | 按 Minor 版本号统计 |
| `--patch` | `String` | `""` | 按 Patch 版本号统计 |
| `--from-file` | `String` | `""` | 从文件读取版本号列表 |

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
