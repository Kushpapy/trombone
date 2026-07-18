package isbnverifier

import "unicode"

func IsValidISBN(isbn string) bool {
var	isbnSlice []rune

    for _,c := range isbn{
        isbnSlice =append(isbnSlice,c)
    }

    

    sum :=0
    count := 10
    num := 0
    for _, c := range isbnSlice {
  var value int

switch {
    case c == '-':
         continue 
    case c == 'X' && count == 1:
    value = 10
    case unicode.IsDigit(c):
    value = int(c - '0')
    default:
    return false
}
        
        sum += value * count
        count = count -1
        num = num + 1
    }

    return num == 10 && sum % 11 == 0
}
