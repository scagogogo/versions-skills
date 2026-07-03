# String

::: info 方法 · `SuffixWeight`
```go
func (w SuffixWeight) String() string
```
:::

## 📖 说明

String 返回后缀权重的可读名称

实现 fmt.Stringer 接口。


#### 返回

- `string`：权重的可读名称，如 "dev"、"alpha"、"beta"、"rc"、"unknown" 等


---

::: details 源码位置
定义于 [`suffix_weight.go`](https://github.com/scagogogo/versions-skills/blob/main/suffix_weight.go)
:::
