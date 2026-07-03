# Union

::: info 函数 · 根包
```go
func Union(a, b []*Version) []*Version
```
:::

## 📖 说明

Union 返回 a 和 b 中所有唯一版本的并集

根据 Raw 字段去重，保持 a 中元素的原始顺序，b 中不重复的元素追加到末尾。


#### 参数

- `a`：版本对象列表
- `b`：版本对象列表


#### 返回

- `[]*Version`：并集版本列表


---

::: details 源码位置
定义于 [`version_utils.go`](https://github.com/scagogogo/versions-skills/blob/main/version_utils.go)
:::
