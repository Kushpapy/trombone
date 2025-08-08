package techpalace
import "strings"
import "fmt"
// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	store := strings.Repeat("*", numStarsPerLine)
    cat := store + "\n" + welcomeMsg + "\n" + store

    return cat
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	newMsg := strings.ReplaceAll(oldMsg, "*","")
    finalNewMsg := strings.TrimSpace(newMsg)

    return finalNewMsg
}
