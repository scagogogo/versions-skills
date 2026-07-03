# VisualizeVersionGroups

::: info 函数 · 根包
```go
func VisualizeVersionGroups(versions []*Version, w io.Writer)
```
:::

## 📖 说明

VisualizeVersionGroups 可视化版本组之间的关系

此函数将版本组集合转换为可视化的树状文本表示，展示其层次关系。
适合用于查看大型版本库的版本组织结构。


#### 参数

- `versions`：要可视化的版本集合
- `w`：输出写入的目标

示例:

versions := ReadVersionsFromFile("versions.txt")
VisualizeVersionGroups(versions, os.Stdout)


---

::: details 源码位置
定义于 [`visualize.go`](https://github.com/scagogogo/versions-skills/blob/main/visualize.go)
:::
