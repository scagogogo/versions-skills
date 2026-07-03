# versions set-numbers

::: info 命令
```bash
versions set-numbers <version-string> <numbers>
```
:::

## 📖 简介

修改版本号的数字部分（不可变修改，返回新版本）

修改版本号的数字部分（VersionNumbers），返回新版本字符串。原版本不变。
numbers 参数以逗号分隔的数字字符串提供。

示例:
  versions set-numbers 1.2.3 4,5,6   # 4.5.6
  versions set-numbers v1.2.3 2,0,0  # v2.0.0

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
