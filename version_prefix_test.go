package versions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestVersionPrefix_IsEmpty 测试版本前缀的空值检查方法
func TestVersionPrefix_IsEmpty(t *testing.T) {
	// 空前缀
	emptyPrefix := VersionPrefix("")
	assert.True(t, emptyPrefix.IsEmpty())
	assert.True(t, EmptyVersionPrefix.IsEmpty())

	// 非空前缀
	nonEmptyPrefix1 := VersionPrefix("v")
	assert.False(t, nonEmptyPrefix1.IsEmpty())

	nonEmptyPrefix2 := VersionPrefix("version-")
	assert.False(t, nonEmptyPrefix2.IsEmpty())

	// 特殊字符前缀
	specialPrefix := VersionPrefix(" ")
	assert.False(t, specialPrefix.IsEmpty())
}

func TestVersionPrefix_String(t *testing.T) {
	p := VersionPrefix("v")
	if p.String() != "v" {
		t.Errorf("String() = %q, want %q", p.String(), "v")
	}
}

func TestVersionPrefix_PurePrefix(t *testing.T) {
	tests := []struct {
		input    VersionPrefix
		expected string
	}{
		{VersionPrefix("v"), "v"},
		{VersionPrefix("curl-"), "curl"},
		{VersionPrefix("lib."), "lib"},
		{VersionPrefix("a-b."), "a-b"},
		{VersionPrefix(""), ""},
	}
	for _, tt := range tests {
		got := tt.input.PurePrefix()
		if got != tt.expected {
			t.Errorf("PurePrefix(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}
