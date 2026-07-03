# Coerce

::: info 函数 · 根包
```go
func Coerce(s string) *Version
```
:::

## 📖 说明

Coerce 从任意字符串中提取版本号

Coerce 尝试在输入字符串中查找第一个符合版本号模式的子串。
如果找不到则返回一个无效的 Version 对象。


#### 参数

- `s`：可能包含版本号的字符串


#### 返回

- `*Version`：提取到的版本对象


```go
v := versions.Coerce("program-1.2.3-linux-amd64")
fmt.Println(v.Raw) // "1.2.3"

v2 := versions.Coerce("download/v2.0.0-beta.tar.gz")
fmt.Println(v2.Raw) // "2.0.0-beta"
```


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
