package microblog


func Truncate(phrase string) string {
	newRune := []rune{}

    oldRune := []rune(phrase)

    for k, v := range oldRune {

       
        if k == 5 {
            break
        }

         newRune = append(newRune, v)
    }

    newPhrase := string(newRune)

    return newPhrase
}
