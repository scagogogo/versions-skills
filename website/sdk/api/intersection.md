# Intersection

::: info 函数 · 根包
```go
func Intersection(a, b []*Version) []*Version
```
:::

## 📖 说明

Intersection 返回同时存在于 a 和 b 中的版本（交集）

根据 Raw 字段判断版本是否相同。返回的版本保持 a 中的原始顺序。


#### 参数

- `a`：版本对象列表
- `b`：版本对象列表


#### 返回

- `[]*Version`：交集版本列表


---

::: details 源码位置
定义于 [`version_utils.go`](https://github.com/scagogogo/versions-skills/blob/main/version_utils.go)
:::
