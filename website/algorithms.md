# 算法详解

本页记录 AI Agent（或人类）在不读源码的情况下预测库行为所需的精确语义。每条规则都可在对应源文件中验证。

## 1. 解析 —— 三段式 + 元数据

源文件：`parser.go`

每个版本字符串被拆为四个字段：

```
 v1.2.3-beta1+build.7
 │ └─┬─┘ └──┬──┘ └──┬──┘
 │  │      │       └─ Metadata    （semver 构建元数据，"+" 之后）
 │  │      └─ Suffix              （数字部分之后的全部内容）
 │  └─ VersionNumbers             （整数段，解析后统一以 "." 拼接）
 └─ Prefix                        （非数字前导，如 "v"、"release-"）
```

:::mermaid
flowchart LR
  IN["'v1.2.3-beta1+build.7'"] --> T["Trim 空白"]
  T --> M{"找 '+'<br/>其后不含 '-'?"}
  M -->|"是"| M2["剥离 Metadata<br/>'build.7'"]
  M -->|"否"| M2X["保留在后缀<br/>(Scala/Maven)"]
  M2 --> D{"含数字?"}
  M2X --> D
  D -->|"否"| PRE["整串作 Prefix<br/>无效版本"]
  D -->|"是"| P["读 Prefix → 'v'"]
  P --> N["读 Numbers → [1,2,3]"]
  N --> SU["读 Suffix → '-beta1'"]
  style IN fill:#eff6ff,stroke:#2563eb
  style PRE fill:#fef2f2,stroke:#dc2626
  style SU fill:#f0fdf4,stroke:#16a34a
:::

算法步骤（顺序固定）：

1. **Trim** 去除首尾空白。空串 → 无效版本，`VersionNumbers` 为空数组。
2. **剥离 metadata**：找最后一个 `+`，其后的内容**仅当不含 `-`** 时才视为 metadata。
3. **纯字母快捷路径**：若剩余串完全不含数字，整个串作为 `Prefix`，`VersionNumbers` 为空（无效版本）。
4. **读前缀**：扫到第一个数字，之前的内容即前缀。`v1.2.3` → `"v"`；`.1` → `""`。
5. **读数字**：从第一个数字起，收集由分隔符（默认 `.`）分隔的数字段。连续分隔符折叠；遇非数字非分隔符即停。`1.2.48` → `[1,2,48]`。
6. **读后缀**：数字部分定位之后的所有内容。先用正则级联处理常见模式（`-snapshot.…`、`-v2.x.x`、`-revN-…`、`+nnn-xxxx`、`-beta1`），都匹配不上则取剩余字符串。

::: warning 关于 metadata 的歧义
`0.9.0+121-bcc5decc`（Scala/Maven 风格）会把 `121-bcc5decc` **留在后缀里**，而不是当成 semver 元数据剥离——因为它含 `-`。而 `1.2.3+build.7` 才把 `build.7` 剥到 `Metadata`。这是为了区分 semver 元数据与合法含 `-` 的预发布标识。
:::

`Coerce(s)` 做同样的事，但先从任意文本里抽出符合版本形态的子串：`"program-1.2.3-linux-amd64"` → `"1.2.3"`。

::: tip 分隔符可统一
解析支持可配置的 `Delimiters`，但 `BuildGroupID()` 重建时一律用 `.` 拼接——所以 `"1/2/3"` 与 `"1.2.3"` 产生相同的规范数字形态。
:::

## 2. 比较 —— 四级优先级

源文件：`version.go` 的 `CompareTo`、`version_numbers.go`、`version_suffix.go`

`a.CompareTo(b)` 返回 `-1/0/1`，按顺序尝试以下键，**第一个不同的胜出**：

:::mermaid
flowchart TD
  START["a.CompareTo(b)"] --> N1{"比较 VersionNumbers<br/>逐位 int"}
  N1 -->|"不等"| R1["返回结果 ✅"]
  N1 -->|"相等"| S1{"后缀对比"}
  S1 -->|"一空一非空<br/>稳定版 > 预发布"| R2["返回结果 ✅"]
  S1 -->|"都非空<br/>按权重比"| W{"后缀权重<br/>+ 子版本号"}
  W -->|"不等"| R3["返回结果 ✅"]
  W -->|"相等"| T1{"PublicTime<br/>两者均非零?"}
  T1 -->|"是 且不等"| R4["晚者胜 ✅"]
  T1 -->|"否"| RAW["原始串字典序<br/>最终兜底 ✅"]
  style R1 fill:#f0fdf4,stroke:#16a34a
  style R2 fill:#f0fdf4,stroke:#16a34a
  style R3 fill:#f0fdf4,stroke:#16a34a
  style R4 fill:#f0fdf4,stroke:#16a34a
  style RAW fill:#fef2f2,stroke:#dc2626
:::

| # | 键 | 规则 | 源文件 |
|:--:|:--|:--|:--|
| 1 | **VersionNumbers** | 从左到右逐位 `int` 比较；共享位全部相等时，**更长的更大** | `version_numbers.go` |
| 2 | **Suffix** | 稳定版（无后缀）**大于**预发布版（有后缀）；两个后缀按权重比 | `version_suffix.go` |
| 3 | **PublicTime** | 两者均非零时，晚者胜（仅在前两级相等时生效） | `version.go` |
| 4 | **Raw** | 最终兜底：原始字符串字典序 | `version.go` |

正是这套排序让 `1.0.0-alpha < 1.0.0-beta < 1.0.0-rc1 < 1.0.0 < 1.0.0-post1`，贴合真实发布阶梯而非 ASCII 序。

::: warning 关于长度
`VersionNumbers.CompareTo` 把 `1.2` 视为**小于** `1.2.0`（共享段相等时短者小）。实际含义：2 段版本会排在其 3 段兄弟之下。如果你指的就是那个发布版，请显式写 `1.2.0`。
:::

## 3. 后缀权重排序

源文件：`suffix_weight.go`

每个后缀（大小写不敏感，可有前导 `-`/`.`）与一张有序模式表匹配。命中的**权重**决定排名；若两者权重相同，再用尾部整数（"子版本号"，如 `-alpha1` 的 `1`）破平。

| 权重 | 后缀模式（示例） | 含义 |
|-------:|:--|:--|
| 50 | `dev`、`dev1`、`.dev.2` | 开发构建 |
| 60 | `snapshot`、`snapshot20201012…` | 快照 |
| 70 | `nightly` | 夜间构建 |
| 100 | `a`、`alpha`、`alpha1`、`.alpha.2` | Alpha |
| 200 | `b`、`beta`、`beta2` | Beta |
| 300 | `m`、`milestone`、`m1` | 里程碑 |
| 400 | `rc`、`rc1` | 候选发布 |
| 410 | `pre`、`pre1` | 预发布 |
| 420 | `cr`、`cr1` | RC 的 CR 变体 |
| 500 | `final`、`release`、`ga` | 正式版（与无后缀等权） |
| 600 | `sp`、`sp1` | 服务包 |
| 700 | `patch`、`patch1` | 补丁 |
| 800 | `post`、`post1` | 后发布（PEP 440） |

规则：
- **未知后缀排在已知后缀之后**：不匹配任何模式的后缀，权重低于任何已知预发布类型；彼此之间退化为字典序。
- `final`/`release`/`ga` 权重都是 500——排序上与无后缀等价，但 `IsStable()` 只在后缀**字面为空**时返回 true。`1.0.0-ga` 是正式权重，但按"空后缀"判稳定版时不算稳定。

## 4. 约束语法 —— 三层

源文件：`constraint.go`

约束表达式是三层语法，OR 拆逗号 AND，AND 拆单约束：

:::mermaid
flowchart TD
  U["Union (OR)<br/>'>=1.0.0,<2.0.0 || >=3.0.0'<br/>以 || 切分"] --> S1["Set₁ (AND)<br/>'>=1.0.0,<2.0.0'<br/>以 , 切分"]
  U --> S2["Set₂ (AND)<br/>'>=3.0.0'"]
  S1 --> C1["Single: >=1.0.0"]
  S1 --> C2["Single: <2.0.0"]
  S2 --> C3["Single: >=3.0.0"]
  style U fill:#eff6ff,stroke:#2563eb
  style S1 fill:#f0fdf4,stroke:#16a34a
  style S2 fill:#f0fdf4,stroke:#16a34a
:::

- `ConstraintUnion.Match(v)`：**任一** Set 命中即 true。
- `ConstraintSet.Match(v)`：**所有** Single 命中才 true。
- 单个 `Constraint` = 操作符 + 目标版本。

支持的操作符（比较一律走 `CompareTo`，因此后缀权重生效）：

| 操作符 | 名称 | base | v 命中条件 |
|:--:|:--|:--|:--|
| `=` | 等于 | `1.2.3` | `v == 1.2.3`（裸写即 `=`） |
| `!=` | 不等 | `1.2.3` | `v != 1.2.3` |
| `>` `<` `>=` `<=` | 范围比较 | `1.2.3` | 直接 `CompareTo` |
| `^` | caret | `^1.2.3` | `v >= 1.2.3` 且 `v < {首个非零段+1, 0…}` |
| `~` | tilde | `~1.2.3` | `v >= 1.2.3` 且 `v < {major, minor+1, 0…}` |
| `x`/`X`/`*` | 通配 | `1.x` | `v >= 1.0.0` 且 `v < {末位非零+1, 0…}` |

### 边界算法详解

**Caret `^`（兼容范围 —— "左起第一个非零位"进一）：**

| 表达式 | 等价范围 | 说明 |
|:--|:--|:--|
| `^1.2.3` | `>=1.2.3, <2.0.0` | 首个非零是 major（1），进一得 2.0.0 上界 |
| `^0.2.3` | `>=0.2.3, <0.3.0` | 首个非零是 minor（2），进一得 0.3.0 |
| `^0.0.3` | `>=0.0.3, <0.0.4` | 首个非零是 patch（3），进一得 0.0.4 |

**Tilde `~`（锁定到 minor）：**

| 表达式 | 等价范围 | 说明 |
|:--|:--|:--|
| `~1.2.3` | `>=1.2.3, <1.3.0` | minor+1 作上界 |
| `~1.2` | `>=1.2.0, <1.3.0` | patch 开放 |

**通配 `x`（进位最后一个指定位）：**

| 表达式 | 等价范围 |
|:--|:--|
| `1.x` | `>=1.0.0, <2.0.0` |
| `1.2.x` | `>=1.2.0, <1.3.0` |

## 5. 范围查询 —— 开/闭区间

源文件：`version_range.go`

`VersionRange{Low, High, LowInclusive, HighInclusive}` 通过四步边界检查判定归属：

- `Low == nil` → 无下界；`High == nil` → 无上界。
- 下界侧：`v < Low` 不通过；`v == Low && !LowInclusive` 不通过（开区间排除端点）。
- 上界侧对称。

| 构造方式 | 区间 | 含义 |
|:--|:--|:--|
| `NewClosedRange(1.0.0, 2.0.0)` | `[1.0.0, 2.0.0]` | 两端都含 |
| `NewOpenRange(1.0.0, 2.0.0)` | `(1.0.0, 2.0.0)` | 两端都不含 |
| `NewVersionRange(1.0.0, 2.0.0, true, false)` | `[1.0.0, 2.0.0)` | 左含右不含 |

`IsEmpty()` 检测退化区间：`Low > High`，或 `Low == High` 但至少一端为开。

## 6. 排序与分组 —— 两阶段、组感知

源文件：`sort.go`、`version_group.go`

`SortVersionSlice` **不是**朴素 `sort.Slice`，而是两阶段：

1. **分组**：按 `BuildGroupID()`（完整数字串）归组。
2. **排序组**：按组数字前缀（`VersionGroup.CompareTo`）排组，**组内再排序**，最后拼接。

收益：`1.10.0` 正确排在 `1.2.0` 之后（数值而非字符串序），同族版本聚在一起。

### 分组粒度 —— 三个 API 别混淆

| 函数 | 分组依据 | 键类型 | 示例桶 |
|:--|:--|:--|:--|
| `Group(versions)` | **完整数字串** | `map[string]*VersionGroup` | `1.2.3` 与 `1.2.4` 分属**不同**组 |
| `GroupByMajor(versions)` | **仅 major 段** | `map[int][]*Version` | `1.2.3`、`1.9.0` 都在组 `1` |
| `GroupByMinor(versions)` | **major.minor** | `map[string][]*Version` | `1.2.3`、`1.2.4` 在组 `"1.2"` |

`Group()` 是排序与范围查询内部用的；`GroupByMajor`/`GroupByMinor` 是便捷分桶。

## 7. 有序索引上的范围查询

源文件：`sorted_version_groups.go`

对大版本集做反复范围查询，先建一次 `SortedVersionGroups`：

```go
sg := versions.NewSortedVersionGroups(allVersions)   // 分组+排序+建索引，O(n log n)
start := tuple.NewTuple2(versions.NewVersion("1.0.0"), versions.ContainsPolicyYes)
end   := tuple.NewTuple2(versions.NewVersion("2.0.0"), versions.ContainsPolicyNo)
hits  := sg.QueryRange(start, end)                    // 跳索引 + 组遍历
```

它预排序所有组并构建 `groupID → 索引` 映射。`QueryRange` 经映射直接跳到起始组（跳过其下一切），随后逐组收集 `QueryRangeVersions` 直到越过结束组。每个 tuple 上的 `ContainsPolicy` 决定边界版本本身是否纳入（`Yes` 含 / `No` 不含）。这是 `version_range_query` MCP 工具的底层引擎，远比每次重新过滤全表便宜。

## 性能摘要

| 操作 | 复杂度 |
|:--|:--|
| 解析 | `O(n)`，n 为版本串长度 |
| 比较 | `O(m)`，m 为数字段数 |
| 排序 | `O(n log n)`，n 为列表长度 |
| 范围查询（有序索引） | O(组数) 扫描 + 索引跳跃 |

→ 想直接调用这些能力，进 [Go SDK API](./sdk)、[CLI 命令](./cli) 或 [MCP 工具](./mcp)。
