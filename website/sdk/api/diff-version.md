# Diff

::: info 方法 · `Version`
```go
func (x *Version) Diff(target *Version) *VersionDiff
```
:::

## 📖 说明

Diff 计算两个版本之间的差异

返回一个 VersionDiff 结构体，包含各版本号段的差异。
如果目标版本为 nil，返回 nil。


#### 参数

- `target`：目标版本


#### 返回

- `*VersionDiff`：版本差异对象


```go
v1 := versions.NewVersion("1.2.3")
v2 := versions.NewVersion("2.0.0")
d := v1.Diff(v2)
fmt.Printf("major diff: %d\n", d.Major) // 1
```


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
