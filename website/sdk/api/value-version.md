# Value

::: info 方法 · `Version`
```go
func (x Version) Value() (driver.Value, error)
```
:::

## 📖 说明

Value 实现 driver.Valuer 接口

返回版本字符串用于数据库存储。


#### 返回

- `driver.Value`：版本原始字符串
- `error`：始终为 nil


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
