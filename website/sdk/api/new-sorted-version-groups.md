# NewSortedVersionGroups

::: info 函数 · 根包
```go
func NewSortedVersionGroups(versions []*Version) *SortedVersionGroups
```
:::

## 📖 说明

NewSortedVersionGroups 为版本号创建有序的分组

该方法接收一个版本对象数组，将其分组并排序，返回一个包含已排序版本组的SortedVersionGroups对象。
处理流程:
1. 首先将版本按照其数字部分分组
2. 然后对所有分组进行排序
3. 最后构建组ID到索引的映射，用于快速查找


#### 参数

- `versions`：需要分组和排序的版本对象数组


#### 返回

- `*SortedVersionGroups`：包含已排序版本组的对象


```go
// 创建版本对象数组
allVersions := versions.NewVersions("1.0.0", "1.1.0", "1.2.0", "2.0.0", "2.1.0")

// 创建已排序的版本组
sortedGroups := versions.NewSortedVersionGroups(allVersions)
```


---

::: details 源码位置
定义于 [`sorted_version_groups.go`](https://github.com/scagogogo/versions-skills/blob/main/sorted_version_groups.go)
:::
