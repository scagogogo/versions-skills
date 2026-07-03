# SuffixWeight

::: info 方法 · `Version`
```go
func (x *Version) SuffixWeight() SuffixWeight
```
:::

## 📖 说明

SuffixWeight 返回版本后缀的语义权重

等价于 GetSuffixWeight(string(x.Suffix))，但作为方法调用更方便。


#### 返回

- `SuffixWeight`：后缀的语义权重值


```go
v := versions.NewVersion("1.0.0-beta")
w := v.SuffixWeight()
fmt.Println(w == versions.SuffixWeightBeta) // true
```


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
