# Patch

::: info 方法 · `Version`
```go
func (x *Version) Patch() int
```
:::

## 📖 说明

Patch 返回修订版本号

如果版本号数字部分少于3位则返回 0。


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
