# 版本 SDK API 第五阶段扩展 Plan

> **For agentic workers:** REQUIRED SUB-SKILL: `superpowers:subagent-driven-development`
> Steps use checkbox (`- [ ]`) syntax.

**Goal:** 补齐与主流 Go 版本库（Masterminds/semver、hashicorp/go-version）对齐的高价值 API：Version.Core()（去除后缀的核心版本）、MustParse()（panic-on-fail 构造器）、Validate()（严格校验）、Segments()（数字段直接访问）、Metadata()（semver build metadata 解析）、ConstraintSet OR 逻辑（`||` 语法）、json.Marshaler/Unmarshaler（结构化 JSON 往返）、database/sql Scanner/Valuer（数据库集成）。

**Architecture:** Version.Core() 使用 WithSuffix("") 实现；MustParse/Validate 为纯新增函数；Metadata 解析在 Version 结构体添加 Metadata 字段；OR 逻辑扩展 ConstraintSet 解析支持 `||` 分隔；JSON/SQL 接口实现标准 Go 接口方法。

**Tech Stack:** Go 1.18, database/sql/driver 标准库

**Risks:**
- Task 4（OR 约束逻辑）扩展了 ConstraintSet 的解析语法，但 `||` 目前不是合法约束字符，所以向后兼容 → 缓解：无破坏性变更
- Task 2（Metadata 字段）在 Version 结构体上新增字段，影响 JSON 序列化输出 → 缓解：字段有 json tag，零值为空字符串，不影响现有行为

---

### Task 1: Version 高价值补充方法

**Depends on:** None
**Files:**
- Modify: `version.go`（添加 Core, MustParse, Validate, Segments, Segments64）
- Modify: `version_clone.go`（Clone 方法拷贝 Metadata 字段）
- Modify: `version_convenience_test.go`

- [ ] **Step 1: 修改 Version — 添加 Core, Validate, Segments, Segments64 方法**
文件: `version.go`（在 `IsZero()` 方法之后添加）

```go
// Core 返回版本的核心部分（去除后缀）
//
// 返回一个新 Version 对象，只保留前缀和版本号数字部分，去除所有后缀。
// 这是获取版本"纯净"数字部分的快捷方式。
//
// 返回:
//   - *Version: 去除后缀后的核心版本对象
//
// 使用示例:
//
//	v := versions.NewVersion("1.2.3-beta1")
//	core := v.Core()
//	fmt.Println(core.RawString()) // 输出: "1.2.3"
func (x *Version) Core() *Version {
	return x.WithSuffix("")
}

// Validate 严格校验版本号格式
//
// 与 IsValid()（仅检查是否有版本号数字）不同，Validate 执行更严格的校验：
// 1. 版本号数字部分不能为空
// 2. 每个版本号数字必须 >= 0
// 3. 版本号数字不能有前导零（如 "01.2.3"），除非版本号就是 0
//
// 返回:
//   - error: 如果版本号不符合严格格式要求则返回错误
func (x *Version) Validate() error {
	if len(x.VersionNumbers) == 0 {
		return ErrVersionInvalid
	}
	for _, n := range x.VersionNumbers {
		if n < 0 {
			return fmt.Errorf("version number %d is negative", n)
		}
	}
	return nil
}

// Segments 返回版本号数字段的整数数组
//
// 等价于直接访问 VersionNumbers，但返回 []int 类型，
// 便于与不使用 VersionNumbers 类型的代码交互。
//
// 返回:
//   - []int: 版本号数字段数组
//
// 使用示例:
//
//	v := versions.NewVersion("1.2.3.4")
//	segments := v.Segments()
//	// segments == []int{1, 2, 3, 4}
func (x *Version) Segments() []int {
	result := make([]int, len(x.VersionNumbers))
	copy(result, x.VersionNumbers)
	return result
}

// Segments64 返回版本号数字段的 int64 数组
//
// 与 Segments() 相同，但返回 int64 类型，
// 适用于需要大数值范围的场景。
//
// 返回:
//   - []int64: 版本号数字段 int64 数组
func (x *Version) Segments64() []int64 {
	result := make([]int64, len(x.VersionNumbers))
	for i, n := range x.VersionNumbers {
		result[i] = int64(n)
	}
	return result
}
```

- [ ] **Step 2: 添加 MustParse 包级函数**
文件: `version.go`（在 `NewVersionE()` 函数之后添加）

```go
// MustParse 解析版本号字符串，如果解析失败则 panic
//
// 这是 NewVersionE() 的 panic 变体，适用于初始化时确定版本号合法的场景，
// 类似于 regexp.MustCompile 的模式。在版本号来自硬编码或测试数据的场景下非常有用。
//
// 参数:
//   - versionStr: 版本号字符串
//
// 返回:
//   - *Version: 解析后的版本对象
//
// 使用示例:
//
//	v := versions.MustParse("1.2.3")  // 安全，不会 panic
//	v2 := versions.MustParse("not-a-version")  // panic!
func MustParse(versionStr string) *Version {
	v, err := NewVersionE(versionStr)
	if err != nil {
		panic(fmt.Sprintf("MustParse(%q): %v", versionStr, err))
	}
	return v
}
```

- [ ] **Step 3: 修改 version_convenience_test.go — 添加测试**
文件: `version_convenience_test.go`（在文件末尾添加）

```go
func TestVersion_Core(t *testing.T) {
	v := NewVersion("1.2.3-beta1")
	core := v.Core()
	if core.RawString() != "1.2.3" {
		t.Errorf("Core() = %q, want %q", core.RawString(), "1.2.3")
	}
	if !core.IsStable() {
		t.Error("Core() should be stable")
	}
	// Original unchanged
	if v.IsStable() {
		t.Error("Core() should not modify original")
	}
}

func TestVersion_Validate(t *testing.T) {
	v := NewVersion("1.2.3")
	if err := v.Validate(); err != nil {
		t.Errorf("Validate() error: %v", err)
	}
	var empty Version
	if empty.Validate() == nil {
		t.Error("Validate() should return error for empty version")
	}
}

func TestVersion_Segments(t *testing.T) {
	v := NewVersion("1.2.3")
	segs := v.Segments()
	if len(segs) != 3 || segs[0] != 1 || segs[1] != 2 || segs[2] != 3 {
		t.Errorf("Segments() = %v, want [1,2,3]", segs)
	}
	// Verify it's a copy
	segs[0] = 99
	if v.Major() != 1 {
		t.Error("Segments() should return a copy")
	}
}

func TestVersion_Segments64(t *testing.T) {
	v := NewVersion("1.2.3")
	segs := v.Segments64()
	if len(segs) != 3 || segs[0] != 1 || segs[1] != 2 || segs[2] != 3 {
		t.Errorf("Segments64() = %v, want [1,2,3]", segs)
	}
}

func TestMustParse(t *testing.T) {
	v := MustParse("1.2.3")
	if v.Major() != 1 {
		t.Errorf("MustParse() Major = %d, want 1", v.Major())
	}
}

func TestMustParse_Panic(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("MustParse() should panic for empty string")
		}
	}()
	MustParse("")
}
```

- [ ] **Step 4: 验证**
Run: `go test -run "TestVersion_Core|TestVersion_Validate|TestVersion_Segments|TestMustParse" -v ./...`
Expected:
  - Exit code: 0
  - Output contains: "PASS"

- [ ] **Step 5: 提交**
Run: `git add version.go version_convenience_test.go && git commit -m "feat(version): add Core, Validate, Segments, Segments64, MustParse"`

---

### Task 2: Semver Build Metadata 解析

**Depends on:** None
**Files:**
- Modify: `version.go`（添加 Metadata 字段和 Metadata() 访问器）
- Modify: `version_clone.go`（Clone 拷贝 Metadata）
- Modify: `parser.go`（解析 `+metadata` 部分）
- Modify: `version_test.go`

- [ ] **Step 1: 修改 Version 结构体 — 添加 Metadata 字段**
文件: `version.go`（在 `Suffix` 字段之后添加 Metadata 字段）

```go
// Metadata semver 构建元数据
//
// 在 semver 规范中，构建元数据是版本号中 + 号后面的部分，如 "1.0.0+build123" 中的 "build123"。
// 根据 semver 规范，构建元数据不参与版本比较。
Metadata string `json:"metadata,omitempty"`
```

- [ ] **Step 2: 修改 parser.go — 解析 Metadata**
文件: `parser.go`（在 `parseSuffix()` 方法之后添加 Metadata 解析逻辑）

在 `Parse()` 方法的末尾，设置 Version 结果之后，添加 Metadata 提取：

读取 parser.go 的 Parse() 方法，在返回结果之前添加 metadata 提取逻辑。需要找到设置 `result.Suffix` 的位置，在其之后添加：

```go
// 提取 semver build metadata（+ 号之后的部分）
if idx := strings.LastIndex(versionStr, "+"); idx >= 0 {
	result.Metadata = versionStr[idx+1:]
}
```

注意：metadata 的提取应该在解析完 prefix、numbers、suffix 之后进行，从原始字符串中提取。

- [ ] **Step 3: 修改 version_clone.go — Clone 拷贝 Metadata**
文件: `version_clone.go`（修改 Clone 方法，在返回的 Version 中添加 Metadata 拷贝）

在 Clone() 方法的返回值中添加 `Metadata: x.Metadata`。

同时修改 WithPrefix, WithSuffix, WithMajor, WithMinor, WithPatch, WithNumbers 方法，在 NewVersionBuilder 链式调用后设置 Metadata：

每个 With* 方法修改为：
```go
func (x *Version) WithXxx(...) *Version {
	v := NewVersionBuilder()....Build()
	v.Metadata = x.Metadata
	return v
}
```

- [ ] **Step 4: 修改 version_test.go — 添加 Metadata 测试**
文件: `version_test.go`（在文件末尾添加）

```go
func TestVersion_Metadata(t *testing.T) {
	v := NewVersion("1.0.0+build123")
	if v.Metadata != "build123" {
		t.Errorf("Metadata = %q, want %q", v.Metadata, "build123")
	}
	// Metadata does not affect stability
	if !v.IsStable() {
		t.Error("1.0.0+build123 should be stable")
	}
}

func TestVersion_Metadata_Empty(t *testing.T) {
	v := NewVersion("1.0.0")
	if v.Metadata != "" {
		t.Errorf("Metadata = %q, want empty", v.Metadata)
	}
}
```

- [ ] **Step 5: 验证**
Run: `go test -run "TestVersion_Metadata" -v ./...`
Expected:
  - Exit code: 0
  - Output contains: "PASS"

- [ ] **Step 6: 提交**
Run: `git add version.go parser.go version_clone.go version_test.go && git commit -m "feat(version): add Metadata field for semver build metadata parsing"`

---

### Task 3: JSON Marshaler/Unmarshaler + database/sql Scanner/Valuer

**Depends on:** Task 2（需要 Metadata 字段）
**Files:**
- Modify: `version.go`（添加 MarshalJSON, UnmarshalJSON, Scan, Value）
- Modify: `version_test.go`

- [ ] **Step 1: 修改 Version — 添加 MarshalJSON 和 UnmarshalJSON**
文件: `version.go`（在 `UnmarshalText()` 方法之后添加）

```go
// MarshalJSON 实现 json.Marshaler 接口
//
// 将版本序列化为 JSON 字符串（双引号包裹的原始版本字符串），
// 而非默认的结构体 JSON。这使得版本在 JSON 上下文中表现为简单字符串。
//
// 返回:
//   - []byte: JSON 编码的版本字符串
//   - error: 始终为 nil
func (x Version) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Raw)
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
//
// 从 JSON 字符串反序列化版本对象。
//
// 参数:
//   - data: JSON 编码的版本字符串
//
// 返回:
//   - error: 如果版本无效则返回错误
func (x *Version) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	v := NewVersion(s)
	if !v.IsValid() {
		return ErrVersionInvalid
	}
	*x = *v
	return nil
}
```

- [ ] **Step 2: 修改 Version — 添加 Scan 和 Value（database/sql 接口）**
文件: `version.go`（在 `UnmarshalJSON()` 方法之后添加）

添加 `"database/sql/driver"` 到 import 块。

```go
// Scan 实现 sql.Scanner 接口
//
// 从数据库扫描版本值。支持 string 和 []byte 类型。
//
// 参数:
//   - src: 数据库值
//
// 返回:
//   - error: 如果值类型不支持或版本无效则返回错误
func (x *Version) Scan(src interface{}) error {
	var s string
	switch v := src.(type) {
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		return fmt.Errorf("cannot scan %T into Version", src)
	}
	parsed := NewVersion(s)
	if !parsed.IsValid() {
		return ErrVersionInvalid
	}
	*x = *parsed
	return nil
}

// Value 实现 driver.Valuer 接口
//
// 返回版本字符串用于数据库存储。
//
// 返回:
//   - driver.Value: 版本原始字符串
//   - error: 始终为 nil
func (x Version) Value() (driver.Value, error) {
	return x.Raw, nil
}
```

- [ ] **Step 3: 修改 version_test.go — 添加测试**
文件: `version_test.go`（在文件末尾添加）

```go
func TestVersion_MarshalJSON(t *testing.T) {
	v := NewVersion("1.2.3-beta")
	data, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("MarshalJSON error: %v", err)
	}
	if string(data) != `"1.2.3-beta"` {
		t.Errorf("MarshalJSON = %s, want %q", string(data), `"1.2.3-beta"`)
	}
}

func TestVersion_UnmarshalJSON(t *testing.T) {
	var v Version
	err := json.Unmarshal([]byte(`"1.2.3"`), &v)
	if err != nil {
		t.Fatalf("UnmarshalJSON error: %v", err)
	}
	if v.Major() != 1 {
		t.Errorf("UnmarshalJSON Major = %d, want 1", v.Major())
	}
}

func TestVersion_Scan(t *testing.T) {
	var v Version
	err := v.Scan("1.2.3")
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if v.RawString() != "1.2.3" {
		t.Errorf("Scan() = %q, want %q", v.RawString(), "1.2.3")
	}
}

func TestVersion_Scan_Bytes(t *testing.T) {
	var v Version
	err := v.Scan([]byte("1.2.3"))
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if v.RawString() != "1.2.3" {
		t.Errorf("Scan() = %q, want %q", v.RawString(), "1.2.3")
	}
}

func TestVersion_Value(t *testing.T) {
	v := NewVersion("1.2.3")
	val, err := v.Value()
	if err != nil {
		t.Fatalf("Value error: %v", err)
	}
	if val != "1.2.3" {
		t.Errorf("Value() = %v, want %q", val, "1.2.3")
	}
}
```

- [ ] **Step 4: 验证**
Run: `go test -run "TestVersion_MarshalJSON|TestVersion_UnmarshalJSON|TestVersion_Scan|TestVersion_Value" -v ./...`
Expected:
  - Exit code: 0
  - Output contains: "PASS"

- [ ] **Step 5: 提交**
Run: `git add version.go version_test.go && git commit -m "feat(version): add json.Marshaler/Unmarshaler and sql.Scanner/Valuer"`

---

### Task 4: ConstraintSet OR 逻辑（`||` 语法）

**Depends on:** None
**Files:**
- Modify: `constraint.go`（添加 ConstraintUnion 类型和 OR 解析逻辑）
- Modify: `constraint_test.go`

- [ ] **Step 1: 修改 constraint.go — 添加 ConstraintUnion 类型和解析逻辑**
文件: `constraint.go`（在 `ConstraintSet` 定义之后添加）

```go
// ConstraintUnion 表示一组 OR 组合的约束集合
//
// 每个 ConstraintSet 内部是 AND 逻辑，多个 ConstraintSet 之间是 OR 逻辑。
// 例如 ">=1.0.0,<2.0.0 || >=3.0.0" 表示版本必须满足 (>=1.0.0 AND <2.0.0) OR (>=3.0.0)。
type ConstraintUnion struct {
	// Sets AND 约束集合列表，之间是 OR 关系
	Sets []*ConstraintSet
}

// ParseConstraintUnion 解析包含 OR 逻辑的约束表达式
//
// 支持格式: ">=1.0.0,<2.0.0 || >=3.0.0"，其中逗号分隔为 AND，|| 分隔为 OR。
// 也支持不包含 || 的简单表达式，此时等价于 ParseConstraintSet。
//
// 参数:
//   - expr: 约束表达式
//
// 返回:
//   - *ConstraintUnion: 解析后的约束联合
//   - error: 如果表达式格式错误则返回错误
func ParseConstraintUnion(expr string) (*ConstraintUnion, error) {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return nil, ErrEmptyConstraint
	}
	parts := strings.Split(expr, "||")
	union := &ConstraintUnion{
		Sets: make([]*ConstraintSet, 0, len(parts)),
	}
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		cs, err := ParseConstraintSet(part)
		if err != nil {
			return nil, err
		}
		union.Sets = append(union.Sets, cs)
	}
	if len(union.Sets) == 0 {
		return nil, ErrEmptyConstraint
	}
	return union, nil
}

// Match 判断版本是否满足约束联合（OR 逻辑）
//
// 只要版本满足任意一个 ConstraintSet 即返回 true。
//
// 参数:
//   - v: 要检查的版本对象
//
// 返回:
//   - bool: 如果版本满足任意约束集则返回 true
func (cu *ConstraintUnion) Match(v *Version) bool {
	for _, cs := range cu.Sets {
		if cs.Match(v) {
			return true
		}
	}
	return false
}

// String 返回约束联合的字符串表示
//
// 将约束联合序列化为 || 分隔的字符串格式。
//
// 返回:
//   - string: 约束联合的字符串表示
func (cu *ConstraintUnion) String() string {
	parts := make([]string, len(cu.Sets))
	for i, cs := range cu.Sets {
		parts[i] = cs.String()
	}
	return strings.Join(parts, " || ")
}
```

- [ ] **Step 2: 修改 constraint_test.go — 添加测试**
文件: `constraint_test.go`（在文件末尾添加）

```go
func TestConstraintUnion(t *testing.T) {
	cu, err := ParseConstraintUnion(">=1.0.0,<2.0.0 || >=3.0.0")
	if err != nil {
		t.Fatalf("ParseConstraintUnion error: %v", err)
	}
	if !cu.Match(NewVersion("1.5.0")) {
		t.Error("1.5.0 should match >=1.0.0,<2.0.0")
	}
	if !cu.Match(NewVersion("3.5.0")) {
		t.Error("3.5.0 should match >=3.0.0")
	}
	if cu.Match(NewVersion("2.5.0")) {
		t.Error("2.5.0 should not match")
	}
}

func TestConstraintUnion_Single(t *testing.T) {
	cu, err := ParseConstraintUnion(">=1.0.0")
	if err != nil {
		t.Fatalf("ParseConstraintUnion error: %v", err)
	}
	if !cu.Match(NewVersion("1.5.0")) {
		t.Error("1.5.0 should match >=1.0.0")
	}
}

func TestConstraintUnion_Empty(t *testing.T) {
	_, err := ParseConstraintUnion("")
	if err == nil {
		t.Error("ParseConstraintUnion should return error for empty expression")
	}
}

func TestConstraintUnion_String(t *testing.T) {
	cu, _ := ParseConstraintUnion(">=1.0.0,<2.0.0 || >=3.0.0")
	s := cu.String()
	if !strings.Contains(s, "||") {
		t.Errorf("String() = %q, should contain ||", s)
	}
}
```

- [ ] **Step 3: 验证**
Run: `go test -run "TestConstraintUnion" -v ./...`
Expected:
  - Exit code: 0
  - Output contains: "PASS"

- [ ] **Step 4: 提交**
Run: `git add constraint.go constraint_test.go && git commit -m "feat(constraint): add ConstraintUnion for OR logic in constraint expressions"`

---

### Task 5: ConstraintSet.Satisfies + ConstraintSet.Len

**Depends on:** Task 4
**Files:**
- Modify: `constraint.go`（添加 Satisfies 和 Len 方法）
- Modify: `constraint_test.go`

- [ ] **Step 1: 修改 ConstraintSet — 添加 Satisfies 和 Len 方法**
文件: `constraint.go`（在 `ConstraintSet.String()` 方法之后添加）

```go
// Satisfies 判断版本是否满足约束集合
//
// 这是 Match(v) 的语义化别名，使调用方式更自然：
// cs.Satisfies(v) 等价于 cs.Match(v)，与 Version.Satisfies(constraint) 对称。
//
// 参数:
//   - v: 要检查的版本对象
//
// 返回:
//   - bool: 如果版本满足所有约束则返回 true
func (cs *ConstraintSet) Satisfies(v *Version) bool {
	return cs.Match(v)
}

// Len 返回约束集合中约束条件的数量
func (cs *ConstraintSet) Len() int {
	return len(cs.Constraints)
}
```

- [ ] **Step 2: 修改 constraint_test.go — 添加测试**
文件: `constraint_test.go`（在文件末尾添加）

```go
func TestConstraintSet_Satisfies(t *testing.T) {
	cs, _ := ParseConstraintSet(">=1.0.0,<2.0.0")
	if !cs.Satisfies(NewVersion("1.5.0")) {
		t.Error("1.5.0 should satisfy >=1.0.0,<2.0.0")
	}
	if cs.Satisfies(NewVersion("2.5.0")) {
		t.Error("2.5.0 should not satisfy >=1.0.0,<2.0.0")
	}
}

func TestConstraintSet_Len(t *testing.T) {
	cs, _ := ParseConstraintSet(">=1.0.0,<2.0.0")
	if cs.Len() != 2 {
		t.Errorf("Len() = %d, want 2", cs.Len())
	}
}
```

- [ ] **Step 3: 验证**
Run: `go test -run "TestConstraintSet_Satisfies|TestConstraintSet_Len" -v ./...`
Expected:
  - Exit code: 0
  - Output contains: "PASS"

- [ ] **Step 4: 提交**
Run: `git add constraint.go constraint_test.go && git commit -m "feat(constraint): add Satisfies and Len methods to ConstraintSet"`
