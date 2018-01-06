package typography

import (
	"bytes"
	"strings"
	"unicode"
)

func Fold(text string, maxLineLength int) string {
	if maxLineLength < 1 {
		return text
	}
	text = strings.Replace(text, "\n", " ", -1)
	text = strings.Replace(text, "\r", " ", -1)
	text = strings.TrimSpace(text)
	runes := []rune(text)

	buf := bytes.NewBufferString("")
	size := len(runes)
	lineBuf := make([]rune, maxLineLength)
	var overflow []rune
	var last rune
	var lastSpace int
	var w int
	for i := 0; i < size; i++ {
		c := runes[i]
		if unicode.IsSpace(c) {
			if unicode.IsSpace(last) || w == 0 {
				last = c
				continue
			} else {
				lastSpace = i
			}
		}
		last = c

		if w < maxLineLength {
			lineBuf[w] = c
			w++
		}

		if w == maxLineLength {
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
				lineBuf = make([]rune, maxLineLength)
				for j := 0; j < len(overflow); j++ {
					lineBuf[j] = overflow[j]
				}
				w = len(overflow)
			} else {
				lineBuf = make([]rune, maxLineLength)
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
