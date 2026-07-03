# SubVersion

::: info 方法 · `Version`
```go
func (x *Version) SubVersion() int
```
:::

## 📖 说明

SubVersion 返回后缀中的子版本号

例如 "-beta2" 返回 2，"-rc1" 返回 1。如果后缀中没有数字则返回 0。
该方法将内部使用的 extractSubVersion 函数暴露为公开 API。


#### 返回

- `int`：后缀中的子版本号数字


```go
v := versions.NewVersion("1.0.0-beta2")
fmt.Println(v.SubVersion()) // 输出: 2
```


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
