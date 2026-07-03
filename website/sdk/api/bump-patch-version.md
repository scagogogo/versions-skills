# BumpPatch

::: info 方法 · `Version`
```go
func (x *Version) BumpPatch() *Version
```
:::

## 📖 说明

BumpPatch 返回一个修订版本号递增的新版本对象

例如 1.2.3 → 1.2.4，后缀被清除。


---

::: details 源码位置
定义于 [`version_builder.go`](https://github.com/scagogogo/versions-skills/blob/main/version_builder.go)
:::
