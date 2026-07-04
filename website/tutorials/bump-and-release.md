# 不可变变更与发布流程

用 `With*` / `Bump*` 链式构造与递增版本号，模拟发布流程。

## 🔄 Bump 递增

```go
v := versions.NewVersion("1.2.3-beta1")

v.BumpPatch().Raw // 1.2.4（Patch+1，后缀清除）
v.BumpMinor().Raw // 1.3.0（Minor+1，Patch 清零，后缀清除）
v.BumpMajor().Raw // 2.0.0（Major+1，其余清零，后缀清除）
```

::: warning Bump 清除后缀
任何 `Bump*` 都会**清除后缀**——预发布版 bump 后即晋升为正式版。这是发布流程的语义：从 `1.2.3-rc1` bump patch 得到的是 `1.2.3`，不是 `1.2.4-rc1`。
:::

## 🎨 With 链式构造

```go
v := versions.NewVersion("1.2.3")

v2 := v.WithSuffix("-rc1")          // 1.2.3-rc1
v3 := v.WithPatch(5)                // 1.2.5
v4 := v.WithMajor(2).WithMinor(0)   // 2.0.3
```

原对象 `v` 不变——每次 `With*` 返回新对象。详见 [不可变性](/concepts/immutability)。

## 🏗 VersionBuilder 流式构造

从零构造一个版本：

```go
v := versions.NewVersionBuilder().
	Prefix("v").
	Major(1).
	Minor(2).
	Patch(3).
	Suffix("-beta1").
	Build()
// v.Raw = "v1.2.3-beta1"
```

## 📋 发布流程示例

```go
// 当前开发版
dev := versions.NewVersion("1.5.0-dev")

// 进入 beta
beta := dev.WithSuffix("-beta1")

// 发布候选
rc := versions.NewVersionBuilder().
	Prefix(dev.Prefix.String()).
	Major(dev.Major()).
	Minor(dev.Minor()).
	Patch(dev.Patch()).
	Suffix("-rc1").
	Build()

// 正式发布（bump 到正式版）
release := dev.BumpMinor() // 或直接 WithSuffix("")
```

发布阶梯——从开发版到正式版的晋升路径：

:::mermaid
flowchart LR
  DEV["dev<br/>1.5.0-dev"] -->|"WithSuffix('-beta1')"| BETA["beta<br/>1.5.0-beta1"]
  BETA -->|"WithSuffix('-rc1')"| RC["rc<br/>1.5.0-rc1"]
  RC -->|"BumpMinor() / WithSuffix('')"| RELEASE["✅ 正式版<br/>1.5.0<br/>（或 1.6.0）"]

  style DEV fill:#fef2f2,stroke:#dc2626
  style BETA fill:#fff7ed,stroke:#ea580c
  style RC fill:#fef9c3,stroke:#ca8a04
  style RELEASE fill:#f0fdf4,stroke:#16a34a,stroke-width:3px
:::

每一步都返回**新对象**，原版本不变（见 [不可变性](/concepts/immutability)）。

## 🚀 下一步

- [文件批处理](/tutorials/file-batch)
- API：[`BumpMajor`](/sdk/api/bump-major-version) · [`WithSuffix`](/sdk/api/with-suffix-version) · [`VersionBuilder`](/sdk/api/version-builder)
- CLI：[`bump`](/cli/commands/bump) · [`build`](/cli/commands/build) · [`set-suffix`](/cli/commands/set-suffix)
