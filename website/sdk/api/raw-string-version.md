# RawString

::: info 方法 · `Version`
```go
func (x *Version) RawString() string
```
:::

## 📖 说明

RawString 返回版本的原始字符串表示

与 String()（返回 JSON 格式）不同，RawString() 返回解析前的原始版本字符串，
如 "v1.2.3-beta1"。这是获取版本字符串最直接的方式。


#### 返回

- `string`：原始版本字符串


```go
v := versions.NewVersion("v1.2.3-beta1")
fmt.Println(v.RawString()) // 输出: v1.2.3-beta1
```


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
