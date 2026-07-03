# FilterByPatch

::: info 函数 · 根包
```go
func FilterByPatch(versions []*Version, patch int) []*Version
```
:::

## 📖 说明

FilterByPatch 过滤指定修订版本号的版本

返回所有修订版本号等于指定值的版本。


#### 参数

- `versions`：版本对象列表
- `patch`：目标修订版本号


#### 返回

- `[]*Version`：满足条件的版本列表


---

::: details 源码位置
定义于 [`version_utils.go`](https://github.com/scagogogo/versions-skills/blob/main/version_utils.go)
:::
