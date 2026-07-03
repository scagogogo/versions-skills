# Format

::: info 方法 · `Version`
```go
func (x *Version) Format(template string) string
```
:::

## 📖 说明

Format 按照模板格式化版本号

支持的占位符:
- %M  主版本号
- %m  次版本号
- %p  修订版本号
- %P  前缀
- %s  后缀
- %r  原始字符串
- %c  规范字符串
- %%  百分号


#### 参数

- `template`：格式化模板字符串


#### 返回

- `string`：格式化后的字符串


```go
v := versions.NewVersion("v1.2.3-beta")
fmt.Println(v.Format("version %M.%m.%p"))       // "version 1.2.3"
fmt.Println(v.Format("prefix=%P major=%M"))      // "prefix=v major=1"
```


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
