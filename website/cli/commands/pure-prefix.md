# versions pure-prefix

::: info 命令
```bash
versions pure-prefix <version-string>
```
:::

## 📖 简介

获取版本号的纯净前缀（去除尾部分隔符）

获取版本号前缀的纯净形式，即去除尾部的 - 或 . 等分隔符。

例如 v1.2.3 的 Prefix 为 "v"，PurePrefix 也为 "v"。
例如 curl-7.85.0 的 Prefix 为 "curl-"，PurePrefix 为 "curl"。

示例:
  versions pure-prefix v1.2.3       # v
  versions pure-prefix curl-7.85.0  # curl

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
