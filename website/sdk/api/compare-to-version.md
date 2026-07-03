# CompareTo

::: info 方法 · `Version`
```go
func (x *Version) CompareTo(target *Version) int
```
:::

## 📖 说明

CompareTo 比较两个版本号

该方法按以下顺序比较两个版本号：
1. 首先比较主版本号数字部分
2. 其次比较发布时间
3. 然后比较后缀
4. 最后比较原始版本号字符串


#### 参数

- `target`：要比较的目标版本对象


#### 返回

- `int`：如果当前版本小于目标版本，返回-1；如果相等，返回0；如果大于，返回1


```go
v1 := versions.NewVersion("1.0.0")
v2 := versions.NewVersion("1.1.0")

switch v1.CompareTo(v2) {
case -1:
fmt.Println("v1 < v2")
case 0:
fmt.Println("v1 = v2")
case 1:
fmt.Println("v1 > v2")
}
```


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
