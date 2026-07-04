# SortVersionStringSlice

::: info 函数 · 根包
```go
func SortVersionStringSlice(versionStringSlice []string) []string
```
:::

## 📖 说明

SortVersionStringSlice 对字符串形式的版本数组进行排序

该函数接收一个字符串形式的版本号数组，将其解析为 Version 对象进行排序，
然后将排序结果转换回字符串数组返回。排序遵循版本号的自然排序规则。


#### 参数

- `versionStringSlice`：待排序的版本号字符串数组


#### 返回

- `[]string`：排序后的版本号字符串数组


```go
versions := []string{"1.2.0", "1.0.0", "1.10.0", "2.0.0"}
sorted := versions.SortVersionStringSlice(vs)
// 结果: ["1.0.0", "1.2.0", "1.10.0", "2.0.0"]
```


---

::: details 源码位置
定义于 [`sort.go`](https://github.com/scagogogo/versions-skills/blob/main/sort.go)
:::
