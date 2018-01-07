package typography

import (
	"bufio"
	"bytes"
	"strings"
)

// Justify aligns the text to the left and right edge by filling the word gaps
// up with spaces until length is reached.
func Justify(text string, length int) string {
	output := bytes.NewBufferString("")
	reader := strings.NewReader(text)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > length {
			// TODO client's problem?
			continue
		}
		words := strings.Fields(line)
		spaces := make([]string, len(words)-1)
		diff := length - len(line) + len(spaces) // spaces where omitted, add them back
		for d, i := diff, 0; d > 0; d-- {
			spaces[i] += " "
			i++
			if i == len(spaces) {
				i = 0
			}
		}
		justified := bytes.NewBufferString("")
		for w, s := 0, 0; w < len(words) || s < len(spaces); w, s = w+1, s+1 {
			if w < len(words) {
				justified.WriteString(words[w])
			}
			if s < len(spaces) {
				justified.WriteString(spaces[s])
			}
		}
		line = justified.String()
		output.WriteString(line)
		output.WriteRune('\n')
	}
	return output.String()
}

func Longest(lines []string) int {
	var longest int
	for _, v := range lines {
		l := len(v)
		if l > longest {
			longest = l
		}
	}
	return longest
}

func reverse(text *string) {
	runes := []rune(*text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*text = string(runes)
}
