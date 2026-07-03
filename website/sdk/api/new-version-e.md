# NewVersionE

::: info 函数 · 根包
```go
func NewVersionE(versionStr string) (*Version, error)
```
:::

## 📖 说明

NewVersionE 从版本字符串创建一个新的 Version 对象，并返回可能的错误

与 NewVersion 不同，该方法会在版本字符串格式不正确时返回错误。


#### 参数

- `versionStr`：要解析的版本号字符串


#### 返回

- `*Version`：解析后的 Version 对象，如果解析失败则为 nil
- `error`：如果版本号无效，则返回 ErrVersionInvalid 错误


```go
version, err := versions.NewVersionE("v1.2.3-beta1")
if err != nil {
log.Fatalf("无效的版本号: %v", err)
}
```


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
