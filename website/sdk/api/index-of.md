# IndexOf

::: info 函数 · 根包
```go
func IndexOf(versions []*Version, target *Version) int
```
:::

## 📖 说明

IndexOf 查找版本在列表中的位置

根据 Raw 字段匹配，返回第一次出现的索引。如果未找到则返回 -1。


#### 参数

- `versions`：版本对象列表
- `target`：目标版本对象


#### 返回

- `int`：版本在列表中的索引，未找到返回 -1


```go
list := versions.NewVersions("1.0.0", "1.1.0", "2.0.0")
idx := versions.IndexOf(list, versions.NewVersion("1.1.0"))
fmt.Println(idx) // 1
```


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
