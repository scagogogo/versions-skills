# FilterByStable

::: info 函数 · 根包
```go
func FilterByStable(versions []*Version) []*Version
```
:::

## 📖 说明

FilterByStable 返回所有稳定版本

稳定版本是指不带任何后缀的版本。等价于 Filter(versions, Version.IsStable)。


#### 参数

- `versions`：版本对象列表


#### 返回

- `[]*Version`：所有稳定版本


---

::: details 源码位置
定义于 [`version_utils.go`](https://github.com/scagogogo/versions-skills/blob/main/version_utils.go)
:::
