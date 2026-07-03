# WithSuffix

::: info 方法 · `Version`
```go
func (x *Version) WithSuffix(suffix string) *Version
```
:::

## 📖 说明

WithSuffix 返回一个修改后缀的新版本对象

原版本对象不变，返回一个新对象，其后缀被替换为指定值。


#### 参数

- `suffix`：新的后缀字符串，如 "-beta1"


#### 返回

- `*Version`：修改后缀后的新版本对象


```go
v := versions.NewVersion("1.2.3")
newV := v.WithSuffix("-rc1")
// newV.Raw == "1.2.3-rc1"
```


---

::: details 源码位置
定义于 [`version_clone.go`](https://github.com/scagogogo/versions-skills/blob/main/version_clone.go)
:::
