package logs
import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, value := range log {
        switch(value) {
            case '❗':
        		return "recommendation"
        	case '🔍':
                return "search"
        	case '☀':
        		return "weather"
        }  
    }

    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    var result []rune
	for _, value := range log {
        switch(value) {
            case oldRune:
        		result = append(result,newRune)
            default:
                result = append(result,value) 
        }        
    }
    return string(result)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    numberOfRunes := utf8.RuneCountInString(log)
	if numberOfRunes <= limit {
        return true
    }else {
        return false
    }
}
