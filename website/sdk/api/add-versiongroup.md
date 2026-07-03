# Add

::: info 方法 · `VersionGroup`
```go
func (x *VersionGroup) Add(v *Version) bool
```
:::

## 📖 说明

Add 把给定的版本添加到本版本组中

该方法将指定的版本添加到版本组中。如果版本已存在，则会被覆盖。


#### 参数

- `v`：要添加的版本对象


#### 返回

- `bool`：如果版本之前已存在于组中则返回 true，否则返回 false


```go
group := versions.NewVersionGroup(versions.NewVersionNumbers([]int{1, 2}))
version := versions.NewVersion("1.2.3")
exists := group.Add(version)
if !exists {
    fmt.Println("添加了新版本")
}
```


---

::: details 源码位置
定义于 [`version_group.go`](https://github.com/scagogogo/versions-skills/blob/main/version_group.go)
:::
