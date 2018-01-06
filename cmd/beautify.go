// Beautify uses the typography package to beautify plain text.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"typography"
)

func main() {
	e := flag.Bool("e", false, "English style: “”")
	d := flag.Bool("d", false, "German style: „“")
	g := flag.Bool("g", false, "Guillemets: «»")
	r := flag.Bool("r", false, "Reverse Guillemets: »«")
	flag.Parse()
	var style typography.QuoteStyle
	if *e {
		style = typography.English
	} else if *d {
		style = typography.German
	} else if *g {
		style = typography.Guillemets
	} else if *r {
		style = typography.ReverseGuillemets
	}
	if style == 0 {
		// default: Guillemets
		style = typography.Guillemets
	}
	// FIXME: line by line processing is a bad idea, because code sections
	// delimited with ``` won't be processed properly. Use io/ioutil.ReadFile
	// until a better idea comes up.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(typography.Beautify(scanner.Text(), style))
	}
}
