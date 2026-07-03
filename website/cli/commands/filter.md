# versions filter

::: info 命令
```bash
versions filter [version-strings...]
```
:::

## 📖 简介

按条件过滤版本号列表

按各种条件过滤版本号列表。可以组合多个过滤条件（AND 逻辑）。

支持通过参数、--from-file 或 stdin 提供版本号列表。
--constraint 支持 single（单个约束）、set（默认，逗号分隔 AND）和 union（|| 分隔 OR）三种类型。

示例:
  versions filter --stable 1.0.0 1.0.0-beta 2.0.0
  versions filter --prerelease 1.0.0-alpha 1.0.0 1.0.0-beta
  versions filter --major 1 1.0.0 2.0.0 1.5.0
  versions filter --constraint "&gt;=1.0.0,&lt;2.0.0" 1.0.0 1.5.0 2.0.0
  versions filter --constraint "&gt;=1.0.0" --constraint-type single 1.0.0 1.5.0 2.0.0
  versions filter --constraint "&gt;=1.0.0 || &gt;=3.0.0" --constraint-type union 1.0.0 2.0.0 3.0.0
  cat versions.txt | versions filter --stable

## ⚙️ Flags

| Flag | 类型 | 默认值 | 说明 |
|:--|:--|:--|:--|
| `--stable` | `Bool` | `false` | 仅保留稳定版本 |
| `--prerelease` | `Bool` | `false` | 仅保留预发布版本 |
| `--major` | `String` | `""` | 按 Major 版本号过滤 |
| `--minor` | `String` | `""` | 按 Minor 版本号过滤 |
| `--patch` | `String` | `""` | 按 Patch 版本号过滤 |
| `--prefix` | `String` | `""` | 按前缀过滤 |
| `--suffix` | `String` | `""` | 按后缀过滤 |
| `--constraint` | `String` | `""` | 按约束表达式过滤 |
| `--constraint-type` | `String` | `"set"` | 约束类型: single|set|union |
| `--from-file` | `String` | `""` | 从文件读取版本号列表 |

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
