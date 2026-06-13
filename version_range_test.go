package versions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionRange_Contains(t *testing.T) {
	low := NewVersion("1.0.0")
	high := NewVersion("2.0.0")

	// 闭区间 [1.0.0, 2.0.0]
	closedRange := NewClosedRange(low, high)
	assert.True(t, closedRange.Contains(NewVersion("1.0.0")))
	assert.True(t, closedRange.Contains(NewVersion("1.5.0")))
	assert.True(t, closedRange.Contains(NewVersion("2.0.0")))
	assert.False(t, closedRange.Contains(NewVersion("0.9.0")))
	assert.False(t, closedRange.Contains(NewVersion("2.1.0")))

	// 开区间 (1.0.0, 2.0.0)
	openRange := NewOpenRange(low, high)
	assert.False(t, openRange.Contains(NewVersion("1.0.0")))
	assert.True(t, openRange.Contains(NewVersion("1.5.0")))
	assert.False(t, openRange.Contains(NewVersion("2.0.0")))

	// 半开区间 [1.0.0, 2.0.0)
	halfOpen := NewVersionRange(low, high, true, false)
	assert.True(t, halfOpen.Contains(NewVersion("1.0.0")))
	assert.True(t, halfOpen.Contains(NewVersion("1.5.0")))
	assert.False(t, halfOpen.Contains(NewVersion("2.0.0")))
}

func TestVersionRange_String(t *testing.T) {
	low := NewVersion("1.0.0")
	high := NewVersion("2.0.0")

	assert.Equal(t, "[1.0.0, 2.0.0]", NewClosedRange(low, high).String())
	assert.Equal(t, "(1.0.0, 2.0.0)", NewOpenRange(low, high).String())
	assert.Equal(t, "[1.0.0, 2.0.0)", NewVersionRange(low, high, true, false).String())
}

func TestVersionRange_Filter(t *testing.T) {
	low := NewVersion("1.0.0")
	high := NewVersion("2.0.0")
	r := NewClosedRange(low, high)

	versions := NewVersions("0.9.0", "1.0.0", "1.5.0", "2.0.0", "2.1.0")
	filtered := r.Filter(versions)
	assert.Equal(t, 3, len(filtered))
	assert.Equal(t, "1.0.0", filtered[0].Raw)
	assert.Equal(t, "1.5.0", filtered[1].Raw)
	assert.Equal(t, "2.0.0", filtered[2].Raw)
}

func TestVersionRange_IsEmpty(t *testing.T) {
	low := NewVersion("1.0.0")
	high := NewVersion("2.0.0")

	assert.False(t, NewClosedRange(low, high).IsEmpty())
	assert.True(t, NewOpenRange(low, low).IsEmpty())    // (1.0.0, 1.0.0) is empty
	assert.False(t, NewClosedRange(low, low).IsEmpty()) // [1.0.0, 1.0.0] is not empty

	reversedLow := NewVersion("2.0.0")
	reversedHigh := NewVersion("1.0.0")
	assert.True(t, NewClosedRange(reversedLow, reversedHigh).IsEmpty())
}

func TestIsSemver(t *testing.T) {
	// Valid semver
	assert.True(t, NewVersion("1.2.3").IsSemver())
	assert.True(t, NewVersion("v1.2.3").IsSemver())
	assert.True(t, NewVersion("1.0.0-alpha").IsSemver())
	assert.True(t, NewVersion("1.0.0-alpha.1").IsSemver())
	assert.True(t, NewVersion("1.0.0-alpha.beta").IsSemver())
	assert.True(t, NewVersion("1.0.0+build.123").IsSemver())
	assert.True(t, NewVersion("1.0.0-alpha+build.123").IsSemver())
	assert.True(t, NewVersion("0.0.0").IsSemver())

	// Invalid semver
	assert.False(t, NewVersion("1.2").IsSemver())    // only 2 segments
	assert.False(t, NewVersion("1").IsSemver())      // only 1 segment
	assert.False(t, NewVersion("01.2.3").IsSemver()) // leading zero
	assert.False(t, NewVersion("1.02.3").IsSemver()) // leading zero
	assert.False(t, NewVersion("1.2.03").IsSemver()) // leading zero
}

func TestValidateSemver(t *testing.T) {
	assert.Nil(t, NewVersion("1.2.3").ValidateSemver())
	assert.Nil(t, NewVersion("1.0.0-alpha.1").ValidateSemver())

	assert.NotNil(t, NewVersion("1.2").ValidateSemver())
	assert.NotNil(t, NewVersion("01.2.3").ValidateSemver())
	assert.NotNil(t, NewVersion("not-a-version").ValidateSemver())
}

func TestPreReleaseType(t *testing.T) {
	assert.Equal(t, "alpha", NewVersion("1.0.0-alpha1").PreReleaseType())
	assert.Equal(t, "beta", NewVersion("1.0.0-beta2").PreReleaseType())
	assert.Equal(t, "rc", NewVersion("1.0.0-rc1").PreReleaseType())
	assert.Equal(t, "dev", NewVersion("1.0.0-dev1").PreReleaseType())
	assert.Equal(t, "snapshot", NewVersion("1.0.0-SNAPSHOT").PreReleaseType())
	assert.Equal(t, "", NewVersion("1.0.0").PreReleaseType())
}

func TestDiff(t *testing.T) {
	v1 := NewVersion("1.2.3")
	v2 := NewVersion("2.0.0")
	d := v1.Diff(v2)
	assert.Equal(t, 1, d.Major)
	assert.Equal(t, -2, d.Minor)
	assert.Equal(t, -3, d.Patch)
	assert.Equal(t, "1.2.3", d.RawFrom)
	assert.Equal(t, "2.0.0", d.RawTo)
	assert.True(t, d.IsUpgrade())
	assert.True(t, d.IsMajorChange())
	assert.False(t, d.IsMinorChange())
	assert.False(t, d.IsPatchChange())

	// Patch-only diff
	v3 := NewVersion("1.2.3")
	v4 := NewVersion("1.2.5")
	d2 := v3.Diff(v4)
	assert.False(t, d2.IsMajorChange())
	assert.False(t, d2.IsMinorChange())
	assert.True(t, d2.IsPatchChange())
	assert.True(t, d2.IsUpgrade())

	// Downgrade
	v5 := NewVersion("2.0.0")
	v6 := NewVersion("1.0.0")
	d3 := v5.Diff(v6)
	assert.True(t, d3.IsDowngrade())

	// nil target
	assert.Nil(t, v1.Diff(nil))
}

func TestDiff_String(t *testing.T) {
	v1 := NewVersion("1.2.3")
	v2 := NewVersion("2.0.0")
	d := v1.Diff(v2)
	assert.Contains(t, d.String(), "1.2.3")
	assert.Contains(t, d.String(), "2.0.0")
}

func TestCoerce(t *testing.T) {
	v1 := Coerce("program-1.2.3-linux-amd64")
	assert.True(t, v1.IsValid())

	v2 := Coerce("download/v2.0.0-beta.tar.gz")
	assert.True(t, v2.IsValid())

	v3 := Coerce("no-version-here")
	assert.False(t, v3.IsValid())

	_, err := CoerceE("no-version-here")
	assert.NotNil(t, err)

	_, err2 := CoerceE("app-1.2.3")
	assert.Nil(t, err2)
}

func TestWithMetadata(t *testing.T) {
	v := NewVersion("1.2.3")
	newV := v.WithMetadata("build.123")
	assert.Equal(t, "build.123", newV.Metadata)
	assert.Equal(t, "", v.Metadata) // original unchanged
}

func TestCanonical(t *testing.T) {
	assert.Equal(t, "1.2.0", NewVersion("1.2").Canonical())
	assert.Equal(t, "1.0.0", NewVersion("1").Canonical())
	assert.Equal(t, "1.2.3", NewVersion("1.2.3").Canonical())
	assert.Equal(t, "v1.2.3-beta", NewVersion("v1.2.3-beta").Canonical())
}

func TestFormat(t *testing.T) {
	v := NewVersion("v1.2.3-beta")
	assert.Equal(t, "1.2.3", v.Format("%M.%m.%p"))
	assert.Equal(t, "v", v.Format("%P"))
	assert.Equal(t, "major=1 minor=2 patch=3", v.Format("major=%M minor=%m patch=%p"))
	assert.Equal(t, "100%", v.Format("100%%"))
}

func TestIncrement(t *testing.T) {
	v := NewVersion("1.2.3.4")
	assert.Equal(t, "2.0.0.0", v.Increment(0).Raw)
	assert.Equal(t, "1.3.0.0", v.Increment(1).Raw)
	assert.Equal(t, "1.2.4.0", v.Increment(2).Raw)
	assert.Equal(t, "1.2.3.5", v.Increment(3).Raw)

	// Negative segment returns clone
	assert.Equal(t, v.Raw, v.Increment(-1).Raw)
}

func TestContainsVersion(t *testing.T) {
	list := NewVersions("1.0.0", "1.1.0", "2.0.0")
	assert.True(t, ContainsVersion(list, NewVersion("1.1.0")))
	assert.False(t, ContainsVersion(list, NewVersion("3.0.0")))
}

func TestIndexOf(t *testing.T) {
	list := NewVersions("1.0.0", "1.1.0", "2.0.0")
	assert.Equal(t, 0, IndexOf(list, NewVersion("1.0.0")))
	assert.Equal(t, 1, IndexOf(list, NewVersion("1.1.0")))
	assert.Equal(t, 2, IndexOf(list, NewVersion("2.0.0")))
	assert.Equal(t, -1, IndexOf(list, NewVersion("3.0.0")))
}

func TestGroupByMajor(t *testing.T) {
	list := NewVersions("1.0.0", "1.1.0", "2.0.0", "2.1.0")
	groups := GroupByMajor(list)
	assert.Equal(t, 2, len(groups[1]))
	assert.Equal(t, 2, len(groups[2]))
}

func TestGroupByMinor(t *testing.T) {
	list := NewVersions("1.0.0", "1.0.1", "1.1.0", "2.0.0")
	groups := GroupByMinor(list)
	assert.Equal(t, 2, len(groups["1.0"]))
	assert.Equal(t, 1, len(groups["1.1"]))
	assert.Equal(t, 1, len(groups["2.0"]))
}

func TestNegateConstraint(t *testing.T) {
	c1, _ := ParseConstraint(">=1.0.0")
	n1 := NegateConstraint(c1)
	assert.Equal(t, ConstraintLessThan, n1.Operator)

	c2, _ := ParseConstraint("=1.0.0")
	n2 := NegateConstraint(c2)
	assert.Equal(t, ConstraintNotEqual, n2.Operator)

	c3, _ := ParseConstraint("<2.0.0")
	n3 := NegateConstraint(c3)
	assert.Equal(t, ConstraintGreaterThanOrEqual, n3.Operator)
}

func TestHash(t *testing.T) {
	v := NewVersion("1.2.3")
	assert.Equal(t, "1.2.3", v.Hash())
}
