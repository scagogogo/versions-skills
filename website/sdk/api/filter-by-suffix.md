# FilterBySuffix

::: info 函数 · 根包
```go
func FilterBySuffix(versions []*Version, suffix string) []*Version
```
:::

## 📖 说明

FilterBySuffix 过滤指定后缀的版本

返回所有后缀字符串等于指定值的版本。


#### 参数

- `versions`：版本对象列表
- `suffix`：目标后缀字符串，如 "-beta"


#### 返回

- `[]*Version`：满足条件的版本列表


---

::: details 源码位置
定义于 [`version_utils.go`](https://github.com/scagogogo/versions-skills/blob/main/version_utils.go)
:::
