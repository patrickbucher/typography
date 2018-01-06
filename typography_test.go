package typography

import (
	"io/ioutil"
	"testing"
)

func TestBeautify(t *testing.T) {
	simpleTests := []struct {
		input      string
		quoteStyle QuoteStyle
		expected   string
	}{
		{"...", 0, "…"},
		{"---", 0, "—"}, // \u2013
		{"--", 0, "–"},  // \u2014
		{"This is... a test.", 0, "This is… a test."},
		{"This is -- a test.", 0, "This is – a test."},
		{"This is---a test.", 0, "This is—a test."},
		{"This... is -- a---test... well---you know.", 0,
			"This… is – a—test… well—you know."},
		{`"What a beautiful day", he said.`, Guillemets,
			`«What a beautiful day», he said.`},
		{`He said: "She said: 'Don't be a fool!'"`, Guillemets,
			`He said: «She said: ‹Don’t be a fool!›»`},
		{`"Ну что...", сказал он -- и молчал.`, Guillemets,
			`«Ну что…», сказал он – и молчал.`},
		{`"Let's 'do' a test."`, English, "“Let’s ‘do’ a test.”"},
		{`"Mach'n ma a 'Test'."`, German, "„Mach’n ma a ‚Test‘.“"},
		{`"Let's 'do' a test."`, Guillemets, "«Let’s ‹do› a test.»"},
		{`"Let's 'do' a test."`, ReverseGuillemets, "»Let’s ›do‹ a test.«"},
		{`"Noch ein Test."`, English, `“Noch ein Test.”`},
		{`"Noch ein Test."`, German, `„Noch ein Test.“`},
		{`"Noch ein Test."`, ReverseGuillemets, `»Noch ein Test.«`},
		{`He said: "Won't this be---the only---'complete' test...?"`, English,
			`He said: “Won’t this be—the only—‘complete’ test…?”`},
		{`Er hat g'sagt: "Sui des net -- d oanzig -- 'komplette' Test...?"`,
			German,
			`Er hat g’sagt: „Sui des net – d oanzig – ‚komplette‘ Test…?“`},
		{`Rock'n'Roll`, English, `Rock’n’Roll`},
		{`Rock'n'Roll`, German, `Rock’n’Roll`},
		{`Rock'n'Roll`, Guillemets, `Rock’n’Roll`},
		{`Rock'n'Roll`, ReverseGuillemets, `Rock’n’Roll`},
		{`Das ist ("war") ein Test.`, Guillemets, `Das ist («war») ein Test.`},
		{`"Das ist ('war') ein Test."`, Guillemets, `«Das ist (‹war›) ein Test.»`},
	}
	advancedTests := []struct {
		inputFile    string
		quoteStyle   QuoteStyle
		expectedFile string
	}{
		{"tests/test-input-markdown.md", Guillemets, "tests/test-expected-markdown.md"},
	}

	for _, test := range simpleTests {
		got := Beautify(test.input, test.quoteStyle)
		if got != test.expected {
			t.Errorf("Beautify(\"%s\") got: %s, expected: %s",
				test.input, got, test.expected)
		}
	}

	readOrFail := func(filename string) string {
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Errorf("Error reading file %s: %v", filename, err)
		}
		return string(content)
	}

	for _, test := range advancedTests {
		input := readOrFail(test.inputFile)
		got := Beautify(input, test.quoteStyle)
		expected := readOrFail(test.expectedFile)
		if got != expected {
			t.Errorf("Beautify\n%s\ngot:\n%s\nexpected:\n%s\n", input, got,
				expected)
		}
	}
}

func BenchmarkBeautify(b *testing.B) {
	str := `He said: "She said: 'Don't be... such a---'"`
	for i := 0; i <= b.N; i++ {
		Beautify(str, 0)
	}
}
