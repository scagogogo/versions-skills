# versions sort-strings

::: info 命令
```bash
versions sort-strings [version-strings...]
```
:::

## 📖 简介

对版本号字符串列表排序（返回原始字符串，非 Version 对象）

对版本号字符串列表排序，返回原始字符串而非结构化 Version 对象。
适用于仅需排序结果不需要解析详情的场景。

示例:
  versions sort-strings 2.0.0 1.0.0 1.5.0
  versions sort-strings --desc 1.0 2.0 1.5
  cat versions.txt | versions sort-strings

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
