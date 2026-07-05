package versions

import "testing"

func TestContainsPolicy_String(t *testing.T) {
	if got := ContainsPolicyNone.String(); got != "none" {
		t.Errorf("ContainsPolicyNone.String() = %q, want \"none\"", got)
	}
	if got := ContainsPolicyYes.String(); got != "yes" {
		t.Errorf("ContainsPolicyYes.String() = %q, want \"yes\"", got)
	}
	if got := ContainsPolicyNo.String(); got != "no" {
		t.Errorf("ContainsPolicyNo.String() = %q, want \"no\"", got)
	}
	// 未知值走 default 分支
	if got := ContainsPolicy(999).String(); got != "none" {
		t.Errorf("ContainsPolicy(999).String() = %q, want \"none\"", got)
	}
}
