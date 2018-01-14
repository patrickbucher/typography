package typography

import (
	"strings"
	"unicode"
)

type quotePair struct {
	opening rune
	closing rune
}

type quoteRule struct {
	single quotePair
	double quotePair
}

// QuoteStyle to be used (English, German, Guillemets or ReverseGuillemets)
type QuoteStyle int

const (
	// English style quotes: “” and ‘’
	English QuoteStyle = iota
	// German style quotes: „“ and ‚‘
	German
	// Guillemets (French style): «» and ‹›
	Guillemets
	// ReverseGuillemets (German book style): »« and ›‹
	ReverseGuillemets
)

var quoteRules = map[QuoteStyle]quoteRule{
	English: quoteRule{
		single: quotePair{opening: '\u2018', closing: '\u2019'}, // ‘’
		double: quotePair{opening: '\u201c', closing: '\u201d'}, // “”
	},
	German: quoteRule{
		single: quotePair{opening: '\u201a', closing: '\u2018'}, // ‚‘
		double: quotePair{opening: '\u201e', closing: '\u201c'}, // „“
	},
	Guillemets: quoteRule{
		single: quotePair{opening: '\u2039', closing: '\u203a'}, // ‹›
		double: quotePair{opening: '\u00ab', closing: '\u00bb'}, // «»
	},
	ReverseGuillemets: quoteRule{
		single: quotePair{opening: '\u203a', closing: '\u2039'}, // ›‹
		double: quotePair{opening: '\u00bb', closing: '\u00ab'}, // »«
	},
}

// Beautify applies replacement rules to enhance the typography of the text
func Beautify(str string, style QuoteStyle) string {
	const fourSpaces = "    "
	var beautified string
	verbatimMode := false
	lines := strings.Split(str, "\n")
	if len(lines) == 1 {
		return replace(str, quoteRules[style])
	}
	for i, line := range lines {
		if strings.HasPrefix(line, "```") {
			verbatimMode = !verbatimMode
		}
		if !verbatimMode && !strings.HasPrefix(line, "\t") &&
			!strings.HasPrefix(line, fourSpaces) {
			beautified += replace(line, quoteRules[style])
		} else {
			beautified += line
		}
		if i < len(lines)-1 {
			beautified += "\n"
		}
	}
	return beautified
}

func replace(input string, rule quoteRule) string {
	const bufSize = 5
	in := []rune(input)
	out := make([]rune, 0)
	buf := make([]rune, 0)
	index := 0
	var last rune
	for {
		for index < len(in) && len(buf) < bufSize {
			buf = append(buf, in[index])
			index++
		}
		nbuf := len(buf)
		if nbuf == 0 {
			break
		}
		if nbuf >= 2 && buf[0] == '-' && buf[1] == '-' {
			// --- to — (mdash) and -- to – (ndash)
			if nbuf >= 3 && buf[2] == '-' {
				out = append(out, '\u2014')
				last = buf[2]
				buf = buf[3:]
			} else {
				out = append(out, '\u2013')
				last = buf[1]
				buf = buf[2:]
			}
		} else if nbuf >= 3 && buf[0] == '.' && buf[1] == '.' && buf[2] == '.' {
			// ... to ellipsis …
			out = append(out, '\u2026')
			buf = buf[3:]
		} else if nbuf >= 1 && buf[0] == '"' {
			// "" to double quotes according to the rule
			if last == 0 || unicode.IsSpace(last) {
				out = append(out, rule.double.opening)
			} else if nbuf >= 2 && unicode.IsLetter(buf[1]) {
				out = append(out, rule.double.opening)
			} else {
				out = append(out, rule.double.closing)
			}
			last = buf[0]
			buf = buf[1:]
		} else if nbuf >= 1 && buf[0] == '\'' {
			// '' to single quotes according to the rule and don't to don’t
			if nbuf >= 2 &&
				(unicode.IsLetter(last) && unicode.IsLetter(buf[1]) ||
					unicode.IsDigit(last) && unicode.IsDigit(buf[1])) {
				out = append(out, '’')
			} else if last == 0 || unicode.IsSpace(last) || last == '"' ||
				last == '-' {
				out = append(out, rule.single.opening)
			} else if nbuf >= 2 && unicode.IsLetter(buf[1]) {
				out = append(out, rule.single.opening)
			} else {
				out = append(out, rule.single.closing)
			}
			last = buf[0]
			buf = buf[1:]
		} else {
			out = append(out, buf[0])
			last = buf[0]
			buf = buf[1:]
		}
	}
	return string(out)
}
