# FilterByPrefix

::: info 函数 · 根包
```go
func FilterByPrefix(versions []*Version, prefix string) []*Version
```
:::

## 📖 说明

FilterByPrefix 过滤指定前缀的版本

返回所有前缀等于指定值的版本。


#### 参数

- `versions`：版本对象列表
- `prefix`：目标前缀字符串，如 "v"


#### 返回

- `[]*Version`：满足条件的版本列表


---

::: details 源码位置
定义于 [`version_utils.go`](https://github.com/scagogogo/versions-skills/blob/main/version_utils.go)
:::
