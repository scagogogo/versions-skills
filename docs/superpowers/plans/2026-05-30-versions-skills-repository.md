# Versions Skills 仓库转换 Plan

> **For agentic workers:** REQUIRED SUB-SKILL: `superpowers:subagent-driven-development`
> Steps use checkbox (`- [ ]`) syntax.

**Goal:** 将 versions Go SDK 项目转换为 Claude Code Skills 仓库，使其既保留原有 Go SDK 功能，又能通过 Claude Code 的插件系统安装和使用版本号处理的专业 Skills。

**Architecture:** 项目保持原有 Go SDK 代码不变，新增 `.claude-plugin/` 目录（插件元数据）和 `skills/` 目录（7 个 Skill 定义）。每个 Skill 对应项目的一个功能模块，SKILL.md 包含 YAML frontmatter（name、description、argument-hint）和 Markdown 正文（使用场景、API 签名、代码示例、注意事项）。用户通过 `claude plugin install` 安装后，在 Claude Code 中输入 `/version-parsing` 等命令即可获得专业指导。

**Tech Stack:** Go 1.18, Claude Code Plugin SDK (markdown + yaml frontmatter), Git

**Risks:**
- Skill 内容必须与 Go 代码 API 完全一致，否则 Claude 会生成错误代码 → 缓解：每个 Skill 中的函数签名和示例代码直接从源码提取
- marketplace.json 格式必须符合 Claude Code 插件规范 → 缓解：严格参照 superpowers-auto 的 marketplace.json 格式编写
- 项目同时作为 Go SDK 和 Skills 仓库，需要确保两种用途不冲突 → 缓解：Skills 目录独立于 Go 代码，不修改任何 .go 文件

---

### Task 1: 插件基础设施

**Depends on:** None
**Files:**
- Create: `.claude-plugin/plugin.json`
- Create: `.claude-plugin/marketplace.json`

- [ ] **Step 1: 创建 plugin.json — 定义插件元数据**

```json
{
  "name": "versions",
  "description": "Go语言版本号解析计算SDK的Claude Code Skills — 提供版本号解析、比较、排序、分组、范围查询、可视化和文件操作的专业技能",
  "version": "0.1.0",
  "author": {
    "name": "scagogogo",
    "url": "https://github.com/scagogogo"
  },
  "homepage": "https://github.com/scagogogo/versions",
  "repository": "https://github.com/scagogogo/versions",
  "license": "MIT",
  "keywords": ["versions", "semver", "version-parsing", "version-sorting", "version-grouping", "go"]
}
```

- [ ] **Step 2: 创建 marketplace.json — 定义市场清单，引用自身作为插件来源**

```json
{
  "name": "versions",
  "description": "Go语言版本号解析计算SDK的Claude Code Skills",
  "owner": {
    "name": "scagogogo",
    "url": "https://github.com/scagogogo"
  },
  "plugins": [
    {
      "name": "versions",
      "description": "版本号解析、比较、排序、分组、范围查询、可视化和文件操作的专业Skills",
      "version": "0.1.0",
      "source": "./",
      "author": {
        "name": "scagogogo",
        "url": "https://github.com/scagogogo"
      }
    }
  ]
}
```

- [ ] **Step 3: 验证插件元数据格式**
Run: `cd /home/cc11001100/github/scagogogo/versions && cat .claude-plugin/plugin.json | python3 -m json.tool > /dev/null && cat .claude-plugin/marketplace.json | python3 -m json.tool > /dev/null && echo "JSON valid"`
Expected:
  - Exit code: 0
  - Output contains: "JSON valid"

- [ ] **Step 4: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add .claude-plugin/ && git commit -m "feat(plugin): add Claude Code plugin metadata"`

---

### Task 2: 核心技能定义（解析、比较、排序、文件操作）

**Depends on:** Task 1
**Files:**
- Create: `skills/version-parsing/SKILL.md`
- Create: `skills/version-comparison/SKILL.md`
- Create: `skills/version-sorting/SKILL.md`
- Create: `skills/version-file-operations/SKILL.md`

- [ ] **Step 1: 创建 version-parsing Skill — 版本号解析技能**

```markdown
---
name: version-parsing
description: Use when parsing, validating, or extracting components from version strings like "1.2.3", "v1.2.3-beta1". Provides expert guidance on using the Go versions SDK for version parsing.
argument-hint: <version-string-or-task>
---

# Version Parsing Skill

## When to Use

- User needs to parse a version string into structured components
- User needs to validate if a string is a valid version number
- User needs to extract prefix, numbers, or suffix from a version string
- User is working with semver, pre-release, or custom version formats in Go

## API Reference

### Core Functions

**NewVersion(versionStr string) *Version**
Creates a Version object from a string. Never panics — returns an object even for invalid strings (check with IsValid()).

**NewVersionE(versionStr string) (*Version, error)**
Creates a Version with error return. Returns ErrVersionInvalid for invalid strings.

**NewVersions(versionStringSlice ...string) []*Version**
Batch-create multiple Version objects.

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

### Validation

**func (v *Version) IsValid() bool**
Returns true if the version contains at least one number in VersionNumbers.

## Code Examples

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions"
)

func main() {
    // Parse a version string
    v := versions.NewVersion("v1.2.3-rc1")
    fmt.Printf("Prefix: %s\n", v.Prefix)           // "v"
    fmt.Printf("Numbers: %v\n", v.VersionNumbers)   // [1 2 3]
    fmt.Printf("Suffix: %s\n", v.Suffix)            // "-rc1"

    // Validate a version
    valid := versions.NewVersion("1.2.3")
    invalid := versions.NewVersion("not-a-version")
    fmt.Println(valid.IsValid())    // true
    fmt.Println(invalid.IsValid())  // false

    // Error-checking parse
    v, err := versions.NewVersionE("")
    if err != nil {
        fmt.Println("Error:", err)  // "version invalid"
    }
}
```

## Important Notes

- NewVersion never returns nil — always check IsValid() for invalid inputs
- Supports arbitrary number of segments: "1.2.3.4.5" works
- Leading zeros are stripped: "1.02.003" becomes [1,2,3]
- Pure alphabetic strings like "abc" return invalid Version (empty VersionNumbers)
- VersionPrefix is a string type — call IsEmpty() to check for no prefix
- VersionSuffix is a string type — call IsEmpty() to check for no suffix
```

- [ ] **Step 2: 创建 version-comparison Skill — 版本号比较技能**

```markdown
---
name: version-comparison
description: Use when comparing version numbers, checking which version is newer/older, or determining version ordering. Provides expert guidance on using the Go versions SDK for version comparison.
argument-hint: <version-comparison-task>
---

# Version Comparison Skill

## When to Use

- User needs to compare two version numbers (which is newer/older)
- User needs to sort versions by their numeric value
- User needs to implement update checking or dependency version logic
- User is working with version comparison in Go

## API Reference

### Version.CompareTo

**func (x *Version) CompareTo(target *Version) int**

Comparison order:
1. VersionNumbers (digit-by-digit, left to right; missing positions treated as 0)
2. PublicTime (earlier = smaller)
3. Suffix (alphabetic comparison)
4. Raw string (alphabetic fallback)

Returns: -1 (x < target), 0 (equal), 1 (x > target)

### VersionNumbers.CompareTo

**func (x VersionNumbers) CompareTo(target []int) int**

Compares number arrays element-by-element. Shorter array is treated as padded with zeros.

Returns: negative (x < target), 0 (equal), positive (x > target)

### VersionSuffix.CompareTo

**func (x VersionSuffix) CompareTo(target VersionSuffix) int**

Compares suffix strings alphabetically. Empty suffix > any non-empty suffix (release > pre-release).

### VersionPrefix

VersionPrefix is a string type with an IsEmpty() method. Prefix comparison is alphabetical and case-sensitive.

## Code Examples

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions"
)

func main() {
    // Basic comparison
    v1 := versions.NewVersion("1.2.3")
    v2 := versions.NewVersion("1.3.0")
    if v1.CompareTo(v2) < 0 {
        fmt.Println("1.2.3 is older than 1.3.0")
    }

    // Pre-release vs release
    beta := versions.NewVersion("1.0.0-beta")
    release := versions.NewVersion("1.0.0")
    if beta.CompareTo(release) < 0 {
        fmt.Println("Pre-release is older than release")
    }

    // Different length versions
    short := versions.NewVersion("1.0")
    full := versions.NewVersion("1.0.0")
    fmt.Println(short.CompareTo(full))  // 0 (treated as equal)

    // Number parts comparison
    nums1 := versions.NewVersionNumbers([]int{1, 2, 0})
    nums2 := versions.NewVersionNumbers([]int{1, 2, 3})
    if nums1.CompareTo(nums2) < 0 {
        fmt.Println("1.2.0 < 1.2.3")
    }
}
```

## Important Notes

- CompareTo treats "1.0" and "1.0.0" as equal (missing digits = 0)
- Prefix comparison is case-sensitive: "v1.0" != "V1.0"
- Empty suffix (release version) is greater than any non-empty suffix (pre-release)
- PublicTime is only compared when both versions have it set (non-zero)
- The comparison implements the compare_anything.Comparable interface
```

- [ ] **Step 3: 创建 version-sorting Skill — 版本号排序技能**

```markdown
---
name: version-sorting
description: Use when sorting a list of version numbers in natural order. Provides expert guidance on using the Go versions SDK for version sorting.
argument-hint: <version-sorting-task>
---

# Version Sorting Skill

## When to Use

- User needs to sort a list of version strings or Version objects
- User needs to find the latest or oldest version in a collection
- User needs to display versions in ascending/descending order
- User is implementing version selection UI or dependency resolution

## API Reference

### SortVersionStringSlice

**func SortVersionStringSlice(versionStringSlice []string) []string**

Sorts a string slice of version numbers. Parses each string, sorts by Version.CompareTo rules, returns sorted strings.

### SortVersionSlice

**func SortVersionSlice(versions []*Version) []*Version**

Sorts Version object slice. Uses group-based algorithm: groups by major version, sorts groups, sorts within each group, merges results.

### SortVersionGroupMap / SortVersionGroupSlice

**func SortVersionGroupMap(versionGroupMap map[string]*VersionGroup) []*VersionGroup**
**func SortVersionGroupSlice(groupSlice []*VersionGroup)**

Utility functions for sorting version group collections.

## Code Examples

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions"
)

func main() {
    // Sort version strings
    unsorted := []string{"2.0.0", "1.0.0", "1.10.0", "1.2.0", "v1.5.0"}
    sorted := versions.SortVersionStringSlice(unsorted)
    for _, v := range sorted {
        fmt.Println(v)
    }
    // Output: 1.0.0, 1.2.0, v1.5.0, 1.10.0, 2.0.0

    // Sort Version objects
    versionList := versions.NewVersions("2.0.0", "1.0.0", "1.10.0")
    sortedVersions := versions.SortVersionSlice(versionList)
    for _, v := range sortedVersions {
        fmt.Println(v.Raw)
    }
    // Output: 1.0.0, 1.10.0, 2.0.0

    // Find latest version
    latest := sortedVersions[len(sortedVersions)-1]
    fmt.Printf("Latest: %s\n", latest.Raw)
}
```

## Important Notes

- SortVersionStringSlice creates new Version objects internally — O(n) extra space
- SortVersionSlice uses a group-based algorithm for stable, semantically correct ordering
- "1.10.0" correctly sorts after "1.2.0" (numeric, not alphabetic)
- Pre-release versions sort before their release counterparts
- Neither function modifies the original input slice
```

- [ ] **Step 4: 创建 version-file-operations Skill — 版本号文件操作技能**

```markdown
---
name: version-file-operations
description: Use when reading version numbers from files, processing version lists stored in text files. Provides expert guidance on using the Go versions SDK for file-based version operations.
argument-hint: <file-path-or-task>
---

# Version File Operations Skill

## When to Use

- User needs to read version numbers from a text file
- User needs to process a list of versions stored line-by-line in a file
- User needs to parse version lists from dependency lock files or release logs
- User is building CI/CD tools that read version files

## API Reference

### ReadVersionsFromFile

**func ReadVersionsFromFile(filepath string) ([]*Version, error)**

Reads a file line-by-line, parses each line as a Version object. Ignores empty lines and lines starting with #.

### ReadVersionsStringFromFile

**func ReadVersionsStringFromFile(filepath string) ([]string, error)**

Reads a file line-by-line, returns raw strings. Does NOT parse into Version objects. Ignores empty lines and lines starting with #.

## Code Examples

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/versions"
)

func main() {
    // Read and parse versions from file
    versionList, err := versions.ReadVersionsFromFile("versions.txt")
    if err != nil {
        log.Fatalf("Failed to read: %v", err)
    }
    fmt.Printf("Read %d versions\n", len(versionList))

    // Read raw version strings
    rawStrings, err := versions.ReadVersionsStringFromFile("versions.txt")
    if err != nil {
        log.Fatalf("Failed to read: %v", err)
    }
    for _, s := range rawStrings {
        fmt.Println(s)
    }

    // Filter valid versions after reading
    validVersions := make([]*versions.Version, 0)
    for _, v := range versionList {
        if v.IsValid() {
            validVersions = append(validVersions, v)
        }
    }
    fmt.Printf("Valid versions: %d\n", len(validVersions))
}
```

## File Format

```
# This is a comment — lines starting with # are ignored
1.0.0
1.0.1

# Blank lines are also ignored
1.1.0-beta
1.1.0
```

## Important Notes

- File format: one version per line, # for comments, blank lines ignored
- Leading/trailing whitespace on each line is trimmed
- ReadVersionsFromFile uses NewVersion (not NewVersionE), so invalid lines become Version objects with IsValid() == false
- To handle invalid versions, filter with IsValid() after reading
- File reading uses os.ReadFile — for very large files, consider streaming
```

- [ ] **Step 5: 验证核心 Skills 文件存在且格式正确**
Run: `cd /home/cc11001100/github/scagogogo/versions && for f in skills/version-parsing/SKILL.md skills/version-comparison/SKILL.md skills/version-sorting/SKILL.md skills/version-file-operations/SKILL.md; do [ -f "$f" ] && head -3 "$f" | grep -q "^---$" && echo "OK: $f" || echo "FAIL: $f"; done`
Expected:
  - Exit code: 0
  - Output contains: "OK: skills/version-parsing/SKILL.md"
  - Output does NOT contain: "FAIL"

- [ ] **Step 6: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add skills/ && git commit -m "feat(skills): add core skills — version-parsing, comparison, sorting, file-operations"`

---

### Task 3: 高级 Skills 定义（分组、范围查询、可视化）

**Depends on:** Task 1
**Files:**
- Create: `skills/version-grouping/SKILL.md`
- Create: `skills/version-range-query/SKILL.md`
- Create: `skills/version-visualization/SKILL.md`

- [ ] **Step 1: 创建 version-grouping Skill — 版本号分组技能**

```markdown
---
name: version-grouping
description: Use when grouping versions by major/minor version number, managing version collections by their version series. Provides expert guidance on using the Go versions SDK for version grouping.
argument-hint: <version-grouping-task>
---

# Version Grouping Skill

## When to Use

- User needs to group versions by their major version number
- User needs to organize a large collection of versions into version series
- User needs to find all versions in a specific major version (e.g., all 1.x versions)
- User is building a version selector UI with hierarchical version choices

## API Reference

### Group Function

**func Group(versions []*Version) map[string]*VersionGroup**

Groups versions by their VersionNumbers' BuildGroupID (full number sequence). Returns a map of groupID → VersionGroup.

### VersionGroup Type

```go
type VersionGroup struct {
    GroupVersionNumbers VersionNumbers
    VersionMap          map[string]*Version
}
```

**Key Methods:**

- **NewVersionGroup(groupVersionNumbers VersionNumbers) *VersionGroup** — create a new group
- **NewVersionGroupFromVersions(versions []*Version) *VersionGroup** — create from version array
- **func (x *VersionGroup) Add(v *Version) bool** — add version to group
- **func (x *VersionGroup) Contains(v *Version) bool** — check if version exists
- **func (x *VersionGroup) ID() string** — get group ID (e.g., "1.2")
- **func (x *VersionGroup) Versions() []*Version** — get all versions (unordered)
- **func (x *VersionGroup) SortVersions() []*Version** — get sorted versions
- **func (x *VersionGroup) QueryRangeVersions(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version** — range query within group

### SortedVersionGroups Type

```go
type SortedVersionGroups struct { ... }
```

**Key Methods:**

- **func NewSortedVersionGroups(versions []*Version) *SortedVersionGroups** — create sorted groups
- **func (x *SortedVersionGroups) GroupIDs() []string** — get sorted group ID list
- **func (x *SortedVersionGroups) QueryRange(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version** — cross-group range query

## Code Examples

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions"
    "github.com/golang-infrastructure/go-tuple"
)

func main() {
    // Group versions
    versionList := versions.NewVersions("1.0.0", "1.1.0", "2.0.0", "2.1.0", "3.0.0")
    groupMap := versions.Group(versionList)
    for groupID, group := range groupMap {
        fmt.Printf("Group %s: %d versions\n", groupID, len(group.Versions()))
    }

    // Create sorted groups and list IDs
    sortedGroups := versions.NewSortedVersionGroups(versionList)
    fmt.Printf("Group IDs: %v\n", sortedGroups.GroupIDs())

    // Find latest in a group
    if group, ok := groupMap["1.0"]; ok {
        sorted := group.SortVersions()
        latest := sorted[len(sorted)-1]
        fmt.Printf("Latest in group 1.0: %s\n", latest.Raw)
    }
}
```

## Important Notes

- Group() uses BuildGroupID() as the key — this is the full version number string (e.g., "1.2.3"), NOT just the major version
- VersionMap is keyed by Raw string — duplicates overwrite
- Versions() returns unordered results; use SortVersions() for ordered
- SortedVersionGroups pre-sorts on construction — efficient for repeated queries
```

- [ ] **Step 2: 创建 version-range-query Skill — 版本号范围查询技能**

```markdown
---
name: version-range-query
description: Use when querying versions within a specific range, implementing version constraints, or checking version compatibility. Provides expert guidance on using the Go versions SDK for range queries.
argument-hint: <version-range-query-task>
---

# Version Range Query Skill

## When to Use

- User needs to find all versions between a start and end version
- User needs to implement version constraint logic (e.g., ">=1.0.0 and <2.0.0")
- User needs to check if a version falls within a specific range
- User is implementing dependency version resolution or compatibility checks

## API Reference

### ContainsPolicy

```go
type ContainsPolicy int

const (
    ContainsPolicyNone ContainsPolicy = iota  // Unspecified
    ContainsPolicyYes                          // Include boundary
    ContainsPolicyNo                           // Exclude boundary
)
```

Used to control whether range boundaries are included or excluded. Mathematically: ContainsPolicyYes = [ or ], ContainsPolicyNo = ( or ).

### VersionGroup.QueryRangeVersions

**func (x *VersionGroup) QueryRangeVersions(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version**

Queries versions within a range inside a single group. Returns sorted results.

### SortedVersionGroups.QueryRange

**func (x *SortedVersionGroups) QueryRange(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version**

Cross-group range query. Searches across all relevant version groups. Returns sorted results.

## Code Examples

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions"
    "github.com/golang-infrastructure/go-tuple"
)

func main() {
    versionList := versions.NewVersions("1.0.0", "1.1.0", "2.0.0", "2.1.0", "3.0.0")
    sortedGroups := versions.NewSortedVersionGroups(versionList)

    // Query [1.0.0, 2.0.0) — include 1.0.0, exclude 2.0.0
    start := tuple.New2(versions.NewVersion("1.0.0"), versions.ContainsPolicyYes)
    end := tuple.New2(versions.NewVersion("2.0.0"), versions.ContainsPolicyNo)
    result := sortedGroups.QueryRange(start, end)
    fmt.Printf("Versions in [1.0.0, 2.0.0): %d\n", len(result))
    for _, v := range result {
        fmt.Println(v.Raw)
    }
    // Output: 1.0.0, 1.1.0

    // Query [2.0.0, ∞) — all versions >= 2.0.0
    start2 := tuple.New2(versions.NewVersion("2.0.0"), versions.ContainsPolicyYes)
    end2 := tuple.New2(versions.NewVersion("999.0.0"), versions.ContainsPolicyNo)
    result2 := sortedGroups.QueryRange(start2, end2)
    fmt.Printf("Versions >= 2.0.0: %d\n", len(result2))

    // Single group range query
    groupMap := versions.Group(versionList)
    for _, group := range groupMap {
        gStart := tuple.New2(versions.NewVersion("0.0.0"), versions.ContainsPolicyYes)
        gEnd := tuple.New2(versions.NewVersion("99.0.0"), versions.ContainsPolicyYes)
        rangeResult := group.QueryRangeVersions(gStart, gEnd)
        fmt.Printf("Group %s: %d versions in range\n", group.ID(), len(rangeResult))
    }
}
```

## Important Notes

- tuple.New2 creates a Tuple2 — import "github.com/golang-infrastructure/go-tuple"
- ContainsPolicyNone defaults to ContainsPolicyYes (inclusive) in most contexts
- QueryRange on SortedVersionGroups returns nil if the start version's group doesn't exist
- For open-ended ranges, use a very large version as the end boundary
- Range queries return pre-sorted results (ascending order)
```

- [ ] **Step 3: 创建 version-visualization Skill — 版本号可视化技能**

```markdown
---
name: version-visualization
description: Use when visualizing version hierarchies, displaying version trees, or showing version group structures in text format. Provides expert guidance on using the Go versions SDK for version visualization.
argument-hint: <version-visualization-task>
---

# Version Visualization Skill

## When to Use

- User needs to display version structure as a text tree
- User needs to visualize how versions are grouped and organized
- User needs to generate a human-readable overview of a version collection
- User is debugging version management logic or presenting version info in CLI

## API Reference

### VisualizeVersions

**func VisualizeVersions(versions []*Version, w io.Writer, maxItemsPerGroup int)**

Renders a detailed tree view of versions grouped by major version. Shows individual versions with optional release times. maxItemsPerGroup controls truncation (0 = show all).

### VisualizeVersionGroups

**func VisualizeVersionGroups(versions []*Version, w io.Writer)**

Renders a summary tree showing only group-level information (group ID and version count). Useful for large collections.

## Code Examples

```go
package main

import (
    "os"
    "time"
    "github.com/scagogogo/versions"
)

func main() {
    versionList := versions.NewVersions(
        "1.0.0", "1.0.1", "1.1.0",
        "2.0.0", "2.0.1", "2.1.0",
        "3.0.0-alpha", "3.0.0-beta",
    )

    // Set release times (optional)
    for i, v := range versionList {
        v.PublicTime = time.Now().AddDate(0, -i, 0)
    }

    // Detailed visualization — max 2 per group
    versions.VisualizeVersions(versionList, os.Stdout, 2)
    // Output:
    // 版本总数: 8
    // 版本组数: 2
    //
    // ┌─ 版本组: 1 (3个版本)
    // ├── 1.0.0 (发布时间: 2024-01-01)
    // ├── 1.0.1 (发布时间: 2024-02-01)
    // └── ...还有1个版本未显示
    //
    // ┌─ 版本组: 2 ...

    // Summary visualization
    versions.VisualizeVersionGroups(versionList, os.Stdout)
    // Output:
    // 版本总数: 8
    // 版本组数: 5
    //
    // ├─ 1 (3个版本组, 共3个版本)
    // │  ├─ 1.0 (1个版本)
    // │  └─ 1.1 (2个版本)
    // └─ 2 (3个版本组, 共3个版本)
}
```

## Important Notes

- Output is in Chinese (中文) — version labels like "版本总数", "版本组" etc.
- Uses Unicode box-drawing characters for tree structure (├──, └──, ┌─)
- Release time is only shown when PublicTime is non-zero
- maxItemsPerGroup = 0 means show all versions (no truncation)
- VisualizeVersionGroups shows group hierarchy, not individual versions
- Write to any io.Writer — os.Stdout, bytes.Buffer, or file
```

- [ ] **Step 4: 验证高级 Skills 文件存在且格式正确**
Run: `cd /home/cc11001100/github/scagogogo/versions && for f in skills/version-grouping/SKILL.md skills/version-range-query/SKILL.md skills/version-visualization/SKILL.md; do [ -f "$f" ] && head -3 "$f" | grep -q "^---$" && echo "OK: $f" || echo "FAIL: $f"; done`
Expected:
  - Exit code: 0
  - Output contains: "OK: skills/version-grouping/SKILL.md"
  - Output does NOT contain: "FAIL"

- [ ] **Step 5: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add skills/ && git commit -m "feat(skills): add advanced skills — version-grouping, range-query, visualization"`

---

### Task 4: 更新 README 和 .gitignore

**Depends on:** Task 2, Task 3
**Files:**
- Modify: `README.md:1-1597`
- Modify: `.gitignore:1-2`

- [ ] **Step 1: 更新 .gitignore — 确保 .claude-plugin 目录不被忽略**
文件: `.gitignore:1-2`

```text
test
test2
```

（.gitignore 无需修改 — 当前规则不涉及 .claude-plugin/ 和 skills/ 目录，它们会被正常跟踪。保持 .gitignore 不变。）

- [ ] **Step 2: 在 README.md 末尾添加 Skills 仓库说明 — 让用户知道如何安装和使用 Skills**
文件: `README.md:1595-1597`（在许可证部分之前插入）

```markdown
---

## 🤖 Claude Code Skills

本项目同时是一个 **Claude Code Skills 仓库**，提供版本号处理的专业技能。安装后，您可以在 Claude Code 中通过斜杠命令调用这些技能。

### 安装方式

```bash
# 添加 marketplace
claude marketplace add versions https://github.com/scagogogo/versions

# 安装插件
claude plugin install versions@versions
```

### 可用 Skills

| Skill | 命令 | 用途 |
|:------|:-----|:-----|
| 版本号解析 | `/version-parsing` | 解析、验证版本字符串 |
| 版本号比较 | `/version-comparison` | 比较两个版本号大小 |
| 版本号排序 | `/version-sorting` | 对版本号列表排序 |
| 版本号分组 | `/version-grouping` | 按主版本号分组管理 |
| 版本范围查询 | `/version-range-query` | 查询指定范围内的版本 |
| 版本号可视化 | `/version-visualization` | 以树形结构展示版本层次 |
| 版本号文件操作 | `/version-file-operations` | 从文件读取版本号列表 |

### 使用示例

在 Claude Code 中输入：

```
/version-parsing v1.2.3-beta1
```

Claude 将基于此 Skill 提供专业的版本号解析指导，包括正确的 Go 代码示例和 API 用法。
```

- [ ] **Step 3: 验证 README 更新**
Run: `cd /home/cc11001100/github/scagogogo/versions && grep -c "Claude Code Skills" README.md && grep -c "version-parsing" README.md && grep -c "version-visualization" README.md`
Expected:
  - Exit code: 0
  - Output contains: "1" (three times, one for each grep)

- [ ] **Step 4: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add README.md && git commit -m "docs: add Claude Code Skills installation guide to README"`

---

### Task 5: 最终验证与全项目测试

**Depends on:** Task 1, Task 2, Task 3, Task 4
**Files:**
- No new files — verification only

- [ ] **Step 1: 验证 Go 测试仍然通过**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test ./...`
Expected:
  - Exit code: 0
  - Output contains: "ok"
  - Output does NOT contain: "FAIL"

- [ ] **Step 2: 验证 Skills 目录结构完整**
Run: `cd /home/cc11001100/github/scagogogo/versions && find skills/ -name "SKILL.md" | sort`
Expected:
  - Exit code: 0
  - Output contains: "skills/version-comparison/SKILL.md"
  - Output contains: "skills/version-file-operations/SKILL.md"
  - Output contains: "skills/version-grouping/SKILL.md"
  - Output contains: "skills/version-parsing/SKILL.md"
  - Output contains: "skills/version-range-query/SKILL.md"
  - Output contains: "skills/version-sorting/SKILL.md"
  - Output contains: "skills/version-visualization/SKILL.md"
  - Total count: 7 files

- [ ] **Step 3: 验证插件元数据结构完整**
Run: `cd /home/cc11001100/github/scagogogo/versions && [ -f .claude-plugin/plugin.json ] && [ -f .claude-plugin/marketplace.json ] && echo "Plugin metadata OK"`
Expected:
  - Exit code: 0
  - Output contains: "Plugin metadata OK"

- [ ] **Step 4: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add -A && git status --porcelain | wc -l | xargs -I{} test {} -eq 0 && echo "Clean working tree" || echo "Unexpected changes"`

Expected:
  - Exit code: 0
  - Output contains: "Clean working tree"
