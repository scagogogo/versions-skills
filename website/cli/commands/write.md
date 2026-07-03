# versions write

::: info 命令
```bash
versions write <filepath> [version-strings...]
```
:::

## 📖 简介

将版本号列表写入文件

将版本号列表写入文件，每行一个版本字符串。版本号会先排序再写入。

示例:
  versions write output.txt 1.0.0 2.0.0 1.5.0
  versions write output.txt --from-file input.txt

## ⚙️ Flags

| Flag | 类型 | 默认值 | 说明 |
|:--|:--|:--|:--|
| `--from-file` | `String` | `""` | 从文件读取版本号列表（用于写入到目标文件） |

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
