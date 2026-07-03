# BumpMajor

::: info 方法 · `Version`
```go
func (x *Version) BumpMajor() *Version
```
:::

## 📖 说明

BumpMajor 返回一个主版本号递增的新版本对象

例如 1.2.3 → 2.0.0，后缀被清除。


---

::: details 源码位置
定义于 [`version_builder.go`](https://github.com/scagogogo/versions-skills/blob/main/version_builder.go)
:::
