package versions

import "testing"

func TestVersion_Clone(t *testing.T) {
	v1 := NewVersion("1.2.3-beta")
	v2 := v1.Clone()
	if v1.Raw != v2.Raw {
		t.Errorf("Clone() Raw = %q, want %q", v2.Raw, v1.Raw)
	}
	// 修改克隆对象不影响原始对象
	v2.Raw = "modified"
	if v1.Raw == "modified" {
		t.Error("Clone() should not affect original")
	}
}

func TestVersion_Clone_Nil(t *testing.T) {
	var v *Version = nil
	if v.Clone() != nil {
		t.Error("Clone(nil) should return nil")
	}
}

func TestVersion_WithPrefix(t *testing.T) {
	v := NewVersion("1.2.3")
	newV := v.WithPrefix("v")
	if newV.Prefix != "v" {
		t.Errorf("WithPrefix() Prefix = %q, want %q", newV.Prefix, "v")
	}
	if v.Prefix == "v" {
		t.Error("WithPrefix() should not modify original")
	}
}

func TestVersion_WithSuffix(t *testing.T) {
	v := NewVersion("1.2.3")
	newV := v.WithSuffix("-rc1")
	if !newV.IsRC() {
		t.Errorf("WithSuffix() should produce RC version, got %s", newV.Suffix)
	}
	if v.IsPrerelease() {
		t.Error("WithSuffix() should not modify original")
	}
}

func TestVersion_WithMajor(t *testing.T) {
	v := NewVersion("1.2.3")
	newV := v.WithMajor(5)
	if newV.Major() != 5 {
		t.Errorf("WithMajor() Major = %d, want 5", newV.Major())
	}
	if newV.Minor() != 2 {
		t.Errorf("WithMajor() Minor = %d, want 2 (preserved)", newV.Minor())
	}
}

func TestVersion_WithMinor(t *testing.T) {
	v := NewVersion("1.2.3")
	newV := v.WithMinor(9)
	if newV.Minor() != 9 {
		t.Errorf("WithMinor() Minor = %d, want 9", newV.Minor())
	}
	if newV.Major() != 1 {
		t.Errorf("WithMinor() Major = %d, want 1 (preserved)", newV.Major())
	}
}

func TestVersion_WithPatch(t *testing.T) {
	v := NewVersion("1.2.3")
	newV := v.WithPatch(7)
	if newV.Patch() != 7 {
		t.Errorf("WithPatch() Patch = %d, want 7", newV.Patch())
	}
	if newV.Major() != 1 || newV.Minor() != 2 {
		t.Errorf("WithPatch() should preserve Major/Minor, got %d.%d", newV.Major(), newV.Minor())
	}
}
