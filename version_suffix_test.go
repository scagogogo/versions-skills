package versions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestVersionSuffix_IsEmpty 测试版本后缀的空值检查方法
func TestVersionSuffix_IsEmpty(t *testing.T) {
	// 空后缀
	emptySuffix := VersionSuffix("")
	assert.True(t, emptySuffix.IsEmpty())
	assert.True(t, EmptyVersionSuffix.IsEmpty())

	// 非空后缀
	nonEmptySuffix1 := VersionSuffix("-beta")
	assert.False(t, nonEmptySuffix1.IsEmpty())

	nonEmptySuffix2 := VersionSuffix(".RC1")
	assert.False(t, nonEmptySuffix2.IsEmpty())

	// 特殊字符后缀
	specialSuffix := VersionSuffix(" ")
	assert.False(t, specialSuffix.IsEmpty())
}

// TestVersionSuffix_CompareTo 测试版本后缀的比较方法
func TestVersionSuffix_CompareTo(t *testing.T) {
	// 相等的后缀
	suffix1 := VersionSuffix("-beta1")
	suffix1Dup := VersionSuffix("-beta1")
	assert.Equal(t, 0, suffix1.CompareTo(suffix1Dup))

	// 不同的后缀 - 字典序比较
	alpha := VersionSuffix("-alpha")
	beta := VersionSuffix("-beta")
	assert.Equal(t, -1, alpha.CompareTo(beta)) // alpha < beta
	assert.Equal(t, 1, beta.CompareTo(alpha))  // beta > alpha

	// 不同长度但有共同前缀
	rc1 := VersionSuffix("-rc1")
	rc2 := VersionSuffix("-rc2")
	assert.Equal(t, -1, rc1.CompareTo(rc2)) // rc1 < rc2
	assert.Equal(t, 1, rc2.CompareTo(rc1))  // rc2 > rc1

	// 空后缀比较
	empty := EmptyVersionSuffix
	nonEmpty := VersionSuffix("-any")
	assert.Equal(t, -1, empty.CompareTo(nonEmpty)) // 空 < 非空
	assert.Equal(t, 1, nonEmpty.CompareTo(empty))  // 非空 > 空
	assert.Equal(t, 0, empty.CompareTo(empty))     // 空 = 空
}

func TestVersionSuffix_String(t *testing.T) {
	s := VersionSuffix("-beta1")
	if s.String() != "-beta1" {
		t.Errorf("String() = %q, want %q", s.String(), "-beta1")
	}
}
