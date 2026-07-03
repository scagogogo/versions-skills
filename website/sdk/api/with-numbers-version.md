# WithNumbers

::: info 方法 · `Version`
```go
func (x *Version) WithNumbers(numbers []int) *Version
```
:::

## 📖 说明

WithNumbers 返回一个修改版本号数字部分的新版本对象

原版本对象不变，返回一个新对象，其版本号数字部分被替换为指定值。
前缀和后缀保持不变。


#### 参数

- `numbers`：新的版本号数字部分


#### 返回

- `*Version`：修改版本号后的新版本对象


```go
v := versions.NewVersion("1.2.3")
newV := v.WithNumbers([]int{2, 0, 0})
// newV.Raw == "2.0.0"
```


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
