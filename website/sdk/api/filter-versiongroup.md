# Filter

::: info 方法 · `VersionGroup`
```go
func (x *VersionGroup) Filter(predicate func(*Version) bool) []*Version
```
:::

## 📖 说明

Filter 根据谓词函数过滤组内版本

返回组内所有满足谓词条件的版本。


#### 参数

- `predicate`：过滤谓词函数


#### 返回

- `[]*Version`：满足条件的版本列表


---

::: details 源码位置
定义于 [`version_group.go`](https://github.com/scagogogo/versions-skills/blob/main/version_group.go)
:::
