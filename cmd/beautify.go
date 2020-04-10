// Beautify uses the typography package to beautify plain text.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/patrickbucher/typography/beautify"
)

func main() {
	e := flag.Bool("e", false, "English style: “”")
	d := flag.Bool("d", false, "German style: „“")
	g := flag.Bool("g", false, "Guillemets: «»")
	r := flag.Bool("r", false, "Reverse Guillemets: »«")
	flag.Parse()
	var style beautify.QuoteStyle
	if *e {
		style = beautify.English
	} else if *d {
		style = beautify.German
	} else if *g {
		style = beautify.Guillemets
	} else if *r {
		style = beautify.ReverseGuillemets
	}
	if style == 0 {
		// default: Guillemets
		style = beautify.Guillemets
	}
	// FIXME: line by line processing is a bad idea, because code sections
	// delimited with ``` won't be processed properly. Use io/ioutil.ReadFile
	// until a better idea comes up.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(beautify.Beautify(scanner.Text(), style))
	}
}
