# GroupByMinor

::: info 函数 · 根包
```go
func GroupByMinor(versions []*Version) map[string][]*Version
```
:::

## 📖 说明

GroupByMinor 按次版本号分组

返回按"主版本号.次版本号"分组的映射。


#### 参数

- `versions`：版本对象列表


#### 返回

- `map[string][]*Version`：按主.次版本号分组的结果，键如 "1.2"


```go
list := versions.NewVersions("1.0.0", "1.0.1", "1.1.0", "2.0.0")
groups := versions.GroupByMinor(list)
for key, vs := range groups {
    fmt.Printf("Minor %s: %d versions\n", key, len(vs))
}
```


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
