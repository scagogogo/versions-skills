package versions

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestNewVersionE 测试带错误返回的版本创建函数
func TestNewVersionE(t *testing.T) {
	// 测试有效的版本字符串
	validVersion, err := NewVersionE("1.2.3")
	assert.Nil(t, err)
	assert.NotNil(t, validVersion)
	assert.Equal(t, "1.2.3", validVersion.Raw)
	assert.Equal(t, VersionNumbers{1, 2, 3}, validVersion.VersionNumbers)

	// 测试无效的版本字符串
	invalidVersion, err := NewVersionE("")
	assert.Equal(t, ErrVersionInvalid, err)
	assert.Nil(t, invalidVersion)

	// 测试另一个无效的版本字符串（没有数字部分）
	invalidVersion2, err := NewVersionE("abc")
	assert.Equal(t, ErrVersionInvalid, err)
	assert.Nil(t, invalidVersion2)
}

// TestVersion_IsValid 测试版本有效性检查方法
func TestVersion_IsValid(t *testing.T) {
	// 有效的版本（有数字部分）
	validVersion := NewVersion("1.2.3")
	assert.True(t, validVersion.IsValid())

	// 无效的版本（没有数字部分）
	invalidVersion := NewVersion("")
	assert.False(t, invalidVersion.IsValid())

	// "abc"应该被解析成无效版本，因为没有数字部分
	weirdVersion := NewVersion("abc")
	t.Logf("'abc'解析结果: %v, 是否有效: %v", weirdVersion.VersionNumbers, weirdVersion.IsValid())
	// 确保已经正确解析为无效版本
	if !weirdVersion.IsValid() {
		t.Log("纯字母版本'abc'被正确识别为无效版本")
	} else {
		t.Error("纯字母版本'abc'应该是无效的，但结果显示它是有效的")
	}

	// 有前缀的有效版本
	validWithPrefix := NewVersion("v1.2.3")
	assert.True(t, validWithPrefix.IsValid())

	// 有后缀的有效版本
	validWithSuffix := NewVersion("1.2.3-beta")
	assert.True(t, validWithSuffix.IsValid())
}

// TestVersion_String 测试版本的字符串表示方法
func TestVersion_String(t *testing.T) {
	// 创建一个完整的版本对象
	v := &Version{
		Raw:            "v1.2.3-beta",
		PublicTime:     time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		VersionNumbers: []int{1, 2, 3},
		Prefix:         "v",
		Suffix:         "-beta",
	}

	// 获取字符串表示并解析为JSON
	str := v.String()
	assert.NotEmpty(t, str)

	// 验证可以被解析回来
	var parsedVersion map[string]interface{}
	err := json.Unmarshal([]byte(str), &parsedVersion)
	assert.Nil(t, err)
	assert.Equal(t, "v1.2.3-beta", parsedVersion["raw"])
}

// TestVersion_CompareTo 全面测试版本比较方法
func TestVersion_CompareTo(t *testing.T) {
	// 基本比较 - 版本号数字部分不同
	v1 := NewVersion("1.2.3")
	v2 := NewVersion("1.3.0")
	assert.Equal(t, -1, v1.CompareTo(v2))
	assert.Equal(t, 1, v2.CompareTo(v1))
	assert.Equal(t, 0, v1.CompareTo(v1)) // 自己和自己比较

	// 发布时间不同，版本号相同
	v3 := &Version{
		Raw:            "1.0.0",
		PublicTime:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		VersionNumbers: []int{1, 0, 0},
	}
	v4 := &Version{
		Raw:            "1.0.0",
		PublicTime:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		VersionNumbers: []int{1, 0, 0},
	}
	assert.Equal(t, -1, v3.CompareTo(v4))
	assert.Equal(t, 1, v4.CompareTo(v3))

	// 后缀不同，版本号和发布时间相同
	v5 := &Version{
		Raw:            "1.0.0-alpha",
		PublicTime:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		VersionNumbers: []int{1, 0, 0},
		Suffix:         "-alpha",
	}
	v6 := &Version{
		Raw:            "1.0.0-beta",
		PublicTime:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		VersionNumbers: []int{1, 0, 0},
		Suffix:         "-beta",
	}
	assert.Equal(t, -1, v5.CompareTo(v6))
	assert.Equal(t, 1, v6.CompareTo(v5))

	// 原始字符串不同，其他都相同（极少见的情况）
	v7 := &Version{
		Raw:            "1.0.0",
		VersionNumbers: []int{1, 0, 0},
	}
	v8 := &Version{
		Raw:            "01.00.00", // 格式不同但语义相同
		VersionNumbers: []int{1, 0, 0},
	}
	assert.NotEqual(t, 0, v7.CompareTo(v8)) // 应该是不相等的

	// 空版本号
	empty1 := NewVersion("")
	empty2 := NewVersion("")
	assert.Equal(t, 0, empty1.CompareTo(empty2))
}

func TestVersion_MarshalText(t *testing.T) {
	v := NewVersion("v1.2.3-beta1")
	text, err := v.MarshalText()
	if err != nil {
		t.Fatalf("MarshalText error: %v", err)
	}
	if string(text) != "v1.2.3-beta1" {
		t.Errorf("MarshalText() = %q, want %q", string(text), "v1.2.3-beta1")
	}
}

func TestVersion_UnmarshalText(t *testing.T) {
	var v Version
	err := v.UnmarshalText([]byte("1.2.3"))
	if err != nil {
		t.Fatalf("UnmarshalText error: %v", err)
	}
	if v.Major() != 1 || v.Minor() != 2 || v.Patch() != 3 {
		t.Errorf("UnmarshalText result = %v, want [1,2,3]", v.VersionNumbers)
	}
}

func TestVersion_UnmarshalText_Invalid(t *testing.T) {
	var v Version
	err := v.UnmarshalText([]byte("not-a-version"))
	if err == nil {
		t.Error("UnmarshalText should return error for invalid version")
	}
}

func TestVersion_Metadata(t *testing.T) {
	v := NewVersion("1.0.0+build123")
	if v.Metadata != "build123" {
		t.Errorf("Metadata = %q, want %q", v.Metadata, "build123")
	}
	if !v.IsStable() {
		t.Error("1.0.0+build123 should be stable (metadata doesn't affect stability)")
	}
}

func TestVersion_Metadata_Empty(t *testing.T) {
	v := NewVersion("1.0.0")
	if v.Metadata != "" {
		t.Errorf("Metadata = %q, want empty", v.Metadata)
	}
}

func TestVersion_Metadata_WithPrerelease(t *testing.T) {
	v := NewVersion("1.0.0-beta+build456")
	if v.Metadata != "build456" {
		t.Errorf("Metadata = %q, want %q", v.Metadata, "build456")
	}
	if !v.IsPrerelease() {
		t.Error("1.0.0-beta+build456 should be prerelease")
	}
}

func TestVersion_Metadata_PreservedInClone(t *testing.T) {
	v := NewVersion("1.0.0+build123")
	cloned := v.Clone()
	if cloned.Metadata != "build123" {
		t.Errorf("Clone() Metadata = %q, want %q", cloned.Metadata, "build123")
	}
}
