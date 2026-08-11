package anagram

import (
    "slices"
    "strings"
    "unicode/utf8"
)

func sortWord(word string) string {
    runes := []rune(strings.ToLower(word))

    slices.Sort(runes)

    return string(runes)
}

func Detect(subject string, candidates []string) []string {
    result := []string{}
	sortedWord := sortWord(subject)
    lowerTarget := strings.ToLower(subject)

    for _, candidate := range candidates {
        lowerCandidate := strings.ToLower(candidate)

        if lowerTarget == lowerCandidate {
            continue
        }

        if utf8.RuneCountInString(subject) != utf8.RuneCountInString(lowerCandidate){
            continue
        }

        if sortedWord == sortWord(candidate){
            result = append(result, candidate)
        }
    }

    return result
}
