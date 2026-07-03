# CompareTo

::: info 方法 · `VersionGroup`
```go
func (x *VersionGroup) CompareTo(target *VersionGroup) int
```
:::

## 📖 说明

CompareTo 比较两个版本组的大小

该方法通过比较版本组的数字部分来确定两个版本组的先后顺序。


#### 参数

- `target`：要比较的目标版本组


#### 返回

- `int`：如果当前版本组小于目标版本组，返回负数；如果相等，返回0；如果大于，返回正数


```go
group1 := versions.NewVersionGroup(versions.NewVersionNumbers([]int{1, 2}))
group2 := versions.NewVersionGroup(versions.NewVersionNumbers([]int{1, 3}))

if group1.CompareTo(group2) < 0 {
fmt.Println("group1 比 group2 旧")
}
```


---

::: details 源码位置
定义于 [`version_group.go`](https://github.com/scagogogo/versions-skills/blob/main/version_group.go)
:::
