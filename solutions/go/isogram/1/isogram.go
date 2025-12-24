package isogram

import "strings"


func IsIsogram(word string) bool {
	var seen = make(map[rune]bool)
    
	for _, v := range strings.ToLower(word) {
        if v == ' ' || v == '-' {
            continue
        }

        if seen[v] {
            return false
        }

        seen[v] = true
    }

    return true
}
