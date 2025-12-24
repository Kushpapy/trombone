package pangram

import "strings"

func IsPangram(input string) bool {
	alphabets := make(map[rune]int)

    for char := 'a'; char <= 'z'; char++ {
        alphabets[char] = 0
    }

    for _, char := range strings.ToLower(input) {
        if _, exists := alphabets[char]; exists {
            alphabets[char]++
        }
    }

    for _,count := range alphabets {
        if count < 1 {
            return false
        }
    }

    return true
}
