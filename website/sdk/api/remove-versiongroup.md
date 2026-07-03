# Remove

::: info 方法 · `VersionGroup`
```go
func (x *VersionGroup) Remove(v *Version) bool
```
:::

## 📖 说明

Remove 从版本组中移除指定的版本

如果版本存在于组中，则删除并返回 true；否则返回 false。


#### 参数

- `v`：要移除的版本对象


#### 返回

- `bool`：如果版本存在并被移除则返回 true


---

::: details 源码位置
定义于 [`version_group.go`](https://github.com/scagogogo/versions-skills/blob/main/version_group.go)
:::
