package versions

import "testing"

func TestGetSuffixWeight(t *testing.T) {
	tests := []struct {
		suffix string
		weight SuffixWeight
	}{
		{"-alpha", SuffixWeightAlpha},
		{"-alpha1", SuffixWeightAlpha},
		{"-ALPHA", SuffixWeightAlpha},
		{"-beta", SuffixWeightBeta},
		{"-beta2", SuffixWeightBeta},
		{"-rc", SuffixWeightRC},
		{"-RC1", SuffixWeightRC},
		{"-cr", SuffixWeightCR},
		{"-dev", SuffixWeightDev},
		{"-snapshot", SuffixWeightSnapshot},
		{"-milestone", SuffixWeightMilestone},
		{"-M1", SuffixWeightMilestone},
		{".Final", SuffixWeightFinal},
		{"-sp1", SuffixWeightSP},
		{"-unknown", SuffixWeightUnknown},
	}
	for _, tt := range tests {
		got := GetSuffixWeight(tt.suffix)
		if got != tt.weight {
			t.Errorf("GetSuffixWeight(%q) = %d, want %d", tt.suffix, got, tt.weight)
		}
	}
}

func TestVersionSuffix_CompareTo_WithWeight(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		expect int
	}{
		// alpha < beta < rc
		{"-alpha", "-beta", -1},
		{"-beta", "-rc", -1},
		{"-alpha1", "-alpha2", -1},
		// 大小写不敏感
		{"-RC1", "-beta1", 1},
		{"-ALPHA", "-beta", -1},
		// 同权重子版本号比较
		{"-beta1", "-beta2", -1},
	}
	for _, tt := range tests {
		a := VersionSuffix(tt.a)
		b := VersionSuffix(tt.b)
		got := a.CompareTo(b)
		if got != tt.expect {
			t.Errorf("VersionSuffix(%q).CompareTo(%q) = %d, want %d", tt.a, tt.b, got, tt.expect)
		}
	}
}

func TestSuffixWeight_String(t *testing.T) {
	if SuffixWeightDev.String() != "dev" {
		t.Errorf("SuffixWeightDev.String() = %q, want %q", SuffixWeightDev.String(), "dev")
	}
	if SuffixWeightAlpha.String() != "alpha" {
		t.Errorf("SuffixWeightAlpha.String() = %q, want %q", SuffixWeightAlpha.String(), "alpha")
	}
	if SuffixWeightUnknown.String() != "unknown" {
		t.Errorf("SuffixWeightUnknown.String() = %q, want %q", SuffixWeightUnknown.String(), "unknown")
	}
}

func TestExtractSubVersion(t *testing.T) {
	tests := []struct {
		suffix string
		expect int
	}{
		{"-alpha1", 1},
		{"-beta2", 2},
		{"-rc10", 10},
		{"-alpha", 0},
	}
	for _, tt := range tests {
		got := extractSubVersion(tt.suffix)
		if got != tt.expect {
			t.Errorf("extractSubVersion(%q) = %d, want %d", tt.suffix, got, tt.expect)
		}
	}
}
