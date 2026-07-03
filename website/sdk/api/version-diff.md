# VersionDiff

::: info 类型 · 根包
```go
type VersionDiff struct {
	// Major 主版本号差值（target - source）
	Major int

	// Minor 次版本号差值（target - source）
	Minor int

	// Patch 修订版本号差值（target - source）
	Patch int

	// RawFrom 源版本原始字符串
	RawFrom string

	// RawTo 目标版本原始字符串
	RawTo string
}
```
:::

## 📖 说明

VersionDiff 表示两个版本之间的差异

VersionDiff 包含各版本号段的差值，以及原始版本字符串。


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
