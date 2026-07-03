# versions partition

::: info 命令
```bash
versions partition [version-strings...]
```
:::

## 📖 简介

将版本号列表分为两组（满足条件和不满足条件）

将版本号列表按条件分成两组：满足条件和不满足条件的版本。

示例:
  versions partition --stable 1.0.0-alpha 1.0.0 1.0.0-beta 2.0.0
  versions partition --prerelease 1.0.0-alpha 1.0.0 2.0.0-rc1

## ⚙️ Flags

| Flag | 类型 | 默认值 | 说明 |
|:--|:--|:--|:--|
| `--stable` | `Bool` | `false` | 按稳定/不稳定分区 |
| `--prerelease` | `Bool` | `false` | 按预发布/非预发布分区 |
| `--from-file` | `String` | `""` | 从文件读取版本号列表 |

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
