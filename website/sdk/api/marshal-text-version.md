# MarshalText

::: info 方法 · `Version`
```go
func (x Version) MarshalText() ([]byte, error)
```
:::

## 📖 说明

MarshalText 实现 encoding.TextMarshaler 接口

将版本序列化为原始版本字符串的字节切片。
这使得 Version 可以被 encoding/json、toml、yaml 等序列化格式自动处理。


#### 返回

- `[]byte`：原始版本字符串的字节切片
- `error`：始终为 nil


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
