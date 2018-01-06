package typography

import "testing"

func TestFold(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		expected string
	}{
		{"a b c d e f", 2, "a\nb\nc\nd\ne\nf"},
		{"ab cd ef gh", 2, "ab\ncd\nef\ngh"},
		{"a bc def ghij", 2, "a\nbc\nde\nf\ngh\nij"},
		{"Das ist ein Test.", 10, "Das ist\nein Test."},
	}
	for _, test := range tests {
		got := Fold(test.input, test.length)
		if got != test.expected {
			t.Errorf("expected: %q\ngot: %q\n", test.expected, got)
		}
	}
}
