# versions check

::: info 命令
```bash
versions check <version-string>
```
:::

## 📖 简介

检查版本号属性（返回布尔结果，exit code 0=真/1=假）

检查版本号的各类属性，返回布尔结果。

使用对应的 flag 检查特定属性，命令返回 JSON 结果，同时以 exit code 表示真假：
  exit code 0 = 属性为真
  exit code 1 = 属性为假

支持的属性检查:
  --prerelease    是否为预发布版本
  --stable        是否为稳定版本
  --dev           是否为开发版
  --alpha         是否为 Alpha 版
  --beta          是否为 Beta 版
  --rc            是否为 RC 版
  --snapshot      是否为快照版
  --milestone     是否为里程碑版
  --nightly       是否为夜间构建版
  --final         是否为 Final 版
  --ga            是否为 GA 版
  --pre           是否为 Pre 版
  --release       是否为 Release 版
  --sp            是否为 SP 版
  --post          是否为 Post 版
  --zero          是否为零值版本
  --newer &lt;v&gt;     是否比指定版本新
  --older &lt;v&gt;     是否比指定版本旧
  --equal &lt;v&gt;     是否与指定版本相等
  --between-low/--between-high  是否在指定范围内

示例:
  versions check --beta v1.2.3-beta1
  versions check --stable 1.2.3
  versions check --stable 1.2.3-beta
  versions check --newer 1.0.0 2.0.0
  versions check --between-low 1.0.0 --between-high 3.0.0 2.0.0

## ⚙️ Flags

| Flag | 类型 | 默认值 | 说明 |
|:--|:--|:--|:--|
| `--prerelease` | `Bool` | `false` | 是否为预发布版本 |
| `--stable` | `Bool` | `false` | 是否为稳定版本 |
| `--dev` | `Bool` | `false` | 是否为开发版 |
| `--alpha` | `Bool` | `false` | 是否为 Alpha 版 |
| `--beta` | `Bool` | `false` | 是否为 Beta 版 |
| `--rc` | `Bool` | `false` | 是否为 RC 版 |
| `--snapshot` | `Bool` | `false` | 是否为快照版 |
| `--milestone` | `Bool` | `false` | 是否为里程碑版 |
| `--nightly` | `Bool` | `false` | 是否为夜间构建版 |
| `--final` | `Bool` | `false` | 是否为 Final 版 |
| `--ga` | `Bool` | `false` | 是否为 GA 版 |
| `--pre` | `Bool` | `false` | 是否为 Pre 版 |
| `--release` | `Bool` | `false` | 是否为 Release 版 |
| `--sp` | `Bool` | `false` | 是否为 SP 版 |
| `--post` | `Bool` | `false` | 是否为 Post 版 |
| `--zero` | `Bool` | `false` | 是否为零值版本 |
| `--newer` | `String` | `""` | 是否比指定版本新（提供目标版本号） |
| `--older` | `String` | `""` | 是否比指定版本旧（提供目标版本号） |
| `--equal` | `String` | `""` | 是否与指定版本相等（提供目标版本号） |
| `--between-low` | `String` | `""` | 范围检查的最低版本 |
| `--between-high` | `String` | `""` | 范围检查的最高版本 |

## 📚 相关

- [SDK API](/sdk/) · [MCP 工具](/mcp/) · [CLI 总览](/cli/)

---

::: details 源码
定义于 `internal/cli/`
:::
