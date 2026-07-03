# UnmarshalJSON

::: info 方法 · `Version`
```go
func (x *Version) UnmarshalJSON(data []byte) error
```
:::

## 📖 说明

UnmarshalJSON 实现 json.Unmarshaler 接口

从 JSON 字符串反序列化版本对象。


#### 参数

- `data`：JSON 编码的版本字符串


#### 返回

- `error`：如果版本无效则返回错误


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
