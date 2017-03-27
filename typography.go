package typography

import "unicode"

func Beautify(str string) string {
    toReplace := make(chan rune)
    fromReplace := make(chan rune)
    go replace(toReplace, fromReplace)
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

func replace(in <-chan rune, out chan<- rune) {
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
        buffered := len(buf)
        if buffered == 0 {
            break
        }
        if buffered >= 2 && buf[0] == '-' && buf[1] == '-' {
            // --- to — (mdash) and -- to – (ndash)
            if buffered >= 3 && buf[2] == '-' {
                out<- '\u2014'
                last = buf[2]
                buf = buf[3:]
            } else {
                out<- '\u2013'
                last = buf[1]
                buf = buf[2:]
            }
        } else if buffered >= 3 && buf[0] == '.' && buf[1] == '.' && buf[2] == '.' {
            // ... to ellipsis …
            out<- '\u2026'
            buf = buf[3:]
        } else if buffered >= 1 && buf[0] == '"' {
            // "" to «»
            if last == 0 || unicode.IsSpace(last) {
                out<- '«'
            } else {
                out<- '»'
            }
            last = buf[0]
            buf = buf[1:]
        } else if buffered >= 1 && buf[0] == '\'' {
            // '' to ‹› or don't to don’t
            if buffered >= 2 && unicode.IsLetter(last) &&
                unicode.IsLetter(buf[1]) {
                out<- '’'
            } else if last == 0 || unicode.IsSpace(last) || last == '"' {
                out<- '‹'
            } else {
                out<- '›'
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
