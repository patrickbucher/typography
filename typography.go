package typography

import "unicode"

type quotePair struct {
    opening rune
    closing rune
}

type quoteRule struct {
    single quotePair
    double quotePair
}

type QuoteStyle int

const (
    English QuoteStyle = iota
    German
    Guillemets
    ReverseGuillemets
)

var quoteRules  = map[QuoteStyle]quoteRule {
    English: quoteRule {
        single: quotePair{opening: '\u2018', closing: '\u2019'}, // ‘’
        double: quotePair{opening: '\u201c', closing: '\u201d'}, // “”
    },
    German: quoteRule {
        single: quotePair{opening: '\u201a', closing: '\u2018'}, // ‚‘
        double: quotePair{opening: '\u201e', closing: '\u201c'}, // „“
    },
    Guillemets: quoteRule {
        single: quotePair{opening: '\u2039', closing: '\u203a'}, // ‹›
        double: quotePair{opening: '\u00ab', closing: '\u00bb'}, // «»
    },
    ReverseGuillemets: quoteRule {
        single: quotePair{opening: '\u203a', closing: '\u2039'}, // ›‹
        double: quotePair{opening: '\u00bb', closing: '\u00ab'}, // »«
    },
}

func Beautify(str string, style QuoteStyle) string {
    toReplace := make(chan rune)
    fromReplace := make(chan rune)
    go replace(toReplace, fromReplace, quoteRules[style])
    go func() {
        for _, r := range str {
            toReplace <- r
        }
        close(toReplace)
    }()
    beautified := make([]rune, 0)
    for r := range fromReplace {
        beautified = append(beautified, r)
    }
    return string(beautified)
}

func replace(in <-chan rune, out chan<- rune, rule quoteRule) {
    const BUF_SIZE = 5
    buf := make([]rune, 0)
    drained := false
    var last rune
    for {
        for !drained && len(buf) < BUF_SIZE {
            r, ok := <-in
            if (ok) {
                buf = append(buf, r)
            } else {
                drained = true
                break
            }
        }
        nbuf := len(buf)
        if nbuf == 0 {
            break
        }
        if nbuf >= 2 && buf[0] == '-' && buf[1] == '-' {
            // --- to — (mdash) and -- to – (ndash)
            if nbuf >= 3 && buf[2] == '-' {
                out<- '\u2014'
                last = buf[2]
                buf = buf[3:]
            } else {
                out<- '\u2013'
                last = buf[1]
                buf = buf[2:]
            }
        } else if nbuf >= 3 && buf[0] == '.' && buf[1] == '.' && buf[2] == '.' {
            // ... to ellipsis …
            out<- '\u2026'
            buf = buf[3:]
        } else if nbuf >= 1 && buf[0] == '"' {
            // "" to «»
            if last == 0 || unicode.IsSpace(last) {
                out<- rule.double.opening
            } else {
                out<- rule.double.closing
            }
            last = buf[0]
            buf = buf[1:]
        } else if nbuf >= 1 && buf[0] == '\'' {
            // '' to ‹› or don't to don’t
            if nbuf >= 2 && unicode.IsLetter(last) &&
                unicode.IsLetter(buf[1]) {
                out<- '’'
            } else if last == 0 || unicode.IsSpace(last) || last == '"' {
                out<- rule.single.opening
            } else {
                out<- rule.single.closing
            }
            last = buf[0]
            buf = buf[1:]
        } else {
            out<- buf[0]
            last = buf[0]
            buf = buf[1:]
        }
    }
    for _, r := range buf {
        out<- r
    }
    close(out)
}
