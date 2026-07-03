# Contains

::: info 方法 · `VersionGroup`
```go
func (x *VersionGroup) Contains(v *Version) bool
```
:::

## 📖 说明

Contains 判断本版本组中是否包含给定的版本

该方法检查指定的版本是否已存在于版本组中。


#### 参数

- `v`：要检查的版本对象


#### 返回

- `bool`：如果版本存在于组中则返回 true，否则返回 false


```go
group := versions.NewVersionGroup(versions.NewVersionNumbers([]int{1, 2}))
version := versions.NewVersion("1.2.3")
if !group.Contains(version) {
    group.Add(version)
}
```


---

::: details 源码位置
定义于 [`version_group.go`](https://github.com/scagogogo/versions-skills/blob/main/version_group.go)
:::
