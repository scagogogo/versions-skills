# NewClosedRange

::: info 函数 · 根包
```go
func NewClosedRange(low, high *Version) *VersionRange
```
:::

## 📖 说明

NewClosedRange 创建一个闭区间版本范围 [low, high]


```go
r := versions.NewClosedRange(versions.NewVersion("1.0.0"), versions.NewVersion("2.0.0"))
```


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
