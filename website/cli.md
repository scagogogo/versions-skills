# CLI 命令

二进制名 `versions`，基于 cobra。安装：

```bash
go install github.com/scagogogo/versions-skills/cmd/versions@latest
# 或一键脚本
curl -sL https://raw.githubusercontent.com/scagogogo/versions-skills/main/install.sh | bash
```

预编译二进制覆盖 Linux/macOS/Windows 等 6 系统 × 多架构，见 [Releases](https://github.com/scagogogo/versions-skills/releases/latest)。所有命令也支持从标准输入读取版本（`cat versions.txt | versions sort`）。

## 命令总览

| 类别 | 命令 |
|:--|:--|
| 解析与验证 | `parse` · `validate` · `info` |
| 比较与检查 | `compare` · `check` · `satisfies` |
| 排序与过滤 | `sort` · `sort-strings` · `filter` · `partition` |
| 分组与范围 | `group` · `group-ids` · `group-id` · `group-contains` · `group-latest` · `group-oldest` · `group-stable` · `group-prerelease` · `group-latest-stable` · `group-latest-prerelease` · `range` |
| 最大/最小 | `min` · `max` · `latest-stable` · `latest-prerelease` |
| 构造与变更 | `build` · `bump` · `core` · `clone` · `set-prefix` · `set-suffix` · `set-major` · `set-minor` · `set-patch` · `set-numbers` |
| 属性查询 | `segments` · `suffix-weight` · `sub-version` · `pure-prefix` · `count` |
| 文件 I/O | `read` · `read-strings` · `write` |
| 可视化 | `visualize` |

## 解析与验证

```bash
versions parse v1.2.3-rc1        # 解析，显示各组成部分
versions validate 1.2.3          # 严格校验
versions info v1.2.3-beta1       # 完整信息
```

## 比较与检查

```bash
versions compare 1.0.0 2.0.0              # 比较
versions check --stable 1.2.3             # 是否稳定版（exit code 0=真/1=假）
versions check --beta 1.2.3-beta1         # 是否 beta
versions check --newer 1.0.0 1.5.0        # 后者是否更新
versions satisfies 1.5.0 ">=1.0.0,<2.0.0" # 是否满足约束
```

`check` 支持的后缀类型 flag：`--stable` `--prerelease` `--dev` `--alpha` `--beta` `--rc` `--snapshot` `--milestone` `--nightly` `--final` `--ga` `--pre` `--release` `--sp` `--post` `--zero`；以及关系 flag：`--newer <v>` `--older <v>` `--equal <v>` `--between-low <v> --between-high <v>`。

## 排序与过滤

```bash
versions sort 3.0.0 1.0.0 2.0.0           # 升序
versions sort --desc 3.0.0 1.0.0 2.0.0    # 降序
versions sort-strings 3.0.0 1.0.0 2.0.0   # 排序但返回原始字符串（非规范化的 Version）
versions filter --stable 1.0.0-alpha 1.0.0 2.0.0-beta 2.0.0
versions filter --constraint ">=1.0.0,<2.0.0" 0.5.0 1.0.0 1.5.0 2.0.0
versions filter --constraint ">=1.0.0 || >=3.0.0" --constraint-type union 1.0.0 2.0.0 3.0.0
versions partition --stable 1.0.0-alpha 1.0.0 2.0.0   # 分两组
```

## 分组与范围

```bash
# 全局分组
versions group 1.0.0 1.1.0 2.0.0          # 按数字部分分组
versions group-ids 1.0.0 1.1.0 2.0.0      # 列出所有分组 ID

# 单版本属性
versions group-id 1.2.3                   # 该版本的分组 ID
versions group-contains 1.0.0 1.0.0 1.0.1 # 检查分组是否含某版本

# 分组内查询
versions group-latest 1.0.0 1.0.1 1.1.0   # 分组最新
versions group-oldest 1.0.0 1.0.1 1.1.0   # 分组最旧
versions group-stable 1.0.0-alpha 1.0.0   # 分组内稳定版
versions group-prerelease 1.0.0-alpha 1.0.0  # 分组内预发布版
versions group-latest-stable 1.0.0-alpha 1.0.0 1.0.1  # 分组最新稳定版
versions group-latest-prerelease 1.0.0-alpha 1.0.0    # 分组最新预发布版

# 范围查询
versions range 1.0.0 2.0.0 1.0.0 1.5.0 2.0.0 3.0.0            # [1.0.0, 2.0.0]，默认含起不含止
versions range 1.0.0 2.0.0 --include-end 1.0.0 1.5.0 2.0.0     # [1.0.0, 2.0.0]，含两端
```

`--include-start`（默认 true）/ `--include-end`（默认 false）控制边界是否纳入，对应 [算法详解 §11](./algorithms#_11-containspolicy-边界纳入策略) 的 `ContainsPolicy`。

## 最大/最小/最新

```bash
versions min 3.0.0 1.0.0 2.0.0
versions max 3.0.0 1.0.0 2.0.0
versions latest-stable 1.0.0-alpha 1.0.0 2.0.0-beta 2.0.0
versions latest-prerelease 1.0.0-alpha 1.0.0 2.0.0-beta
```

## 构造与变更

`set-*` 系列与 `bump` 都遵循**不可变**语义——返回新版本，原版本不变（见 [算法详解 §8](./algorithms#_8-不可变变更-builder-重建)）。

```bash
versions build --prefix v --major 1 --minor 2 --patch 3          # → v1.2.3
versions build --major 1 --minor 2 --patch 3 --suffix -beta1     # → 1.2.3-beta1
versions build --numbers 1,2,3,4                                  # → 1.2.3.4（任意段数）
versions bump 1.2.3 --patch              # 递增 patch → 1.2.4（后缀脱落）
versions core 1.2.3-beta                 # 去后缀核心版本 → 1.2.3
versions clone 1.2.3                     # 深拷贝
versions set-suffix 1.2.3 -rc1           # 改后缀（不可变）
versions set-prefix 1.2.3 v              # 改前缀（不可变）
versions set-major 1.2.3 5               # 改 major（不可变）
versions set-patch 1.2.3 9               # 改 patch（不可变）
```

## 属性查询

```bash
versions segments 1.2.3                  # 数字段列表 → [1 2 3]
versions suffix-weight 1.2.3-beta1       # 后缀权重 → 200
versions sub-version 1.2.3-beta1         # 后缀子版本号 → 1
versions pure-prefix v1.2.3              # 纯净前缀（去尾部分隔符）→ v
versions count --stable 1.0.0-alpha 1.0.0 2.0.0  # 统计满足条件的数量
```

## 文件 I/O

```bash
versions read versions.txt               # 从文件读取版本列表（解析为 Version）
versions read-strings versions.txt       # 读原始字符串（不解析）
versions write output.txt 1.0.0 2.0.0 3.0.0  # 写入文件（写前自动排序）
```

文件格式与坑见 [算法详解 §9](./algorithms#_9-文件-io-行式格式)。

## 可视化

```bash
versions visualize 1.0.0 1.1.0 2.0.0 --groups        # 仅显示分组摘要树
versions visualize 1.0.0 1.1.0 2.0.0 --max-items 5    # 每组最多显示 5 个版本
```

可视化输出形态见 [算法详解 §10](./algorithms#_10-可视化-组感知文本树)。

---

::: tip exit code 约定
`check` 类命令用 exit code 表示布尔结果：`0` = 真，`1` = 假。便于在 shell 脚本和 CI 中用 `&&` / `||` 串联：

```bash
versions check --stable "$VER" && echo "已是稳定版，可发布"
versions check --newer 1.0.0 1.5.0 && echo "1.5.0 比 1.0.0 新"
```
:::

::: tip 管道用法
所有接受 `[version-strings...]` 的命令都支持从标准输入读取（每行一个版本）：

```bash
cat versions.txt | versions sort | versions filter --stable
```
:::

::: tip 从文件读取
不想用管道时，`sort` / `filter` / `visualize` / `count` / `min` / `max` 等命令都支持 `--from-file <path>` 直接读版本列表文件：

```bash
versions sort --from-file versions.txt --desc
versions latest-stable --from-file versions.txt
```
:::

→ SDK 等价 API 见 [Go SDK](./sdk)，MCP 等价工具见 [MCP](./mcp)。
