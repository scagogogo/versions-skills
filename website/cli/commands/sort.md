# versions sort

::: info 命令
```bash
versions sort [version-strings...]
```
:::

## 📖 简介

对版本号列表进行排序

对版本号列表进行排序，默认升序排列。

支持通过参数、--from-file 或 stdin 提供版本号列表。

示例:
  versions sort 2.0.0 1.0.0 1.5.0
  versions sort --desc 1.0 2.0 1.5
  cat versions.txt | versions sort
  versions sort --from-file versions.txt

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
