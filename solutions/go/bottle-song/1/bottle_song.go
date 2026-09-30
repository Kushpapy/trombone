package bottlesong

import (
	"fmt"
	"strings"
)

var numberWords = []string{
	"no", "one", "two", "three", "four", "five",
	"six", "seven", "eight", "nine", "ten",
}

func bottleWord(n int) string {
	if n == 1 {
		return "bottle"
	}
	return "bottles"
}

func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func verse(n int) []string {
	opening := fmt.Sprintf("%s green %s hanging on the wall,", capitalize(numberWords[n]), bottleWord(n))
	return []string{
		opening,
		opening,
		"And if one green bottle should accidentally fall,",
		fmt.Sprintf("There'll be %s green %s hanging on the wall.", numberWords[n-1], bottleWord(n-1)),
	}
}

func Recite(startBottles, takeDown int) []string {
	var lines []string
	for i := 0; i < takeDown; i++ {
		if i > 0 {
			lines = append(lines, "")
		}
		lines = append(lines, verse(startBottles-i)...)
	}
	return lines
}