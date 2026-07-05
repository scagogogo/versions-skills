package versions

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionSlice_Sort(t *testing.T) {
	slice := VersionSlice(NewVersions("2.0.0", "1.0.0", "1.5.0", "1.10.0"))
	sort.Sort(slice)
	if slice[0].Raw != "1.0.0" {
		t.Errorf("After sort, first = %s, want 1.0.0", slice[0].Raw)
	}
	if slice[len(slice)-1].Raw != "2.0.0" {
		t.Errorf("After sort, last = %s, want 2.0.0", slice[len(slice)-1].Raw)
	}
}

func TestVersionSlice_Len(t *testing.T) {
	assert.Equal(t, 0, VersionSlice{}.Len())
	assert.Equal(t, 3, VersionSlice(NewVersions("1.0.0", "2.0.0", "3.0.0")).Len())
}

func TestVersionSlice_Less(t *testing.T) {
	s := VersionSlice(NewVersions("1.0.0", "2.0.0"))
	assert.True(t, s.Less(0, 1))
	assert.False(t, s.Less(1, 0))
}

func TestVersionSlice_Swap(t *testing.T) {
	s := VersionSlice(NewVersions("1.0.0", "2.0.0"))
	s.Swap(0, 1)
	assert.Equal(t, "2.0.0", s[0].Raw)
	assert.Equal(t, "1.0.0", s[1].Raw)
}

func TestVersionSlice_Min(t *testing.T) {
	assert.Nil(t, VersionSlice{}.Min())
	s := VersionSlice(NewVersions("2.0.0", "1.0.0", "1.5.0"))
	assert.Equal(t, "1.0.0", s.Min().Raw)
}

func TestVersionSlice_Max(t *testing.T) {
	assert.Nil(t, VersionSlice{}.Max())
	s := VersionSlice(NewVersions("1.0.0", "2.0.0", "1.5.0"))
	assert.Equal(t, "2.0.0", s.Max().Raw)
}

func TestVersionSlice_Filter(t *testing.T) {
	s := VersionSlice(NewVersions("1.0.0", "2.0.0", "1.5.0", "3.0.0"))
	got := s.Filter(func(v *Version) bool { return v.Major() >= 2 })
	assert.Equal(t, []string{"2.0.0", "3.0.0"}, raws(got))
	// 空切片
	assert.Equal(t, 0, len(VersionSlice{}.Filter(func(v *Version) bool { return true })))
}

func TestVersionSlice_Contains(t *testing.T) {
	s := VersionSlice(NewVersions("1.0.0", "2.0.0"))
	assert.True(t, s.Contains(NewVersion("1.0.0")))
	assert.False(t, s.Contains(NewVersion("3.0.0")))
	assert.False(t, VersionSlice{}.Contains(NewVersion("1.0.0")))
}

func TestVersionSlice_IndexOf(t *testing.T) {
	s := VersionSlice(NewVersions("1.0.0", "2.0.0", "3.0.0"))
	assert.Equal(t, 0, s.IndexOf(NewVersion("1.0.0")))
	assert.Equal(t, 2, s.IndexOf(NewVersion("3.0.0")))
	assert.Equal(t, -1, s.IndexOf(NewVersion("9.9.9")))
	assert.Equal(t, -1, VersionSlice{}.IndexOf(NewVersion("1.0.0")))
}

func TestVersionSlice_Unique(t *testing.T) {
	s := VersionSlice(NewVersions("1.0.0", "2.0.0", "1.0.0", "3.0.0", "2.0.0"))
	got := s.Unique()
	assert.Equal(t, []string{"1.0.0", "2.0.0", "3.0.0"}, raws(got))
	assert.Equal(t, 0, len(VersionSlice{}.Unique()))
}

func TestVersionSlice_SortMethod(t *testing.T) {
	// 直接调用 Sort 方法（冒泡实现路径）
	s := VersionSlice(NewVersions("3.0.0", "1.0.0", "2.0.0"))
	s.Sort()
	assert.Equal(t, []string{"1.0.0", "2.0.0", "3.0.0"}, raws(s))
	// 空 / 单元素
	VersionSlice{}.Sort()
	VersionSlice(NewVersions("1.0.0")).Sort()
}

func TestVersionSlice_Sorted(t *testing.T) {
	s := VersionSlice(NewVersions("3.0.0", "1.0.0", "2.0.0"))
	got := s.Sorted()
	assert.Equal(t, []string{"1.0.0", "2.0.0", "3.0.0"}, raws(got))
	// 原切片不变
	assert.Equal(t, "3.0.0", s[0].Raw)
	// 空切片
	assert.Equal(t, 0, len(VersionSlice{}.Sorted()))
}

// raws 提取 VersionSlice 的 Raw 字段切片，便于断言
func raws(s VersionSlice) []string {
	out := make([]string, 0, len(s))
	for _, v := range s {
		out = append(out, v.Raw)
	}
	return out
}
