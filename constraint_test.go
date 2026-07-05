package versions

import (
	"strings"
	"testing"
)

func TestParseConstraint(t *testing.T) {
	tests := []struct {
		expr     string
		op       ConstraintOperator
		version  string
		hasError bool
	}{
		{">=1.0.0", ConstraintGreaterThanOrEqual, "1.0.0", false},
		{"<2.0.0", ConstraintLessThan, "2.0.0", false},
		{"^1.2.3", ConstraintCaret, "1.2.3", false},
		{"~1.2", ConstraintTilde, "1.2", false},
		{"1.0.0", ConstraintEqual, "1.0.0", false},
		{"!=1.0.0", ConstraintNotEqual, "1.0.0", false},
		{"", "", "", true},
	}
	for _, tt := range tests {
		c, err := ParseConstraint(tt.expr)
		if tt.hasError {
			if err == nil {
				t.Errorf("ParseConstraint(%q) should return error", tt.expr)
			}
			continue
		}
		if err != nil {
			t.Errorf("ParseConstraint(%q) unexpected error: %v", tt.expr, err)
			continue
		}
		if c.Operator != tt.op {
			t.Errorf("ParseConstraint(%q) operator = %q, want %q", tt.expr, c.Operator, tt.op)
		}
		if c.Version.Raw != tt.version {
			t.Errorf("ParseConstraint(%q) version = %q, want %q", tt.expr, c.Version.Raw, tt.version)
		}
	}
}

func TestConstraint_Match_Comparison(t *testing.T) {
	c, _ := ParseConstraint(">=1.0.0")
	if !c.Match(NewVersion("1.5.0")) {
		t.Error("1.5.0 should match >=1.0.0")
	}
	if c.Match(NewVersion("0.9.0")) {
		t.Error("0.9.0 should not match >=1.0.0")
	}
}

func TestConstraint_Match_Caret(t *testing.T) {
	c, _ := ParseConstraint("^1.2.3")
	if !c.Match(NewVersion("1.9.9")) {
		t.Error("1.9.9 should match ^1.2.3")
	}
	if c.Match(NewVersion("2.0.0")) {
		t.Error("2.0.0 should not match ^1.2.3")
	}
	if !c.Match(NewVersion("1.2.3")) {
		t.Error("1.2.3 should match ^1.2.3")
	}
}

func TestConstraint_Match_Caret_ZeroMajor(t *testing.T) {
	c, _ := ParseConstraint("^0.2.3")
	if !c.Match(NewVersion("0.2.9")) {
		t.Error("0.2.9 should match ^0.2.3")
	}
	if c.Match(NewVersion("0.3.0")) {
		t.Error("0.3.0 should not match ^0.2.3")
	}
}

func TestConstraint_Match_Tilde(t *testing.T) {
	c, _ := ParseConstraint("~1.2.3")
	if !c.Match(NewVersion("1.2.9")) {
		t.Error("1.2.9 should match ~1.2.3")
	}
	if c.Match(NewVersion("1.3.0")) {
		t.Error("1.3.0 should not match ~1.2.3")
	}
}

func TestConstraint_Match_LessThan(t *testing.T) {
	c, _ := ParseConstraint("<2.0.0")
	if !c.Match(NewVersion("1.9.9")) {
		t.Error("1.9.9 should match <2.0.0")
	}
	if c.Match(NewVersion("2.0.0")) {
		t.Error("2.0.0 should not match <2.0.0")
	}
}

func TestConstraint_Match_NotEqual(t *testing.T) {
	c, _ := ParseConstraint("!=1.0.0")
	if c.Match(NewVersion("1.0.0")) {
		t.Error("1.0.0 should not match !=1.0.0")
	}
	if !c.Match(NewVersion("1.0.1")) {
		t.Error("1.0.1 should match !=1.0.0")
	}
}

func TestConstraintSet_Match(t *testing.T) {
	cs, _ := ParseConstraintSet(">=1.0.0,<2.0.0")
	if !cs.Match(NewVersion("1.5.0")) {
		t.Error("1.5.0 should match >=1.0.0,<2.0.0")
	}
	if cs.Match(NewVersion("2.0.0")) {
		t.Error("2.0.0 should not match >=1.0.0,<2.0.0")
	}
	if cs.Match(NewVersion("0.9.0")) {
		t.Error("0.9.0 should not match >=1.0.0,<2.0.0")
	}
}

func TestConstraintSet_SingleConstraint(t *testing.T) {
	cs, _ := ParseConstraintSet("^1.2.3")
	if !cs.Match(NewVersion("1.5.0")) {
		t.Error("1.5.0 should match ^1.2.3 via ConstraintSet")
	}
	if cs.Match(NewVersion("2.0.0")) {
		t.Error("2.0.0 should not match ^1.2.3 via ConstraintSet")
	}
}

func TestConstraint_String(t *testing.T) {
	tests := []struct {
		expr     string
		expected string
	}{
		{">=1.0.0", ">=1.0.0"},
		{"<2.0.0", "<2.0.0"},
		{"=1.5.0", "=1.5.0"},
		{"!=0.9.0", "!=0.9.0"},
		{"^1.2.3", "^1.2.3"},
		{"~1.2", "~1.2"},
	}
	for _, tt := range tests {
		t.Run(tt.expr, func(t *testing.T) {
			c, err := ParseConstraint(tt.expr)
			if err != nil {
				t.Fatalf("ParseConstraint(%q) error: %v", tt.expr, err)
			}
			if got := c.String(); got != tt.expected {
				t.Errorf("Constraint.String() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestConstraintSet_String(t *testing.T) {
	cs, err := ParseConstraintSet(">=1.0.0,<2.0.0")
	if err != nil {
		t.Fatalf("ParseConstraintSet error: %v", err)
	}
	got := cs.String()
	if got != ">=1.0.0,<2.0.0" {
		t.Errorf("ConstraintSet.String() = %q, want %q", got, ">=1.0.0,<2.0.0")
	}
}

func TestConstraintUnion(t *testing.T) {
	cu, err := ParseConstraintUnion(">=1.0.0,<2.0.0 || >=3.0.0")
	if err != nil {
		t.Fatalf("ParseConstraintUnion error: %v", err)
	}
	if !cu.Match(NewVersion("1.5.0")) {
		t.Error("1.5.0 should match >=1.0.0,<2.0.0")
	}
	if !cu.Match(NewVersion("3.5.0")) {
		t.Error("3.5.0 should match >=3.0.0")
	}
	if cu.Match(NewVersion("2.5.0")) {
		t.Error("2.5.0 should not match")
	}
}

func TestConstraintUnion_Single(t *testing.T) {
	cu, err := ParseConstraintUnion(">=1.0.0")
	if err != nil {
		t.Fatalf("ParseConstraintUnion error: %v", err)
	}
	if !cu.Match(NewVersion("1.5.0")) {
		t.Error("1.5.0 should match >=1.0.0")
	}
}

func TestConstraintUnion_Empty(t *testing.T) {
	_, err := ParseConstraintUnion("")
	if err == nil {
		t.Error("ParseConstraintUnion should return error for empty expression")
	}
}

func TestConstraintUnion_String(t *testing.T) {
	cu, _ := ParseConstraintUnion(">=1.0.0,<2.0.0 || >=3.0.0")
	s := cu.String()
	if !strings.Contains(s, "||") {
		t.Errorf("String() = %q, should contain ||", s)
	}
}

func TestConstraintSet_Satisfies(t *testing.T) {
	cs, _ := ParseConstraintSet(">=1.0.0,<2.0.0")
	if !cs.Satisfies(NewVersion("1.5.0")) {
		t.Error("1.5.0 should satisfy >=1.0.0,<2.0.0")
	}
	if cs.Satisfies(NewVersion("2.5.0")) {
		t.Error("2.5.0 should not satisfy >=1.0.0,<2.0.0")
	}
}

func TestConstraintSet_Len(t *testing.T) {
	cs, _ := ParseConstraintSet(">=1.0.0,<2.0.0")
	if cs.Len() != 2 {
		t.Errorf("Len() = %d, want 2", cs.Len())
	}
}

func TestConstraintWildcard(t *testing.T) {
	// 1.x := >=1.0.0, <2.0.0
	c, err := ParseConstraint("1.x")
	if err != nil {
		t.Fatalf("ParseConstraint 1.x: %v", err)
	}
	if c.Operator != ConstraintWildcard {
		t.Errorf("1.x operator = %v, want ConstraintWildcard", c.Operator)
	}
	if !c.Match(NewVersion("1.0.0")) {
		t.Error("1.0.0 should match 1.x")
	}
	if !c.Match(NewVersion("1.9.9")) {
		t.Error("1.9.9 should match 1.x")
	}
	if c.Match(NewVersion("2.0.0")) {
		t.Error("2.0.0 should not match 1.x")
	}
	if c.Match(NewVersion("0.9.0")) {
		t.Error("0.9.0 should not match 1.x")
	}

	// 1.2.x := >=1.2.0, <1.3.0
	c2, _ := ParseConstraint("1.2.x")
	if !c2.Match(NewVersion("1.2.0")) {
		t.Error("1.2.0 should match 1.2.x")
	}
	if !c2.Match(NewVersion("1.2.9")) {
		t.Error("1.2.9 should match 1.2.x")
	}
	if c2.Match(NewVersion("1.3.0")) {
		t.Error("1.3.0 should not match 1.2.x")
	}

	// 大写 X 和 * 通配符
	c3, _ := ParseConstraint("1.2.X")
	if !c3.Match(NewVersion("1.2.5")) {
		t.Error("1.2.5 should match 1.2.X")
	}
	c4, _ := ParseConstraint("1.*")
	if !c4.Match(NewVersion("1.5.0")) {
		t.Error("1.5.0 should match 1.*")
	}
}

func TestReplaceWildcardWithZero(t *testing.T) {
	if got := replaceWildcardWithZero("1.x"); got != "1.0" {
		t.Errorf("replaceWildcardWithZero(1.x) = %s, want 1.0", got)
	}
	if got := replaceWildcardWithZero("1.2.X"); got != "1.2.0" {
		t.Errorf("replaceWildcardWithZero(1.2.X) = %s, want 1.2.0", got)
	}
	if got := replaceWildcardWithZero("1.*"); got != "1.0" {
		t.Errorf("replaceWildcardWithZero(1.*) = %s, want 1.0", got)
	}
	if got := replaceWildcardWithZero("1.2.3"); got != "1.2.3" {
		t.Errorf("replaceWildcardWithZero(1.2.3) = %s, want 1.2.3", got)
	}
}

func TestConstraintUnion_Satisfies(t *testing.T) {
	cu, err := ParseConstraintUnion(">=1.0.0,<2.0.0 || >=3.0.0,<4.0.0")
	if err != nil {
		t.Fatalf("ParseConstraintUnion: %v", err)
	}
	// Satisfies 是 Match 的包装
	if !cu.Satisfies(NewVersion("1.5.0")) {
		t.Error("1.5.0 should satisfy union")
	}
	if !cu.Satisfies(NewVersion("3.5.0")) {
		t.Error("3.5.0 should satisfy union")
	}
	if cu.Satisfies(NewVersion("2.5.0")) {
		t.Error("2.5.0 should not satisfy union")
	}
}
