// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob
import (
    "strings"
    "unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

    remark = strings.TrimSpace(remark)

    hasLetters := func () bool {
        for _, ch := range remark {
            if unicode.IsLetter(ch){
                return true
            }
        }
        return false
    }()

    isQuestion := len(remark) > 0 && remark[len(remark) - 1] == '?'
    isYelling := hasLetters && strings.ToUpper(remark) == remark

    
	switch  {
        case remark == "":
        return "Fine. Be that way!"
        case isQuestion && isYelling:
        return "Calm down, I know what I'm doing!"
        case isQuestion:
        return "Sure."
        case isYelling:
        return "Whoa, chill out!"
        default:
        return "Whatever."
    }    
	return ""
}
