# Clone

::: info 方法 · `Version`
```go
func (x *Version) Clone() *Version
```
:::

## 📖 说明

Clone 创建版本的深拷贝

返回一个与原版本完全相同的新 Version 对象，修改拷贝不会影响原版本。
对于不可变的 Version 对象，Clone 主要用于与 With* 方法配合使用。


#### 返回

- `*Version`：版本的深拷贝


```go
v1 := versions.NewVersion("1.2.3")
v2 := v1.Clone()
v2.Raw = "modified"
fmt.Println(v1.Raw) // 仍然是 "1.2.3"
```


---

::: details 源码位置
定义于 [`version_clone.go`](https://github.com/scagogogo/versions-skills/blob/main/version_clone.go)
:::
