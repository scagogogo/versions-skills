# SCA 核心功能增强 Plan (Phase 1)

> **For agentic workers:** REQUIRED SUB-SKILL: `superpowers:subagent-driven-development`
> Steps use checkbox (`- [ ]`) syntax.

**Goal:** 修复版本比较核心 bug，引入后缀语义权重系统，支持可配置分隔符解析，新增版本约束表达式解析器 — 使 SDK 达到 SCA 产品底层库的最低可用标准。

**Architecture:** 修复 CompareTo 流程：当前 VersionNumbers → PublicTime → Suffix → Raw 存在 bug（空后缀 vs 有后缀比较被跳过），改为 VersionNumbers → Suffix（含语义权重） → PublicTime → Raw。新增 `suffix_weight.go` 实现后缀权重映射。新增 `parser_option.go` 支持可配置分隔符。新增 `constraint.go` 实现约束表达式解析和匹配，支持 `>=`, `<=`, `^`, `~`, `x` 等操作符及 AND 组合约束。

**Tech Stack:** Go 1.18, stretchr/testify v1.8.2, 现有依赖不变

**Risks:**
- Task 1 修改 CompareTo 核心逻辑可能破坏现有测试 → 缓解：保持现有测试通过，增加新测试覆盖修复场景
- Task 2 后缀权重系统可能改变 suffix 比较的现有行为 → 缓解：未知后缀 fallback 到字典序，权重系统仅增强已知后缀
- Task 3 多分隔符支持可能改变已有版本的解析结果 → 缓解：通过 ParserOption 控制，默认仅 `.`，不改变 NewVersion 行为
- Task 4 约束表达式解析器是新功能，不影响现有代码 → 无风险

---

### Task 1: 修复 CompareTo 核心比较 bug

**Depends on:** None
**Files:**
- Modify: `version.go:198-240`（CompareTo 方法 — 修复空后缀跳过 bug）
- Modify: `version_numbers.go:82-93`（CompareTo — 修复整数溢出）
- Create: `version_compareto_test.go`

- [ ] **Step 1: 修改 Version.CompareTo — 修复空后缀 vs 有后缀比较被跳过的 bug**

当前 bug（`version.go:222`）：当一方有后缀另一方没有时，后缀比较被跳过，导致 `1.0.0` 和 `1.0.0-beta` 无法正确区分。语义上，无后缀（release）应大于有后缀（pre-release）。

文件: `version.go:198-240`（替换整个 CompareTo 方法）

```go
func (x *Version) CompareTo(target *Version) int {

	// 1. 先按照主版本号排序，仅当两个的主版本号都存在的时候才会进行比较，它们的长度不必相等，但是不能有为空的
	if len(x.VersionNumbers) != 0 && len(target.VersionNumbers) != 0 {
		r := x.VersionNumbers.CompareTo(target.VersionNumbers)
		if r != 0 {
			return r
		}
	}

	// 2. 然后按照后缀排序，修复：空后缀(release)应大于有后缀(pre-release)
	switch {
	case x.Suffix.IsEmpty() && target.Suffix.IsEmpty():
		// 两者都是正式版，后缀相等，继续比较
	case x.Suffix.IsEmpty() && !target.Suffix.IsEmpty():
		// 当前是正式版，目标是预发布版，当前更大
		return 1
	case !x.Suffix.IsEmpty() && target.Suffix.IsEmpty():
		// 当前是预发布版，目标是正式版，当前更小
		return -1
	default:
		// 两者都有后缀，比较后缀
		r := x.Suffix.CompareTo(target.Suffix)
		if r != 0 {
			return r
		}
	}

	// 3. 然后按照发布时间排序
	if !target.PublicTime.IsZero() && !x.PublicTime.IsZero() {
		r2 := x.PublicTime.UnixMilli() - target.PublicTime.UnixMilli()
		if r2 != 0 {
			if r2 > 0 {
				return 1
			} else {
				return -1
			}
		}
	}

	// 4. 最后比较原始版本号的字典序
	if x.Raw == target.Raw {
		return 0
	} else if x.Raw < target.Raw {
		return -1
	} else if x.Raw > target.Raw {
		return 1
	}

	// unreachable
	return 0
}
```

- [ ] **Step 2: 修改 VersionNumbers.CompareTo — 修复整数溢出风险**

当前 `version_numbers.go:86-87` 使用 `x[i] - target[i]` 可能整数溢出。

文件: `version_numbers.go:82-93`（替换 CompareTo 方法）

```go
func (x VersionNumbers) CompareTo(target []int) int {
	minLen := len(x)
	if len(target) < minLen {
		minLen = len(target)
	}
	for i := 0; i < minLen; i++ {
		if x[i] < target[i] {
			return -1
		} else if x[i] > target[i] {
			return 1
		}
	}
	// 共有部分相等，比较长度
	if len(x) < len(target) {
		return -1
	} else if len(x) > len(target) {
		return 1
	}
	return 0
}
```

- [ ] **Step 3: 创建 version_compareto_test.go — 覆盖修复场景**

```go
package versions

import (
	"testing"
	"time"
)

func TestVersion_CompareTo_ReleaseVsPreRelease(t *testing.T) {
	release := NewVersion("1.0.0")
	beta := NewVersion("1.0.0-beta")
	alpha := NewVersion("1.0.0-alpha")

	if release.CompareTo(beta) <= 0 {
		t.Errorf("1.0.0 (release) should be greater than 1.0.0-beta, got %d", release.CompareTo(beta))
	}
	if release.CompareTo(alpha) <= 0 {
		t.Errorf("1.0.0 (release) should be greater than 1.0.0-alpha, got %d", release.CompareTo(alpha))
	}
	if beta.CompareTo(alpha) <= 0 {
		t.Errorf("1.0.0-beta should be greater than 1.0.0-alpha (alphabetic), got %d", beta.CompareTo(alpha))
	}
}

func TestVersion_CompareTo_DifferentLengths(t *testing.T) {
	short := NewVersion("1.0")
	full := NewVersion("1.0.0")
	if short.CompareTo(full) != 0 {
		t.Errorf("1.0 and 1.0.0 should be equal, got %d", short.CompareTo(full))
	}
}

func TestVersion_CompareTo_TimeBased(t *testing.T) {
	v1 := NewVersion("1.0.0")
	v2 := NewVersion("1.0.0")
	v1.PublicTime = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	v2.PublicTime = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	if v1.CompareTo(v2) >= 0 {
		t.Errorf("older version should be less than newer version")
	}
}

func TestVersionNumbers_CompareTo_NoOverflow(t *testing.T) {
	big := NewVersionNumbers([]int{2147483647, 0})
	bigger := NewVersionNumbers([]int{2147483648, 0})
	if big.CompareTo(bigger) >= 0 {
		t.Errorf("2147483647.0 should be less than 2147483648.0")
	}
}
```

- [ ] **Step 4: 验证 CompareTo 修复**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test ./...`
Expected:
  - Exit code: 0
  - Output contains: "ok"
  - Output does NOT contain: "FAIL"

- [ ] **Step 5: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add version.go version_numbers.go version_compareto_test.go && git commit -m "fix(compare): fix CompareTo release vs pre-release ordering and integer overflow"`

---

### Task 2: 后缀语义权重系统

**Depends on:** Task 1
**Files:**
- Create: `suffix_weight.go`
- Create: `suffix_weight_test.go`
- Modify: `version_suffix.go:82-90`（CompareTo 使用权重）

- [ ] **Step 1: 创建 suffix_weight.go — 定义后缀语义权重映射**

```go
package versions

import (
	"regexp"
	"strings"
)

// SuffixWeight 表示版本后缀的语义权重
//
// SuffixWeight 用于在版本比较时为不同类型的后缀分配语义化的权重值，
// 使得预发布版本的排序符合实际发布顺序，而非简单的字典序。
//
// 权重规则（从低到高）:
//   - dev/snapshot < alpha/a < beta/b < milestone/m < rc/cr/pre < 正式版(无后缀)
//
// 使用示例:
//
//	weight := GetSuffixWeight("-alpha1")
//	fmt.Println(weight) // 100
type SuffixWeight int

const (
	SuffixWeightUnknown  SuffixWeight = 0
	SuffixWeightDev      SuffixWeight = 50
	SuffixWeightSnapshot SuffixWeight = 60
	SuffixWeightNightly  SuffixWeight = 70
	SuffixWeightAlpha    SuffixWeight = 100
	SuffixWeightBeta     SuffixWeight = 200
	SuffixWeightMilestone SuffixWeight = 300
	SuffixWeightRC       SuffixWeight = 400
	SuffixWeightPre      SuffixWeight = 410
	SuffixWeightCR       SuffixWeight = 420
	SuffixWeightFinal    SuffixWeight = 500
	SuffixWeightRelease  SuffixWeight = 500
	SuffixWeightGA       SuffixWeight = 500
	SuffixWeightSP       SuffixWeight = 600
	SuffixWeightPatch    SuffixWeight = 700
	SuffixWeightPost     SuffixWeight = 800
)

// suffixWeightPatterns 后缀权重匹配规则，按优先级排列
var suffixWeightPatterns = []struct {
	Pattern *regexp.Regexp
	Weight  SuffixWeight
}{
	{regexp.MustCompile(`(?i)^[-.]?dev\.?\d*$`), SuffixWeightDev},
	{regexp.MustCompile(`(?i)^[-.]?snapshot.*$`), SuffixWeightSnapshot},
	{regexp.MustCompile(`(?i)^[-.]?nightly.*$`), SuffixWeightNightly},
	{regexp.MustCompile(`(?i)^[-.]?a(?:lpha)?\.?\d*$`), SuffixWeightAlpha},
	{regexp.MustCompile(`(?i)^[-.]?b(?:eta)?\.?\d*$`), SuffixWeightBeta},
	{regexp.MustCompile(`(?i)^[-.]?m(?:ilestone)?\.?\d*$`), SuffixWeightMilestone},
	{regexp.MustCompile(`(?i)^[-.]?rc\.?\d*$`), SuffixWeightRC},
	{regexp.MustCompile(`(?i)^[-.]?cr\.?\d*$`), SuffixWeightCR},
	{regexp.MustCompile(`(?i)^[-.]?pre\.?\d*$`), SuffixWeightPre},
	{regexp.MustCompile(`(?i)^[-.]?final$`), SuffixWeightFinal},
	{regexp.MustCompile(`(?i)^[-.]?release\.?\d*$`), SuffixWeightRelease},
	{regexp.MustCompile(`(?i)^[-.]?ga$`), SuffixWeightGA},
	{regexp.MustCompile(`(?i)^[-.]?sp\.?\d*$`), SuffixWeightSP},
	{regexp.MustCompile(`(?i)^[-.]?patch\.?\d*$`), SuffixWeightPatch},
	{regexp.MustCompile(`(?i)^[-.]?post\.?\d*$`), SuffixWeightPost},
}

// GetSuffixWeight 获取后缀的语义权重
func GetSuffixWeight(suffix string) SuffixWeight {
	normalized := strings.ToLower(strings.TrimSpace(suffix))
	for _, rule := range suffixWeightPatterns {
		if rule.Pattern.MatchString(normalized) {
			return rule.Weight
		}
	}
	return SuffixWeightUnknown
}

// extractSubVersion 从后缀中提取子版本号，如 "-alpha1" 提取出 1
func extractSubVersion(suffix string) int {
	re := regexp.MustCompile(`(\d+)$`)
	matches := re.FindStringSubmatch(suffix)
	if len(matches) > 1 {
		n := 0
		for _, c := range matches[1] {
			n = n*10 + int(c-'0')
		}
		return n
	}
	return 0
}
```

- [ ] **Step 2: 修改 VersionSuffix.CompareTo — 使用语义权重**

文件: `version_suffix.go:82-90`（替换 CompareTo 方法）

```go
func (x VersionSuffix) CompareTo(target VersionSuffix) int {
	weightX := GetSuffixWeight(string(x))
	weightTarget := GetSuffixWeight(string(target))

	// 如果两个后缀都能匹配到权重规则，则按权重比较
	if weightX != SuffixWeightUnknown && weightTarget != SuffixWeightUnknown {
		if weightX < weightTarget {
			return -1
		} else if weightX > weightTarget {
			return 1
		}
		// 权重相同，比较子版本号（如 alpha1 vs alpha2）
		subX := extractSubVersion(string(x))
		subTarget := extractSubVersion(string(target))
		if subX < subTarget {
			return -1
		} else if subX > subTarget {
			return 1
		}
		return 0
	}

	// 如果其中一个匹配权重规则，另一个不匹配
	if weightX != SuffixWeightUnknown && weightTarget == SuffixWeightUnknown {
		return -1 // 未知后缀排在已知后缀之后
	}
	if weightX == SuffixWeightUnknown && weightTarget != SuffixWeightUnknown {
		return 1
	}

	// 都无法匹配权重规则，退化为字典序比较
	if x > target {
		return 1
	} else if x < target {
		return -1
	}
	return 0
}
```

- [ ] **Step 3: 创建 suffix_weight_test.go**

```go
package versions

import "testing"

func TestGetSuffixWeight(t *testing.T) {
	tests := []struct {
		suffix string
		weight SuffixWeight
	}{
		{"-alpha", SuffixWeightAlpha},
		{"-alpha1", SuffixWeightAlpha},
		{"-ALPHA", SuffixWeightAlpha},
		{"-beta", SuffixWeightBeta},
		{"-beta2", SuffixWeightBeta},
		{"-rc", SuffixWeightRC},
		{"-RC1", SuffixWeightRC},
		{"-cr", SuffixWeightCR},
		{"-dev", SuffixWeightDev},
		{"-snapshot", SuffixWeightSnapshot},
		{"-milestone", SuffixWeightMilestone},
		{"-M1", SuffixWeightMilestone},
		{".Final", SuffixWeightFinal},
		{"-sp1", SuffixWeightSP},
		{"-unknown", SuffixWeightUnknown},
	}
	for _, tt := range tests {
		got := GetSuffixWeight(tt.suffix)
		if got != tt.weight {
			t.Errorf("GetSuffixWeight(%q) = %d, want %d", tt.suffix, got, tt.weight)
		}
	}
}

func TestVersionSuffix_CompareTo_WithWeight(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		expect int
	}{
		{"-alpha", "-beta", -1},
		{"-beta", "-rc", -1},
		{"-alpha1", "-alpha2", -1},
		{"-RC1", "-beta1", 1},
		{"-ALPHA", "-beta", -1},
		{"-beta1", "-beta2", -1},
	}
	for _, tt := range tests {
		a := VersionSuffix(tt.a)
		b := VersionSuffix(tt.b)
		got := a.CompareTo(b)
		if got != tt.expect {
			t.Errorf("VersionSuffix(%q).CompareTo(%q) = %d, want %d", tt.a, tt.b, got, tt.expect)
		}
	}
}

func TestExtractSubVersion(t *testing.T) {
	tests := []struct {
		suffix string
		expect int
	}{
		{"-alpha1", 1},
		{"-beta2", 2},
		{"-rc10", 10},
		{"-alpha", 0},
	}
	for _, tt := range tests {
		got := extractSubVersion(tt.suffix)
		if got != tt.expect {
			t.Errorf("extractSubVersion(%q) = %d, want %d", tt.suffix, got, tt.expect)
		}
	}
}
```

- [ ] **Step 4: 验证后缀权重系统**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test ./...`
Expected:
  - Exit code: 0
  - Output contains: "ok"
  - Output does NOT contain: "FAIL"

- [ ] **Step 5: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add suffix_weight.go suffix_weight_test.go version_suffix.go && git commit -m "feat(suffix): add semantic weight system for version suffix comparison"`

---

### Task 3: 多分隔符解析支持

**Depends on:** Task 1
**Files:**
- Create: `parser_option.go`
- Modify: `parser.go:33-44`（VersionStringParser 添加 option 字段）
- Modify: `parser.go:329-331`（IsVersionNumberDelimiter 使用选项）
- Create: `parser_option_test.go`

- [ ] **Step 1: 创建 parser_option.go — 定义解析器选项**

```go
package versions

// ParserOption 配置版本号解析器的行为
//
// ParserOption 允许调用者自定义解析器支持的数字分隔符，
// 以适应不同语言生态系统的版本号格式差异。
//
// 使用示例:
//
//	// 支持 underscore 分隔（Python/RPM 生态）
//	v := versions.NewVersionWithOption("1_2_3",
//	    versions.ParserOption{Delimiters: ".-_"})
type ParserOption struct {
	// Delimiters 版本号数字部分的分隔符集合
	// 默认值: "." (仅点号)
	// 常见扩展: ".-_" (支持 RPM/Debian 的连字符和 Python 的下划线)
	Delimiters string
}

// DefaultParserOption 返回默认的解析器选项
func DefaultParserOption() ParserOption {
	return ParserOption{
		Delimiters: ".",
	}
}

// NewVersionWithOption 使用指定选项创建版本对象
func NewVersionWithOption(versionStr string, option ParserOption) *Version {
	return NewVersionStringParserWithOptions(versionStr, option).Parse()
}

// NewVersionStringParserWithOptions 创建带选项的版本号解析器
func NewVersionStringParserWithOptions(versionStr string, option ParserOption) *VersionStringParser {
	p := NewVersionStringParser(versionStr)
	p.option = option
	return p
}
```

- [ ] **Step 2: 修改 VersionStringParser — 添加 option 字段**

文件: `parser.go:33-44`（在 VersionStringParser 结构体末尾 `v *Version` 之后添加 `option` 字段）

在 `v *Version` 行之后添加：

```go
	// option 解析器选项
	option ParserOption
```

- [ ] **Step 3: 修改 IsVersionNumberDelimiter — 使用选项中的分隔符**

文件: `parser.go:329-331`（替换整个方法）

```go
func (x *VersionStringParser) IsVersionNumberDelimiter(c rune) bool {
	delimiters := x.option.Delimiters
	if delimiters == "" {
		delimiters = DefaultParserOption().Delimiters
	}
	for _, d := range delimiters {
		if c == d {
			return true
		}
	}
	return false
}
```

- [ ] **Step 4: 创建 parser_option_test.go**

```go
package versions

import "testing"

func TestParserOption_UnderscoreDelimiter(t *testing.T) {
	v := NewVersionWithOption("1_2_3", ParserOption{Delimiters: ".-_"})
	if !v.IsValid() {
		t.Fatalf("1_2_3 should be valid with underscore delimiter")
	}
}

func TestParserOption_HyphenDelimiter(t *testing.T) {
	v := NewVersionWithOption("1-2-3", ParserOption{Delimiters: ".-_"})
	if !v.IsValid() {
		t.Fatalf("1-2-3 should be valid with hyphen delimiter")
	}
}

func TestParserOption_DefaultOnlyDot(t *testing.T) {
	v := NewVersion("1.2.3")
	if !v.IsValid() {
		t.Fatalf("1.2.3 should be valid with default options")
	}
	if v.VersionNumbers[0] != 1 || v.VersionNumbers[1] != 2 || v.VersionNumbers[2] != 3 {
		t.Fatalf("1.2.3 should parse as [1,2,3], got %v", v.VersionNumbers)
	}
}

func TestParserOption_PythonStyleVersion(t *testing.T) {
	v := NewVersionWithOption("2023_09_15", ParserOption{Delimiters: ".-_"})
	if !v.IsValid() {
		t.Fatalf("2023_09_15 should be valid with underscore delimiter")
	}
}
```

- [ ] **Step 5: 验证多分隔符支持**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test ./...`
Expected:
  - Exit code: 0
  - Output contains: "ok"
  - Output does NOT contain: "FAIL"

- [ ] **Step 6: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add parser_option.go parser_option_test.go parser.go && git commit -m "feat(parser): add configurable delimiter support via ParserOption"`

---

### Task 4: 版本约束表达式解析器

**Depends on:** Task 1
**Files:**
- Create: `constraint.go`
- Create: `constraint_test.go`

- [ ] **Step 1: 创建 constraint.go — 版本约束表达式解析和匹配**

```go
package versions

import (
	"fmt"
	"strings"
)

// ConstraintOperator 约束操作符类型
type ConstraintOperator string

const (
	ConstraintEqual                ConstraintOperator = "="
	ConstraintNotEqual             ConstraintOperator = "!="
	ConstraintGreaterThan          ConstraintOperator = ">"
	ConstraintGreaterThanOrEqual   ConstraintOperator = ">="
	ConstraintLessThan             ConstraintOperator = "<"
	ConstraintLessThanOrEqual      ConstraintOperator = "<="
	ConstraintCaret                ConstraintOperator = "^"
	ConstraintTilde                ConstraintOperator = "~"
	ConstraintWildcard             ConstraintOperator = "x"
)

// Constraint 表示一个版本约束条件
//
// 使用示例:
//
//	c, err := versions.ParseConstraint(">=1.0.0")
//	v := versions.NewVersion("1.5.0")
//	if c.Match(v) {
//	    fmt.Println("1.5.0 satisfies >=1.0.0")
//	}
type Constraint struct {
	Operator ConstraintOperator
	Version  *Version
}

// ConstraintSet 表示一组 AND 组合的约束条件
type ConstraintSet struct {
	Constraints []Constraint
}

// ParseConstraint 解析单个版本约束表达式
//
// 支持的操作符: =, !=, >, >=, <, <=, ^, ~
// 支持的通配符: x, X, * (如 1.x, 1.2.*)
func ParseConstraint(expr string) (*Constraint, error) {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return nil, fmt.Errorf("empty constraint expression")
	}

	// 按操作符长度降序匹配，避免 >= 被 > 先匹配
	operators := []struct {
		op   ConstraintOperator
		len  int
	}{
		{ConstraintGreaterThanOrEqual, 2},
		{ConstraintLessThanOrEqual, 2},
		{ConstraintNotEqual, 2},
		{ConstraintCaret, 1},
		{ConstraintTilde, 1},
		{ConstraintGreaterThan, 1},
		{ConstraintLessThan, 1},
		{ConstraintEqual, 1},
	}

	for _, o := range operators {
		if strings.HasPrefix(expr, string(o.op)) {
			versionStr := strings.TrimSpace(expr[o.len:])
			if versionStr == "" {
				return nil, fmt.Errorf("missing version after operator %s", o.op)
			}
			v := NewVersion(versionStr)
			if !v.IsValid() {
				return nil, fmt.Errorf("invalid version in constraint: %s", versionStr)
			}
			return &Constraint{Operator: o.op, Version: v}, nil
		}
	}

	// 检查通配符
	if isWildcardVersion(expr) {
		return &Constraint{Operator: ConstraintWildcard, Version: NewVersion(replaceWildcardWithZero(expr))}, nil
	}

	// 无操作符前缀，视为等于
	v := NewVersion(expr)
	if !v.IsValid() {
		return nil, fmt.Errorf("invalid version in constraint: %s", expr)
	}
	return &Constraint{Operator: ConstraintEqual, Version: v}, nil
}

// ParseConstraintSet 解析逗号分隔的 AND 组合约束
//
// 支持格式: ">=1.0.0,<2.0.0", "^1.2.3", "~1.2"
func ParseConstraintSet(expr string) (*ConstraintSet, error) {
	parts := strings.Split(expr, ",")
	cs := &ConstraintSet{Constraints: make([]Constraint, 0, len(parts))}
	for _, part := range parts {
		c, err := ParseConstraint(part)
		if err != nil {
			return nil, fmt.Errorf("parse constraint %q: %w", part, err)
		}
		cs.Constraints = append(cs.Constraints, *c)
	}
	return cs, nil
}

// Match 判断版本是否满足约束条件
func (c *Constraint) Match(v *Version) bool {
	switch c.Operator {
	case ConstraintEqual:
		return v.CompareTo(c.Version) == 0
	case ConstraintNotEqual:
		return v.CompareTo(c.Version) != 0
	case ConstraintGreaterThan:
		return v.CompareTo(c.Version) > 0
	case ConstraintGreaterThanOrEqual:
		return v.CompareTo(c.Version) >= 0
	case ConstraintLessThan:
		return v.CompareTo(c.Version) < 0
	case ConstraintLessThanOrEqual:
		return v.CompareTo(c.Version) <= 0
	case ConstraintCaret:
		return matchCaret(c.Version, v)
	case ConstraintTilde:
		return matchTilde(c.Version, v)
	case ConstraintWildcard:
		return matchWildcard(c.Version, v)
	default:
		return false
	}
}

// Match 判断版本是否满足所有约束（AND 逻辑）
func (cs *ConstraintSet) Match(v *Version) bool {
	for _, c := range cs.Constraints {
		if !c.Match(v) {
			return false
		}
	}
	return true
}

// matchCaret 实现 ^ 操作符：兼容左起第一个非零版本号
//
// ^1.2.3 := >=1.2.3, <2.0.0
// ^0.2.3 := >=0.2.3, <0.3.0
// ^0.0.3 := >=0.0.3, <0.0.4
func matchCaret(base, v *Version) bool {
	if v.CompareTo(base) < 0 {
		return false
	}
	if len(base.VersionNumbers) == 0 {
		return true
	}
	// 找到第一个非零位
	firstNonZero := -1
	for i, n := range base.VersionNumbers {
		if n != 0 {
			firstNonZero = i
			break
		}
	}
	if firstNonZero == -1 {
		// 全零，如 ^0.0.0，匹配任何版本
		return true
	}
	// 上界：第一个非零位+1，后面全0
	upper := make([]int, len(base.VersionNumbers))
	upper[firstNonZero] = base.VersionNumbers[firstNonZero] + 1
	upperVersion := &Version{
		Raw:            "upper",
		VersionNumbers: upper,
	}
	return v.CompareTo(upperVersion) < 0
}

// matchTilde 实现 ~ 操作符：兼容到次版本号
//
// ~1.2.3 := >=1.2.3, <1.3.0
// ~1.2   := >=1.2.0, <1.3.0
func matchTilde(base, v *Version) bool {
	if v.CompareTo(base) < 0 {
		return false
	}
	if len(base.VersionNumbers) < 2 {
		return true
	}
	upper := make([]int, len(base.VersionNumbers))
	copy(upper, base.VersionNumbers)
	upper[0] = base.VersionNumbers[0]
	upper[1] = base.VersionNumbers[1] + 1
	for i := 2; i < len(upper); i++ {
		upper[i] = 0
	}
	upperVersion := &Version{
		Raw:            "upper",
		VersionNumbers: upper,
	}
	return v.CompareTo(upperVersion) < 0
}

// matchWildcard 实现 x/X/* 通配符
//
// 1.x := >=1.0.0, <2.0.0
// 1.2.x := >=1.2.0, <1.3.0
func matchWildcard(base, v *Version) bool {
	if v.CompareTo(base) < 0 {
		return false
	}
	// 最后一个有效数字位+1
	lastNonZero := -1
	for i, n := range base.VersionNumbers {
		if n != 0 {
			lastNonZero = i
		}
	}
	if lastNonZero == -1 {
		return true
	}
	upper := make([]int, len(base.VersionNumbers))
	copy(upper, base.VersionNumbers)
	upper[lastNonZero] = base.VersionNumbers[lastNonZero] + 1
	for i := lastNonZero + 1; i < len(upper); i++ {
		upper[i] = 0
	}
	upperVersion := &Version{
		Raw:            "upper",
		VersionNumbers: upper,
	}
	return v.CompareTo(upperVersion) < 0
}

// isWildcardVersion 检查版本字符串是否包含通配符
func isWildcardVersion(s string) bool {
	return strings.ContainsAny(s, "xX*")
}

// replaceWildcardWithZero 将通配符替换为0
func replaceWildcardWithZero(s string) string {
	s = strings.ReplaceAll(s, "x", "0")
	s = strings.ReplaceAll(s, "X", "0")
	s = strings.ReplaceAll(s, "*", "0")
	return s
}
```

- [ ] **Step 2: 创建 constraint_test.go**

```go
package versions

import "testing"

func TestParseConstraint(t *testing.T) {
	tests := []struct {
		expr     string
		op       ConstraintOperator
		version  string
		hasError bool
	}{
		{">=1.0.0", ConstraintGreaterThanOrEqual, "1.0.0", false},
		{"<2.0.0", ConstraintLessThan, "2.0.0", false},
		{"^1.2.3", ConstraintCaret, "1.2.3", false},
		{"~1.2", ConstraintTilde, "1.2", false},
		{"1.0.0", ConstraintEqual, "1.0.0", false},
		{"!=1.0.0", ConstraintNotEqual, "1.0.0", false},
		{"", "", "", true},
	}
	for _, tt := range tests {
		c, err := ParseConstraint(tt.expr)
		if tt.hasError {
			if err == nil {
				t.Errorf("ParseConstraint(%q) should return error", tt.expr)
			}
			continue
		}
		if err != nil {
			t.Errorf("ParseConstraint(%q) unexpected error: %v", tt.expr, err)
			continue
		}
		if c.Operator != tt.op {
			t.Errorf("ParseConstraint(%q) operator = %q, want %q", tt.expr, c.Operator, tt.op)
		}
		if c.Version.Raw != tt.version {
			t.Errorf("ParseConstraint(%q) version = %q, want %q", tt.expr, c.Version.Raw, tt.version)
		}
	}
}

func TestConstraint_Match_Comparison(t *testing.T) {
	c, _ := ParseConstraint(">=1.0.0")
	if !c.Match(NewVersion("1.5.0")) {
		t.Error("1.5.0 should match >=1.0.0")
	}
	if c.Match(NewVersion("0.9.0")) {
		t.Error("0.9.0 should not match >=1.0.0")
	}
}

func TestConstraint_Match_Caret(t *testing.T) {
	c, _ := ParseConstraint("^1.2.3")
	if !c.Match(NewVersion("1.9.9")) {
		t.Error("1.9.9 should match ^1.2.3")
	}
	if c.Match(NewVersion("2.0.0")) {
		t.Error("2.0.0 should not match ^1.2.3")
	}
	if !c.Match(NewVersion("1.2.3")) {
		t.Error("1.2.3 should match ^1.2.3")
	}
}

func TestConstraint_Match_Tilde(t *testing.T) {
	c, _ := ParseConstraint("~1.2.3")
	if !c.Match(NewVersion("1.2.9")) {
		t.Error("1.2.9 should match ~1.2.3")
	}
	if c.Match(NewVersion("1.3.0")) {
		t.Error("1.3.0 should not match ~1.2.3")
	}
}

func TestConstraintSet_Match(t *testing.T) {
	cs, _ := ParseConstraintSet(">=1.0.0,<2.0.0")
	if !cs.Match(NewVersion("1.5.0")) {
		t.Error("1.5.0 should match >=1.0.0,<2.0.0")
	}
	if cs.Match(NewVersion("2.0.0")) {
		t.Error("2.0.0 should not match >=1.0.0,<2.0.0")
	}
	if cs.Match(NewVersion("0.9.0")) {
		t.Error("0.9.0 should not match >=1.0.0,<2.0.0")
	}
}
```

- [ ] **Step 3: 验证约束表达式解析器**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test ./...`
Expected:
  - Exit code: 0
  - Output contains: "ok"
  - Output does NOT contain: "FAIL"

- [ ] **Step 4: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add constraint.go constraint_test.go && git commit -m "feat(constraint): add version constraint expression parser with ^, ~, >=, <=, wildcard support"`
