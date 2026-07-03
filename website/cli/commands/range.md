# versions range

::: info 命令
```bash
versions range <start> <end> [version-strings...]
```
:::

## 📖 简介

查询指定范围内的版本号

查询指定范围内的版本号列表。

边界包含策略:
  --include-start (默认 true): 包含起始版本
  --include-end (默认 false): 包含结束版本

示例:
  versions range 1.0.0 3.0.0 1.0.0 1.5.0 2.0.0 3.0.0 4.0.0
  versions range 1.0.0 3.0.0 --include-end 1.0.0 3.0.0 4.0.0
  cat versions.txt | versions range 1.0.0 2.0.0

## ⚙️ Flags

| Flag | 类型 | 默认值 | 说明 |
|:--|:--|:--|:--|
| `--include-start` | `Bool` | `true` | 包含起始版本 |
| `--include-end` | `Bool` | `false` | 包含结束版本 |
| `--from-file` | `String` | `""` | 从文件读取版本号列表 |

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
