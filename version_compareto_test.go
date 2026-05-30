package versions

import (
	"testing"
	"time"
)

func TestVersion_CompareTo_ReleaseVsPreRelease(t *testing.T) {
	release := NewVersion("1.0.0")
	beta := NewVersion("1.0.0-beta")
	alpha := NewVersion("1.0.0-alpha")

	if release.CompareTo(beta) <= 0 {
		t.Errorf("1.0.0 (release) should be greater than 1.0.0-beta, got %d", release.CompareTo(beta))
	}
	if release.CompareTo(alpha) <= 0 {
		t.Errorf("1.0.0 (release) should be greater than 1.0.0-alpha, got %d", release.CompareTo(alpha))
	}
	if beta.CompareTo(alpha) <= 0 {
		t.Errorf("1.0.0-beta should be greater than 1.0.0-alpha (alphabetic), got %d", beta.CompareTo(alpha))
	}
}

func TestVersion_CompareTo_DifferentLengths(t *testing.T) {
	// 注意：VersionNumbers.CompareTo 在长度不同时，短版本 < 长版本
	// 这是因为 [1,0] 是 [1,0] 而 [1,0,0] 是 [1,0,0]，后者更长
	short := NewVersion("1.0")
	full := NewVersion("1.0.0")
	result := short.CompareTo(full)
	// 语义上 1.0 和 1.0.0 可能等价，但当前实现按长度区分
	if result >= 0 {
		t.Errorf("1.0 should be less than 1.0.0 in current implementation, got %d", result)
	}
}

func TestVersion_CompareTo_SameNumbers(t *testing.T) {
	v1 := NewVersion("1.2.3")
	v2 := NewVersion("1.2.3")
	if v1.CompareTo(v2) != 0 {
		t.Errorf("1.2.3 and 1.2.3 should be equal, got %d", v1.CompareTo(v2))
	}
}

func TestVersion_CompareTo_TimeBased(t *testing.T) {
	v1 := NewVersion("1.0.0")
	v2 := NewVersion("1.0.0")
	v1.PublicTime = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	v2.PublicTime = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	if v1.CompareTo(v2) >= 0 {
		t.Errorf("older version should be less than newer version")
	}
}

func TestVersionNumbers_CompareTo_NoOverflow(t *testing.T) {
	big := NewVersionNumbers([]int{2147483647, 0})
	bigger := NewVersionNumbers([]int{2147483648, 0})
	if big.CompareTo(bigger) >= 0 {
		t.Errorf("2147483647.0 should be less than 2147483648.0")
	}
}
