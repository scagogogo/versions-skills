# UnmarshalText

::: info 方法 · `Version`
```go
func (x *Version) UnmarshalText(text []byte) error
```
:::

## 📖 说明

UnmarshalText 实现 encoding.TextUnmarshaler 接口

从字节切片反序列化版本对象。解析给定的版本字符串，
如果版本无效则返回 ErrVersionInvalid。


#### 参数

- `text`：版本字符串的字节切片


#### 返回

- `error`：如果版本无效则返回 ErrVersionInvalid


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
