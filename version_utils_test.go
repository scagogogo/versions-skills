package versions

import "testing"

func TestMin(t *testing.T) {
	versions := NewVersions("2.0.0", "1.0.0", "1.5.0")
	min := Min(versions)
	if min.Raw != "1.0.0" {
		t.Errorf("Min() = %s, want 1.0.0", min.Raw)
	}
}

func TestMin_Empty(t *testing.T) {
	if Min(nil) != nil {
		t.Error("Min(nil) should return nil")
	}
}

func TestMax(t *testing.T) {
	versions := NewVersions("1.0.0", "3.0.0", "2.0.0")
	max := Max(versions)
	if max.Raw != "3.0.0" {
		t.Errorf("Max() = %s, want 3.0.0", max.Raw)
	}
}

func TestLatestStable(t *testing.T) {
	versions := NewVersions("1.0.0-alpha", "1.0.0", "1.1.0-beta", "1.1.0")
	latest := LatestStable(versions)
	if latest.Raw != "1.1.0" {
		t.Errorf("LatestStable() = %s, want 1.1.0", latest.Raw)
	}
}

func TestLatestStable_None(t *testing.T) {
	versions := NewVersions("1.0.0-alpha", "1.0.0-beta")
	if LatestStable(versions) != nil {
		t.Error("LatestStable() should return nil when no stable versions")
	}
}

func TestLatestPrerelease(t *testing.T) {
	versions := NewVersions("1.0.0", "1.1.0-alpha", "1.0.0-beta")
	latest := LatestPrerelease(versions)
	if latest.Raw != "1.1.0-alpha" {
		t.Errorf("LatestPrerelease() = %s, want 1.1.0-alpha", latest.Raw)
	}
}

func TestFilter(t *testing.T) {
	versions := NewVersions("1.0.0", "1.0.0-beta", "2.0.0-alpha", "2.0.0")
	stable := Filter(versions, func(v *Version) bool { return v.IsStable() })
	if len(stable) != 2 {
		t.Errorf("Filter stable = %d, want 2", len(stable))
	}
}

func TestFilterByConstraint(t *testing.T) {
	versions := NewVersions("0.9.0", "1.0.0", "1.5.0", "2.0.0")
	c, _ := ParseConstraint(">=1.0.0")
	result := FilterByConstraint(versions, c)
	if len(result) != 3 {
		t.Errorf("FilterByConstraint >=1.0.0 = %d, want 3", len(result))
	}
}

func TestFilterByConstraintSet(t *testing.T) {
	versions := NewVersions("0.9.0", "1.0.0", "1.5.0", "2.0.0")
	cs, _ := ParseConstraintSet(">=1.0.0,<2.0.0")
	result := FilterByConstraintSet(versions, cs)
	if len(result) != 2 {
		t.Errorf("FilterByConstraintSet >=1.0.0,<2.0.0 = %d, want 2", len(result))
	}
}

func TestUnique(t *testing.T) {
	versions := NewVersions("1.0.0", "1.0.0", "2.0.0", "2.0.0")
	result := Unique(versions)
	if len(result) != 2 {
		t.Errorf("Unique() = %d, want 2", len(result))
	}
}

func TestFilterByMajor(t *testing.T) {
	versions := NewVersions("1.0.0", "1.1.0", "2.0.0", "2.1.0")
	result := FilterByMajor(versions, 2)
	if len(result) != 2 {
		t.Errorf("FilterByMajor(2) = %d, want 2", len(result))
	}
}

func TestCount(t *testing.T) {
	versions := NewVersions("1.0.0", "1.0.0-beta", "2.0.0-alpha")
	n := Count(versions, func(v *Version) bool { return v.IsPrerelease() })
	if n != 2 {
		t.Errorf("Count prerelease = %d, want 2", n)
	}
}
