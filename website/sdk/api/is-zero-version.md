# IsZero

::: info 方法 · `Version`
```go
func (x *Version) IsZero() bool
```
:::

## 📖 说明

IsZero 判断版本是否为零值

零值版本是未初始化的 Version{} 结构体，其所有字段都是默认值。
与 IsValid()（检查是否有版本号数字）不同，IsZero 检查是否完全没有被设置。


#### 返回

- `bool`：如果是零值则返回 true


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
