# MarshalJSON

::: info 方法 · `Version`
```go
func (x Version) MarshalJSON() ([]byte, error)
```
:::

## 📖 说明

MarshalJSON 实现 json.Marshaler 接口

将版本序列化为 JSON 字符串（双引号包裹的原始版本字符串），
而非默认的结构体 JSON。这使得版本在 JSON 上下文中表现为简单字符串。


#### 返回

- `[]byte`：JSON 编码的版本字符串
- `error`：始终为 nil


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
