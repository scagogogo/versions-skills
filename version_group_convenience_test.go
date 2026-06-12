package versions

import (
	"strings"
	"testing"
)

func TestVersionGroup_GetLatest(t *testing.T) {
	vg := NewVersionGroupFromVersions(NewVersions("1.0.0", "1.0.1", "1.0.2"))
	latest := vg.GetLatest()
	if latest == nil || latest.Raw != "1.0.2" {
		t.Errorf("GetLatest() = %v, want 1.0.2", latest)
	}
}

func TestVersionGroup_GetOldest(t *testing.T) {
	vg := NewVersionGroupFromVersions(NewVersions("1.0.0", "1.0.1", "1.0.2"))
	oldest := vg.GetOldest()
	if oldest == nil || oldest.Raw != "1.0.0" {
		t.Errorf("GetOldest() = %v, want 1.0.0", oldest)
	}
}

func TestVersionGroup_Count(t *testing.T) {
	vg := NewVersionGroupFromVersions(NewVersions("1.0.0", "1.0.1"))
	if vg.Count() != 2 {
		t.Errorf("Count() = %d, want 2", vg.Count())
	}
}

func TestVersionGroup_StableVersions(t *testing.T) {
	vg := NewVersionGroupFromVersions(NewVersions("1.0.0", "1.0.0-beta", "1.0.1"))
	stable := vg.StableVersions()
	if len(stable) != 2 {
		t.Errorf("StableVersions() = %d, want 2", len(stable))
	}
}

func TestVersionGroup_LatestStable(t *testing.T) {
	vg := NewVersionGroupFromVersions(NewVersions("1.0.0-beta", "1.0.0", "1.0.1-alpha"))
	latest := vg.LatestStable()
	if latest == nil || latest.Raw != "1.0.0" {
		t.Errorf("LatestStable() = %v, want 1.0.0", latest)
	}
}

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
	// Verify it's a prerelease version
	if !latest.IsPrerelease() {
		t.Error("LatestPrerelease() should return a prerelease version")
	}
}

func TestVersionGroup_LatestPrerelease_None(t *testing.T) {
	vg := NewVersionGroupFromVersions(NewVersions("1.0.0", "1.0.1"))
	if vg.LatestPrerelease() != nil {
		t.Error("LatestPrerelease() should return nil when no prerelease versions")
	}
}

func TestVersionGroup_String(t *testing.T) {
	vg := NewVersionGroupFromVersions(NewVersions("1.0.0", "1.0.1"))
	s := vg.String()
	if !strings.Contains(s, "1.0") || !strings.Contains(s, "2") {
		t.Errorf("String() = %q, should contain group ID and count", s)
	}
}

func TestVersionGroup_Filter(t *testing.T) {
	vg := NewVersionGroupFromVersions(NewVersions("1.0.0", "1.0.0-beta", "1.0.1"))
	stable := vg.Filter(func(v *Version) bool { return v.IsStable() })
	if len(stable) != 2 {
		t.Errorf("Filter(stable) = %d, want 2", len(stable))
	}
}
