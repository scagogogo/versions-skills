# 算法详解

本页记录 AI Agent（或人类）在不读源码的情况下预测库行为所需的精确语义。每条规则都可在对应源文件中验证。

## 本页导览

| # | 主题 | 核心问题 |
|:--:|:--|:--|
| 1 | [解析](#_1-解析-——-三段式-元数据) | 字符串如何拆成 Prefix / Numbers / Suffix / Metadata |
| 2 | [比较](#_2-比较-——-四级优先级) | 两个版本谁更大，按什么顺序判定 |
| 3 | [后缀权重](#_3-后缀权重排序) | alpha / beta / rc / 正式版如何排序 |
| 4 | [约束语法](#_4-约束语法-——-三层) | `^` `~` `1.x` `\|\|` 如何展开成范围 |
| 5 | [范围查询](#_5-范围查询-——-开-闭区间) | `[1.0.0, 2.0.0)` 这类区间如何判定 |
| 6 | [排序与分组](#_6-排序与分组-——-两阶段、组感知) | 为何 `1.10.0` 能正确排在 `1.2.0` 后 |
| 7 | [有序索引范围查询](#_7-有序索引上的范围查询) | 大列表反复查询如何加速 |
| 8 | [不可变变更](#_8-不可变变更-——-builder-重建) | Bump / With 为何返回新对象 |
| 9 | [文件 I/O](#_9-文件-i-o-——-行式格式) | 版本列表文件的格式与坑 |
| 10 | [可视化](#_10-可视化-——-组感知文本树) | 文本树如何绘制版本层次 |
| 11 | [ContainsPolicy](#_11-containspolicy-——-边界纳入策略) | 边界版本是否纳入结果 |

## 1. 解析 —— 三段式 + 元数据

源文件：`parser.go`

每个版本字符串被拆为四个字段：

| 字段 | 含义 | 示例值（对 `v1.2.3-beta1+build.7`） |
|:--|:--|:--|
| `Prefix` | 非数字前导，如 `"v"`、`"release-"` | `"v"` |
| `VersionNumbers` | 整数段，解析后统一以 `.` 拼接 | `[1, 2, 3]` |
| `Suffix` | 数字部分之后的全部内容 | `"-beta1"` |
| `Metadata` | semver 构建元数据，`+` 之后 | `"build.7"` |

字符串上四段的边界：

```text
 v   1.2.3   -beta1   +build.7
 └─  └──┬──┘  └──┬──┘  └───┬───┘
Prefix  Numbers   Suffix   Metadata
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
  START["a.CompareTo(b)"] --> N1{"1️⃣ 比较 VersionNumbers<br/>逐位 int"}
  N1 -->|"不等"| DONE1["返回 ✅"]
  N1 -->|"相等"| S1{"2️⃣ 后缀对比"}
  S1 -->|"一空一非空<br/>稳定版 > 预发布"| DONE2["返回 ✅"]
  S1 -->|"都非空<br/>按权重比"| W{"3️⃣ 后缀权重<br/>+ 子版本号"}
  W -->|"不等"| DONE3["返回 ✅"]
  W -->|"相等"| T1{"4️⃣ PublicTime<br/>两者均非零?"}
  T1 -->|"是 且不等"| DONE4["晚者胜 ✅"]
  T1 -->|"否"| RAW["原始串字典序<br/>兜底"]
  style START fill:#eff6ff,stroke:#2563eb
  style DONE1 fill:#f0fdf4,stroke:#16a34a
  style DONE2 fill:#f0fdf4,stroke:#16a34a
  style DONE3 fill:#f0fdf4,stroke:#16a34a
  style DONE4 fill:#f0fdf4,stroke:#16a34a
  style RAW fill:#fefce8,stroke:#ca8a04
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

| 权重 | 后缀模式（正则示例） | 含义 |
|-------:|:--|:--|
| 0 | 不匹配任何模式 | 未知后缀（见下方规则） |
| 50 | `dev`、`dev1`、`.dev.2` | 开发构建 |
| 60 | `snapshot`、`snapshot20201012` | 快照 |
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

- **未知后缀权重为 0**：不匹配任何模式的后缀，权重 `SuffixWeightUnknown = 0`，低于任何已知预发布类型；彼此之间退化为字典序。所以一个拼写怪异的后缀会排在所有已知后缀之后。
- `final`/`release`/`ga` 权重都是 500——排序上与无后缀等价，但 `IsStable()` 只在后缀**字面为空**时返回 true。`1.0.0-ga` 是正式权重，但按"空后缀"判稳定版时不算稳定。
- 匹配对大小写不敏感：`-RC1` 与 `-rc1` 同权。

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

| 操作符 | 名称 | 示例 | `v` 命中条件 |
|:--:|:--|:--|:--|
| `=` | 等于 | `=1.2.3` | `v == 1.2.3`（裸写 `1.2.3` 即 `=`） |
| `!=` | 不等 | `!=1.2.3` | `v != 1.2.3` |
| `>` | 大于 | `>1.2.3` | `v > 1.2.3` |
| `<` | 小于 | `<2.0.0` | `v < 2.0.0` |
| `>=` | 大于等于 | `>=1.0.0` | `v >= 1.0.0` |
| `<=` | 小于等于 | `<=2.0.0` | `v <= 2.0.0` |
| `^` | caret（兼容） | `^1.2.3` | `v >= 1.2.3` 且 `v < 2.0.0`（首个非零段进一作上界） |
| `~` | tilde（锁 minor） | `~1.2.3` | `v >= 1.2.3` 且 `v < 1.3.0`（minor+1 作上界） |
| `x`/`X`/`*` | 通配 | `1.x` | `v >= 1.0.0` 且 `v < 2.0.0`（末位非零进一） |

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

`VersionRange` 由两个 `*Version` 端点（`Low` / `High`）和两个 `bool` 闭性标志（`LowInclusive` / `HighInclusive`）组成。`Contains(v)` 按四步边界检查判定归属：

1. **无界检测**：`Low == nil` → 无下界；`High == nil` → 无上界。两端皆 nil 则全包含。
2. **下界检查**：`v < Low` → 不通过；`v == Low && !LowInclusive` → 不通过（开区间排除端点）。
3. **上界检查**：与下界对称，`v > High` 或 `v == High && !HighInclusive` → 不通过。
4. **通过**：两端都未拦下 → 命中。

构造器：

| 构造方式 | 区间 | 含义 |
|:--|:--|:--|
| `NewClosedRange(low, high)` | `[low, high]` | 两端都含 |
| `NewOpenRange(low, high)` | `(low, high)` | 两端都不含 |
| `NewVersionRange(low, high, true, false)` | `[low, high)` | 左含右不含 |

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

## 8. 不可变变更 —— Builder 重建

源文件：`version_builder.go`、`version_clone.go`

`Version` 设计为**不可变**：所有 `Bump*` / `With*` 都返回新对象，原对象绝不被修改。它们不是简单的字段赋值，而是走 **Builder 重建** 流程：

:::mermaid
flowchart LR
  SRC["原版本 1.2.3"] --> DECIDE{"操作类型"}
  DECIDE -->|"Bump*"| BUMP["数字段 +1<br/>低位置零<br/>后缀清除"]
  DECIDE -->|"With*"| WITH["替换指定段<br/>保留其余"]
  BUMP --> B["NewVersionBuilder"]
  WITH --> B
  B --> REBUILD["buildRawString<br/>prefix+numbers+suffix"]
  REBUILD --> PARSE["NewVersion(raw)<br/>重新解析"]
  PARSE --> OUT["新 Version 对象"]
  style SRC fill:#eff6ff,stroke:#2563eb
  style OUT fill:#f0fdf4,stroke:#16a34a
:::

关键细节：

1. **Bump 清除后缀**：`BumpMajor` / `BumpMinor` / `BumpPatch` 都通过 Builder 重建，**只设数字段**，不传 Suffix —— 所以 `v := NewVersion("1.2.3-beta1"); v.BumpPatch()` 得到 `1.2.4`（后缀没了），不是 `1.2.4-beta1`。这符合"发布新版本时预发布标记应脱落"的语义。
2. **Bump 的零填充**：`BumpMajor` 显式设 `Minor(0).Patch(0)`，保证 `1.2.3 → 2.0.0` 而非 `2.2.3`。空数字段的 Version（无效解析结果）Bump 会退化为 `"1"` / `"0.1"` / `"0.0.1"`。
3. **With\* 保留前缀与后缀**：`WithMajor(5)` 只改 `numbers[0]`，Prefix、Suffix、Metadata 原样带入 Builder。`WithMinor` / `WithPatch` 在数字段不足时会**自动补零扩长**（`WithPatch` 于 2 段版本上会扩成 3 段）。
4. **WithPublicTime 走 Clone**：它不重建 raw 字符串（发布时间不影响 Raw），而是 `Clone()` 后改 `PublicTime` 字段。
5. **Clone 是深拷贝**：`VersionNumbers` 是 `[]int`，`Clone` 会 `make + copy` 新切片，所以改拷贝的数字段不影响原对象。

::: warning Bump 会丢弃 Metadata
`WithPrefix` / `WithSuffix` / `WithMajor` 等在 Builder 构造后**手动回填** `v.Metadata = x.Metadata`，所以 metadata 保留。但 `Bump*` 系列没有回填——对 `NewVersion("1.2.3+build.7").BumpPatch()` 得到的 `1.2.4` **会丢失** `+build.7`。如果你的工作流依赖 metadata，Bump 后需用 `WithMetadata(...)` 补回。
:::

`VersionBuilder` 本身是流式 API，也可独立用于程序化构造版本：

```go
v := versions.NewVersionBuilder().
    Prefix("v").
    Major(1).Minor(2).Patch(3).
    Suffix("-beta1").
    Build()                 // v.Raw == "v1.2.3-beta1"
```

`Build()` 内部先 `buildRawString()` 拼出 `"v1.2.3-beta1"`，再交给 `NewVersion` **重新解析**——所以 Builder 产物与直接解析等价字符串的结果一致。

## 9. 文件 I/O —— 行式格式

源文件：`file.go`

versions-skills 用一个极简的**行式文本格式**承载版本列表，便于 diff、人工编辑、管道处理：

```text
1.1.28
1.1.29
1.1.31.sec01     ← Maven 风格后缀，照常解析
# 这一行会被当作版本号解析，不是注释！
```

:::mermaid
flowchart LR
  F["文件 versions.txt"] --> READ["os.ReadFile<br/>整文件读入"]
  READ --> SPLIT["按 \\n 切分"]
  SPLIT --> LOOP{"逐行处理"}
  LOOP -->|"TrimSpace 后为空"| SKIP["跳过"]
  LOOP -->|"非空"| PARSE["NewVersionStringParser<br/>.Parse()"]
  PARSE --> APPEND["追加到结果"]
  APPEND --> LOOP
  style SKIP fill:#fef2f2,stroke:#dc2626
  style PARSE fill:#fff7ed,stroke:#ea580c
:::

**读**（`ReadVersionsFromFile` / `ReadVersionsFromReader`）：

1. 整文件读入，按 `\n` 切分。
2. 每行 `TrimSpace`；**空行跳过**。
3. 非空行**原样交给解析器**——没有注释剥离、没有字段切分。

::: warning 三个易踩的坑
1. **没有 `#` 注释支持**：`#` 开头的行会被当作版本字符串解析。由于它不含数字，走"纯字母快捷路径"成为**无效版本**（VersionNumbers 为空）。它不会报错，但会污染结果集——用前请自行预处理注释行。
2. **带计数前缀的行解析失真**：`maven_versions_count.txt` 里 `71270\t1.0.0` 这种行，解析器从第一个数字 `7` 开始读，得到 `numbers=[71270]`、`suffix="\t1.0.0"`——**不是** `1.0.0`。要处理这种格式，先用 `Coerce` 或自行切 tab 取版本列。
3. **CRLF**：`TrimSpace` 会吃掉 `\r`，所以 Windows 换行也能正确处理。
:::

**写**（`WriteVersionsToFile`）：

```go
func WriteVersionsToFile(versions []*Version, filepath string) error
```

- **写入前先排序**：内部调用 `SortVersionSlice`，保证输出有序（即使输入乱序）。
- **写 `Raw` 字段**：每行写版本的原始字符串，`v.Raw`，而非重建的规范形态。
- **无尾随换行**：行间用 `\n` 连接，最后一行后不补 `\n`。

::: tip 读 ↔ 写 的对称性破缺
`ReadVersionsFromFile` 解析每行得到 Version，其 `Raw` 等于该行原文；`WriteVersionsToFile` 写的也是 `Raw`。所以"读 → 写"往返保持字符串一致——**但仅在输入已是规范形态时**。若输入是 `v1.2.3`（带前缀），读出的 `Raw` 是 `v1.2.3`，写回仍是 `v1.2.3`。排序发生在往返之间，故乱序输入经"读 → 写"后变有序。
:::

`ReadVersionsStringFromFile` 是只读字符串、不解析的变体——当你只需要原始行、不想承担解析开销时用。

## 10. 可视化 —— 组感知文本树

源文件：`visualize.go`

"一图抵千言"在 CLI/MCP 场景下落地为**Unicode 文本树**。两个函数都建立在第 6 节的分组排序之上：

### VisualizeVersions —— 组内明细树

```text
版本总数: 6
版本组数: 2

┌─ 版本组: 1.2 (3个版本)
├── 1.2.0
├── 1.2.0-beta
└── 1.2.1 (发布时间: 2026-06-01)

┌─ 版本组: 2.0 (1个版本)
└── 2.0.0
```

算法：

1. `Group(versions)` 按完整数字串建组（同第 6 节）。
2. `NewSortedVersionGroups` 给组排序。
3. 逐组 `group.SortVersions()` 排组内版本。
4. `maxItems > 0` 时截断，末尾打印 `...还有 N 个版本未显示`。
5. 每个版本若 `PublicTime` 非零，附 `(发布时间: YYYY-MM-DD)`。

### VisualizeVersionGroups —— 主版本概览树

```text
版本总数: 6
版本组数: 2

├─ 1 (2个版本组, 共5个版本)
│  ├─ 1.2 (3个版本)
│  └─ 1.3 (2个版本)
└─ 2 (1个版本组, 共1个版本)
   └─ 2.0 (1个版本)
```

它**不调** `SortedVersionGroups`，而是自己 `sort.Strings(groupIDs)`（字符串序，注意：这对 `1.10` vs `1.2` 是字典序，**不**保证数值序——这是概览树的已知取舍）。然后按 groupID 的首段（`.` 之前）二次归并为"主版本 → 子组"两级树，用 `├─` / `└─` / `│ ` / `  ` 绘制缩进。

::: tip 两个函数的选择
- 想看**每个具体版本** → `VisualizeVersions`（组内明细，可截断）。
- 想看**版本库整体结构** → `VisualizeVersionGroups`（主版本概览，紧凑）。
- 两者都写到 `io.Writer`，可落地 stdout、文件、HTTP 响应、MCP 工具返回。
:::

## 11. ContainsPolicy —— 边界纳入策略

源文件：`contains_policy.go`

`ContainsPolicy` 是个三值枚举，控制范围查询时**边界版本本身是否纳入结果**：

| 值 | 常量 | 语义 | 在范围查询中的效果 |
|:--:|:--|:--|:--|
| 0 | `ContainsPolicyNone` | 未指定 | 等同 `Yes`（纳入边界） |
| 1 | `ContainsPolicyYes` | 纳入 | 闭端点，边界版本**保留** |
| 2 | `ContainsPolicyNo` | 排除 | 开端点，边界版本**剔除** |

### 与 VersionRange 的对应

第 5 节的 `VersionRange` 用 `LowInclusive` / `HighInclusive` 两个 `bool` 表达开闭；`SortedVersionGroups.QueryRange` 把同样的语义收纳进 `ContainsPolicy`，挂在起止 tuple 上：

```go
start := tuple.NewTuple2(versions.NewVersion("1.0.0"), versions.ContainsPolicyYes)  // 含 1.0.0
end   := tuple.NewTuple2(versions.NewVersion("2.0.0"), versions.ContainsPolicyNo)   // 不含 2.0.0
hits  := sg.QueryRange(start, end)   // 结果区间 = [1.0.0, 2.0.0)
```

两者是同一概念的不同载体：`bool` 适合构造单个区间，`ContainsPolicy` 三值枚举适合在有序索引上链式传递（`None` 表"这一端未由调用方指定，走默认纳入"）。

### 底层判定

`VersionGroup.QueryRangeVersions` 在逐版本收集时，对恰好等于端点的版本按策略分流（`version_group.go`）：

- `ContainsPolicyNone`、`ContainsPolicyYes` → 纳入边界版本；
- `ContainsPolicyNo` → 剔除边界版本。

`String()` 返回 `"none"` / `"yes"` / `"no"`，便于日志与 MCP 工具的入参校验。CLI 的 `range` 命令与 MCP 的 `version_range_query` 工具都通过 `--include-start/--include-end` 之类的开关映射到这两个枚举值。

## 性能摘要

| 操作 | 复杂度 | 说明 |
|:--|:--|:--|
| 解析 `Parse` | `O(n)` | n 为版本串长度 |
| 比较 `CompareTo` | `O(m)` | m 为数字段数 |
| 排序 `SortVersionSlice` | `O(n log n)` | n 为列表长度（含分组） |
| 范围查询 `QueryRange` | `O(g + k)` | g 为组数扫描、k 为命中数；经有序索引跳跃 |
| Bump / With | `O(m) + O(n)` | Builder 拼接 + 重新解析 |
| 文件读取 | `O(Σ 行长)` | 整文件读入 + 逐行解析 |
| 可视化 | `O(n log n)` | 内含分组排序 |

→ 想直接调用这些能力，进 [Go SDK API](/sdk/)、[CLI 命令](/cli/) 或 [MCP 工具](/mcp/)。
