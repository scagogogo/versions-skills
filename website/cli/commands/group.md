# versions group

::: info 命令
```bash
versions group [version-strings...]
```
:::

## 📖 简介

按版本号数字部分分组

按版本号数字部分（VersionNumbers）对版本列表进行分组。
数字部分相同的版本归入同一组，如 1.2.3 和 1.2.3-beta 属于同一组。

示例:
  versions group 1.0.0 1.0.0-alpha 2.0.0 1.0.0-beta 2.0.0-rc1
  versions group --id 1.0.0 1.0.0 1.0.0-alpha 2.0.0
  cat versions.txt | versions group

## ⚙️ Flags

| Flag | 类型 | 默认值 | 说明 |
|:--|:--|:--|:--|
| `--id` | `String` | `""` | 仅显示指定组 ID 的详情 |
| `--from-file` | `String` | `""` | 从文件读取版本号列表 |

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
