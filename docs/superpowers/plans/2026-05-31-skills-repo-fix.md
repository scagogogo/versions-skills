# Skills 仓库结构修复 Plan

> **For agentic workers:** REQUIRED SUB-SKILL: `superpowers:subagent-driven-development`
> Steps use checkbox (`- [ ]`) syntax.

**Goal:** 修复当前仓库的 Claude Code Skills 插件结构，使其能被其他 AI 通过 Claude Code 插件系统正确安装和使用。核心问题是缺少 `.claude/plugins/versions.js` 入口文件（导致 Skills 无法被发现），以及 Skill 内容未包含最新 API。

**Architecture:** 用户执行 `claude plugin install` → Claude Code 读取 `.claude-plugin/plugin.json` 获取插件元数据 → 加载 `.claude/plugins/versions.js` → JS 入口通过 `config` hook 注册 `skills/` 目录路径 → Claude Code 扫描 `skills/*/SKILL.md` → 用户输入 `/version-parsing` 等命令即可调用。参照 superpowers-auto 插件的实际工作结构。

**Tech Stack:** Claude Code Plugin SDK (JS entry + SKILL.md + plugin.json), Go 1.18

**Risks:**
- `.claude/plugins/versions.js` 的 hook API 必须与 Claude Code 插件运行时匹配 → 缓解：直接参照已安装可用的 superpowers-auto.js 格式
- Skill 内容必须与实际 Go API 一致 → 缓解：从源码提取函数签名

---

### Task 1: 创建插件入口文件

**Depends on:** None
**Files:**
- Create: `.claude/plugins/versions.js`

- [ ] **Step 1: 创建 .claude/plugins/versions.js — 注册 skills 目录路径到 Claude Code 插件系统**

```javascript
/**
 * Versions SDK Plugin for Claude Code
 * Registers skills directory path via config hook so Claude Code discovers SKILL.md files.
 */

import path from 'path';
import { fileURLToPath } from 'url';

const __dirname = path.dirname(fileURLToPath(import.meta.url));

export const VersionsPlugin = async () => {
  const skillsDir = path.resolve(__dirname, '../../skills');

  return {
    config: async (config) => {
      config.skills = config.skills || {};
      config.skills.paths = config.skills.paths || [];
      if (!config.skills.paths.includes(skillsDir)) {
        config.skills.paths.push(skillsDir);
      }
    }
  };
};

export default VersionsPlugin;
```

- [ ] **Step 2: 验证入口文件存在且语法正确**
Run: `cd /home/cc11001100/github/scagogogo/versions && node -e "import('.claude/plugins/versions.js').then(m => console.log('OK:', Object.keys(m))).catch(e => console.error('FAIL:', e.message))" 2>&1 || echo "ESM check failed, checking syntax only" && node --check --input-type=module < .claude/plugins/versions.js 2>&1 && echo "Syntax OK"`
Expected:
  - Exit code: 0
  - Output does NOT contain: "SyntaxError"

- [ ] **Step 3: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add .claude/ && git commit -m "feat(plugin): add .claude/plugins/versions.js entry file for skills directory registration"`

---

### Task 2: 更新 SKILL.md 内容 — 添加最新 API 和修正元数据

**Depends on:** Task 1
**Files:**
- Modify: `skills/version-parsing/SKILL.md`（添加 NewVersionBuilder、Major/Minor/Patch、IsPrerelease 等）
- Modify: `skills/version-comparison/SKILL.md`（添加 IsNewerThan、IsOlderThan、Equals、IsBetween）
- Modify: `skills/version-sorting/SKILL.md`（添加 Min、Max、LatestStable、Filter、FilterByConstraint）

- [ ] **Step 1: 修改 version-parsing SKILL.md — 添加新版便捷 API 和构建器**

文件: `skills/version-parsing/SKILL.md`（全文替换）

```markdown
---
name: version-parsing
description: Use when parsing, validating, building, or extracting components from version strings. Provides expert guidance on using the Go versions SDK for version parsing, construction, and property access.
argument-hint: <version-string-or-task>
---

# Version Parsing & Construction Skill

## When to Use

- User needs to parse a version string into structured components
- User needs to validate if a string is a valid version number
- User needs to extract prefix, numbers, or suffix from a version string
- User needs to build a version programmatically (fluent API)
- User needs to bump version numbers (major/minor/patch)
- User needs to access major, minor, patch numbers individually
- User is working with semver, pre-release, or custom version formats in Go

## API Reference

### Core Parsing Functions

**NewVersion(versionStr string) *Version**
Creates a Version object from a string. Never panics — returns an object even for invalid strings (check with IsValid()).

**NewVersionE(versionStr string) (*Version, error)**
Creates a Version with error return. Returns ErrVersionInvalid for invalid strings.

**NewVersions(versionStringSlice ...string) []*Version**
Batch-create multiple Version objects.

**NewVersionWithOption(versionStr string, option ParserOption) *Version**
Parse with custom delimiters. E.g., `ParserOption{Delimiters: ".-_"}` allows underscores as separators.

### Version Struct

```go
type Version struct {
    Raw            string         // Original string, e.g. "v1.2.3-beta1"
    PublicTime     time.Time      // Release time
    VersionNumbers VersionNumbers // Number part, e.g. [1,2,3]
    Prefix         VersionPrefix  // Prefix, e.g. "v"
    Suffix         VersionSuffix  // Suffix, e.g. "-beta1"
}
```

### Property Accessors

**func (x *Version) Major() int** — Returns the major version number (0 if missing)
**func (x *Version) Minor() int** — Returns the minor version number (0 if missing)
**func (x *Version) Patch() int** — Returns the patch version number (0 if missing)

### Version Type Checks

**func (x *Version) IsValid() bool** — True if VersionNumbers is non-empty
**func (x *Version) IsPrerelease() bool** — True if version has a suffix (alpha/beta/rc/etc)
**func (x *Version) IsStable() bool** — True if version has no suffix (release version)
**func (x *Version) IsDev() bool** — True if suffix is a dev variant
**func (x *Version) IsAlpha() bool** — True if suffix is an alpha variant
**func (x *Version) IsBeta() bool** — True if suffix is a beta variant
**func (x *Version) IsRC() bool** — True if suffix is rc or CR variant
**func (x *Version) IsSnapshot() bool** — True if suffix is a snapshot variant

### Version Builder (Fluent API)

**NewVersionBuilder() *VersionBuilder** — Create a new builder

```go
v := versions.NewVersionBuilder().
    Prefix("v").
    Major(1).
    Minor(2).
    Patch(3).
    Suffix("-beta1").
    Build()
```

### Version Bumping

**func (x *Version) BumpMajor() *Version** — e.g., 1.2.3 → 2.0.0 (clears suffix)
**func (x *Version) BumpMinor() *Version** — e.g., 1.2.3 → 1.3.0 (clears suffix)
**func (x *Version) BumpPatch() *Version** — e.g., 1.2.3 → 1.2.4 (clears suffix)

## Code Examples

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
)

func main() {
    // Parse a version string
    v := versions.NewVersion("v1.2.3-rc1")
    fmt.Printf("Prefix: %s\n", v.Prefix)           // "v"
    fmt.Printf("Numbers: %v\n", v.VersionNumbers)   // [1 2 3]
    fmt.Printf("Suffix: %s\n", v.Suffix)            // "-rc1"

    // Property accessors
    fmt.Printf("Major: %d, Minor: %d, Patch: %d\n",
        v.Major(), v.Minor(), v.Patch())  // 1, 2, 3

    // Type checks
    fmt.Println(v.IsPrerelease())  // true
    fmt.Println(v.IsRC())          // true

    // Build a version programmatically
    v2 := versions.NewVersionBuilder().
        Major(2).
        Minor(0).
        Patch(0).
        Build()

    // Bump versions
    bumped := versions.NewVersion("1.2.3").BumpPatch()  // 1.2.4
    fmt.Println(bumped.Patch())  // 4
}
```

## Important Notes

- NewVersion never returns nil — always check IsValid() for invalid inputs
- Supports arbitrary number of segments: "1.2.3.4.5" works
- Leading zeros are stripped: "1.02.003" becomes [1,2,3]
- Pure alphabetic strings like "abc" return invalid Version (empty VersionNumbers)
- VersionPrefix is a string type — call IsEmpty() to check for no prefix
- VersionSuffix is a string type — call IsEmpty() to check for no suffix
- BumpMajor/Minor/Patch always clear the suffix — the result is always a stable version
- VersionBuilder.Build() creates a new Version by constructing and parsing a string
```

- [ ] **Step 2: 修改 version-comparison SKILL.md — 添加便捷比较方法**

文件: `skills/version-comparison/SKILL.md`（全文替换）

```markdown
---
name: version-comparison
description: Use when comparing version numbers, checking which version is newer/older, or determining version ordering. Provides expert guidance on using the Go versions SDK for version comparison.
argument-hint: <version-comparison-task>
---

# Version Comparison Skill

## When to Use

- User needs to compare two version numbers (which is newer/older)
- User needs to check if a version falls within a range
- User needs to sort versions by their numeric value
- User needs to implement update checking or dependency version logic
- User needs constraint matching (semver ranges like ^1.2.3, ~1.2.3, >=1.0.0)

## API Reference

### Core Comparison

**func (x *Version) CompareTo(target *Version) int**

Comparison order:
1. VersionNumbers (digit-by-digit, left to right; missing positions treated as 0)
2. PublicTime (earlier = smaller)
3. Suffix (semantic weight: dev < alpha < beta < rc < release)
4. Raw string (alphabetic fallback)

Returns: -1 (x < target), 0 (equal), 1 (x > target)

### Convenience Comparison Methods

**func (x *Version) IsNewerThan(target *Version) bool** — Equivalent to CompareTo(target) > 0
**func (x *Version) IsOlderThan(target *Version) bool** — Equivalent to CompareTo(target) < 0
**func (x *Version) Equals(target *Version) bool** — Equivalent to CompareTo(target) == 0
**func (x *Version) IsBetween(low, high *Version) bool** — True if low <= x <= high (nil boundary = unbounded)

### VersionNumbers.CompareTo

**func (x VersionNumbers) CompareTo(target []int) int**
Compares number arrays element-by-element. Shorter array is treated as padded with zeros.

### VersionSuffix.CompareTo

**func (x VersionSuffix) CompareTo(target VersionSuffix) int**
Compares suffix strings using semantic weight. Empty suffix > any non-empty suffix (release > pre-release).

### Constraint Matching

**func ParseConstraint(expr string) (*Constraint, error)** — Parse a single constraint expression
**func ParseConstraintSet(expr string) (*ConstraintSet, error)** — Parse comma-separated AND constraints
**func (c *Constraint) Match(v *Version) bool** — Check if version satisfies constraint
**func (cs *ConstraintSet) Match(v *Version) bool** — Check if version satisfies all constraints

Supported operators: `=`, `!=`, `>`, `>=`, `<`, `<=`, `^` (caret), `~` (tilde), `x/X/*` (wildcard)

## Code Examples

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
)

func main() {
    // Convenience comparison methods
    v1 := versions.NewVersion("1.2.3")
    v2 := versions.NewVersion("1.3.0")
    fmt.Println(v2.IsNewerThan(v1))  // true
    fmt.Println(v1.IsOlderThan(v2))  // true
    fmt.Println(v1.Equals(v1))       // true

    // Range check
    v := versions.NewVersion("1.5.0")
    low := versions.NewVersion("1.0.0")
    high := versions.NewVersion("2.0.0")
    fmt.Println(v.IsBetween(low, high))  // true

    // Pre-release vs release
    beta := versions.NewVersion("1.0.0-beta")
    release := versions.NewVersion("1.0.0")
    fmt.Println(beta.CompareTo(release) < 0)  // true

    // Constraint matching
    c, _ := versions.ParseConstraint(">=1.0.0")
    fmt.Println(c.Match(versions.NewVersion("1.5.0")))  // true
    fmt.Println(c.Match(versions.NewVersion("0.9.0")))  // false

    // Caret constraint (^1.2.3 = >=1.2.3, <2.0.0)
    cc, _ := versions.ParseConstraint("^1.2.3")
    fmt.Println(cc.Match(versions.NewVersion("1.9.0")))  // true
    fmt.Println(cc.Match(versions.NewVersion("2.0.0")))  // false

    // Constraint set (AND logic)
    cs, _ := versions.ParseConstraintSet(">=1.0.0,<2.0.0")
    fmt.Println(cs.Match(versions.NewVersion("1.5.0")))  // true
}
```

## Important Notes

- CompareTo treats "1.0" and "1.0.0" as equal (missing digits = 0)
- Prefix comparison is case-sensitive: "v1.0" != "V1.0"
- Suffix comparison uses semantic weight: -alpha < -beta < -rc < (release)
- IsBetween is inclusive on both boundaries; pass nil to ignore a boundary
- Constraint caret: ^0.2.3 means >=0.2.3, <0.3.0 (zero-major special case)
```

- [ ] **Step 3: 修改 version-sorting SKILL.md — 添加批量操作 API**

文件: `skills/version-sorting/SKILL.md`（全文替换）

```markdown
---
name: version-sorting
description: Use when sorting, filtering, or querying a list of version numbers. Provides expert guidance on using the Go versions SDK for version sorting and batch operations.
argument-hint: <version-sorting-task>
---

# Version Sorting & Batch Operations Skill

## When to Use

- User needs to sort a list of version strings or Version objects
- User needs to find the latest or oldest version in a collection
- User needs to filter versions by type (stable, prerelease) or constraint
- User needs to deduplicate a version list
- User needs to count versions matching a condition

## API Reference

### Sorting Functions

**func SortVersionStringSlice(versionStringSlice []string) []string**
Sorts a string slice of version numbers. Parses each string, sorts by Version.CompareTo rules, returns sorted strings.

**func SortVersionSlice(versions []*Version) []*Version**
Sorts Version object slice. Uses group-based algorithm for stable, semantically correct ordering.

### Batch Query Functions

**func Min(versions []*Version) *Version** — Find smallest version (nil if empty)
**func Max(versions []*Version) *Version** — Find largest version (nil if empty)
**func LatestStable(versions []*Version) *Version** — Find latest stable (no suffix) version
**func LatestPrerelease(versions []*Version) *Version** — Find latest prerelease version

### Filter Functions

**func Filter(versions []*Version, predicate func(*Version) bool) []*Version** — Generic filter
**func FilterByConstraint(versions []*Version, constraint *Constraint) []*Version** — Filter by constraint
**func FilterByConstraintSet(versions []*Version, cs *ConstraintSet) []*Version** — Filter by constraint set
**func FilterByMajor(versions []*Version, major int) []*Version** — Filter by major version number

### Utility Functions

**func Unique(versions []*Version) []*Version** — Remove duplicates (by Raw string)
**func Count(versions []*Version, predicate func(*Version) bool) int** — Count matching versions

## Code Examples

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
)

func main() {
    // Sort version strings
    unsorted := []string{"2.0.0", "1.0.0", "1.10.0", "1.2.0"}
    sorted := versions.SortVersionStringSlice(unsorted)
    // Result: ["1.0.0", "1.2.0", "1.10.0", "2.0.0"]

    // Find min and max
    versionList := versions.NewVersions("2.0.0", "1.0.0", "1.10.0", "1.0.0-beta")
    fmt.Println(versions.Min(versionList).Raw)  // "1.0.0-beta"
    fmt.Println(versions.Max(versionList).Raw)  // "2.0.0"

    // Find latest stable and prerelease
    fmt.Println(versions.LatestStable(versionList).Raw)      // "2.0.0"
    fmt.Println(versions.LatestPrerelease(versionList).Raw)  // "1.0.0-beta"

    // Filter by constraint
    c, _ := versions.ParseConstraint(">=1.0.0,<2.0.0")
    result := versions.FilterByConstraint(versionList, c)
    // Result: [1.0.0, 1.10.0]

    // Filter by major version
    v2 := versions.FilterByMajor(versionList, 2)
    fmt.Println(len(v2))  // 1

    // Unique versions
    dupes := versions.NewVersions("1.0.0", "1.0.0", "2.0.0")
    fmt.Println(len(versions.Unique(dupes)))  // 2

    // Count
    n := versions.Count(versionList, func(v *versions.Version) bool {
        return v.IsPrerelease()
    })
    fmt.Println(n)  // 1
}
```

## Important Notes

- "1.10.0" correctly sorts after "1.2.0" (numeric, not alphabetic)
- Pre-release versions sort before their release counterparts
- Min/Max consider ALL versions including pre-release
- LatestStable ignores pre-release versions entirely
- Filter functions return new slices, never modify the input
- Unique deduplicates by Raw string, preserving first occurrence
```

- [ ] **Step 4: 验证 SKILL.md 文件存在且格式正确**
Run: `cd /home/cc11001100/github/scagogogo/versions && for f in skills/version-parsing/SKILL.md skills/version-comparison/SKILL.md skills/version-sorting/SKILL.md; do [ -f "$f" ] && head -3 "$f" | grep -q "^---$" && echo "OK: $f" || echo "FAIL: $f"; done`
Expected:
  - Exit code: 0
  - Output contains: "OK: skills/version-parsing/SKILL.md"
  - Output does NOT contain: "FAIL"

- [ ] **Step 5: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add skills/version-parsing/SKILL.md skills/version-comparison/SKILL.md skills/version-sorting/SKILL.md && git commit -m "feat(skills): update parsing, comparison, sorting skills with new API — builder, bump, convenience methods, batch ops"`

---

### Task 3: 更新其余 SKILL.md — 添加新增 API 引用

**Depends on:** Task 2
**Files:**
- Modify: `skills/version-grouping/SKILL.md`（添加 GetLatest/GetOldest/Count/StableVersions）
- Modify: `skills/version-range-query/SKILL.md`（添加 FilterByConstraint 相关）
- Modify: `skills/version-visualization/SKILL.md`（微调说明）
- Modify: `skills/version-file-operations/SKILL.md`（微调说明）

- [ ] **Step 1: 修改 version-grouping SKILL.md — 添加 VersionGroup 便捷方法**

文件: `skills/version-grouping/SKILL.md`（全文替换）

```markdown
---
name: version-grouping
description: Use when grouping versions by major/minor version number, finding latest/oldest in a group, or managing version collections by their version series.
argument-hint: <version-grouping-task>
---

# Version Grouping Skill

## When to Use

- User needs to group versions by their major version number
- User needs to find the latest or oldest version in a group
- User needs to organize a large collection of versions into version series
- User needs to find all stable or prerelease versions within a group

## API Reference

### Group Function

**func Group(versions []*Version) map[string]*VersionGroup**
Groups versions by their VersionNumbers' BuildGroupID.

### VersionGroup Type

```go
type VersionGroup struct {
    GroupVersionNumbers VersionNumbers
    VersionMap          map[string]*Version
}
```

**Key Methods:**

- **NewVersionGroupFromVersions(versions []*Version) *VersionGroup** — create from version array
- **func (x *VersionGroup) Add(v *Version) bool** — add version to group
- **func (x *VersionGroup) Contains(v *Version) bool** — check if version exists
- **func (x *VersionGroup) ID() string** — get group ID (e.g., "1.2")
- **func (x *VersionGroup) Count() int** — number of versions in group
- **func (x *VersionGroup) Versions() []*Version** — get all versions (unordered)
- **func (x *VersionGroup) SortVersions() []*Version** — get sorted versions

### VersionGroup Convenience Methods

- **func (x *VersionGroup) GetLatest() *Version** — newest version in group (nil if empty)
- **func (x *VersionGroup) GetOldest() *Version** — oldest version in group (nil if empty)
- **func (x *VersionGroup) StableVersions() []*Version** — all stable (no suffix) versions
- **func (x *VersionGroup) PrereleaseVersions() []*Version** — all prerelease versions
- **func (x *VersionGroup) LatestStable() *Version** — latest stable version in group

### SortedVersionGroups Type

- **func NewSortedVersionGroups(versions []*Version) *SortedVersionGroups** — create sorted groups
- **func (x *SortedVersionGroups) GroupIDs() []string** — get sorted group ID list

## Code Examples

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
)

func main() {
    versionList := versions.NewVersions("1.0.0", "1.0.1", "1.1.0-beta", "1.1.0", "2.0.0")
    groupMap := versions.Group(versionList)

    for _, group := range groupMap {
        fmt.Printf("Group %s: %d versions\n", group.ID(), group.Count())
        fmt.Printf("  Latest: %s\n", group.GetLatest().Raw)
        fmt.Printf("  Oldest: %s\n", group.GetOldest().Raw)
        fmt.Printf("  Stable: %d, Prerelease: %d\n",
            len(group.StableVersions()),
            len(group.PrereleaseVersions()))
        if latest := group.LatestStable(); latest != nil {
            fmt.Printf("  Latest stable: %s\n", latest.Raw)
        }
    }
}
```

## Important Notes

- Group() uses BuildGroupID() as the key — this is the full version number string
- GetLatest/GetOldest call SortVersions() internally — O(n log n)
- LatestStable returns nil if no stable versions exist in the group
- StableVersions/PrereleaseVersions use the Filter utility function
```

- [ ] **Step 2: 修改 version-range-query SKILL.md — 添加约束过滤 API**

文件: `skills/version-range-query/SKILL.md`（全文替换）

```markdown
---
name: version-range-query
description: Use when querying versions within a specific range, implementing version constraints, or checking version compatibility.
argument-hint: <version-range-query-task>
---

# Version Range Query Skill

## When to Use

- User needs to find all versions between a start and end version
- User needs to implement version constraint logic (e.g., ">=1.0.0 and <2.0.0")
- User needs to check if a version falls within a specific range
- User needs to filter a version list by semver constraints

## API Reference

### Constraint Parsing

**func ParseConstraint(expr string) (*Constraint, error)** — Parse single constraint
**func ParseConstraintSet(expr string) (*ConstraintSet, error)** — Parse comma-separated AND constraints

Supported operators:
- `=1.0.0` — exact match
- `!=1.0.0` — not equal
- `>1.0.0`, `>=1.0.0`, `<1.0.0`, `<=1.0.0` — comparisons
- `^1.2.3` — caret (compatible with leftmost non-zero: >=1.2.3, <2.0.0)
- `~1.2.3` — tilde (compatible with minor: >=1.2.3, <1.3.0)
- `1.x`, `1.2.*` — wildcard (any sub-version)

### Constraint Matching

**func (c *Constraint) Match(v *Version) bool** — Check single constraint
**func (cs *ConstraintSet) Match(v *Version) bool** — Check all constraints (AND)

### Batch Filter by Constraint

**func FilterByConstraint(versions []*Version, constraint *Constraint) []*Version**
**func FilterByConstraintSet(versions []*Version, cs *ConstraintSet) []*Version**

### Range Query (Group-based)

**func (x *VersionGroup) QueryRangeVersions(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version**
**func (x *SortedVersionGroups) QueryRange(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version**

### ContainsPolicy

```go
const (
    ContainsPolicyNone ContainsPolicy = iota  // Unspecified
    ContainsPolicyYes                          // Include boundary
    ContainsPolicyNo                           // Exclude boundary
)
```

### Point-in-Range Check

**func (x *Version) IsBetween(low, high *Version) bool** — True if low <= x <= high

## Code Examples

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
    "github.com/golang-infrastructure/go-tuple"
)

func main() {
    versionList := versions.NewVersions("0.9.0", "1.0.0", "1.5.0", "2.0.0", "2.1.0")

    // Simple constraint matching
    c, _ := versions.ParseConstraint(">=1.0.0")
    fmt.Println(c.Match(versions.NewVersion("1.5.0")))  // true

    // Caret constraint
    caret, _ := versions.ParseConstraint("^1.2.3")
    fmt.Println(caret.Match(versions.NewVersion("1.9.9")))  // true
    fmt.Println(caret.Match(versions.NewVersion("2.0.0")))  // false

    // Batch filter by constraint
    cs, _ := versions.ParseConstraintSet(">=1.0.0,<2.0.0")
    result := versions.FilterByConstraintSet(versionList, cs)
    fmt.Printf("Versions in [1.0.0, 2.0.0): %d\n", len(result))  // 2

    // Point-in-range check
    v := versions.NewVersion("1.5.0")
    fmt.Println(v.IsBetween(versions.NewVersion("1.0.0"), versions.NewVersion("2.0.0")))  // true

    // Group-based range query
    sortedGroups := versions.NewSortedVersionGroups(versionList)
    start := tuple.New2(versions.NewVersion("1.0.0"), versions.ContainsPolicyYes)
    end := tuple.New2(versions.NewVersion("2.0.0"), versions.ContainsPolicyNo)
    rangeResult := sortedGroups.QueryRange(start, end)
    fmt.Printf("Range query result: %d versions\n", len(rangeResult))
}
```

## Important Notes

- ParseConstraintSet uses comma as AND separator
- Caret with zero major: ^0.2.3 means >=0.2.3, <0.3.0
- IsBetween is inclusive on both boundaries; pass nil for unbounded
- FilterByConstraintSet is more efficient than nested FilterByConstraint calls
```

- [ ] **Step 3: 修改 version-visualization 和 version-file-operations SKILL.md**

文件: `skills/version-visualization/SKILL.md`（仅修改 frontmatter description 行）

将第 2 行替换为：
```yaml
description: Use when visualizing version hierarchies, displaying version trees, or showing version group structures in text format.
```

文件: `skills/version-file-operations/SKILL.md`（仅修改 frontmatter description 行）

将第 2 行替换为：
```yaml
description: Use when reading version numbers from files, processing version lists stored in text files.
```

- [ ] **Step 4: 验证所有 SKILL.md 文件格式正确**
Run: `cd /home/cc11001100/github/scagogogo/versions && find skills/ -name "SKILL.md" | sort | while read f; do head -3 "$f" | grep -q "^---$" && echo "OK: $f" || echo "FAIL: $f"; done`
Expected:
  - Exit code: 0
  - Output does NOT contain: "FAIL"
  - Total count: 7 OK lines

- [ ] **Step 5: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add skills/ && git commit -m "feat(skills): update grouping, range-query, visualization, file-operations skills with new API"`

---

### Task 4: 修正 README 安装命令和最终验证

**Depends on:** Task 2, Task 3
**Files:**
- Modify: `README.md:1596-1601`（修正安装命令）

- [ ] **Step 1: 修改 README 安装命令 — 添加 .git 后缀和新增 Skills 说明**

文件: `README.md:1596-1601`（替换安装命令区块和可用 Skills 表格）

```markdown
### 安装方式

```bash
# 添加 marketplace
claude marketplace add versions https://github.com/scagogogo/versions-skills.git

# 安装插件
claude plugin install versions@versions
```

### 可用 Skills

| Skill | 命令 | 用途 |
|:------|:-----|:-----|
| 版本号解析与构建 | `/version-parsing` | 解析、验证、构建版本号，属性访问 |
| 版本号比较 | `/version-comparison` | 比较版本号大小，约束匹配 |
| 版本号排序与批量操作 | `/version-sorting` | 排序、过滤、去重、批量查询 |
| 版本号分组 | `/version-grouping` | 按主版本号分组，查找最新/最旧 |
| 版本范围查询 | `/version-range-query` | 约束解析，范围过滤 |
| 版本号可视化 | `/version-visualization` | 以树形结构展示版本层次 |
| 版本号文件操作 | `/version-file-operations` | 从文件读取版本号列表 |
```

- [ ] **Step 2: 最终验证 — Go 测试通过 + 插件结构完整**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test ./... && [ -f .claude/plugins/versions.js ] && [ -f .claude-plugin/plugin.json ] && find skills/ -name "SKILL.md" | wc -l | grep -q "7" && echo "All checks passed"`
Expected:
  - Exit code: 0
  - Output contains: "ok"
  - Output contains: "All checks passed"

- [ ] **Step 3: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add README.md && git commit -m "docs: fix Skills install command (.git suffix) and update skill descriptions in README"`
