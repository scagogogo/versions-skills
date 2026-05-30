package versions

import "testing"

func TestParserOption_UnderscoreDelimiter(t *testing.T) {
	v := NewVersionWithOption("1_2_3", ParserOption{Delimiters: ".-_"})
	if !v.IsValid() {
		t.Fatalf("1_2_3 should be valid with underscore delimiter")
	}
}

func TestParserOption_HyphenDelimiter(t *testing.T) {
	v := NewVersionWithOption("1-2-3", ParserOption{Delimiters: ".-_"})
	if !v.IsValid() {
		t.Fatalf("1-2-3 should be valid with hyphen delimiter")
	}
}

func TestParserOption_DefaultOnlyDot(t *testing.T) {
	v := NewVersion("1.2.3")
	if !v.IsValid() {
		t.Fatalf("1.2.3 should be valid with default options")
	}
	if v.VersionNumbers[0] != 1 || v.VersionNumbers[1] != 2 || v.VersionNumbers[2] != 3 {
		t.Fatalf("1.2.3 should parse as [1,2,3], got %v", v.VersionNumbers)
	}
}

func TestParserOption_PythonStyleVersion(t *testing.T) {
	v := NewVersionWithOption("2023_09_15", ParserOption{Delimiters: ".-_"})
	if !v.IsValid() {
		t.Fatalf("2023_09_15 should be valid with underscore delimiter")
	}
}
