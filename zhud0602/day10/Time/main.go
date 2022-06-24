// Golang program to illustrate
// the strings.TrimFunc() Function
package main

// importing fmt, unicode and strings
import (
	"fmt"
	"strings"
	"unicode"
)

// Calling Main method
func main() {

	// Here we have a string. The function
	// returns true for the letters
	// and all other will trim out
	// from the string
	fmt.Print(strings.TrimFunc("77 GeeksForGeeks!!!", func(r rune) bool {
		return !unicode.IsLetter(r)
	}))
}
