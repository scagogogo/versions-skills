# SuffixWeight

::: info 类型 · 根包
```go
type SuffixWeight int
```
:::

## 📖 说明

SuffixWeight 表示版本后缀的语义权重

SuffixWeight 用于在版本比较时为不同类型的后缀分配语义化的权重值，
使得预发布版本的排序符合实际发布顺序，而非简单的字典序。

### 完整权重表（从低到高）

| 权重值 | 常量 | 匹配后缀 | 含义 |
|--------|------|----------|------|
| 0 | `SuffixWeightUnknown` | （未匹配） | 未知后缀 |
| 50 | `SuffixWeightDev` | `dev`、`dev.1` | 开发版 |
| 60 | `SuffixWeightSnapshot` | `snapshot`、`SNAPSHOT` | 快照版（Maven 生态） |
| 70 | `SuffixWeightNightly` | `nightly` | 夜间构建版 |
| 100 | `SuffixWeightAlpha` | `alpha`、`a1` | Alpha 版 |
| 200 | `SuffixWeightBeta` | `beta`、`b2` | Beta 版 |
| 300 | `SuffixWeightMilestone` | `milestone`、`m1` | 里程碑版 |
| 400 | `SuffixWeightRC` | `rc1`、`RC2` | 候选发布版 |
| 410 | `SuffixWeightPre` | `pre1` | 预发布版 |
| 420 | `SuffixWeightCR` | `cr1` | 候选发布版（CR 变体） |
| 500 | `SuffixWeightFinal` / `Release` / `GA` | `final`、`release`、`ga` | 正式版（三个常量同值） |
| 600 | `SuffixWeightSP` | `sp1` | 服务包（高于正式版） |
| 700 | `SuffixWeightPatch` | `patch1` | 补丁版（高于正式版） |
| 800 | `SuffixWeightPost` | `post1` | Post 发布版（PEP 440，高于正式版） |

::: warning 反直觉：sp/patch/post 高于正式版
注意 `sp(600)`、`patch(700)`、`post(800)` 的权重**大于**正式版 `final/release/ga(500)`。这是因为在语义上，`1.0.0-sp1` 是在 `1.0.0` 正式版之后发布的服务包，`1.0.0-post1` 是 PEP 440 中正式版之后的修订。因此 `1.0.0-post1 > 1.0.0 > 1.0.0-rc1`。

源码 doc 注释只列了 `dev/snapshot < alpha < beta < milestone < rc < 正式版`，遗漏了 sp/patch/post 三档，**以本表为准**。
:::

```go
weight := GetSuffixWeight("-alpha1")
fmt.Println(weight) // 100

fmt.Println(GetSuffixWeight("-post1") > GetSuffixWeight("")) // true，post 高于正式版
```


---

::: details 源码位置
定义于 [`suffix_weight.go`](https://github.com/scagogogo/versions-skills/blob/main/suffix_weight.go)
:::
