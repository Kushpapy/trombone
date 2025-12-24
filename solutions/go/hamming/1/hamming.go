package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
 //num :=	strings.Compare(a, b)

 firstStr := []rune(a) 
 secondStr := []rune(b)

 if len(firstStr) != len(secondStr) {
	return 0, errors.New("Length of string is not equal")
 }

 //t1 := 0;
 t2 := 0
 count := 0

 for _, v := range firstStr {
	if v != secondStr[t2]{
		count++;
	}
	t2++
 }
 
 return count,nil
}
