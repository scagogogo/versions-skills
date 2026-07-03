# WithPrefix

::: info 方法 · `Version`
```go
func (x *Version) WithPrefix(prefix string) *Version
```
:::

## 📖 说明

WithPrefix 返回一个修改前缀的新版本对象

原版本对象不变，返回一个新对象，其前缀被替换为指定值。


#### 参数

- `prefix`：新的前缀字符串


#### 返回

- `*Version`：修改前缀后的新版本对象


```go
v := versions.NewVersion("1.2.3")
newV := v.WithPrefix("v")
// newV.Raw == "v1.2.3"
```


---

::: details 源码位置
定义于 [`version_clone.go`](https://github.com/scagogogo/versions-skills/blob/main/version_clone.go)
:::
