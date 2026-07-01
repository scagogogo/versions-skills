# CLI 命令

二进制名 `versions`，基于 cobra。安装：

```bash
go install github.com/scagogogo/versions-skills/cmd/versions@latest
# 或一键脚本
curl -sL https://raw.githubusercontent.com/scagogogo/versions-skills/main/install.sh | bash
```

预编译二进制覆盖 Linux/macOS/Windows 等 6 系统 × 多架构，见 [Releases](https://github.com/scagogogo/versions-skills/releases/latest)。

## 解析与验证

```bash
versions parse v1.2.3-rc1        # 解析，显示各组成部分
versions validate 1.2.3          # 校验
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

## 排序与过滤

```bash
versions sort 3.0.0 1.0.0 2.0.0           # 升序
versions sort --desc 3.0.0 1.0.0 2.0.0    # 降序
versions filter --stable 1.0.0-alpha 1.0.0 2.0.0-beta 2.0.0
versions filter --constraint ">=1.0.0,<2.0.0" 0.5.0 1.0.0 1.5.0 2.0.0
versions partition --stable 1.0.0-alpha 1.0.0 2.0.0   # 分两组
```

## 分组与范围

```bash
versions group 1.0.0 1.1.0 2.0.0          # 按数字部分分组
versions range 1.0.0 2.0.0 1.0.0 1.5.0 2.0.0 3.0.0  # 范围内版本
versions group-list 1.0.0 1.1.0 2.0.0     # 列出分组 ID
```

## 最大/最小/最新

```bash
versions min 3.0.0 1.0.0 2.0.0
versions max 3.0.0 1.0.0 2.0.0
versions latest-stable 1.0.0-alpha 1.0.0 2.0.0-beta 2.0.0
versions latest-prerelease 1.0.0-alpha 1.0.0 2.0.0-beta
```

## 构造与变更

```bash
versions build --prefix v --major 1 --minor 2 --patch 3
versions bump 1.2.3 --patch              # 递增 patch → 1.2.4
versions core 1.2.3-beta                 # 去后缀核心版本 → 1.2.3
versions clone 1.2.3                     # 深拷贝
versions with-suffix 1.2.3 -rc1          # 改后缀（不可变）
```

## 属性查询

```bash
versions prefix v1.2.3                   # 纯净前缀
versions group-id 1.2.3                  # 分组 ID
versions numbers 1.2.3                   # 数字段列表
versions suffix-weight 1.2.3-beta1       # 后缀权重
versions sub-version 1.2.3-beta1         # 后缀子版本号
```

## 文件 I/O

```bash
versions read versions.txt               # 从文件读取版本列表
versions read-raw versions.txt            # 读原始字符串（不解析）
versions write output.txt 1.0.0 2.0.0 3.0.0  # 写入文件
```

## 可视化

```bash
versions visualize 1.0.0 1.1.0 2.0.0 --groups   # 树形层次
```

## 统计

```bash
versions count --stable 1.0.0-alpha 1.0.0 2.0.0  # 统计满足条件的数量
```

---

::: tip exit code 约定
`check` 类命令用 exit code 表示布尔结果：`0` = 真，`1` = 假。便于在 shell 脚本和 CI 中用 `&&` / `||` 串联。
:::

→ SDK 等价 API 见 [Go SDK](./sdk)，MCP 等价工具见 [MCP](./mcp)。
