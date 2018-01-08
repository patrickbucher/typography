package typography

import (
	"bytes"
	"strings"
	"unicode"
)

// Fold joins the line of the given text and inserts new line breaks at spaces
// after length is reached.
func Fold(text string, length int) string {
	if length < 1 {
		// TODO: consider throwing error instead
		return text
	}
	runes := []rune(SquashSpaces(joinLines(text)))
	buf := bytes.NewBufferString("")
	lineBuf := make([]rune, length)
	var overflow []rune
	var w, lastSpace int
	for i := 0; i < len(runes); i++ {
		c := runes[i]
		if unicode.IsSpace(c) {
			if w == 0 {
				continue
			} else {
				lastSpace = w
			}
		}
		if w < length {
			lineBuf[w] = c
			w++
		}
		if w == length {
			cutoff := lastSpace
			if cutoff >= len(lineBuf) || cutoff == 0 {
				cutoff = len(lineBuf)
			}
			if unicode.IsSpace(lineBuf[cutoff-1]) && cutoff > 1 {
				cutoff--
			}
			buf.WriteString(string(lineBuf[:cutoff]))
			buf.WriteRune('\n')
			bufEnd := w
			if cutoff+1 < bufEnd {
				overflow = lineBuf[cutoff+1 : bufEnd]
				lineBuf = make([]rune, length)
				for j := 0; j < len(overflow); j++ {
					lineBuf[j] = overflow[j]
				}
				w = len(overflow)
			} else {
				lineBuf = make([]rune, length)
				w = 0
			}
			lastSpace = 0
		}
	}
	rest := string(lineBuf[:w])
	if len(rest) > 0 {
		buf.WriteString(rest)
		buf.WriteRune('\n')
	}
	return strings.TrimSpace(buf.String())
}

func SimpleFold(text string, length int) string {
	buf := bytes.NewBufferString("")
	var n int
	for _, word := range strings.Fields(text) {
		l := len([]rune(word))
		if n+l+1 <= length {
			if n > 0 {
				buf.WriteRune(' ')
				n++
			}
		} else {
			if n > 0 {
				buf.WriteRune('\n')
				n = 0
			}
		}
		buf.WriteString(word)
		n += l
	}
	return buf.String()
}

func joinLines(text string) string {
	text = strings.Replace(text, "\n", " ", -1)
	text = strings.Replace(text, "\r", " ", -1)
	return strings.TrimSpace(text)
}
