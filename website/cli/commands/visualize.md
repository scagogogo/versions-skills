# versions visualize

::: info 命令
```bash
versions visualize [version-strings...]
```
:::

## 📖 简介

可视化版本号层级结构

以文本树状图展示版本号的层级结构，便于理解版本间的关系。

示例:
  versions visualize 1.0.0 1.0.0-alpha 1.0.0-beta 2.0.0 2.0.0-rc1
  versions visualize --max-items 5 1.0.0 1.0.1 1.0.2 2.0.0
  versions visualize --groups 1.0.0 2.0.0
  cat versions.txt | versions visualize

## ⚙️ Flags

| Flag | 类型 | 默认值 | 说明 |
|:--|:--|:--|:--|
| `--max-items` | `Int` | `0` | 每组最多显示的版本数（0=不限） |
| `--groups` | `Bool` | `false` | 仅显示分组摘要 |
| `--from-file` | `String` | `""` | 从文件读取版本号列表 |

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
