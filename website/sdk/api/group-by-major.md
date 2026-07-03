# GroupByMajor

::: info 函数 · 根包
```go
func GroupByMajor(versions []*Version) map[int][]*Version
```
:::

## 📖 说明

GroupByMajor 按主版本号分组

返回按主版本号分组的映射，键为主版本号，值为对应的版本组。


#### 参数

- `versions`：版本对象列表


#### 返回

- `map[int][]*Version`：按主版本号分组的结果


```go
list := versions.NewVersions("1.0.0", "1.1.0", "2.0.0", "2.1.0")
groups := versions.GroupByMajor(list)
for major, vs := range groups {
fmt.Printf("Major %d: %d versions\n", major, len(vs))
}
```


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
