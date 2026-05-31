package versions

import "testing"

func TestVersion_IsPrerelease(t *testing.T) {
	if NewVersion("1.0.0").IsPrerelease() {
		t.Error("1.0.0 should not be prerelease")
	}
	if !NewVersion("1.0.0-beta").IsPrerelease() {
		t.Error("1.0.0-beta should be prerelease")
	}
}

func TestVersion_IsStable(t *testing.T) {
	if !NewVersion("1.0.0").IsStable() {
		t.Error("1.0.0 should be stable")
	}
	if NewVersion("1.0.0-rc1").IsStable() {
		t.Error("1.0.0-rc1 should not be stable")
	}
}

func TestVersion_IsBeta(t *testing.T) {
	if !NewVersion("1.0.0-beta").IsBeta() {
		t.Error("1.0.0-beta should be beta")
	}
	if NewVersion("1.0.0-alpha").IsBeta() {
		t.Error("1.0.0-alpha should not be beta")
	}
}

func TestVersion_IsRC(t *testing.T) {
	if !NewVersion("1.0.0-rc1").IsRC() {
		t.Error("1.0.0-rc1 should be RC")
	}
	if !NewVersion("1.0.0-CR1").IsRC() {
		t.Error("1.0.0-CR1 should be RC (CR variant)")
	}
}

func TestVersion_IsNewerThan(t *testing.T) {
	v1 := NewVersion("1.1.0")
	v2 := NewVersion("1.0.0")
	if !v1.IsNewerThan(v2) {
		t.Error("1.1.0 should be newer than 1.0.0")
	}
	if v2.IsNewerThan(v1) {
		t.Error("1.0.0 should not be newer than 1.1.0")
	}
}

func TestVersion_IsOlderThan(t *testing.T) {
	v1 := NewVersion("1.0.0")
	v2 := NewVersion("1.1.0")
	if !v1.IsOlderThan(v2) {
		t.Error("1.0.0 should be older than 1.1.0")
	}
}

func TestVersion_Equals(t *testing.T) {
	v1 := NewVersion("1.0.0")
	v2 := NewVersion("1.0.0")
	if !v1.Equals(v2) {
		t.Error("1.0.0 should equal 1.0.0")
	}
}

func TestVersion_IsBetween(t *testing.T) {
	v := NewVersion("1.5.0")
	low := NewVersion("1.0.0")
	high := NewVersion("2.0.0")
	if !v.IsBetween(low, high) {
		t.Error("1.5.0 should be between 1.0.0 and 2.0.0")
	}
	if NewVersion("0.9.0").IsBetween(low, high) {
		t.Error("0.9.0 should not be between 1.0.0 and 2.0.0")
	}
	if NewVersion("2.1.0").IsBetween(low, high) {
		t.Error("2.1.0 should not be between 1.0.0 and 2.0.0")
	}
}

func TestVersion_MajorMinorPatch(t *testing.T) {
	v := NewVersion("1.2.3")
	if v.Major() != 1 {
		t.Errorf("Major() = %d, want 1", v.Major())
	}
	if v.Minor() != 2 {
		t.Errorf("Minor() = %d, want 2", v.Minor())
	}
	if v.Patch() != 3 {
		t.Errorf("Patch() = %d, want 3", v.Patch())
	}
}

func TestVersion_MajorMinorPatch_Zero(t *testing.T) {
	v := NewVersion("5")
	if v.Major() != 5 {
		t.Errorf("Major() = %d, want 5", v.Major())
	}
	if v.Minor() != 0 {
		t.Errorf("Minor() = %d, want 0", v.Minor())
	}
	if v.Patch() != 0 {
		t.Errorf("Patch() = %d, want 0", v.Patch())
	}
}

func TestVersion_IsMilestone(t *testing.T) {
	if !NewVersion("1.0.0-milestone1").IsMilestone() {
		t.Error("1.0.0-milestone1 should be milestone")
	}
	if !NewVersion("1.0.0-m1").IsMilestone() {
		t.Error("1.0.0-m1 should be milestone (short form)")
	}
	if NewVersion("1.0.0-beta").IsMilestone() {
		t.Error("1.0.0-beta should not be milestone")
	}
}

func TestVersion_IsNightly(t *testing.T) {
	if !NewVersion("1.0.0-nightly").IsNightly() {
		t.Error("1.0.0-nightly should be nightly")
	}
	if NewVersion("1.0.0-beta").IsNightly() {
		t.Error("1.0.0-beta should not be nightly")
	}
}

func TestVersion_IsFinal(t *testing.T) {
	if !NewVersion("1.0.0-final").IsFinal() {
		t.Error("1.0.0-final should be final")
	}
	if NewVersion("1.0.0").IsFinal() {
		t.Error("1.0.0 should not be final (no suffix)")
	}
}

func TestVersion_IsGA(t *testing.T) {
	if !NewVersion("1.0.0-ga").IsGA() {
		t.Error("1.0.0-ga should be GA")
	}
	if NewVersion("1.0.0").IsGA() {
		t.Error("1.0.0 should not be GA (no suffix)")
	}
}
