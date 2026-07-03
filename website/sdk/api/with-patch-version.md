# WithPatch

::: info 方法 · `Version`
```go
func (x *Version) WithPatch(patch int) *Version
```
:::

## 📖 说明

WithPatch 返回一个修改修订版本号的新版本对象

原版本对象不变，返回一个新对象，其修订版本号被替换为指定值。
前缀和后缀保持不变。


#### 参数

- `patch`：新的修订版本号


#### 返回

- `*Version`：修改修订版本号后的新版本对象


---

::: details 源码位置
定义于 [`version_clone.go`](https://github.com/scagogogo/versions-skills/blob/main/version_clone.go)
:::
