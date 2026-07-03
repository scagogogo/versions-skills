# CompareTo

::: info 方法 · `VersionNumbers`
```go
func (x VersionNumbers) CompareTo(target []int) int
```
:::

## 📖 说明

CompareTo 比较两个版本号数字部分的大小

该方法按照版本号比较规则，从左到右逐位比较两个版本号数字部分的大小。
对于不同长度的版本号，首先比较共有部分，如果共有部分相等，则较长的版本号更大。


#### 参数

- `target`：要比较的目标版本号数字部分


#### 返回

- `int`：如果当前版本小于目标版本，返回负数；如果相等，返回0；如果大于，返回正数


```go
v1 := versions.NewVersionNumbers([]int{1, 0, 0})
v2 := versions.NewVersionNumbers([]int{1, 1, 0})

result := v1.CompareTo(v2)
if result < 0 {
    fmt.Println("v1 比 v2 旧")
} else if result > 0 {
    fmt.Println("v1 比 v2 新")
} else {
    fmt.Println("v1 和 v2 相等")
}
```


## 🔗 同类方法

- [`VersionNumbers.BuildGroupID`](/sdk/api/build-group-id-versionnumbers)
- [`VersionNumbers.Len`](/sdk/api/len-versionnumbers)
- [`VersionNumbers.At`](/sdk/api/at-versionnumbers)
- [`VersionNumbers.String`](/sdk/api/string-versionnumbers)
- [`VersionNumbers.Equals`](/sdk/api/equals-versionnumbers)


---

::: details 源码位置
定义于 [`version_numbers.go`](https://github.com/scagogogo/versions-skills/blob/main/version_numbers.go)
:::
