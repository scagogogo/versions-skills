# Difference

::: info 函数 · 根包
```go
func Difference(a, b []*Version) []*Version
```
:::

## 📖 说明

Difference 返回在 a 中但不在 b 中的版本（差集）

根据 Raw 字段判断版本是否相同。返回的版本保持 a 中的原始顺序。


#### 参数

- `a`：版本对象列表
- `b`：要排除的版本对象列表


#### 返回

- `[]*Version`：差集版本列表


---

::: details 源码位置
定义于 [`version_utils.go`](https://github.com/scagogogo/versions-skills/blob/main/version_utils.go)
:::
