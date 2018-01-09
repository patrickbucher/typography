package typography

import "testing"

var tests = []struct {
	input    string
	length   int
	expected string
}{
	{"a b c d e f", 2, "a\nb\nc\nd\ne\nf"},
	{"ab cd ef gh", 2, "ab\ncd\nef\ngh"},
	{"Das ist ein Test.", 10, "Das ist\nein Test."},
	{"Das   ist   ein   Test.", 10, "Das ist\nein Test."},
	{"Dies ist ein beinahe realistischer Test.", 21, "Dies ist ein beinahe\nrealistischer Test."},
}

func TestFold(t *testing.T) {
	for _, test := range tests {
		got := Fold(test.input, test.length)
		if got != test.expected {
			t.Errorf("expected: %q\ngot: %q\n", test.expected, got)
		}
	}
}

func BenchmarkFold(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fold("This is another test used for benchmarking.", 16)
	}
}
