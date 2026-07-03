# NewVersionRange

::: info 函数 · 根包
```go
func NewVersionRange(low, high *Version, lowInclusive, highInclusive bool) *VersionRange
```
:::

## 📖 说明

NewVersionRange 创建一个新的版本范围


#### 参数

- `low`：下界版本，nil 表示无下界
- `high`：上界版本，nil 表示无上界
- `lowInclusive`：下界是否为闭区间
- `highInclusive`：上界是否为闭区间


#### 返回

- `*VersionRange`：新的版本范围对象


```go
r := versions.NewVersionRange(
versions.NewVersion("1.0.0"),
versions.NewVersion("2.0.0"),
true,  // [1.0.0
false, // 2.0.0)
)
```


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
