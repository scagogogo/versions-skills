package versions

import (
	"testing"
	"time"
)

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

func TestVersion_IsPre(t *testing.T) {
	if !NewVersion("1.0.0-pre1").IsPre() {
		t.Error("1.0.0-pre1 should be pre")
	}
	if NewVersion("1.0.0-beta").IsPre() {
		t.Error("1.0.0-beta should not be pre")
	}
}

func TestVersion_IsRelease(t *testing.T) {
	if !NewVersion("1.0.0-release").IsRelease() {
		t.Error("1.0.0-release should be release")
	}
	if NewVersion("1.0.0").IsRelease() {
		t.Error("1.0.0 should not be release (no suffix)")
	}
}

func TestVersion_IsSP(t *testing.T) {
	if !NewVersion("1.0.0-sp1").IsSP() {
		t.Error("1.0.0-sp1 should be SP")
	}
	if NewVersion("1.0.0").IsSP() {
		t.Error("1.0.0 should not be SP")
	}
}

func TestVersion_IsPost(t *testing.T) {
	if !NewVersion("1.0.0-post1").IsPost() {
		t.Error("1.0.0-post1 should be post")
	}
	if NewVersion("1.0.0").IsPost() {
		t.Error("1.0.0 should not be post")
	}
}

func TestVersion_Satisfies(t *testing.T) {
	c, _ := ParseConstraint(">=1.0.0")
	v := NewVersion("1.5.0")
	if !v.Satisfies(c) {
		t.Error("1.5.0 should satisfy >=1.0.0")
	}
	v2 := NewVersion("0.9.0")
	if v2.Satisfies(c) {
		t.Error("0.9.0 should not satisfy >=1.0.0")
	}
}

func TestVersion_Matches(t *testing.T) {
	v := NewVersion("1.5.0")
	ok, err := v.Matches(">=1.0.0,<2.0.0")
	if err != nil {
		t.Fatalf("Matches() error: %v", err)
	}
	if !ok {
		t.Error("1.5.0 should match >=1.0.0,<2.0.0")
	}

	ok2, err2 := v.Matches(">=2.0.0")
	if err2 != nil {
		t.Fatalf("Matches() error: %v", err2)
	}
	if ok2 {
		t.Error("1.5.0 should not match >=2.0.0")
	}

	_, err3 := v.Matches("not-valid")
	if err3 == nil {
		t.Error("Matches() should return error for invalid expression")
	}
}

func TestVersion_RawString(t *testing.T) {
	v := NewVersion("v1.2.3-beta1")
	if v.RawString() != "v1.2.3-beta1" {
		t.Errorf("RawString() = %q, want %q", v.RawString(), "v1.2.3-beta1")
	}
	// Verify it's different from String() which returns JSON
	if v.String() == v.RawString() {
		t.Error("String() and RawString() should differ — String returns JSON")
	}
}

func TestVersion_WithNumbers(t *testing.T) {
	v := NewVersion("1.2.3")
	newV := v.WithNumbers([]int{2, 0, 0})
	if newV.Major() != 2 || newV.Minor() != 0 || newV.Patch() != 0 {
		t.Errorf("WithNumbers() = %v, want [2,0,0]", newV.VersionNumbers)
	}
	// Original should not be modified
	if v.Major() != 1 {
		t.Error("WithNumbers() should not modify original")
	}
}

func TestVersion_SubVersion(t *testing.T) {
	v1 := NewVersion("1.0.0-beta2")
	if v1.SubVersion() != 2 {
		t.Errorf("SubVersion() = %d, want 2", v1.SubVersion())
	}
	v2 := NewVersion("1.0.0-beta")
	if v2.SubVersion() != 0 {
		t.Errorf("SubVersion() for no number = %d, want 0", v2.SubVersion())
	}
}

func TestVersion_SuffixWeight(t *testing.T) {
	v := NewVersion("1.0.0-beta")
	if v.SuffixWeight() != SuffixWeightBeta {
		t.Errorf("SuffixWeight() = %d, want %d (SuffixWeightBeta)", v.SuffixWeight(), SuffixWeightBeta)
	}
	v2 := NewVersion("1.0.0")
	if v2.SuffixWeight() != SuffixWeightUnknown {
		t.Errorf("SuffixWeight() for stable = %d, want %d (SuffixWeightUnknown)", v2.SuffixWeight(), SuffixWeightUnknown)
	}
}

func TestVersion_WithPublicTime(t *testing.T) {
	v := NewVersion("1.2.3")
	tt := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	newV := v.WithPublicTime(tt)
	if newV.PublicTime != tt {
		t.Errorf("WithPublicTime() time mismatch")
	}
	if !v.PublicTime.IsZero() {
		t.Error("WithPublicTime() should not modify original's time")
	}
}

func TestVersion_IsZero(t *testing.T) {
	var v Version
	if !v.IsZero() {
		t.Error("Zero value Version should be IsZero()")
	}
	v2 := NewVersion("1.0.0")
	if v2.IsZero() {
		t.Error("Parsed version should not be IsZero()")
	}
}

func TestVersion_Core(t *testing.T) {
	v := NewVersion("1.2.3-beta1")
	core := v.Core()
	if core.RawString() != "1.2.3" {
		t.Errorf("Core() = %q, want %q", core.RawString(), "1.2.3")
	}
	if !core.IsStable() {
		t.Error("Core() should be stable")
	}
	if v.IsStable() {
		t.Error("Core() should not modify original")
	}
}

func TestVersion_Validate(t *testing.T) {
	v := NewVersion("1.2.3")
	if err := v.Validate(); err != nil {
		t.Errorf("Validate() error: %v", err)
	}
	var empty Version
	if empty.Validate() == nil {
		t.Error("Validate() should return error for empty version")
	}
}

func TestVersion_Segments(t *testing.T) {
	v := NewVersion("1.2.3")
	segs := v.Segments()
	if len(segs) != 3 || segs[0] != 1 || segs[1] != 2 || segs[2] != 3 {
		t.Errorf("Segments() = %v, want [1,2,3]", segs)
	}
	segs[0] = 99
	if v.Major() != 1 {
		t.Error("Segments() should return a copy")
	}
}

func TestVersion_Segments64(t *testing.T) {
	v := NewVersion("1.2.3")
	segs := v.Segments64()
	if len(segs) != 3 || segs[0] != 1 || segs[1] != 2 || segs[2] != 3 {
		t.Errorf("Segments64() = %v, want [1,2,3]", segs)
	}
}

func TestMustParse(t *testing.T) {
	v := MustParse("1.2.3")
	if v.Major() != 1 {
		t.Errorf("MustParse() Major = %d, want 1", v.Major())
	}
}

func TestMustParse_Panic(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("MustParse() should panic for empty string")
		}
	}()
	MustParse("")
}
