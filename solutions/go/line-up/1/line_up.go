package lineup

import (
    "fmt"
    "strconv"
)

func Format(name string, number int) string {
	suffix := "th"

    lastTwo := number % 100
    lastDigit := number % 10 

    if lastTwo != 11 && lastTwo != 12 && lastTwo != 13 {
        switch lastDigit {
            case 1:
            suffix = "st"
            case 2:
            suffix = "nd"
            case 3:
            suffix = "rd"
        }
    }

    ordinal := strconv.Itoa(number) + suffix

    return fmt.Sprintf("%s, you are the %s customer we serve today. Thank you!", name, ordinal)
}
