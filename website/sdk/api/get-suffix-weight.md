# GetSuffixWeight

::: info 函数 · 根包
```go
func GetSuffixWeight(suffix string) SuffixWeight
```
:::

## 📖 说明

GetSuffixWeight 获取后缀的语义权重

根据后缀字符串匹配预定义的权重规则，返回对应的语义权重。
如果没有匹配到任何规则，返回 SuffixWeightUnknown。


#### 参数

- `suffix`：版本后缀字符串，如 "-alpha1"、"-beta2"、"-rc1"


#### 返回

- `SuffixWeight`：后缀的语义权重值


---

::: details 源码位置
定义于 [`suffix_weight.go`](https://github.com/scagogogo/versions-skills/blob/main/suffix_weight.go)
:::
