# At

::: info 方法 · `SortedVersionGroups`
```go
func (x *SortedVersionGroups) At(index int) *VersionGroup
```
:::

## 📖 说明

At 根据索引获取版本组

返回排序后指定位置的版本组。如果索引越界则返回 nil。


#### 参数

- `index`：从 0 开始的索引位置


#### 返回

- `*VersionGroup`：对应的版本组，越界则返回 nil


---

::: details 源码位置
定义于 [`sorted_version_groups.go`](https://github.com/scagogogo/versions-skills/blob/main/sorted_version_groups.go)
:::
