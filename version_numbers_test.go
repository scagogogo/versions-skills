package versions

import "testing"

func TestVersionNumbers_Len(t *testing.T) {
	vn := NewVersionNumbers([]int{1, 2, 3})
	if vn.Len() != 3 {
		t.Errorf("Len() = %d, want 3", vn.Len())
	}
	empty := NewVersionNumbers([]int{})
	if empty.Len() != 0 {
		t.Errorf("Len() = %d, want 0", empty.Len())
	}
}

func TestVersionNumbers_At(t *testing.T) {
	vn := NewVersionNumbers([]int{1, 2, 3})
	if vn.At(0) != 1 {
		t.Errorf("At(0) = %d, want 1", vn.At(0))
	}
	if vn.At(2) != 3 {
		t.Errorf("At(2) = %d, want 3", vn.At(2))
	}
	if vn.At(5) != 0 {
		t.Errorf("At(5) = %d, want 0 (out of bounds)", vn.At(5))
	}
	if vn.At(-1) != 0 {
		t.Errorf("At(-1) = %d, want 0 (out of bounds)", vn.At(-1))
	}
}

func TestVersionNumbers_String(t *testing.T) {
	vn := NewVersionNumbers([]int{1, 2, 3})
	if vn.String() != "1.2.3" {
		t.Errorf("String() = %q, want %q", vn.String(), "1.2.3")
	}
}

func TestVersionNumbers_Equals(t *testing.T) {
	a := NewVersionNumbers([]int{1, 2, 3})
	b := NewVersionNumbers([]int{1, 2, 3})
	c := NewVersionNumbers([]int{1, 2})
	d := NewVersionNumbers([]int{1, 2, 4})
	if !a.Equals(b) {
		t.Error("a should equal b")
	}
	if a.Equals(c) {
		t.Error("a should not equal c (different length)")
	}
	if a.Equals(d) {
		t.Error("a should not equal d (different values)")
	}
}
