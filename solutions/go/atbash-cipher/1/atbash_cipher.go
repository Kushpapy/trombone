package atbashcipher

import "strings"

func group(s string, n int) string {
	runes := []rune(s)
	result := ""
	for i, r := range runes {
		if i > 0 && i%n == 0 {
			result += " "
		}
		result += string(r)
	}
	return result
}

func Atbash(s string) string {
	first := map[rune]rune{
	'a': 'z', 'b': 'y', 'c': 'x', 'd': 'w', 'e': 'v', 'f': 'u', 
	'g': 't', 'h': 's', 'i': 'r', 'j': 'q', 'k': 'p', 'l': 'o', 
	'm': 'n', 'n': 'm', 'o': 'l', 'p': 'k', 'q': 'j', 'r': 'i', 
	's': 'h', 't': 'g', 'u': 'f', 'v': 'e', 'w': 'd', 'x': 'c', 
	'y': 'b', 'z': 'a', '0': '0', '1': '1', '2': '2', '3': '3', '4': '4',
		'5': '5', '6': '6', '7': '7', '8': '8', '9': '9',
}


    cleanedStr := strings.ReplaceAll(strings.ToLower(s), " ", "")
	newstr := ""
    for _,v := range cleanedStr {
        val, exists := first[v]
        if exists {
            newstr += string(val)
        }
    }

    return group(newstr, 5)
}
