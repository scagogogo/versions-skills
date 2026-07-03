# SortVersionSlice

::: info 函数 · 根包
```go
func SortVersionSlice(versions []*Version) []*Version
```
:::

## 📖 说明

SortVersionSlice 对版本号对象数组进行排序

该函数实现了版本号的分组排序算法：
1. 首先将版本号按照主版本号分组
2. 对版本组进行排序
3. 在每个组内对版本号进行排序
4. 最后将所有组中的版本号按顺序合并


#### 参数

- `versions`：待排序的版本号对象数组


#### 返回

- `[]*Version`：排序后的版本号对象数组


```go
versions := versions.NewVersions("1.2.0", "1.0.0", "1.10.0", "2.0.0")
sorted := versions.SortVersionSlice(versions)
for _, v := range sorted {
    fmt.Println(v.Raw)
}
```


---

::: details 源码位置
定义于 [`sort.go`](https://github.com/scagogogo/versions-skills/blob/main/sort.go)
:::
