# versions bump

::: info 命令
```bash
versions bump <version-string>
```
:::

## 📖 简介

递增版本号

递增版本号的指定部分，并清除后缀。

  --major: 递增 Major，Minor 和 Patch 清零，后缀清除 (1.2.3 → 2.0.0)
  --minor: 递增 Minor，Patch 清零，后缀清除 (1.2.3 → 1.3.0)
  --patch: 递增 Patch，后缀清除 (1.2.3 → 1.2.4)

必须指定其中一个递增类型。

示例:
  versions bump 1.2.3 --major    # 2.0.0
  versions bump 1.2.3 --minor    # 1.3.0
  versions bump 1.2.3 --patch    # 1.2.4

## ⚙️ Flags

| Flag | 类型 | 默认值 | 说明 |
|:--|:--|:--|:--|
| `--major` | `Bool` | `false` | 递增 Major 版本号 |
| `--minor` | `Bool` | `false` | 递增 Minor 版本号 |
| `--patch` | `Bool` | `false` | 递增 Patch 版本号 |

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
