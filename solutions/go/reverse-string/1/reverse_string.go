package reverse

import "strings"

func Reverse(input string) string {
    strSlice := strings.Split(input,"")
	start := 0
    end := len(strSlice) - 1

    for start < end {
        test := strSlice[start] 
        strSlice[start] = strSlice[end]
        strSlice[end] = test
        start++
        end--
    }

    return strings.Join(strSlice,"")
}
