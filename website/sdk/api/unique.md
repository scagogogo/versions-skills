# Unique

::: info 函数 · 根包
```go
func Unique(versions []*Version) []*Version
```
:::

## 📖 说明

Unique 去除版本列表中的重复版本

根据 Raw 字段去重，保留第一次出现的版本。


---

::: details 源码位置
定义于 [`version_utils.go`](https://github.com/scagogogo/versions-skills/blob/main/version_utils.go)
:::
