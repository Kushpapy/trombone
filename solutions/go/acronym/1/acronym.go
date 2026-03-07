// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym
import (
    "strings"
    "unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	result := strings.TrimSpace(s)
	runes := []rune(result)
    acronym := string(runes[0])

    for i := 0; i < len(runes); i++{
        if runes[i] == ' ' || runes[i] == '-' || runes[i] == '_'{

            for i+1 < len(runes) && !unicode.IsLetter(runes[i+1]){
                i++
            }
            
            if i+1 < len(runes){
                acronym += string(runes[i + 1])
            }
        }
    }
    

        



    
	return strings.ToUpper(acronym)
}
