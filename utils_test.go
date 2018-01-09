package typography

import "testing"

func TestSquashSpaces(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Das ist  ein   Test.", "Das ist ein Test."},
		{"Ein            Test.", "Ein Test."},
	}
	for _, test := range tests {
		if got := SquashSpaces(test.input); got != test.expected {
			t.Errorf("expected: \"%s\"\n, got: \"%s\"\n", test.expected, got)
		}
	}
}

func TestLongestLine(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"a\nbc\ndef\nghij", 4},
		{"Das ist\nein Test.", 9},
	}
	for _, test := range tests {
		if got := LongestLine(test.input); got != test.expected {
			t.Errorf("expected: \"%s\"\n,  got: \"%d\"\n", test.expected, got)
		}
	}
}
