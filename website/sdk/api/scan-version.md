# Scan

::: info 方法 · `Version`
```go
func (x *Version) Scan(src interface
```
:::

## 📖 说明

Scan 实现 sql.Scanner 接口

从数据库扫描版本值。支持 string 和 []byte 类型。


#### 参数

- `src`：数据库值


#### 返回

- `error`：如果值类型不支持或版本无效则返回错误


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
