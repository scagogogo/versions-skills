# versions suffix-weight

::: info 命令
```bash
versions suffix-weight <version-string>
```
:::

## 📖 简介

获取版本号后缀的语义权重

获取版本号后缀的语义权重（SuffixWeight）。

权重排序: dev(50) &lt; snapshot(60) &lt; nightly(70) &lt; alpha(100) &lt; beta(200) &lt; milestone(300) &lt; rc(400) &lt; final/release/ga(500) &lt; sp(600) &lt; patch(700) &lt; post(800)

示例:
  versions suffix-weight 1.2.3-beta1  # beta (200)
  versions suffix-weight 1.2.3        # unknown (0)

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
