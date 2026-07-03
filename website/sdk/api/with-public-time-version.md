# WithPublicTime

::: info 方法 · `Version`
```go
func (x *Version) WithPublicTime(t time.Time) *Version
```
:::

## 📖 说明

WithPublicTime 返回一个修改发布时间的新版本对象

原版本对象不变，返回一个新对象，其发布时间被替换为指定值。


#### 参数

- `t`：新的发布时间


#### 返回

- `*Version`：修改发布时间后的新版本对象


---

::: details 源码位置
定义于 [`version_clone.go`](https://github.com/scagogogo/versions-skills/blob/main/version_clone.go)
:::
