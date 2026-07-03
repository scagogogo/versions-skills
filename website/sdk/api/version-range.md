# VersionRange

::: info 类型 · 根包
```go
type VersionRange struct {
	// Low 下界版本
	Low *Version

	// High 上界版本
	High *Version

	// LowInclusive 下界是否为闭区间（包含下界）
	LowInclusive bool

	// HighInclusive 上界是否为闭区间（包含上界）
	HighInclusive bool
}
```
:::

## 📖 说明

VersionRange 表示一个版本范围

VersionRange 是一个包含下界和上界的版本区间，支持开区间和闭区间。
它可以用于判断版本是否落在某个范围内，以及过滤版本列表。


```go
low := versions.NewVersion("1.0.0")
high := versions.NewVersion("2.0.0")
r := versions.NewVersionRange(low, high, true, true)
v := versions.NewVersion("1.5.0")
if r.Contains(v) {
    fmt.Println("1.5.0 in [1.0.0, 2.0.0]")
}
```


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
