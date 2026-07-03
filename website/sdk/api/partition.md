# Partition

::: info 函数 · 根包
```go
func Partition(versions []*Version, predicate func(*Version) bool) ([]*Version, []*Version)
```
:::

## 📖 说明

Partition 根据谓词将版本列表分为两组

返回两个切片：第一个包含满足谓词的版本，第二个包含不满足谓词的版本。
保持原始顺序。


#### 参数

- `versions`：版本对象列表
- `predicate`：分区谓词函数


#### 返回

- `[]*Version`：满足谓词的版本列表
- `[]*Version`：不满足谓词的版本列表


---

::: details 源码位置
定义于 [`version_utils.go`](https://github.com/scagogogo/versions-skills/blob/main/version_utils.go)
:::
