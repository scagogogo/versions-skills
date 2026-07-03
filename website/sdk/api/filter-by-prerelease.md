# FilterByPrerelease

::: info 函数 · 根包
```go
func FilterByPrerelease(versions []*Version) []*Version
```
:::

## 📖 说明

FilterByPrerelease 返回所有预发布版本

预发布版本是指带有后缀的版本。等价于 Filter(versions, Version.IsPrerelease)。


#### 参数

- `versions`：版本对象列表


#### 返回

- `[]*Version`：所有预发布版本


---

::: details 源码位置
定义于 [`version_utils.go`](https://github.com/scagogogo/versions-skills/blob/main/version_utils.go)
:::
