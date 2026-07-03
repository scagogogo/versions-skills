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

权重规则（从低到高）:
- dev/snapshot < alpha/a < beta/b < milestone/m < rc/cr/pre < 正式版(无后缀)


```go
weight := GetSuffixWeight("-alpha1")
fmt.Println(weight) // 100
```


---

::: details 源码位置
定义于 [`suffix_weight.go`](https://github.com/scagogogo/versions-skills/blob/main/suffix_weight.go)
:::
