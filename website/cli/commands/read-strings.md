# versions read-strings

::: info 命令
```bash
versions read-strings <filepath>
```
:::

## 📖 简介

从文件读取原始版本字符串列表（不解析）

从文件中读取版本字符串列表，每行一个字符串，不做任何解析或验证。

与 read 命令不同，此命令仅返回原始字符串列表。

示例:
  versions read-strings versions.txt

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
