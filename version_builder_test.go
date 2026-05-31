package versions

import "testing"

func TestVersionBuilder_FullBuild(t *testing.T) {
	v := NewVersionBuilder().
		Prefix("v").
		Major(1).
		Minor(2).
		Patch(3).
		Suffix("-beta1").
		Build()
	if !v.IsValid() {
		t.Fatal("builder result should be valid")
	}
	if v.Major() != 1 || v.Minor() != 2 || v.Patch() != 3 {
		t.Errorf("builder numbers = %v, want [1,2,3]", v.VersionNumbers)
	}
}

func TestVersionBuilder_Minimal(t *testing.T) {
	v := NewVersionBuilder().Major(5).Build()
	if v.Major() != 5 {
		t.Errorf("Major() = %d, want 5", v.Major())
	}
}

func TestVersionBuilder_Numbers(t *testing.T) {
	v := NewVersionBuilder().Numbers([]int{2, 0, 1}).Build()
	if v.Major() != 2 || v.Minor() != 0 || v.Patch() != 1 {
		t.Errorf("numbers = %v, want [2,0,1]", v.VersionNumbers)
	}
}

func TestVersionBuilder_DoubleDigitNumbers(t *testing.T) {
	v := NewVersionBuilder().Major(10).Minor(20).Patch(30).Build()
	if v.Major() != 10 || v.Minor() != 20 || v.Patch() != 30 {
		t.Errorf("numbers = %v, want [10,20,30]", v.VersionNumbers)
	}
}

func TestVersion_BumpMajor(t *testing.T) {
	v := NewVersion("1.2.3")
	bumped := v.BumpMajor()
	if bumped.Major() != 2 {
		t.Errorf("BumpMajor() Major = %d, want 2", bumped.Major())
	}
	if bumped.Minor() != 0 {
		t.Errorf("BumpMajor() Minor = %d, want 0", bumped.Minor())
	}
	if bumped.Patch() != 0 {
		t.Errorf("BumpMajor() Patch = %d, want 0", bumped.Patch())
	}
}

func TestVersion_BumpMinor(t *testing.T) {
	v := NewVersion("1.2.3")
	bumped := v.BumpMinor()
	if bumped.Major() != 1 {
		t.Errorf("BumpMinor() Major = %d, want 1", bumped.Major())
	}
	if bumped.Minor() != 3 {
		t.Errorf("BumpMinor() Minor = %d, want 3", bumped.Minor())
	}
	if bumped.Patch() != 0 {
		t.Errorf("BumpMinor() Patch = %d, want 0", bumped.Patch())
	}
}

func TestVersion_BumpPatch(t *testing.T) {
	v := NewVersion("1.2.3")
	bumped := v.BumpPatch()
	if bumped.Patch() != 4 {
		t.Errorf("BumpPatch() Patch = %d, want 4", bumped.Patch())
	}
}

func TestVersion_BumpPatch_ClearsSuffix(t *testing.T) {
	v := NewVersion("1.2.3-beta")
	bumped := v.BumpPatch()
	if bumped.IsPrerelease() {
		t.Error("BumpPatch() should clear suffix")
	}
}
