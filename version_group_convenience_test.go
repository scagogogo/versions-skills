package versions

import "testing"

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
