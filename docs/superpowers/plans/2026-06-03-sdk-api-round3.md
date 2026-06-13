# 版本 SDK API 第三阶段扩展 Plan

> **For agentic workers:** REQUIRED SUB-SKILL: `superpowers:subagent-driven-development`
> Steps use checkbox (`- [ ]`) syntax.

**Goal:** 完善版本 SDK 的 API 覆盖率：补全 Is* 便捷方法（IsPre/IsRelease/IsSP/IsPost）、添加 Version 约束检查方法（Satisfies/Matches）、SortedVersionGroups 集合访问器（Len/Get/At/Contains）、VersionGroup 增强方法（Remove/LatestPrerelease）、I/O 补全（WriteVersionsToFile/ReadVersionsFromReader）、集合运算扩展（Union/Partition）和 VersionNumbers 辅助方法（Len/String/At/Equals）。

**Architecture:** 所有改动均为纯新增方法或函数，不修改现有方法签名。Task 1-6 互相独立，可按任意顺序执行。Version.Satisfies/Matches 在 Version 上添加方法调用已有 Constraint API；SortedVersionGroups 方法基于已有的 groupSlice 索引；VersionNumbers 方法为简单的切片包装。

**Tech Stack:** Go 1.18, 现有依赖不变

**Risks:**
- 所有 Task 均为纯新增，无兼容性风险
- Task 6 修改 version_numbers.go，该类型被多个模块引用，但新增方法不影响现有代码

---

### Task 1: 补全 Is* 便捷方法（IsPre/IsRelease/IsSP/IsPost）

**Depends on:** None
**Files:**
- Modify: `version.go`（在 IsGA() 方法之后添加）
- Modify: `version_convenience_test.go`（添加新测试）

- [ ] **Step 1: 修改 Version — 添加 IsPre/IsRelease/IsSP/IsPost 方法**
文件: `version.go`（在 `IsGA()` 方法之后添加以下方法）

```go
// IsPre 判断版本是否为预发布版（pre 标识）
//
// 预发布版是指后缀包含 pre 标识的版本，如 "1.0.0-pre1"。
// 注意：这是显式的 pre 后缀，与 IsPrerelease()（判断是否有任何后缀）不同。
func (x *Version) IsPre() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightPre
}

// IsRelease 判断版本是否带有 release 后缀
//
// 带 release 后缀的版本如 "1.0.0-release"。
// 注意：这与 IsStable()（无后缀）不同，IsRelease 检测的是显式的 "-release" 后缀。
func (x *Version) IsRelease() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightRelease
}

// IsSP 判断版本是否为服务包版本
//
// 服务包版本是指后缀包含 sp 标识的版本，如 "1.0.0-sp1"。
func (x *Version) IsSP() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightSP
}

// IsPost 判断版本是否为 Post 发布版
//
// Post 发布版是指后缀包含 post 标识的版本（PEP 440 规范），如 "1.0.0-post1"。
func (x *Version) IsPost() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightPost
}
```

- [ ] **Step 2: 修改 version_convenience_test.go — 添加新测试**
文件: `version_convenience_test.go`（在文件末尾添加）

```go
func TestVersion_IsPre(t *testing.T) {
	if !NewVersion("1.0.0-pre1").IsPre() {
		t.Error("1.0.0-pre1 should be pre")
	}
	if NewVersion("1.0.0-beta").IsPre() {
		t.Error("1.0.0-beta should not be pre")
	}
}

func TestVersion_IsRelease(t *testing.T) {
	if !NewVersion("1.0.0-release").IsRelease() {
		t.Error("1.0.0-release should be release")
	}
	if NewVersion("1.0.0").IsRelease() {
		t.Error("1.0.0 should not be release (no suffix)")
	}
}

func TestVersion_IsSP(t *testing.T) {
	if !NewVersion("1.0.0-sp1").IsSP() {
		t.Error("1.0.0-sp1 should be SP")
	}
	if NewVersion("1.0.0").IsSP() {
		t.Error("1.0.0 should not be SP")
	}
}

func TestVersion_IsPost(t *testing.T) {
	if !NewVersion("1.0.0-post1").IsPost() {
		t.Error("1.0.0-post1 should be post")
	}
	if NewVersion("1.0.0").IsPost() {
		t.Error("1.0.0 should not be post")
	}
}
```

- [ ] **Step 3: 验证**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test -run "TestVersion_IsPre|TestVersion_IsRelease|TestVersion_IsSP|TestVersion_IsPost" -v ./...`
Expected:
  - Exit code: 0
  - Output contains: "PASS"

- [ ] **Step 4: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add version.go version_convenience_test.go && git commit -m "feat(version): add IsPre, IsRelease, IsSP, IsPost convenience methods"`

---

### Task 2: Version 约束检查方法（Satisfies/Matches）

**Depends on:** None
**Files:**
- Modify: `version.go`（在 IsPost() 方法之后添加）
- Modify: `version_convenience_test.go`（添加新测试）

- [ ] **Step 1: 修改 Version — 添加 Satisfies 和 Matches 方法**
文件: `version.go`（在最后一个 Is* 方法之后添加）

```go
// Satisfies 判断版本是否满足给定的约束条件
//
// 这是 Constraint.Match(v) 的反向调用方式，语义更自然：
// v.Satisfies(constraint) 等价于 constraint.Match(v)。
//
// 参数:
//   - constraint: 版本约束条件
//
// 返回:
//   - bool: 如果版本满足约束则返回 true
//
// 使用示例:
//
//	c, _ := versions.ParseConstraint(">=1.0.0")
//	v := versions.NewVersion("1.5.0")
//	if v.Satisfies(c) {
//	    fmt.Println("版本满足约束")
//	}
func (x *Version) Satisfies(constraint *Constraint) bool {
	return constraint.Match(x)
}

// Matches 判断版本是否满足约束表达式字符串
//
// 该方法是 ParseConstraint + Match 的便捷组合，
// 适用于需要从字符串快速判断版本是否满足约束的场景。
//
// 参数:
//   - expr: 约束表达式字符串，如 ">=1.0.0"
//
// 返回:
//   - bool: 如果版本满足约束则返回 true
//   - error: 如果约束表达式解析失败则返回错误
//
// 使用示例:
//
//	v := versions.NewVersion("1.5.0")
//	ok, err := v.Matches(">=1.0.0,<2.0.0")
//	if ok {
//	    fmt.Println("版本在范围内")
//	}
func (x *Version) Matches(expr string) (bool, error) {
	cs, err := ParseConstraintSet(expr)
	if err != nil {
		return false, err
	}
	return cs.Match(x), nil
}
```

- [ ] **Step 2: 修改 version_convenience_test.go — 添加测试**
文件: `version_convenience_test.go`（在文件末尾添加）

```go
func TestVersion_Satisfies(t *testing.T) {
	c, _ := ParseConstraint(">=1.0.0")
	v := NewVersion("1.5.0")
	if !v.Satisfies(c) {
		t.Error("1.5.0 should satisfy >=1.0.0")
	}
	v2 := NewVersion("0.9.0")
	if v2.Satisfies(c) {
		t.Error("0.9.0 should not satisfy >=1.0.0")
	}
}

func TestVersion_Matches(t *testing.T) {
	v := NewVersion("1.5.0")
	ok, err := v.Matches(">=1.0.0,<2.0.0")
	if err != nil {
		t.Fatalf("Matches() error: %v", err)
	}
	if !ok {
		t.Error("1.5.0 should match >=1.0.0,<2.0.0")
	}

	ok2, err2 := v.Matches(">=2.0.0")
	if err2 != nil {
		t.Fatalf("Matches() error: %v", err2)
	}
	if ok2 {
		t.Error("1.5.0 should not match >=2.0.0")
	}

	_, err3 := v.Matches("not-valid")
	if err3 == nil {
		t.Error("Matches() should return error for invalid expression")
	}
}
```

- [ ] **Step 3: 验证**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test -run "TestVersion_Satisfies|TestVersion_Matches" -v ./...`
Expected:
  - Exit code: 0
  - Output contains: "PASS"

- [ ] **Step 4: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add version.go version_convenience_test.go && git commit -m "feat(version): add Satisfies and Matches constraint check methods"`

---

### Task 3: SortedVersionGroups 集合访问器

**Depends on:** None
**Files:**
- Modify: `sorted_version_groups.go`（在 QueryRange 方法之后添加）
- Modify: `sorted_version_groups_test.go`（添加新测试）

- [ ] **Step 1: 修改 SortedVersionGroups — 添加 Len/Get/At/Contains 方法**
文件: `sorted_version_groups.go`（在 `QueryRange()` 方法之后添加）

```go
// Len 返回版本组的数量
func (x *SortedVersionGroups) Len() int {
	return len(x.groupSlice)
}

// Get 根据组 ID 获取版本组
//
// 如果组 ID 不存在则返回 nil。
//
// 参数:
//   - groupID: 版本组 ID，如 "1.2"
//
// 返回:
//   - *VersionGroup: 对应的版本组，不存在则返回 nil
func (x *SortedVersionGroups) Get(groupID string) *VersionGroup {
	idx, exists := x.groupIdToIndexMap[groupID]
	if !exists {
		return nil
	}
	return x.groupSlice[idx]
}

// At 根据索引获取版本组
//
// 返回排序后指定位置的版本组。如果索引越界则返回 nil。
//
// 参数:
//   - index: 从 0 开始的索引位置
//
// 返回:
//   - *VersionGroup: 对应的版本组，越界则返回 nil
func (x *SortedVersionGroups) At(index int) *VersionGroup {
	if index < 0 || index >= len(x.groupSlice) {
		return nil
	}
	return x.groupSlice[index]
}

// Contains 检查是否包含指定组 ID 的版本组
//
// 参数:
//   - groupID: 版本组 ID
//
// 返回:
//   - bool: 如果存在则返回 true
func (x *SortedVersionGroups) Contains(groupID string) bool {
	_, exists := x.groupIdToIndexMap[groupID]
	return exists
}
```

- [ ] **Step 2: 修改 sorted_version_groups_test.go — 添加新测试**
文件: `sorted_version_groups_test.go`（在文件末尾添加）

```go
func TestSortedVersionGroups_Len(t *testing.T) {
	versions := NewVersions("1.0.0", "1.1.0", "2.0.0")
	svg := NewSortedVersionGroups(versions)
	if svg.Len() != 2 {
		t.Errorf("Len() = %d, want 2", svg.Len())
	}
}

func TestSortedVersionGroups_Get(t *testing.T) {
	versions := NewVersions("1.0.0", "1.1.0", "2.0.0")
	svg := NewSortedVersionGroups(versions)
	g := svg.Get("1.0")
	if g == nil {
		t.Fatal("Get(\"1.0\") returned nil")
	}
	if g.ID() != "1.0" {
		t.Errorf("Get().ID() = %q, want %q", g.ID(), "1.0")
	}
	if svg.Get("99.99") != nil {
		t.Error("Get(\"99.99\") should return nil for non-existent group")
	}
}

func TestSortedVersionGroups_At(t *testing.T) {
	versions := NewVersions("2.0.0", "1.0.0")
	svg := NewSortedVersionGroups(versions)
	first := svg.At(0)
	if first == nil {
		t.Fatal("At(0) returned nil")
	}
	// 排序后第一个应该是较小的版本组
	if first.ID() != "1.0" {
		t.Errorf("At(0).ID() = %q, want %q", first.ID(), "1.0")
	}
	if svg.At(-1) != nil {
		t.Error("At(-1) should return nil")
	}
	if svg.At(99) != nil {
		t.Error("At(99) should return nil")
	}
}

func TestSortedVersionGroups_Contains(t *testing.T) {
	versions := NewVersions("1.0.0", "2.0.0")
	svg := NewSortedVersionGroups(versions)
	if !svg.Contains("1.0") {
		t.Error("Should contain group 1.0")
	}
	if svg.Contains("99.99") {
		t.Error("Should not contain group 99.99")
	}
}
```

- [ ] **Step 3: 验证**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test -run "TestSortedVersionGroups_(Len|Get|At|Contains)" -v ./...`
Expected:
  - Exit code: 0
  - Output contains: "PASS"

- [ ] **Step 4: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add sorted_version_groups.go sorted_version_groups_test.go && git commit -m "feat(sorted_groups): add Len, Get, At, Contains collection accessors"`

---

### Task 4: VersionGroup 增强方法

**Depends on:** None
**Files:**
- Modify: `version_group.go`（在 LatestStable() 方法之后添加）
- Modify: `version_group_convenience_test.go`（添加新测试）

- [ ] **Step 1: 修改 VersionGroup — 添加 Remove 和 LatestPrerelease 方法**
文件: `version_group.go`（在 `LatestStable()` 方法之后添加）

```go
// Remove 从版本组中移除指定的版本
//
// 如果版本存在于组中，则删除并返回 true；否则返回 false。
//
// 参数:
//   - v: 要移除的版本对象
//
// 返回:
//   - bool: 如果版本存在并被移除则返回 true
func (x *VersionGroup) Remove(v *Version) bool {
	if _, exists := x.VersionMap[v.Raw]; exists {
		delete(x.VersionMap, v.Raw)
		return true
	}
	return false
}

// LatestPrerelease 获取版本组中最新预发布版本
//
// 返回按排序后最新的预发布版本，如果组中无预发布版本则返回 nil。
func (x *VersionGroup) LatestPrerelease() *Version {
	return LatestPrerelease(x.Versions())
}
```

- [ ] **Step 2: 修改 version_group_convenience_test.go — 添加测试**
文件: `version_group_convenience_test.go`（在文件末尾添加）

```go
func TestVersionGroup_Remove(t *testing.T) {
	vg := NewVersionGroupFromVersions(NewVersions("1.0.0", "1.0.1", "1.0.2"))
	if vg.Count() != 3 {
		t.Fatalf("Count() = %d, want 3", vg.Count())
	}
	removed := vg.Remove(NewVersion("1.0.1"))
	if !removed {
		t.Error("Remove should return true for existing version")
	}
	if vg.Count() != 2 {
		t.Errorf("Count() after remove = %d, want 2", vg.Count())
	}
	removed2 := vg.Remove(NewVersion("9.9.9"))
	if removed2 {
		t.Error("Remove should return false for non-existing version")
	}
}

func TestVersionGroup_LatestPrerelease(t *testing.T) {
	vg := NewVersionGroupFromVersions(NewVersions("1.0.0", "1.0.1-alpha", "1.0.1-beta", "1.0.0-rc1"))
	latest := vg.LatestPrerelease()
	if latest == nil {
		t.Fatal("LatestPrerelease() returned nil")
	}
	// rc1 > beta > alpha in suffix weight
	if latest.Raw != "1.0.1-rc1" && latest.Raw != "1.0.0-rc1" {
		t.Errorf("LatestPrerelease() = %q, want rc1", latest.Raw)
	}
}

func TestVersionGroup_LatestPrerelease_None(t *testing.T) {
	vg := NewVersionGroupFromVersions(NewVersions("1.0.0", "1.0.1"))
	if vg.LatestPrerelease() != nil {
		t.Error("LatestPrerelease() should return nil when no prerelease versions")
	}
}
```

- [ ] **Step 3: 验证**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test -run "TestVersionGroup_(Remove|LatestPrerelease)" -v ./...`
Expected:
  - Exit code: 0
  - Output contains: "PASS"

- [ ] **Step 4: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add version_group.go version_group_convenience_test.go && git commit -m "feat(group): add Remove and LatestPrerelease methods to VersionGroup"`

---

### Task 5: I/O 增强（WriteVersionsToFile/ReadVersionsFromReader）

**Depends on:** None
**Files:**
- Modify: `file.go`（在文件末尾添加新函数）
- Modify: `file_test.go`（添加新测试）

- [ ] **Step 1: 修改 file.go — 添加 WriteVersionsToFile 和 ReadVersionsFromReader**
文件: `file.go`（在 `ReadVersionsStringFromFile()` 函数之后添加）

```go
// WriteVersionsToFile 将版本列表写入文件
//
// 每个版本号占一行，写入版本号数字部分拼接而成的字符串。
// 该函数会先对版本进行排序，确保输出有序。
//
// 参数:
//   - versions: 要写入的版本对象列表
//   - filepath: 输出文件路径
//
// 返回:
//   - error: 如果文件写入失败则返回错误
//
// 使用示例:
//
//	versions := versions.NewVersions("2.0.0", "1.0.0", "1.1.0")
//	err := versions.WriteVersionsToFile(versions, "./output.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
func WriteVersionsToFile(versions []*Version, filepath string) error {
	sorted := SortVersionSlice(versions)
	var sb strings.Builder
	for i, v := range sorted {
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(v.Raw)
	}
	return os.WriteFile(filepath, []byte(sb.String()), 0644)
}

// ReadVersionsFromReader 从 io.Reader 读取版本号并解析
//
// 该函数从任意的 io.Reader 中读取版本号列表，每行一个版本号，
// 并将其解析为 Version 对象数组。适用于从网络连接、字符串缓冲区等读取版本。
//
// 参数:
//   - reader: 实现 io.Reader 接口的读取器
//
// 返回:
//   - []*Version: 解析后的 Version 对象数组
//   - error: 如果读取失败则返回错误
//
// 使用示例:
//
//	data := strings.NewReader("1.0.0\n1.1.0\n2.0.0\n")
//	versions, err := versions.ReadVersionsFromReader(data)
//	if err != nil {
//	    log.Fatal(err)
//	}
func ReadVersionsFromReader(reader io.Reader) ([]*Version, error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	versions := make([]*Version, 0)
	for _, line := range strings.Split(string(bytes), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		v := NewVersionStringParser(line).Parse()
		versions = append(versions, v)
	}
	return versions, nil
}
```

- [ ] **Step 2: 添加 io import**
文件: `file.go:3-5`（更新 import 块）

```go
import (
	"io"
	"os"
	"strings"
)
```

- [ ] **Step 3: 修改 file_test.go — 添加新测试**
文件: `file_test.go`（在文件末尾添加）

```go
func TestWriteVersionsToFile(t *testing.T) {
	versions := NewVersions("2.0.0", "1.0.0", "1.1.0")
	tmpFile := filepath.Join(t.TempDir(), "versions.txt")
	err := WriteVersionsToFile(versions, tmpFile)
	if err != nil {
		t.Fatalf("WriteVersionsToFile error: %v", err)
	}
	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}
	content := string(data)
	if !strings.Contains(content, "1.0.0") {
		t.Error("Output should contain 1.0.0")
	}
	if !strings.Contains(content, "1.1.0") {
		t.Error("Output should contain 1.1.0")
	}
	if !strings.Contains(content, "2.0.0") {
		t.Error("Output should contain 2.0.0")
	}
}

func TestReadVersionsFromReader(t *testing.T) {
	data := strings.NewReader("1.0.0\n1.1.0\n2.0.0\n")
	versions, err := ReadVersionsFromReader(data)
	if err != nil {
		t.Fatalf("ReadVersionsFromReader error: %v", err)
	}
	if len(versions) != 3 {
		t.Errorf("len(versions) = %d, want 3", len(versions))
	}
	if versions[0].Raw != "1.0.0" {
		t.Errorf("versions[0].Raw = %q, want %q", versions[0].Raw, "1.0.0")
	}
}
```

- [ ] **Step 4: 添加 import**
文件: `file_test.go`（确保 import 包含以下）

```go
import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)
```

- [ ] **Step 5: 验证**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test -run "TestWriteVersionsToFile|TestReadVersionsFromReader" -v ./...`
Expected:
  - Exit code: 0
  - Output contains: "PASS"

- [ ] **Step 6: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add file.go file_test.go && git commit -m "feat(file): add WriteVersionsToFile and ReadVersionsFromReader"`

---

### Task 6: 集合运算扩展 + VersionNumbers 辅助方法

**Depends on:** None
**Files:**
- Modify: `version_utils.go`（添加 Union、Partition）
- Modify: `version_utils_test.go`（添加测试）
- Modify: `version_numbers.go`（添加 Len、String、At、Equals）
- Create: `version_numbers_test.go`（新测试文件）

- [ ] **Step 1: 修改 version_utils.go — 添加 Union 和 Partition**
文件: `version_utils.go`（在 `Intersection()` 函数之后添加）

```go
// Union 返回 a 和 b 中所有唯一版本的并集
//
// 根据 Raw 字段去重，保持 a 中元素的原始顺序，b 中不重复的元素追加到末尾。
//
// 参数:
//   - a: 版本对象列表
//   - b: 版本对象列表
//
// 返回:
//   - []*Version: 并集版本列表
func Union(a, b []*Version) []*Version {
	result := Unique(a)
	aSet := make(map[string]bool, len(a))
	for _, v := range a {
		aSet[v.Raw] = true
	}
	for _, v := range b {
		if !aSet[v.Raw] {
			result = append(result, v)
		}
	}
	return result
}

// Partition 根据谓词将版本列表分为两组
//
// 返回两个切片：第一个包含满足谓词的版本，第二个包含不满足谓词的版本。
// 保持原始顺序。
//
// 参数:
//   - versions: 版本对象列表
//   - predicate: 分区谓词函数
//
// 返回:
//   - []*Version: 满足谓词的版本列表
//   - []*Version: 不满足谓词的版本列表
func Partition(versions []*Version, predicate func(*Version) bool) ([]*Version, []*Version) {
	matched := make([]*Version, 0)
	unmatched := make([]*Version, 0)
	for _, v := range versions {
		if predicate(v) {
			matched = append(matched, v)
		} else {
			unmatched = append(unmatched, v)
		}
	}
	return matched, unmatched
}
```

- [ ] **Step 2: 修改 version_utils_test.go — 添加测试**
文件: `version_utils_test.go`（在文件末尾添加）

```go
func TestUnion(t *testing.T) {
	a := NewVersions("1.0.0", "1.1.0")
	b := NewVersions("1.1.0", "2.0.0")
	result := Union(a, b)
	if len(result) != 3 {
		t.Errorf("Union() = %d, want 3", len(result))
	}
}

func TestUnion_Empty(t *testing.T) {
	a := NewVersions("1.0.0")
	result := Union(a, nil)
	if len(result) != 1 {
		t.Errorf("Union(a, nil) = %d, want 1", len(result))
	}
}

func TestPartition(t *testing.T) {
	versions := NewVersions("1.0.0", "1.0.0-beta", "1.1.0", "1.1.0-alpha")
	stable, pre := Partition(versions, func(v *Version) bool {
		return v.IsStable()
	})
	if len(stable) != 2 {
		t.Errorf("Partition stable = %d, want 2", len(stable))
	}
	if len(pre) != 2 {
		t.Errorf("Partition prerelease = %d, want 2", len(pre))
	}
}

func TestPartition_Empty(t *testing.T) {
	stable, pre := Partition(nil, func(v *Version) bool {
		return v.IsStable()
	})
	if len(stable) != 0 || len(pre) != 0 {
		t.Errorf("Partition(nil) = (%d, %d), want (0, 0)", len(stable), len(pre))
	}
}
```

- [ ] **Step 3: 修改 version_numbers.go — 添加 Len/String/At/Equals**
文件: `version_numbers.go`（在 `BuildGroupID()` 方法之后添加）

```go
// Len 返回版本号数字部分的长度
func (x VersionNumbers) Len() int {
	return len(x)
}

// At 返回指定索引位置的版本号数字
//
// 如果索引越界则返回 0。
//
// 参数:
//   - index: 从 0 开始的索引
//
// 返回:
//   - int: 对应位置的版本号数字，越界返回 0
func (x VersionNumbers) At(index int) int {
	if index < 0 || index >= len(x) {
		return 0
	}
	return x[index]
}

// String 返回版本号数字部分的字符串表示
//
// 等价于 BuildGroupID()，提供更符合 Go 惯例的 String() 方法。
//
// 返回:
//   - string: 如 "1.2.3"
func (x VersionNumbers) String() string {
	return x.BuildGroupID()
}

// Equals 判断两个版本号数字部分是否相等
//
// 参数:
//   - target: 目标版本号数字部分
//
// 返回:
//   - bool: 如果完全相等则返回 true
func (x VersionNumbers) Equals(target VersionNumbers) bool {
	if len(x) != len(target) {
		return false
	}
	for i := range x {
		if x[i] != target[i] {
			return false
		}
	}
	return true
}
```

- [ ] **Step 4: 创建 version_numbers_test.go**

```go
package versions

import "testing"

func TestVersionNumbers_Len(t *testing.T) {
	vn := NewVersionNumbers([]int{1, 2, 3})
	if vn.Len() != 3 {
		t.Errorf("Len() = %d, want 3", vn.Len())
	}
	empty := NewVersionNumbers([]int{})
	if empty.Len() != 0 {
		t.Errorf("Len() = %d, want 0", empty.Len())
	}
}

func TestVersionNumbers_At(t *testing.T) {
	vn := NewVersionNumbers([]int{1, 2, 3})
	if vn.At(0) != 1 {
		t.Errorf("At(0) = %d, want 1", vn.At(0))
	}
	if vn.At(2) != 3 {
		t.Errorf("At(2) = %d, want 3", vn.At(2))
	}
	if vn.At(5) != 0 {
		t.Errorf("At(5) = %d, want 0 (out of bounds)", vn.At(5))
	}
	if vn.At(-1) != 0 {
		t.Errorf("At(-1) = %d, want 0 (out of bounds)", vn.At(-1))
	}
}

func TestVersionNumbers_String(t *testing.T) {
	vn := NewVersionNumbers([]int{1, 2, 3})
	if vn.String() != "1.2.3" {
		t.Errorf("String() = %q, want %q", vn.String(), "1.2.3")
	}
}

func TestVersionNumbers_Equals(t *testing.T) {
	a := NewVersionNumbers([]int{1, 2, 3})
	b := NewVersionNumbers([]int{1, 2, 3})
	c := NewVersionNumbers([]int{1, 2})
	d := NewVersionNumbers([]int{1, 2, 4})
	if !a.Equals(b) {
		t.Error("a should equal b")
	}
	if a.Equals(c) {
		t.Error("a should not equal c (different length)")
	}
	if a.Equals(d) {
		t.Error("a should not equal d (different values)")
	}
}
```

- [ ] **Step 5: 验证**
Run: `cd /home/cc11001100/github/scagogogo/versions && go test -run "TestUnion|TestPartition|TestVersionNumbers" -v ./...`
Expected:
  - Exit code: 0
  - Output contains: "PASS"

- [ ] **Step 6: 提交**
Run: `cd /home/cc11001100/github/scagogogo/versions && git add version_utils.go version_utils_test.go version_numbers.go version_numbers_test.go && git commit -m "feat(utils): add Union, Partition collection ops and VersionNumbers Len/String/At/Equals helpers"`
