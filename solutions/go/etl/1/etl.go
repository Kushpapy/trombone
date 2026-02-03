package etl
import "strings"

func Transform(in map[int][]string) map[string]int {
	newDataStr := make(map[string]int)

    for key, values := range in {
        for _, value := range values {
            newDataStr[strings.ToLower(value)] = key
        }
    }

    return newDataStr
}
