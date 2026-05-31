package versions

import (
	"strconv"
	"strings"
)

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
		sb.WriteString(strconv.Itoa(n))
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
