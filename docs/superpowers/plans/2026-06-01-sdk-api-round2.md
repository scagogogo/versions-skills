# 版本 SDK API 第二阶段扩展 Plan

> **For agentic workers:** REQUIRED SUB-SKILL: `superpowers:subagent-driven-development`
> Steps use checkbox (`- [ ]`) syntax.

**Goal:** 修复可视化模块硬编码 Bug，补全缺失的便捷判断方法（IsMilestone/IsNightly/IsFinal/IsGA），添加 Constraint 的 String() 序列化、版本不可变修改器（Clone/With*）、版本集合运算（Difference/Intersection）和 FilterByMinor/FilterByPrefix 工具函数，使 SDK API 更完整。

**Architecture:** 修复 visualize.go 中硬编码的组数和省略提示 → 在 version.go 上补充 IsMilestone/IsNightly/IsFinal/IsGA 便捷方法 → 为 Constraint/ConstraintSet 添加 String() 方法 → 新增 version_clone.go 提供 Clone/With* 不可变修改器 → 在 version_utils.go 中添加 FilterByMinor/FilterByPrefix/Difference/Intersection → 所有新增方法均使用纯函数模式，不修改现有方法签名。

**Tech Stack:** Go 1.18, 现有依赖不变（stretchr/testify, go-tuple, go-compare-anything）

**Risks:**
- Task 1 修改了 visualize.go 和 visualize_test.go 中的硬编码值，需确保修复后动态计算值与现有测试期望一致 → 缓解：测试数据不变，动态计算结果与之前硬编码值相同
- Task 3 的 Constraint.String() 需要能正确还原原始约束表达式 → 缓解：使用 Operator + Version.Raw 重建，通配符单独处理
- Task 4 的 With* 方法返回新对象，需确保深拷贝不遗漏字段 → 缓解：基于 VersionBuilder 构建，字段不多

---

### Task 1: 修复可视化模块硬编码 Bug

**Depends on:** None
**Files:**
- Modify: `visualize.go:45,83-86,134`
- Modify: `visualize_test.go:52,109`

- [ ] **Step 1: 修改 VisualizeVersions — 替换硬编码的组数和省略提示为动态计算**
文件: `visualize.go:44-45`（替换硬编码的版本组数）

```go
// 写入总览信息
fmt.Fprintf(w, "版本总数: %d\n", len(versions))
fmt.Fprintf(w, "版本组数: %d\n\n", len(groups))
```

文件: `visualize.go:82-87`（替换硬编码的省略提示）

```go
// 如果有更多版本未显示，提示省略情况
if maxItems > 0 && len(sortedVersions) > maxItems {
	fmt.Fprintf(w, "└── ...还有%d个版本未显示\n", len(sortedVersions)-maxItems)
}
```

- [ ] **Step 2: 修改 VisualizeVersionGroups — 替换硬编码的组数为动态计算**
文件: `visualize.go:133-134`（替换硬编码的版本组数）

```go
// 写入总览信息
fmt.Fprintf(w, "版本总数: %d\n", len(versions))
fmt.Fprintf(w, "版本组数: %d\n\n", len(groups))
```

- [ ] **Step 3: 修改 visualize_test.go — 更新测试断言使用动态值**
文件: `visualize_test.go:52`（替换硬编码断言）

```go
assert.Contains(t, output, "版本组数: 2")
```

此处无需修改，因为测试数据确实有 2 个版本组（1.x 和 2.x），动态计算结果也是 2。

文件: `visualize_test.go:109`（替换硬编码断言）

```go
assert.Contains(t, output, "版本组数: 5")
```

此处无需修改，因为测试数据确实有 5 个版本组（1.0, 1.1, 2.0, 2.1, 10），动态计算结果也是 5。

- [ ] **Step 4: 验证可视化修复**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test -run TestVisualize -v ./...`
Expected:
  - Exit code: 0
  - Output contains: "PASS"
  - Output does NOT contain: "FAIL"

- [ ] **Step 5: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add visualize.go && git commit -m "fix(visualize): replace hardcoded group counts with dynamic calculation"`

---

### Task 2: 补全 Version 便捷判断方法

**Depends on:** None
**Files:**
- Modify: `version.go:329-331`（在 IsSnapshot 后添加新方法）
- Modify: `version_convenience_test.go`（添加新测试）

- [ ] **Step 1: 修改 Version — 添加 IsMilestone/IsNightly/IsFinal/IsGA 便捷方法**
文件: `version.go`（在 `IsSnapshot()` 方法之后添加以下方法）

```go
// IsMilestone 判断版本是否为里程碑版
//
// 里程碑版是指后缀包含 milestone/m 标识的版本，如 "1.0.0-m1"、"1.0.0-milestone1"。
func (x *Version) IsMilestone() bool {
	w := GetSuffixWeight(string(x.Suffix))
	return w == SuffixWeightMilestone
}

// IsNightly 判断版本是否为夜间构建版
//
// 夜间构建版是指后缀包含 nightly 标识的版本，如 "1.0.0-nightly"。
func (x *Version) IsNightly() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightNightly
}

// IsFinal 判断版本是否为 Final 版本
//
// Final 版本是指后缀包含 final 标识的版本（Maven 生态常见），如 "1.0.0-final"。
// 注意：Final 后缀与无后缀的正式版语义相同，但 IsFinal 专用于检测显式的 final 后缀。
func (x *Version) IsFinal() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightFinal
}

// IsGA 判断版本是否为 GA（Generally Available）版本
//
// GA 版本是指后缀包含 ga 标识的版本，如 "1.0.0-ga"。
func (x *Version) IsGA() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightGA
}
```

- [ ] **Step 2: 修改 version_convenience_test.go — 添加新方法测试**
文件: `version_convenience_test.go`（在文件末尾添加以下测试函数）

```go
func TestVersion_IsMilestone(t *testing.T) {
	if !NewVersion("1.0.0-milestone1").IsMilestone() {
		t.Error("1.0.0-milestone1 should be milestone")
	}
	if !NewVersion("1.0.0-m1").IsMilestone() {
		t.Error("1.0.0-m1 should be milestone (short form)")
	}
	if NewVersion("1.0.0-beta").IsMilestone() {
		t.Error("1.0.0-beta should not be milestone")
	}
}

func TestVersion_IsNightly(t *testing.T) {
	if !NewVersion("1.0.0-nightly").IsNightly() {
		t.Error("1.0.0-nightly should be nightly")
	}
	if NewVersion("1.0.0-beta").IsNightly() {
		t.Error("1.0.0-beta should not be nightly")
	}
}

func TestVersion_IsFinal(t *testing.T) {
	if !NewVersion("1.0.0-final").IsFinal() {
		t.Error("1.0.0-final should be final")
	}
	if NewVersion("1.0.0").IsFinal() {
		t.Error("1.0.0 should not be final (no suffix)")
	}
}

func TestVersion_IsGA(t *testing.T) {
	if !NewVersion("1.0.0-ga").IsGA() {
		t.Error("1.0.0-ga should be GA")
	}
	if NewVersion("1.0.0").IsGA() {
		t.Error("1.0.0 should not be GA (no suffix)")
	}
}
```

- [ ] **Step 3: 验证便捷方法**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test -run "TestVersion_Is(Milestone|Nightly|Final|GA)" -v ./...`
Expected:
  - Exit code: 0
  - Output contains: "PASS"

- [ ] **Step 4: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add version.go version_convenience_test.go && git commit -m "feat(version): add IsMilestone, IsNightly, IsFinal, IsGA convenience methods"`

---

### Task 3: Constraint 和 ConstraintSet 添加 String() 序列化方法

**Depends on:** None
**Files:**
- Modify: `constraint.go`（在 Match 方法后添加 String() 方法）
- Modify: `constraint_test.go`（添加 String() 测试）

- [ ] **Step 1: 修改 Constraint — 添加 String() 方法**
文件: `constraint.go`（在 `Constraint.Match()` 方法之后添加以下方法）

```go
// String 返回约束条件的字符串表示
//
// 将约束条件序列化为可解析的字符串格式，如 ">=1.0.0"、"^1.2.3"、"~1.2"。
// 对于通配符约束（1.x），返回原始版本字符串形式。
//
// 返回:
//   - string: 约束条件的字符串表示
//
// 使用示例:
//
//	c, _ := versions.ParseConstraint(">=1.0.0")
//	fmt.Println(c.String()) // 输出: ">=1.0.0"
func (c *Constraint) String() string {
	switch c.Operator {
	case ConstraintWildcard:
		// 通配符需要特殊处理：将版本号数字中对应位置的 0 还原为 x
		parts := make([]string, len(c.Version.VersionNumbers))
		for i, n := range c.Version.VersionNumbers {
			parts[i] = strconv.Itoa(n)
		}
		// 最后一个 0 替换为 x（通配符位置）
		for i := len(parts) - 1; i >= 0; i-- {
			if parts[i] == "0" {
				parts[i] = "x"
				break
			}
		}
		return strings.Join(parts, ".")
	default:
		return string(c.Operator) + c.Version.Raw
	}
}
```

- [ ] **Step 2: 修改 ConstraintSet — 添加 String() 方法**
文件: `constraint.go`（在 `ConstraintSet.Match()` 方法之后添加以下方法）

```go
// String 返回约束集合的字符串表示
//
// 将约束集合序列化为逗号分隔的字符串格式，如 ">=1.0.0,<2.0.0"。
//
// 返回:
//   - string: 约束集合的字符串表示
//
// 使用示例:
//
//	cs, _ := versions.ParseConstraintSet(">=1.0.0,<2.0.0")
//	fmt.Println(cs.String()) // 输出: ">=1.0.0,<2.0.0"
func (cs *ConstraintSet) String() string {
	parts := make([]string, len(cs.Constraints))
	for i, c := range cs.Constraints {
		parts[i] = c.String()
	}
	return strings.Join(parts, ",")
}
```

- [ ] **Step 3: 添加 strconv import**
文件: `constraint.go:3-6`（更新 import 块）

```go
import (
	"fmt"
	"strconv"
	"strings"
)
```

- [ ] **Step 4: 修改 constraint_test.go — 添加 String() 测试**
文件: `constraint_test.go`（在文件末尾添加以下测试函数）

```go
func TestConstraint_String(t *testing.T) {
	tests := []struct {
		expr     string
		expected string
	}{
		{">=1.0.0", ">=1.0.0"},
		{"<2.0.0", "<2.0.0"},
		{"=1.5.0", "=1.5.0"},
		{"!=0.9.0", "!=0.9.0"},
		{"^1.2.3", "^1.2.3"},
		{"~1.2", "~1.2"},
	}
	for _, tt := range tests {
		t.Run(tt.expr, func(t *testing.T) {
			c, err := ParseConstraint(tt.expr)
			if err != nil {
				t.Fatalf("ParseConstraint(%q) error: %v", tt.expr, err)
			}
			if got := c.String(); got != tt.expected {
				t.Errorf("Constraint.String() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestConstraintSet_String(t *testing.T) {
	cs, err := ParseConstraintSet(">=1.0.0,<2.0.0")
	if err != nil {
		t.Fatalf("ParseConstraintSet error: %v", err)
	}
	got := cs.String()
	if got != ">=1.0.0,<2.0.0" {
		t.Errorf("ConstraintSet.String() = %q, want %q", got, ">=1.0.0,<2.0.0")
	}
}
```

- [ ] **Step 5: 验证 Constraint String()**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test -run "TestConstraint_String|TestConstraintSet_String" -v ./...`
Expected:
  - Exit code: 0
  - Output contains: "PASS"

- [ ] **Step 6: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add constraint.go constraint_test.go && git commit -m "feat(constraint): add String() serialization for Constraint and ConstraintSet"`

---

### Task 4: Version 不可变修改器（Clone/With*）

**Depends on:** None
**Files:**
- Create: `version_clone.go`
- Create: `version_clone_test.go`

- [ ] **Step 1: 创建 version_clone.go — 提供 Clone 和 With* 不可变修改方法**

```go
package versions

// Clone 创建版本的深拷贝
//
// 返回一个与原版本完全相同的新 Version 对象，修改拷贝不会影响原版本。
// 对于不可变的 Version 对象，Clone 主要用于与 With* 方法配合使用。
//
// 返回:
//   - *Version: 版本的深拷贝
//
// 使用示例:
//
//	v1 := versions.NewVersion("1.2.3")
//	v2 := v1.Clone()
//	v2.Raw = "modified"
//	fmt.Println(v1.Raw) // 仍然是 "1.2.3"
func (x *Version) Clone() *Version {
	if x == nil {
		return nil
	}
	numbers := make([]int, len(x.VersionNumbers))
	copy(numbers, x.VersionNumbers)
	return &Version{
		Raw:            x.Raw,
		PublicTime:     x.PublicTime,
		VersionNumbers: numbers,
		Prefix:         x.Prefix,
		Suffix:         x.Suffix,
	}
}

// WithPrefix 返回一个修改前缀的新版本对象
//
// 原版本对象不变，返回一个新对象，其前缀被替换为指定值。
//
// 参数:
//   - prefix: 新的前缀字符串
//
// 返回:
//   - *Version: 修改前缀后的新版本对象
//
// 使用示例:
//
//	v := versions.NewVersion("1.2.3")
//	newV := v.WithPrefix("v")
//	// newV.Raw == "v1.2.3"
func (x *Version) WithPrefix(prefix string) *Version {
	return NewVersionBuilder().
		Prefix(prefix).
		Numbers(x.VersionNumbers).
		Suffix(string(x.Suffix)).
		Build()
}

// WithSuffix 返回一个修改后缀的新版本对象
//
// 原版本对象不变，返回一个新对象，其后缀被替换为指定值。
//
// 参数:
//   - suffix: 新的后缀字符串，如 "-beta1"
//
// 返回:
//   - *Version: 修改后缀后的新版本对象
//
// 使用示例:
//
//	v := versions.NewVersion("1.2.3")
//	newV := v.WithSuffix("-rc1")
//	// newV.Raw == "1.2.3-rc1"
func (x *Version) WithSuffix(suffix string) *Version {
	return NewVersionBuilder().
		Prefix(string(x.Prefix)).
		Numbers(x.VersionNumbers).
		Suffix(suffix).
		Build()
}

// WithMajor 返回一个修改主版本号的新版本对象
//
// 原版本对象不变，返回一个新对象，其主版本号被替换为指定值。
// 后缀和前缀保持不变。
//
// 参数:
//   - major: 新的主版本号
//
// 返回:
//   - *Version: 修改主版本号后的新版本对象
func (x *Version) WithMajor(major int) *Version {
	numbers := make([]int, len(x.VersionNumbers))
	copy(numbers, x.VersionNumbers)
	if len(numbers) == 0 {
		numbers = []int{major}
	} else {
		numbers[0] = major
	}
	return NewVersionBuilder().
		Prefix(string(x.Prefix)).
		Numbers(numbers).
		Suffix(string(x.Suffix)).
		Build()
}

// WithMinor 返回一个修改次版本号的新版本对象
//
// 原版本对象不变，返回一个新对象，其次版本号被替换为指定值。
// 前缀和后缀保持不变。
//
// 参数:
//   - minor: 新的次版本号
//
// 返回:
//   - *Version: 修改次版本号后的新版本对象
func (x *Version) WithMinor(minor int) *Version {
	numbers := make([]int, len(x.VersionNumbers))
	copy(numbers, x.VersionNumbers)
	for len(numbers) < 2 {
		numbers = append(numbers, 0)
	}
	numbers[1] = minor
	return NewVersionBuilder().
		Prefix(string(x.Prefix)).
		Numbers(numbers).
		Suffix(string(x.Suffix)).
		Build()
}

// WithPatch 返回一个修改修订版本号的新版本对象
//
// 原版本对象不变，返回一个新对象，其修订版本号被替换为指定值。
// 前缀和后缀保持不变。
//
// 参数:
//   - patch: 新的修订版本号
//
// 返回:
//   - *Version: 修改修订版本号后的新版本对象
func (x *Version) WithPatch(patch int) *Version {
	numbers := make([]int, len(x.VersionNumbers))
	copy(numbers, x.VersionNumbers)
	for len(numbers) < 3 {
		numbers = append(numbers, 0)
	}
	numbers[2] = patch
	return NewVersionBuilder().
		Prefix(string(x.Prefix)).
		Numbers(numbers).
		Suffix(string(x.Suffix)).
		Build()
}
```

- [ ] **Step 2: 创建 version_clone_test.go**

```go
package versions

import "testing"

func TestVersion_Clone(t *testing.T) {
	v1 := NewVersion("1.2.3-beta")
	v2 := v1.Clone()
	if v1.Raw != v2.Raw {
		t.Errorf("Clone() Raw = %q, want %q", v2.Raw, v1.Raw)
	}
	// 修改克隆对象不影响原始对象
	v2.Raw = "modified"
	if v1.Raw == "modified" {
		t.Error("Clone() should not affect original")
	}
}

func TestVersion_Clone_Nil(t *testing.T) {
	var v *Version = nil
	if v.Clone() != nil {
		t.Error("Clone(nil) should return nil")
	}
}

func TestVersion_WithPrefix(t *testing.T) {
	v := NewVersion("1.2.3")
	newV := v.WithPrefix("v")
	if newV.Prefix != "v" {
		t.Errorf("WithPrefix() Prefix = %q, want %q", newV.Prefix, "v")
	}
	if v.Prefix == "v" {
		t.Error("WithPrefix() should not modify original")
	}
}

func TestVersion_WithSuffix(t *testing.T) {
	v := NewVersion("1.2.3")
	newV := v.WithSuffix("-rc1")
	if !newV.IsRC() {
		t.Errorf("WithSuffix() should produce RC version, got %s", newV.Suffix)
	}
	if v.IsPrerelease() {
		t.Error("WithSuffix() should not modify original")
	}
}

func TestVersion_WithMajor(t *testing.T) {
	v := NewVersion("1.2.3")
	newV := v.WithMajor(5)
	if newV.Major() != 5 {
		t.Errorf("WithMajor() Major = %d, want 5", newV.Major())
	}
	if newV.Minor() != 2 {
		t.Errorf("WithMajor() Minor = %d, want 2 (preserved)", newV.Minor())
	}
}

func TestVersion_WithMinor(t *testing.T) {
	v := NewVersion("1.2.3")
	newV := v.WithMinor(9)
	if newV.Minor() != 9 {
		t.Errorf("WithMinor() Minor = %d, want 9", newV.Minor())
	}
	if newV.Major() != 1 {
		t.Errorf("WithMinor() Major = %d, want 1 (preserved)", newV.Major())
	}
}

func TestVersion_WithPatch(t *testing.T) {
	v := NewVersion("1.2.3")
	newV := v.WithPatch(7)
	if newV.Patch() != 7 {
		t.Errorf("WithPatch() Patch = %d, want 7", newV.Patch())
	}
	if newV.Major() != 1 || newV.Minor() != 2 {
		t.Errorf("WithPatch() should preserve Major/Minor, got %d.%d", newV.Major(), newV.Minor())
	}
}
```

- [ ] **Step 3: 验证不可变修改器**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test -run "TestVersion_Clone|TestVersion_With" -v ./...`
Expected:
  - Exit code: 0
  - Output contains: "PASS"

- [ ] **Step 4: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add version_clone.go version_clone_test.go && git commit -m "feat(version): add Clone and With* immutable modifier methods"`

---

### Task 5: 版本集合运算和过滤工具函数扩展

**Depends on:** None
**Files:**
- Modify: `version_utils.go`（在文件末尾添加新函数）
- Modify: `version_utils_test.go`（在文件末尾添加新测试）

- [ ] **Step 1: 修改 version_utils.go — 添加 FilterByMinor/FilterByPrefix/Difference/Intersection**
文件: `version_utils.go`（在 `Count()` 函数之后添加以下函数）

```go
// FilterByMinor 过滤指定次版本号的版本
//
// 返回所有次版本号等于指定值的版本。
//
// 参数:
//   - versions: 版本对象列表
//   - minor: 目标次版本号
//
// 返回:
//   - []*Version: 满足条件的版本列表
func FilterByMinor(versions []*Version, minor int) []*Version {
	return Filter(versions, func(v *Version) bool {
		return v.Minor() == minor
	})
}

// FilterByPrefix 过滤指定前缀的版本
//
// 返回所有前缀等于指定值的版本。
//
// 参数:
//   - versions: 版本对象列表
//   - prefix: 目标前缀字符串，如 "v"
//
// 返回:
//   - []*Version: 满足条件的版本列表
func FilterByPrefix(versions []*Version, prefix string) []*Version {
	return Filter(versions, func(v *Version) bool {
		return string(v.Prefix) == prefix
	})
}

// Difference 返回在 a 中但不在 b 中的版本（差集）
//
// 根据 Raw 字段判断版本是否相同。返回的版本保持 a 中的原始顺序。
//
// 参数:
//   - a: 版本对象列表
//   - b: 要排除的版本对象列表
//
// 返回:
//   - []*Version: 差集版本列表
func Difference(a, b []*Version) []*Version {
	bSet := make(map[string]bool, len(b))
	for _, v := range b {
		bSet[v.Raw] = true
	}
	return Filter(a, func(v *Version) bool {
		return !bSet[v.Raw]
	})
}

// Intersection 返回同时存在于 a 和 b 中的版本（交集）
//
// 根据 Raw 字段判断版本是否相同。返回的版本保持 a 中的原始顺序。
//
// 参数:
//   - a: 版本对象列表
//   - b: 版本对象列表
//
// 返回:
//   - []*Version: 交集版本列表
func Intersection(a, b []*Version) []*Version {
	bSet := make(map[string]bool, len(b))
	for _, v := range b {
		bSet[v.Raw] = true
	}
	return Filter(a, func(v *Version) bool {
		return bSet[v.Raw]
	})
}
```

- [ ] **Step 2: 修改 version_utils_test.go — 添加新函数测试**
文件: `version_utils_test.go`（在文件末尾添加以下测试函数）

```go
func TestFilterByMinor(t *testing.T) {
	versions := NewVersions("1.0.0", "1.1.0", "2.1.0", "2.2.0")
	result := FilterByMinor(versions, 1)
	if len(result) != 2 {
		t.Errorf("FilterByMinor(1) = %d, want 2", len(result))
	}
}

func TestFilterByPrefix(t *testing.T) {
	versions := NewVersions("1.0.0", "v1.0.0", "v2.0.0")
	result := FilterByPrefix(versions, "v")
	if len(result) != 2 {
		t.Errorf("FilterByPrefix(\"v\") = %d, want 2", len(result))
	}
}

func TestDifference(t *testing.T) {
	a := NewVersions("1.0.0", "1.1.0", "1.2.0")
	b := NewVersions("1.1.0", "2.0.0")
	result := Difference(a, b)
	if len(result) != 2 {
		t.Errorf("Difference() = %d, want 2", len(result))
	}
	for _, v := range result {
		if v.Raw == "1.1.0" {
			t.Error("Difference() should not contain 1.1.0")
		}
	}
}

func TestDifference_Empty(t *testing.T) {
	a := NewVersions("1.0.0")
	result := Difference(a, nil)
	if len(result) != 1 {
		t.Errorf("Difference(a, nil) = %d, want 1", len(result))
	}
}

func TestIntersection(t *testing.T) {
	a := NewVersions("1.0.0", "1.1.0", "1.2.0")
	b := NewVersions("1.1.0", "1.2.0", "2.0.0")
	result := Intersection(a, b)
	if len(result) != 2 {
		t.Errorf("Intersection() = %d, want 2", len(result))
	}
}

func TestIntersection_Empty(t *testing.T) {
	a := NewVersions("1.0.0")
	result := Intersection(a, nil)
	if len(result) != 0 {
		t.Errorf("Intersection(a, nil) = %d, want 0", len(result))
	}
}
```

- [ ] **Step 3: 验证集合运算和过滤函数**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test -run "TestFilterByMinor|TestFilterByPrefix|TestDifference|TestIntersection" -v ./...`
Expected:
  - Exit code: 0
  - Output contains: "PASS"

- [ ] **Step 4: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add version_utils.go version_utils_test.go && git commit -m "feat(utils): add FilterByMinor, FilterByPrefix, Difference, Intersection collection operations"`
