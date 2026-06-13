# 版本 SDK API 第四阶段扩展 Plan

> **For agentic workers:** REQUIRED SUB-SKILL: `superpowers:subagent-driven-development`
> Steps use checkbox (`- [ ]`) syntax.

**Goal:** 补全高优先级 API 缺口：Version.RawString()、WithNumbers()、TextMarshaler/TextUnmarshaler、String() 方法补全（VersionSuffix/VersionPrefix/SuffixWeight/ContainsPolicy/VersionGroup）、VersionSlice 排序接口、Version.SubVersion()/SuffixWeight()、FilterByPatch/FilterBySuffix/FilterByStable/FilterByPrerelease、VersionBuilder.PublicTime()、Version.WithPublicTime()、Version.IsZero()、visualize 死代码清理、VersionGroup.String()/Filter()、SortedVersionGroups.Versions()。

**Architecture:** 所有改动均为纯新增方法或函数，不修改现有方法签名。Task 间互相独立。

**Tech Stack:** Go 1.18, 现有依赖不变

**Risks:**
- Task 3 添加 TextMarshaler/TextUnmarshaler 会影响 json.Marshal 行为（json 包优先调用 MarshalText）→ 缓解：MarshalText 返回 Raw 字符串，这是最自然的序列化形式
- Task 2 添加 VersionSlice 类型和 sort.Interface，这是新类型，不影响现有代码
- 所有其他 Task 均为纯新增，无兼容性风险

---

### Task 1: Version 高优先级补充方法

**Depends on:** None
**Files:**
- Modify: `version.go`
- Modify: `version_convenience_test.go`

添加: RawString(), WithNumbers(), SubVersion(), SuffixWeight(), WithPublicTime(), IsZero()

---

### Task 2: VersionSlice 排序类型 + String() 方法补全

**Depends on:** None
**Files:**
- Create: `version_slice.go` (VersionSlice 类型 + sort.Interface)
- Create: `version_slice_test.go`
- Modify: `version_suffix.go` (添加 String())
- Modify: `version_prefix.go` (添加 String() + PurePrefix())
- Modify: `suffix_weight.go` (添加 String())
- Modify: `contains_policy.go` (添加 String())
- Modify: `version_group.go` (添加 String())
- Modify: `sorted_version_groups.go` (添加 Versions())
- 对应测试文件

---

### Task 3: Version 编码接口 + Builder 增强

**Depends on:** None
**Files:**
- Modify: `version.go` (MarshalText/UnmarshalText)
- Modify: `version_builder.go` (PublicTime setter)
- Modify: `version_clone.go` (WithPublicTime)
- 对应测试文件

---

### Task 4: 过滤工具函数扩展 + SortedVersionGroups.Versions

**Depends on:** None
**Files:**
- Modify: `version_utils.go` (FilterByPatch/FilterBySuffix/FilterByStable/FilterByPrerelease)
- Modify: `version_utils_test.go`
- Modify: `sorted_version_groups.go` (Versions())
- Modify: `sorted_version_groups_test.go`

---

### Task 5: Visualize 死代码清理

**Depends on:** None
**Files:**
- Modify: `visualize.go` (移除 majorGroups 死代码和误导注释)
