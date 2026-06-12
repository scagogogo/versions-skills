package versions

import (
	"sort"
	"testing"
)

func TestVersionSlice_Sort(t *testing.T) {
	slice := VersionSlice(NewVersions("2.0.0", "1.0.0", "1.5.0", "1.10.0"))
	sort.Sort(slice)
	if slice[0].Raw != "1.0.0" {
		t.Errorf("After sort, first = %s, want 1.0.0", slice[0].Raw)
	}
	if slice[len(slice)-1].Raw != "2.0.0" {
		t.Errorf("After sort, last = %s, want 2.0.0", slice[len(slice)-1].Raw)
	}
}
