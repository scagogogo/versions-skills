# 解析与检查

深入解析版本号字符串，提取各组成部分，并判断版本类型。

```mermaid
flowchart TB
  INPUT["字符串<br/>v1.2.3-beta1+build.7"]
  INPUT --> PARSE["NewVersion<br/>宽松解析"]
  PARSE --> PARTS["拆解 5 段<br/>Prefix / Numbers / Suffix / Metadata / Raw"]
  PARTS --> CHECK{"类型判断"}
  CHECK --> T1["IsPrerelease? 有后缀"]
  CHECK --> T2["IsStable? 无后缀"]
  CHECK --> T3["IsBeta / IsRC / IsAlpha..."]
  PARTS --> VLD{"校验"}
  VLD --> V1["Validate 基本校验"]
  VLD --> V2["IsSemver 严格 SemVer"]
  INPUT -.->|"Coerce 从任意文本提取"| COERCE["Coerce<br/>download/app-1.2.3.tar.gz → 1.2.3"]

  style INPUT fill:#f8fafc,stroke:#475569
  style PARSE fill:#eff6ff,stroke:#2563eb
  style PARTS fill:#eff6ff,stroke:#2563eb
  style CHECK fill:#fff7ed,stroke:#ea580c
  style VLD fill:#fff7ed,stroke:#ea580c
  style COERCE fill:#f0fdf4,stroke:#16a34a,stroke-dasharray:4 3
```

## 🧱 拆解结构

```go
v := versions.NewVersion("v1.2.3-beta1+build.7")

fmt.Println(v.Prefix)         // v
fmt.Println(v.VersionNumbers) // [1 2 3]
fmt.Println(v.Suffix)         // -beta1
fmt.Println(v.Metadata)       // build.7
fmt.Println(v.Segments())     // [1 2 3]
fmt.Println(v.SubVersion())   // 1（后缀里的子版本号）
```

字段语义见 [版本号结构](/concepts/version-anatomy)。

## 🏷 类型判断

`Version` 提供一族 `Is*` 谓词：

```go
v := versions.NewVersion("1.0.0-beta2")

v.IsStable()      // false
v.IsPrerelease()  // true
v.IsBeta()        // true
v.IsAlpha()       // false
v.IsRC()          // false
v.PreReleaseType() // "beta"
v.SuffixWeight()  // 200 (SuffixWeightBeta)
```

完整谓词见 [谓词方法索引](/sdk/predicates)。

## ✅ 校验

```go
v, _ := versions.NewVersionE("not-a-version")
fmt.Println(v.IsValid()) // false

err := v.Validate()      // ErrVersionInvalid
err = v.ValidateSemver() // 严格 SemVer 校验
```

## 🎯 从文本提取（Coerce）

```go
v := versions.Coerce("download/app-1.2.3-linux-amd64.tar.gz")
fmt.Println(v.Raw) // 1.2.3
```

`Coerce` 在任意字符串中找第一个版本号模式子串。

## 🚀 下一步

- [排序与极值](/tutorials/sort-and-minmax)
- 概念：[后缀权重](/concepts/suffix-weight) · [SemVer](/concepts/semver)
- API：[`Version`](/sdk/api/version) · [`Coerce`](/sdk/api/coerce)
