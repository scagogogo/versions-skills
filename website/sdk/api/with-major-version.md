# WithMajor

::: info 方法 · `Version`
```go
func (x *Version) WithMajor(major int) *Version
```
:::

## 📖 说明

WithMajor 返回一个修改主版本号的新版本对象

原版本对象不变，返回一个新对象，其主版本号被替换为指定值。
后缀和前缀保持不变。


#### 参数

- `major`：新的主版本号


#### 返回

- `*Version`：修改主版本号后的新版本对象


---

::: details 源码位置
定义于 [`version_clone.go`](https://github.com/scagogogo/versions-skills/blob/main/version_clone.go)
:::
