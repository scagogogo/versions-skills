# versions max

::: info 命令
```bash
versions max [version-strings...]
```
:::

## 📖 简介

查找最大（最新）版本

从版本号列表中查找最大（最新）的版本。

示例:
  versions max 2.0.0 1.0.0 1.5.0
  cat versions.txt | versions max

## ⚙️ Flags

| Flag | 类型 | 默认值 | 说明 |
|:--|:--|:--|:--|
| `--from-file` | `String` | `""` | 从文件读取版本号列表 |

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
