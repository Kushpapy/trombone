package resistorcolor
import (
	"strings" // Import the strings package
)
// Colors returns the list of all colors.
func Colors() []string {
		return []string{
		"black",
		"brown",
		"red",
		"orange",
		"yellow",
		"green",
		"blue",
		"violet",
		"grey",
		"white",
	}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	newStr := strings.ToLower(color)

     
    for i,v := range Colors() {
        if newStr == v {
            return i
        } 
    }

    return -1
}
