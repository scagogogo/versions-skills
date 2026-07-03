# Min

::: info 函数 · 根包
```go
func Min(versions []*Version) *Version
```
:::

## 📖 说明

Min 从版本列表中找到最小的版本

如果列表为空则返回 nil。


#### 参数

- `versions`：版本对象列表


#### 返回

- `*Version`：最小的版本对象，列表为空时返回 nil


---

::: details 源码位置
定义于 [`version_utils.go`](https://github.com/scagogogo/versions-skills/blob/main/version_utils.go)
:::
