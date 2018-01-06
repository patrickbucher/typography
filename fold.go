package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(fold("Das ist ein Test.", 10))
	fmt.Println(fold("abc def ghi jkl mno.", 5))
}

func fold(text string, maxLineLength int) (string, error) {
	if maxLineLength < 1 {
		return text, errors.New("maxLineLength must be at least 1")
	}
	text = strings.Replace(text, "\n", " ", -1)
	text = strings.Replace(text, "\r", " ", -1)
	text = strings.TrimSpace(text)
	runes := []rune(text)
	size := len(runes)
	buf := bytes.NewBufferString("")
	lineBuf := make([]rune, maxLineLength)
	var overflow []rune
	var last rune
	var lastSpace, w int
	for i := 0; i < size; i++ {
		c := runes[i]
		if unicode.IsSpace(c) {
			if !unicode.IsSpace(last) {
				lastSpace = i
			} else {
				continue
			}
		}
		if w < maxLineLength {
			lineBuf[w] = c
			w++
		}
		if w == maxLineLength {
			lineEnd := lastSpace + 1
			if lineEnd > len(lineBuf) {
				buf.WriteString(string(lineBuf))
				lineBuf = make([]rune, maxLineLength)
			} else {
				buf.WriteString(string(lineBuf[:lineEnd]) + "\n")
				bufEnd := w
				overflow = lineBuf[lineEnd:bufEnd]
				lineBuf = make([]rune, maxLineLength)
				for j := 0; j < len(overflow); j++ {
					lineBuf[j] = overflow[j]
				}
				w = len(overflow)
			}
		}
	}
	rest := string(lineBuf[:w])
	if len(rest) > 0 {
		buf.WriteString(rest + "\n")
	}
	return buf.String(), nil
}
