# versions min

::: info 命令
```bash
versions min [version-strings...]
```
:::

## 📖 简介

查找最小（最旧）版本

从版本号列表中查找最小（最旧）的版本。

示例:
  versions min 2.0.0 1.0.0 1.5.0
  cat versions.txt | versions min

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
