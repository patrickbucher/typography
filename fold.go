package typography

import (
	"bytes"
	"strings"
)

// Fold joins the line of the given text and inserts new line breaks at spaces
// after length is reached.
func Fold(text string, length int) string {
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
