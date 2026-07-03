# SortVersionGroupMap

::: info 函数 · 根包
```go
func SortVersionGroupMap(versionGroupMap map[string]*VersionGroup) []*VersionGroup
```
:::

## 📖 说明

SortVersionGroupMap 对版本组映射进行排序

该函数将版本组映射转换为切片，并对切片中的版本组进行排序。


#### 参数

- `versionGroupMap`：版本组映射，键为版本组ID，值为版本组对象


#### 返回

- `[]*VersionGroup`：排序后的版本组切片


```go
groupMap := Group(versions)
sortedGroups := SortVersionGroupMap(groupMap)
for _, group := range sortedGroups {
    fmt.Printf("组: %s, 版本数: %d\n", group.GroupID, len(group.Versions))
}
```


---

::: details 源码位置
定义于 [`sort.go`](https://github.com/scagogogo/versions-skills/blob/main/sort.go)
:::
