package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"typography"
)

func main() {
	bytes, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)
	text = typography.Beautify(text, typography.German)
	text = typography.SimpleFold(text, 45)
	text = typography.Justify(text, typography.LongestLine(text))
	fmt.Println(text)
}
