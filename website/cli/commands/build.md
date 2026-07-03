# versions build

::: info 命令
```bash
versions build
```
:::

## 📖 简介

构建版本号字符串

通过指定各组成部分构建版本号字符串。

可使用 --numbers 直接指定完整的数字部分（逗号分隔），或使用 --major/--minor/--patch 分别指定。
--numbers 与 --major/--minor/--patch 同时指定时，--numbers 优先生效。

示例:
  versions build --major 1 --minor 2 --patch 3
  versions build --prefix v --major 1 --minor 0 --suffix -alpha1
  versions build --numbers 1,2,3,4 --prefix v

## ⚙️ Flags

| Flag | 类型 | 默认值 | 说明 |
|:--|:--|:--|:--|
| `--prefix` | `String` | `""` | 版本前缀（如 'v'） |
| `--major` | `String` | `""` | Major 版本号 |
| `--minor` | `String` | `""` | Minor 版本号 |
| `--patch` | `String` | `""` | Patch 版本号 |
| `--suffix` | `String` | `""` | 版本后缀（如 '-beta1'） |
| `--numbers` | `String` | `""` | 完整数字部分，逗号分隔（如 '1,2,3,4'） |

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
