package utils

import (
	"bytes"
	"strings"
	"unicode"
)

// SquashSpaces removes excess whitespace between words.
func SquashSpaces(text string) string {
	buf := bytes.NewBufferString("")
	var last rune
	for _, r := range []rune(text) {
		if !unicode.IsSpace(r) || unicode.IsSpace(r) && !unicode.IsSpace(last) {
			buf.WriteRune(r)
		}
		last = r
	}
	return buf.String()
}

// LongestLine counts the rune length of the longest line.
func LongestLine(text string) int {
	var longest int
	lines := strings.Split(text, "\n")
	for _, v := range lines {
		l := len([]rune(v))
		if l > longest {
			longest = l
		}
	}
	return longest
}
