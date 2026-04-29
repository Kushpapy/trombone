package microblog


func Truncate(phrase string) string {
	newRune := []rune{}

    oldRune := []rune(phrase)

    if len(oldRune) < 5 {
        return string(oldRune)
    }

    for k, v := range oldRune {

       
        if k == 5 {
            break
        }

         newRune = append(newRune, v)
    }

    newPhrase := string(newRune)

    return newPhrase
}
