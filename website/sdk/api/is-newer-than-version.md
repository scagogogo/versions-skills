# IsNewerThan

::: info 方法 · `Version`
```go
func (x *Version) IsNewerThan(target *Version) bool
```
:::

## 📖 说明

IsNewerThan 判断当前版本是否比目标版本更新

等价于 CompareTo(target) > 0，但语义更清晰。


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
