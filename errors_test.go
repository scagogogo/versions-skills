package versions

import (
	"errors"
	"testing"
)

func TestErrConstraintInvalid(t *testing.T) {
	_, err := ParseConstraint("")
	if !errors.Is(err, ErrEmptyConstraint) {
		t.Errorf("ParseConstraint(\"\") error = %v, want ErrEmptyConstraint", err)
	}
}

func TestErrMissingVersionInConstraint(t *testing.T) {
	_, err := ParseConstraint(">=")
	if !errors.Is(err, ErrMissingVersionInConstraint) {
		t.Errorf("ParseConstraint(\">=\") error = %v, want ErrMissingVersionInConstraint", err)
	}
}

func TestErrInvalidVersionInConstraint(t *testing.T) {
	_, err := ParseConstraint(">=not-a-version")
	if !errors.Is(err, ErrInvalidVersionInConstraint) {
		t.Errorf("ParseConstraint(\">=not-a-version\") error = %v, want ErrInvalidVersionInConstraint", err)
	}
}

func TestErrVersionInvalid(t *testing.T) {
	_, err := NewVersionE("not-a-version")
	if !errors.Is(err, ErrVersionInvalid) {
		t.Errorf("NewVersionE error = %v, want ErrVersionInvalid", err)
	}
}
