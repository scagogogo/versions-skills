# 用 Builder 构造版本

::: tip 问题
从各组件构造 `v1.2.3-beta1`。
:::

## 🛠 SDK (Go)

```go
v := versions.NewVersionBuilder().Prefix("v").Major(1).Minor(2).Patch(3).Suffix("-beta1").Build()
```

## 🐚 CLI

```bash
versions build --prefix v --major 1 --minor 2 --patch 3 --suffix -beta1
```

## 🤖 MCP

```json
{
  "tool": "version_build",
  "arguments": {
    "prefix": "v",
    "suffix": "-beta1",
    "major": 1,
    "minor": 2,
    "patch": 3
  }
}
```

## 📚 参考

- [version-builder](/sdk/api/version-builder)
- [build](/cli/commands/build)

---

[← 返回配方索引](/recipes/)
