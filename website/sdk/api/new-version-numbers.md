# NewVersionNumbers

::: info 函数 · 根包
```go
func NewVersionNumbers(versionNumbers []int) VersionNumbers
```
:::

## 📖 说明

NewVersionNumbers 创建一个新的 VersionNumbers 对象

该方法将整数切片转换为 VersionNumbers 类型，便于进行版本号相关操作。


#### 参数

- `versionNumbers`：表示版本号的整数切片，如 [1,2,3] 表示版本号 "1.2.3"


#### 返回

- `VersionNumbers`：一个新的 VersionNumbers 对象


```go
numbers := versions.NewVersionNumbers([]int{1, 2, 3})
```


---

::: details 源码位置
定义于 [`version_numbers.go`](https://github.com/scagogogo/versions-skills/blob/main/version_numbers.go)
:::
