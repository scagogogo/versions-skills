# SortVersions

::: info 方法 · `VersionGroup`
```go
func (x *VersionGroup) SortVersions() []*Version
```
:::

## 📖 说明

SortVersions 对组下的所有版本进行排序返回

该方法返回版本组中所有版本的有序数组，排序遵循版本号的自然排序规则。


#### 返回

- `[]*Version`：排序后的版本数组


```go
group := versions.NewVersionGroupFromVersions(versions.NewVersions("1.2.2", "1.2.0", "1.2.1"))
sortedVersions := group.SortVersions()
// 结果顺序: ["1.2.0", "1.2.1", "1.2.2"]
```


---

::: details 源码位置
定义于 [`version_group.go`](https://github.com/scagogogo/versions-skills/blob/main/version_group.go)
:::
