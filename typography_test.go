package typography

import "testing"

func TestBeautify(t *testing.T) {
    tests := []struct {
        input string
        expected string
    }{
        {"...", "…"},
        {"---", "—"}, // \u2013
        {"--", "–"}, // \u2014
        {"This is... a test.", "This is… a test."},
        {"This is -- a test.", "This is – a test."},
        {"This is---a test.", "This is—a test."},
        {"This... is -- a---test... well---you know.",
            "This… is – a—test… well—you know."},
        {`"What a beautiful day", he said.`,
            `«What a beautiful day», he said.`},
        {`He said: "She said: 'Don't be a fool!'"`,
            `He said: «She said: ‹Don’t be a fool!›»`},
        {`"Ну что...", сказал он -- и молчал.`,
            `«Ну что…», сказал он – и молчал.`},
    }
    for _, test := range tests {
        got := Beautify(test.input)
        if  got != test.expected {
            t.Errorf("Beautify(\"%s\") got: %s, expected: %s",
                test.input, got, test.expected)
        }
    }
}

func BenchmarkBeautify(b *testing.B) {
    str := `He said: "She said: 'Don't be... such a---'"`
    for i := 0; i <= b.N; i++ {
        Beautify(str)
    }
}
