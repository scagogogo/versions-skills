# Increment

::: info 方法 · `Version`
```go
func (x *Version) Increment(segment int) *Version
```
:::

## 📖 说明

Increment 按位置递增版本号数字段

与 BumpMajor/BumpMinor/BumpPatch 不同，Increment 可以递增任意位置的版本号段，
并且将更高位置的段重置为零。


#### 参数

- `segment`：版本号段的位置索引（0=主版本号，1=次版本号，2=修订版本号，...）


#### 返回

- `*Version`：递增后的新版本对象


```go
v := versions.NewVersion("1.2.3.4")
newV := v.Increment(2)  // 递增修订版本号
fmt.Println(newV.Raw)   // "1.2.4.0"
```


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
