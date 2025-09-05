package scrabble
import "strings"

func Score(word string) int {
	store := make(map[rune]int)

	store['A'] = 1
	store['E'] = 1
	store['I'] = 1
	store['O'] = 1
	store['U'] = 1
	store['L'] = 1
	store['N'] = 1
	store['R'] = 1
	store['S'] = 1
	store['T'] = 1

	store['D'] = 2
	store['G'] = 2

	store['B'] = 3
	store['C'] = 3
	store['M'] = 3
	store['P'] = 3

	store['F'] = 4
	store['H'] = 4
	store['V'] = 4
	store['W'] = 4
	store['Y'] = 4

	store['K'] = 5

	store['J'] = 8
	store['X'] = 8

	store['Q'] = 10
	store['Z'] = 10

	text := strings.ToUpper(word)

	var count int

	for _, v := range text {
		count += store[v]
	}

	return count
}
