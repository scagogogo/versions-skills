# CoerceE

::: info 函数 · 根包
```go
func CoerceE(s string) (*Version, error)
```
:::

## 📖 说明

CoerceE 从任意字符串中提取版本号，找不到则返回错误


#### 参数

- `s`：可能包含版本号的字符串


#### 返回

- `*Version`：提取到的版本对象
- `error`：如果找不到版本号则返回错误


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
