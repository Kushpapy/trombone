package luhn

import (
    "strings"
    "unicode"
)

func Valid(id string) bool {
    // Remove all spaces
    id = strings.ReplaceAll(id, " ", "")

    // Must be more than 1 character
    if len(id) <= 1 {
        return false
    }

    // Extract digits and ensure no invalid characters
    digits := []int{}
    for _, r := range id {
        if !unicode.IsDigit(r) {
            return false
        }
        digits = append(digits, int(r-'0'))
    }

    // Apply Luhn algorithm
    sum := 0
    double := false

    // Traverse digits from right to left
    for i := len(digits) - 1; i >= 0; i-- {
        d := digits[i]
        if double {
            d *= 2
            if d > 9 {
                d -= 9
            }
        }
        sum += d
        double = !double
    }

    return sum%10 == 0
}
