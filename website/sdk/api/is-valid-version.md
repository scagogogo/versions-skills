# IsValid

::: info 方法 · `Version`
```go
func (v *Version) IsValid() bool
```
:::

## 📖 说明

IsValid 检查版本号是否有效

判断依据是版本号中是否包含数字部分。只有当解析到了版本号数字时才认为是有效的版本号。


#### 返回

- `bool`：如果版本号有效则返回 true，否则返回 false


```go
version := versions.NewVersion("not-a-version")
if !version.IsValid() {
fmt.Println("无效的版本号")
}
```


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
