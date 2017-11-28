// TODO: document
package main

import (
	"github.com/idagio/poethe/poethe"
	"flag"
	"strings"
	"fmt"
)
var in string
var gen string

//.... -in="#fffe32, #65de43"
func init() {
	flag.StringVar(&in, "in", "", "type your hex color")
	flag.StringVar(&gen, "gen", "", "type your color generator")
}

// function
// args = [string]
// convert to poethe colors
// -in ""
// -gen ""
func main () {
	flag.Parse()
	colorAsStrings:= strings.Split(in, ", ")
	poethe.Parse("#001000")
	poethe.Parse("#ff4")
	var colorsList []poethe.Color
	for _, color := range colorAsStrings {
		newColor, _ := poethe.Parse(color)
		colorsList = append(colorsList, newColor)
	}
	fmt.Println(colorsList)
}