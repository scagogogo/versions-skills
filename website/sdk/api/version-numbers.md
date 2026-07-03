# VersionNumbers

::: info 类型 · 根包
```go
type VersionNumbers []int
```
:::

## 📖 说明

VersionNumbers 表示版本号中的数字部分

VersionNumbers 是一个整数切片类型，用于表示和操作版本号的数字部分。
例如对于版本号 "v1.2.3-beta1"，其 VersionNumbers 为 [1,2,3]。
它实现了 Comparable 接口，支持版本号数字部分的比较操作。


```go
// 创建一个版本号数字部分
numbers := versions.NewVersionNumbers([]int{1, 2, 3})

// 获取版本组ID
groupID := numbers.BuildGroupID() // 返回 "1.2.3"

// 比较两个版本号数字部分
other := versions.NewVersionNumbers([]int{1, 3, 0})
if numbers.CompareTo(other) < 0 {
fmt.Println("1.2.3 比 1.3.0 旧")
}
```


---

::: details 源码位置
定义于 [`version_numbers.go`](https://github.com/scagogogo/versions-skills/blob/main/version_numbers.go)
:::
