# 版本 SDK API 扩展 Plan

> **For agentic workers:** REQUIRED SUB-SKILL: `superpowers:subagent-driven-development`
> Steps use checkbox (`- [ ]`) syntax.

**Goal:** 为版本 SDK 添加常用的便捷方法、查询/过滤 API、构建器模式、版本递增方法，使其作为 SCA 底层库的 API 更完整易用。

**Architecture:** 在现有 `Version` 结构体上添加便捷判断方法（IsPrerelease/IsStable/IsBeta 等）和比较方法（IsNewerThan/IsOlderThan/Equals/IsBetween）；新增 `version_utils.go` 提供批量操作（Min/Max/Filter/Unique/FilterByConstraint）；新增 `version_builder.go` 提供流式 API（VersionBuilder）；为 `VersionGroup` 添加 GetLatest/GetOldest 便捷方法；添加 `ErrConstraintInvalid` 等错误类型。

**Tech Stack:** Go 1.18, 现有依赖不变

**Risks:**
- 所有 Task 都是纯新增（不修改现有方法签名），不存在兼容性风险
- Task 2 的 FilterByConstraint 依赖 constraint.go 的 Constraint 类型，已存在

---

### Task 1: Version 便捷判断和比较方法

**Depends on:** None
**Files:**
- Modify: `version.go`（在 String() 方法后添加新方法）
- Create: `version_convenience_test.go`

- [ ] **Step 1: 修改 Version — 添加 IsPrerelease/IsStable/IsBeta/IsRC 等便捷判断方法**

文件: `version.go`（在 `String()` 方法之后添加以下方法）

```go
// IsPrerelease 判断版本是否为预发布版本
//
// 预发布版本是指带有后缀（如 -alpha, -beta, -rc, -SNAPSHOT 等）的版本。
// 正式版本（无后缀）返回 false。
//
// 返回:
//   - bool: 如果是预发布版本则返回 true
//
// 使用示例:
//
//	v := versions.NewVersion("1.0.0-beta")
//	if v.IsPrerelease() {
//	    fmt.Println("这是预发布版本")
//	}
func (x *Version) IsPrerelease() bool {
	return !x.Suffix.IsEmpty()
}

// IsStable 判断版本是否为正式稳定版本
//
// 正式稳定版本是指不带任何后缀的版本，如 "1.0.0"。
//
// 返回:
//   - bool: 如果是正式稳定版本则返回 true
func (x *Version) IsStable() bool {
	return x.Suffix.IsEmpty()
}

// IsDev 判断版本是否为开发版
//
// 开发版是指后缀包含 dev 标识的版本，如 "1.0.0-dev1"。
func (x *Version) IsDev() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightDev
}

// IsAlpha 判断版本是否为 Alpha 版
func (x *Version) IsAlpha() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightAlpha
}

// IsBeta 判断版本是否为 Beta 版
func (x *Version) IsBeta() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightBeta
}

// IsRC 判断版本是否为候选发布版(RC)
func (x *Version) IsRC() bool {
	w := GetSuffixWeight(string(x.Suffix))
	return w == SuffixWeightRC || w == SuffixWeightCR
}

// IsSnapshot 判断版本是否为快照版
func (x *Version) IsSnapshot() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightSnapshot
}

// IsNewerThan 判断当前版本是否比目标版本更新
//
// 等价于 CompareTo(target) > 0，但语义更清晰。
func (x *Version) IsNewerThan(target *Version) bool {
	return x.CompareTo(target) > 0
}

// IsOlderThan 判断当前版本是否比目标版本更旧
//
// 等价于 CompareTo(target) < 0，但语义更清晰。
func (x *Version) IsOlderThan(target *Version) bool {
	return x.CompareTo(target) < 0
}

// Equals 判断当前版本是否与目标版本相等
//
// 等价于 CompareTo(target) == 0，但语义更清晰。
func (x *Version) Equals(target *Version) bool {
	return x.CompareTo(target) == 0
}

// IsBetween 判断当前版本是否在两个版本之间（包含边界）
//
// 如果 low <= x <= high 则返回 true。
// 如果 low 或 high 为 nil，则忽略对应的边界检查。
func (x *Version) IsBetween(low, high *Version) bool {
	if low != nil && x.CompareTo(low) < 0 {
		return false
	}
	if high != nil && x.CompareTo(high) > 0 {
		return false
	}
	return true
}

// Major 返回主版本号
//
// 如果版本号数字部分为空则返回 0。
func (x *Version) Major() int {
	if len(x.VersionNumbers) > 0 {
		return x.VersionNumbers[0]
	}
	return 0
}

// Minor 返回次版本号
//
// 如果版本号数字部分少于2位则返回 0。
func (x *Version) Minor() int {
	if len(x.VersionNumbers) > 1 {
		return x.VersionNumbers[1]
	}
	return 0
}

// Patch 返回修订版本号
//
// 如果版本号数字部分少于3位则返回 0。
func (x *Version) Patch() int {
	if len(x.VersionNumbers) > 2 {
		return x.VersionNumbers[2]
	}
	return 0
}
```

- [ ] **Step 2: 创建 version_convenience_test.go**

```go
package versions

import "testing"

func TestVersion_IsPrerelease(t *testing.T) {
	if NewVersion("1.0.0").IsPrerelease() {
		t.Error("1.0.0 should not be prerelease")
	}
	if !NewVersion("1.0.0-beta").IsPrerelease() {
		t.Error("1.0.0-beta should be prerelease")
	}
}

func TestVersion_IsStable(t *testing.T) {
	if !NewVersion("1.0.0").IsStable() {
		t.Error("1.0.0 should be stable")
	}
	if NewVersion("1.0.0-rc1").IsStable() {
		t.Error("1.0.0-rc1 should not be stable")
	}
}

func TestVersion_IsBeta(t *testing.T) {
	if !NewVersion("1.0.0-beta").IsBeta() {
		t.Error("1.0.0-beta should be beta")
	}
	if NewVersion("1.0.0-alpha").IsBeta() {
		t.Error("1.0.0-alpha should not be beta")
	}
}

func TestVersion_IsRC(t *testing.T) {
	if !NewVersion("1.0.0-rc1").IsRC() {
		t.Error("1.0.0-rc1 should be RC")
	}
	if !NewVersion("1.0.0-CR1").IsRC() {
		t.Error("1.0.0-CR1 should be RC (CR variant)")
	}
}

func TestVersion_IsNewerThan(t *testing.T) {
	v1 := NewVersion("1.1.0")
	v2 := NewVersion("1.0.0")
	if !v1.IsNewerThan(v2) {
		t.Error("1.1.0 should be newer than 1.0.0")
	}
	if v2.IsNewerThan(v1) {
		t.Error("1.0.0 should not be newer than 1.1.0")
	}
}

func TestVersion_IsOlderThan(t *testing.T) {
	v1 := NewVersion("1.0.0")
	v2 := NewVersion("1.1.0")
	if !v1.IsOlderThan(v2) {
		t.Error("1.0.0 should be older than 1.1.0")
	}
}

func TestVersion_Equals(t *testing.T) {
	v1 := NewVersion("1.0.0")
	v2 := NewVersion("1.0.0")
	if !v1.Equals(v2) {
		t.Error("1.0.0 should equal 1.0.0")
	}
}

func TestVersion_IsBetween(t *testing.T) {
	v := NewVersion("1.5.0")
	low := NewVersion("1.0.0")
	high := NewVersion("2.0.0")
	if !v.IsBetween(low, high) {
		t.Error("1.5.0 should be between 1.0.0 and 2.0.0")
	}
	if NewVersion("0.9.0").IsBetween(low, high) {
		t.Error("0.9.0 should not be between 1.0.0 and 2.0.0")
	}
	if NewVersion("2.1.0").IsBetween(low, high) {
		t.Error("2.1.0 should not be between 1.0.0 and 2.0.0")
	}
}

func TestVersion_MajorMinorPatch(t *testing.T) {
	v := NewVersion("1.2.3")
	if v.Major() != 1 {
		t.Errorf("Major() = %d, want 1", v.Major())
	}
	if v.Minor() != 2 {
		t.Errorf("Minor() = %d, want 2", v.Minor())
	}
	if v.Patch() != 3 {
		t.Errorf("Patch() = %d, want 3", v.Patch())
	}
}

func TestVersion_MajorMinorPatch_Zero(t *testing.T) {
	v := NewVersion("5")
	if v.Major() != 5 {
		t.Errorf("Major() = %d, want 5", v.Major())
	}
	if v.Minor() != 0 {
		t.Errorf("Minor() = %d, want 0", v.Minor())
	}
	if v.Patch() != 0 {
		t.Errorf("Patch() = %d, want 0", v.Patch())
	}
}
```

- [ ] **Step 3: 验证便捷方法**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test ./...`
Expected:
  - Exit code: 0
  - Output contains: "ok"
  - Output does NOT contain: "FAIL"

- [ ] **Step 4: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add version.go version_convenience_test.go && git commit -m "feat(version): add convenience methods — IsPrerelease, IsStable, IsBeta, IsRC, IsNewerThan, Equals, IsBetween, Major/Minor/Patch"`

---

### Task 2: 批量查询和过滤工具函数

**Depends on:** Task 1
**Files:**
- Create: `version_utils.go`
- Create: `version_utils_test.go`

- [ ] **Step 1: 创建 version_utils.go — 批量版本操作工具函数**

```go
package versions

// Min 从版本列表中找到最小的版本
//
// 如果列表为空则返回 nil。
//
// 参数:
//   - versions: 版本对象列表
//
// 返回:
//   - *Version: 最小的版本对象，列表为空时返回 nil
func Min(versions []*Version) *Version {
	if len(versions) == 0 {
		return nil
	}
	min := versions[0]
	for _, v := range versions[1:] {
		if v.CompareTo(min) < 0 {
			min = v
		}
	}
	return min
}

// Max 从版本列表中找到最大的版本
//
// 如果列表为空则返回 nil。
func Max(versions []*Version) *Version {
	if len(versions) == 0 {
		return nil
	}
	max := versions[0]
	for _, v := range versions[1:] {
		if v.CompareTo(max) > 0 {
			max = v
		}
	}
	return max
}

// LatestStable 从版本列表中找到最新的稳定版本
//
// 稳定版本是指不带后缀的版本。如果不存在稳定版本则返回 nil。
func LatestStable(versions []*Version) *Version {
	var latest *Version
	for _, v := range versions {
		if v.IsStable() {
			if latest == nil || v.CompareTo(latest) > 0 {
				latest = v
			}
		}
	}
	return latest
}

// LatestPrerelease 从版本列表中找到最新的预发布版本
//
// 如果不存在预发布版本则返回 nil。
func LatestPrerelease(versions []*Version) *Version {
	var latest *Version
	for _, v := range versions {
		if v.IsPrerelease() {
			if latest == nil || v.CompareTo(latest) > 0 {
				latest = v
			}
		}
	}
	return latest
}

// Filter 根据谓词函数过滤版本列表
//
// 返回所有满足谓词条件的版本。
//
// 参数:
//   - versions: 版本对象列表
//   - predicate: 过滤谓词函数，返回 true 表示保留该版本
//
// 返回:
//   - []*Version: 满足条件的版本列表
func Filter(versions []*Version, predicate func(*Version) bool) []*Version {
	result := make([]*Version, 0)
	for _, v := range versions {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// FilterByConstraint 根据约束条件过滤版本列表
//
// 返回所有满足约束条件的版本。
//
// 参数:
//   - versions: 版本对象列表
//   - constraint: 版本约束条件
//
// 返回:
//   - []*Version: 满足约束的版本列表
func FilterByConstraint(versions []*Version, constraint *Constraint) []*Version {
	return Filter(versions, func(v *Version) bool {
		return constraint.Match(v)
	})
}

// FilterByConstraintSet 根据约束集合过滤版本列表
//
// 返回所有满足约束集合中所有条件的版本。
func FilterByConstraintSet(versions []*Version, cs *ConstraintSet) []*Version {
	return Filter(versions, func(v *Version) bool {
		return cs.Match(v)
	})
}

// Unique 去除版本列表中的重复版本
//
// 根据 Raw 字段去重，保留第一次出现的版本。
func Unique(versions []*Version) []*Version {
	seen := make(map[string]bool)
	result := make([]*Version, 0)
	for _, v := range versions {
		if !seen[v.Raw] {
			seen[v.Raw] = true
			result = append(result, v)
		}
	}
	return result
}

// FilterByMajor 过滤指定主版本号的版本
func FilterByMajor(versions []*Version, major int) []*Version {
	return Filter(versions, func(v *Version) bool {
		return v.Major() == major
	})
}

// Count 统计版本列表中满足谓词的版本数量
func Count(versions []*Version, predicate func(*Version) bool) int {
	n := 0
	for _, v := range versions {
		if predicate(v) {
			n++
		}
	}
	return n
}
```

- [ ] **Step 2: 创建 version_utils_test.go**

```go
package versions

import "testing"

func TestMin(t *testing.T) {
	versions := NewVersions("2.0.0", "1.0.0", "1.5.0")
	min := Min(versions)
	if min.Raw != "1.0.0" {
		t.Errorf("Min() = %s, want 1.0.0", min.Raw)
	}
}

func TestMin_Empty(t *testing.T) {
	if Min(nil) != nil {
		t.Error("Min(nil) should return nil")
	}
}

func TestMax(t *testing.T) {
	versions := NewVersions("1.0.0", "3.0.0", "2.0.0")
	max := Max(versions)
	if max.Raw != "3.0.0" {
		t.Errorf("Max() = %s, want 3.0.0", max.Raw)
	}
}

func TestLatestStable(t *testing.T) {
	versions := NewVersions("1.0.0-alpha", "1.0.0", "1.1.0-beta", "1.1.0")
	latest := LatestStable(versions)
	if latest.Raw != "1.1.0" {
		t.Errorf("LatestStable() = %s, want 1.1.0", latest.Raw)
	}
}

func TestLatestStable_None(t *testing.T) {
	versions := NewVersions("1.0.0-alpha", "1.0.0-beta")
	if LatestStable(versions) != nil {
		t.Error("LatestStable() should return nil when no stable versions")
	}
}

func TestLatestPrerelease(t *testing.T) {
	versions := NewVersions("1.0.0", "1.1.0-alpha", "1.0.0-beta")
	latest := LatestPrerelease(versions)
	if latest.Raw != "1.1.0-alpha" {
		t.Errorf("LatestPrerelease() = %s, want 1.1.0-alpha", latest.Raw)
	}
}

func TestFilter(t *testing.T) {
	versions := NewVersions("1.0.0", "1.0.0-beta", "2.0.0-alpha", "2.0.0")
	stable := Filter(versions, func(v *Version) bool { return v.IsStable() })
	if len(stable) != 2 {
		t.Errorf("Filter stable = %d, want 2", len(stable))
	}
}

func TestFilterByConstraint(t *testing.T) {
	versions := NewVersions("0.9.0", "1.0.0", "1.5.0", "2.0.0")
	c, _ := ParseConstraint(">=1.0.0")
	result := FilterByConstraint(versions, c)
	if len(result) != 3 {
		t.Errorf("FilterByConstraint >=1.0.0 = %d, want 3", len(result))
	}
}

func TestFilterByConstraintSet(t *testing.T) {
	versions := NewVersions("0.9.0", "1.0.0", "1.5.0", "2.0.0")
	cs, _ := ParseConstraintSet(">=1.0.0,<2.0.0")
	result := FilterByConstraintSet(versions, cs)
	if len(result) != 2 {
		t.Errorf("FilterByConstraintSet >=1.0.0,<2.0.0 = %d, want 2", len(result))
	}
}

func TestUnique(t *testing.T) {
	versions := NewVersions("1.0.0", "1.0.0", "2.0.0", "2.0.0")
	result := Unique(versions)
	if len(result) != 2 {
		t.Errorf("Unique() = %d, want 2", len(result))
	}
}

func TestFilterByMajor(t *testing.T) {
	versions := NewVersions("1.0.0", "1.1.0", "2.0.0", "2.1.0")
	result := FilterByMajor(versions, 2)
	if len(result) != 2 {
		t.Errorf("FilterByMajor(2) = %d, want 2", len(result))
	}
}

func TestCount(t *testing.T) {
	versions := NewVersions("1.0.0", "1.0.0-beta", "2.0.0-alpha")
	n := Count(versions, func(v *Version) bool { return v.IsPrerelease() })
	if n != 2 {
		t.Errorf("Count prerelease = %d, want 2", n)
	}
}
```

- [ ] **Step 3: 验证工具函数**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test ./...`
Expected:
  - Exit code: 0
  - Output contains: "ok"
  - Output does NOT contain: "FAIL"

- [ ] **Step 4: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add version_utils.go version_utils_test.go && git commit -m "feat(utils): add batch operations — Min, Max, LatestStable, Filter, FilterByConstraint, Unique"`

---

### Task 3: Version 构建器和版本递增

**Depends on:** Task 1
**Files:**
- Create: `version_builder.go`
- Create: `version_builder_test.go`

- [ ] **Step 1: 创建 version_builder.go — 流式 API 和版本递增方法**

```go
package versions

import "strings"

// VersionBuilder 提供流式 API 构建版本对象
//
// VersionBuilder 允许通过方法链的方式逐步构建版本对象，
// 适用于需要程序化生成版本号的场景。
//
// 使用示例:
//
//	v := versions.NewVersionBuilder().
//	    Prefix("v").
//	    Major(1).
//	    Minor(2).
//	    Patch(3).
//	    Suffix("-beta1").
//	    Build()
//	// v.Raw == "v1.2.3-beta1"
type VersionBuilder struct {
	prefix  string
	numbers []int
	suffix  string
}

// NewVersionBuilder 创建一个新的版本构建器
func NewVersionBuilder() *VersionBuilder {
	return &VersionBuilder{
		numbers: make([]int, 0),
	}
}

// Prefix 设置版本前缀
func (b *VersionBuilder) Prefix(prefix string) *VersionBuilder {
	b.prefix = prefix
	return b
}

// Major 设置主版本号
func (b *VersionBuilder) Major(major int) *VersionBuilder {
	b.ensureLength(1)
	b.numbers[0] = major
	return b
}

// Minor 设置次版本号
func (b *VersionBuilder) Minor(minor int) *VersionBuilder {
	b.ensureLength(2)
	b.numbers[1] = minor
	return b
}

// Patch 设置修订版本号
func (b *VersionBuilder) Patch(patch int) *VersionBuilder {
	b.ensureLength(3)
	b.numbers[2] = patch
	return b
}

// Numbers 设置版本号数字部分
func (b *VersionBuilder) Numbers(numbers []int) *VersionBuilder {
	b.numbers = make([]int, len(numbers))
	copy(b.numbers, numbers)
	return b
}

// Suffix 设置版本后缀
func (b *VersionBuilder) Suffix(suffix string) *VersionBuilder {
	b.suffix = suffix
	return b
}

// Build 构建并返回版本对象
func (b *VersionBuilder) Build() *Version {
	raw := b.buildRawString()
	return NewVersion(raw)
}

// buildRawString 从构建器组件重建版本字符串
func (b *VersionBuilder) buildRawString() string {
	var sb strings.Builder
	sb.WriteString(b.prefix)
	for i, n := range b.numbers {
		if i > 0 {
			sb.WriteString(DefaultVersionDelimiter)
		}
		sb.WriteString(strings.TrimLeft(strings.TrimSpace(intToString(n)), "0"))
		if n == 0 {
			sb.WriteString("0")
		}
	}
	sb.WriteString(b.suffix)
	return sb.String()
}

// ensureLength 确保版本号数组至少有指定长度
func (b *VersionBuilder) ensureLength(minLen int) {
	for len(b.numbers) < minLen {
		b.numbers = append(b.numbers, 0)
	}
}

// intToString 将整数转换为字符串（避免导入 strconv）
func intToString(n int) string {
	if n == 0 {
		return "0"
	}
	negative := n < 0
	if negative {
		n = -n
	}
	var digits []byte
	for n > 0 {
		digits = append([]byte{byte('0' + n%10)}, digits...)
		n /= 10
	}
	if negative {
		digits = append([]byte{'-'}, digits...)
	}
	return string(digits)
}

// BumpMajor 返回一个主版本号递增的新版本对象
//
// 例如 1.2.3 → 2.0.0，后缀被清除。
func (x *Version) BumpMajor() *Version {
	if len(x.VersionNumbers) == 0 {
		return NewVersion("1")
	}
	return NewVersionBuilder().
		Prefix(string(x.Prefix)).
		Major(x.Major()+1).
		Minor(0).
		Patch(0).
		Build()
}

// BumpMinor 返回一个次版本号递增的新版本对象
//
// 例如 1.2.3 → 1.3.0，后缀被清除。
func (x *Version) BumpMinor() *Version {
	if len(x.VersionNumbers) == 0 {
		return NewVersion("0.1")
	}
	return NewVersionBuilder().
		Prefix(string(x.Prefix)).
		Major(x.Major()).
		Minor(x.Minor()+1).
		Patch(0).
		Build()
}

// BumpPatch 返回一个修订版本号递增的新版本对象
//
// 例如 1.2.3 → 1.2.4，后缀被清除。
func (x *Version) BumpPatch() *Version {
	if len(x.VersionNumbers) == 0 {
		return NewVersion("0.0.1")
	}
	return NewVersionBuilder().
		Prefix(string(x.Prefix)).
		Major(x.Major()).
		Minor(x.Minor()).
		Patch(x.Patch()+1).
		Build()
}
```

- [ ] **Step 2: 创建 version_builder_test.go**

```go
package versions

import "testing"

func TestVersionBuilder_FullBuild(t *testing.T) {
	v := NewVersionBuilder().
		Prefix("v").
		Major(1).
		Minor(2).
		Patch(3).
		Suffix("-beta1").
		Build()
	if !v.IsValid() {
		t.Fatal("builder result should be valid")
	}
	if v.Major() != 1 || v.Minor() != 2 || v.Patch() != 3 {
		t.Errorf("builder numbers = %v, want [1,2,3]", v.VersionNumbers)
	}
}

func TestVersionBuilder_Minimal(t *testing.T) {
	v := NewVersionBuilder().Major(5).Build()
	if v.Major() != 5 {
		t.Errorf("Major() = %d, want 5", v.Major())
	}
}

func TestVersionBuilder_Numbers(t *testing.T) {
	v := NewVersionBuilder().Numbers([]int{2, 0, 1}).Build()
	if v.Major() != 2 || v.Minor() != 0 || v.Patch() != 1 {
		t.Errorf("numbers = %v, want [2,0,1]", v.VersionNumbers)
	}
}

func TestVersion_BumpMajor(t *testing.T) {
	v := NewVersion("1.2.3")
	bumped := v.BumpMajor()
	if bumped.Major() != 2 {
		t.Errorf("BumpMajor() Major = %d, want 2", bumped.Major())
	}
	if bumped.Minor() != 0 {
		t.Errorf("BumpMajor() Minor = %d, want 0", bumped.Minor())
	}
	if bumped.Patch() != 0 {
		t.Errorf("BumpMajor() Patch = %d, want 0", bumped.Patch())
	}
}

func TestVersion_BumpMinor(t *testing.T) {
	v := NewVersion("1.2.3")
	bumped := v.BumpMinor()
	if bumped.Major() != 1 {
		t.Errorf("BumpMinor() Major = %d, want 1", bumped.Major())
	}
	if bumped.Minor() != 3 {
		t.Errorf("BumpMinor() Minor = %d, want 3", bumped.Minor())
	}
	if bumped.Patch() != 0 {
		t.Errorf("BumpMinor() Patch = %d, want 0", bumped.Patch())
	}
}

func TestVersion_BumpPatch(t *testing.T) {
	v := NewVersion("1.2.3")
	bumped := v.BumpPatch()
	if bumped.Patch() != 4 {
		t.Errorf("BumpPatch() Patch = %d, want 4", bumped.Patch())
	}
}

func TestVersion_BumpPatch_ClearsSuffix(t *testing.T) {
	v := NewVersion("1.2.3-beta")
	bumped := v.BumpPatch()
	if bumped.IsPrerelease() {
		t.Error("BumpPatch() should clear suffix")
	}
}
```

- [ ] **Step 3: 验证构建器**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test ./...`
Expected:
  - Exit code: 0
  - Output contains: "ok"
  - Output does NOT contain: "FAIL"

- [ ] **Step 4: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add version_builder.go version_builder_test.go && git commit -m "feat(builder): add VersionBuilder fluent API and BumpMajor/Minor/Patch methods"`

---

### Task 4: VersionGroup 便捷方法

**Depends on:** Task 1
**Files:**
- Modify: `version_group.go`（在 SortVersions 后添加新方法）
- Create: `version_group_convenience_test.go`

- [ ] **Step 1: 修改 VersionGroup — 添加 GetLatest/GetOldest/Count/StableVersions 方法**

文件: `version_group.go`（在 `SortVersions()` 方法之后添加以下方法）

```go
// GetLatest 获取版本组中最新的版本
//
// 返回按排序后最新的版本，如果组为空则返回 nil。
func (x *VersionGroup) GetLatest() *Version {
	sorted := x.SortVersions()
	if len(sorted) == 0 {
		return nil
	}
	return sorted[len(sorted)-1]
}

// GetOldest 获取版本组中最旧的版本
//
// 返回按排序后最旧的版本，如果组为空则返回 nil。
func (x *VersionGroup) GetOldest() *Version {
	sorted := x.SortVersions()
	if len(sorted) == 0 {
		return nil
	}
	return sorted[0]
}

// Count 返回版本组中的版本数量
func (x *VersionGroup) Count() int {
	return len(x.VersionMap)
}

// StableVersions 返回版本组中所有稳定版本
func (x *VersionGroup) StableVersions() []*Version {
	return Filter(x.Versions(), func(v *Version) bool {
		return v.IsStable()
	})
}

// PrereleaseVersions 返回版本组中所有预发布版本
func (x *VersionGroup) PrereleaseVersions() []*Version {
	return Filter(x.Versions(), func(v *Version) bool {
		return v.IsPrerelease()
	})
}

// LatestStable 获取版本组中最新稳定版本
func (x *VersionGroup) LatestStable() *Version {
	return LatestStable(x.Versions())
}
```

- [ ] **Step 2: 创建 version_group_convenience_test.go**

```go
package versions

import "testing"

func TestVersionGroup_GetLatest(t *testing.T) {
	vg := NewVersionGroupFromVersions(NewVersions("1.0.0", "1.0.1", "1.0.2"))
	latest := vg.GetLatest()
	if latest == nil || latest.Raw != "1.0.2" {
		t.Errorf("GetLatest() = %v, want 1.0.2", latest)
	}
}

func TestVersionGroup_GetOldest(t *testing.T) {
	vg := NewVersionGroupFromVersions(NewVersions("1.0.0", "1.0.1", "1.0.2"))
	oldest := vg.GetOldest()
	if oldest == nil || oldest.Raw != "1.0.0" {
		t.Errorf("GetOldest() = %v, want 1.0.0", oldest)
	}
}

func TestVersionGroup_Count(t *testing.T) {
	vg := NewVersionGroupFromVersions(NewVersions("1.0.0", "1.0.1"))
	if vg.Count() != 2 {
		t.Errorf("Count() = %d, want 2", vg.Count())
	}
}

func TestVersionGroup_StableVersions(t *testing.T) {
	vg := NewVersionGroupFromVersions(NewVersions("1.0.0", "1.0.0-beta", "1.0.1"))
	stable := vg.StableVersions()
	if len(stable) != 2 {
		t.Errorf("StableVersions() = %d, want 2", len(stable))
	}
}

func TestVersionGroup_LatestStable(t *testing.T) {
	vg := NewVersionGroupFromVersions(NewVersions("1.0.0-beta", "1.0.0", "1.0.1-alpha"))
	latest := vg.LatestStable()
	if latest == nil || latest.Raw != "1.0.0" {
		t.Errorf("LatestStable() = %v, want 1.0.0", latest)
	}
}
```

- [ ] **Step 3: 验证 VersionGroup 便捷方法**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test ./...`
Expected:
  - Exit code: 0
  - Output contains: "ok"
  - Output does NOT contain: "FAIL"

- [ ] **Step 4: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add version_group.go version_group_convenience_test.go && git commit -m "feat(group): add GetLatest, GetOldest, Count, StableVersions, LatestStable convenience methods"`

---

### Task 5: 错误类型增强

**Depends on:** None
**Files:**
- Modify: `constraint.go:83-87,107,111,125`（用错误变量替换 fmt.Errorf）
- Modify: `version.go:12-16`（添加新错误变量）
- Create: `errors_test.go`

- [ ] **Step 1: 修改 version.go — 添加新错误变量**

文件: `version.go:12-16`（在现有 ErrVersionInvalid 后添加）

```go
var (
	// ErrVersionInvalid 表示版本号格式无效的错误
	ErrVersionInvalid = errors.New("version invalid")

	// ErrConstraintInvalid 表示版本约束表达式格式无效的错误
	ErrConstraintInvalid = errors.New("constraint invalid")

	// ErrEmptyConstraint 表示约束表达式为空的错误
	ErrEmptyConstraint = errors.New("empty constraint expression")

	// ErrMissingVersionInConstraint 表示约束表达式中缺少版本号的错误
	ErrMissingVersionInConstraint = errors.New("missing version in constraint")

	// ErrInvalidVersionInConstraint 表示约束表达式中版本号无效的错误
	ErrInvalidVersionInConstraint = errors.New("invalid version in constraint")
)
```

- [ ] **Step 2: 修改 constraint.go — 使用错误变量**

文件: `constraint.go:83-87`（替换 ParseConstraint 中的 fmt.Errorf 调用）

替换 `return nil, fmt.Errorf("empty constraint expression")` 为：
```go
return nil, ErrEmptyConstraint
```

替换 `return nil, fmt.Errorf("missing version after operator %s", o.op)` 为：
```go
return nil, ErrMissingVersionInConstraint
```

替换 `return nil, fmt.Errorf("invalid version in constraint: %s", versionStr)` 为：
```go
return nil, ErrInvalidVersionInConstraint
```

替换 `return nil, fmt.Errorf("invalid version in constraint: %s", expr)` 为：
```go
return nil, ErrInvalidVersionInConstraint
```

- [ ] **Step 3: 创建 errors_test.go**

```go
package versions

import (
	"errors"
	"testing"
)

func TestErrConstraintInvalid(t *testing.T) {
	_, err := ParseConstraint("")
	if !errors.Is(err, ErrEmptyConstraint) {
		t.Errorf("ParseConstraint(\"\") error = %v, want ErrEmptyConstraint", err)
	}
}

func TestErrMissingVersionInConstraint(t *testing.T) {
	_, err := ParseConstraint(">=")
	if !errors.Is(err, ErrMissingVersionInConstraint) {
		t.Errorf("ParseConstraint(\">=\") error = %v, want ErrMissingVersionInConstraint", err)
	}
}

func TestErrInvalidVersionInConstraint(t *testing.T) {
	_, err := ParseConstraint(">=not-a-version")
	if !errors.Is(err, ErrInvalidVersionInConstraint) {
		t.Errorf("ParseConstraint(\">=not-a-version\") error = %v, want ErrInvalidVersionInConstraint", err)
	}
}

func TestErrVersionInvalid(t *testing.T) {
	_, err := NewVersionE("not-a-version")
	if !errors.Is(err, ErrVersionInvalid) {
		t.Errorf("NewVersionE error = %v, want ErrVersionInvalid", err)
	}
}
```

- [ ] **Step 4: 验证错误类型**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test ./...`
Expected:
  - Exit code: 0
  - Output contains: "ok"
  - Output does NOT contain: "FAIL"

- [ ] **Step 5: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add version.go constraint.go errors_test.go && git commit -m "feat(errors): add typed error variables for constraint parsing and version validation"`
