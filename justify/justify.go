package typography

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/patrickbucher/typography/utils"
)

// Justify aligns the text to the left and right edge by filling the word gaps
// up with spaces until length is reached.
func Justify(text string, length int) string {
	output := bytes.NewBufferString("")
	reader := strings.NewReader(text)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	switchFillDirection := false
	for scanner.Scan() {
		line := scanner.Text()
		nRunes := len([]rune(line))
		if nRunes > length {
			// TODO client's problem?
			continue
		}
		words := strings.Fields(line)
		var gaps int
		if len(words) <= 1 {
			gaps = 1
		} else {
			gaps = len(words) - 1
		}
		spaces := make([]string, gaps)
		diff := length - nRunes + len(spaces) // spaces where omitted, add them back
		for d, i := diff, 0; d > 0; d-- {
			spaces[i] += " "
			i++
			if i == len(spaces) {
				i = 0
			}
		}
		// alternate fill direction
		if switchFillDirection {
			reverse(&spaces)
			switchFillDirection = false
		} else {
			switchFillDirection = true
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
	// left-align last line
	out := strings.TrimSpace(output.String())
	from := strings.LastIndex(out, "\n")
	to := len(out)
	return out[:from] + utils.SquashSpaces(out[from:to])
}

func reverse(s *[]string) {
	strings := *s
	for i, j := 0, len(strings)-1; i < j; i, j = i+1, j-1 {
		strings[i], strings[j] = strings[j], strings[i]
	}
	*s = strings
}
