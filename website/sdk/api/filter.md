# Filter

::: info 函数 · 根包
```go
func Filter(versions []*Version, predicate func(*Version) bool) []*Version
```
:::

## 📖 说明

Filter 根据谓词函数过滤版本列表

返回所有满足谓词条件的版本。


#### 参数

- `versions`：版本对象列表
- `predicate`：过滤谓词函数，返回 true 表示保留该版本


#### 返回

- `[]*Version`：满足条件的版本列表


---

::: details 源码位置
定义于 [`version_utils.go`](https://github.com/scagogogo/versions-skills/blob/main/version_utils.go)
:::
