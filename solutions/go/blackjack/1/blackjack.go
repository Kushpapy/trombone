package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
case "ace":
    return 11
case "eight":
    return 8
case "two":
    return 2
case "nine":
    return 9
case "three":
    return 3
case "ten":
    return 10
case "jack":
    return 10
case "queen":
    return 10
case "king":
    return 10
case "seven":
    return 7
case "four":
    return 4
case "five":
    return 5
case "six":
    return 6
default:
    return 0
}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string { 
	playerScore := ParseCard(card1) + ParseCard(card2)
    dealerScore := ParseCard(dealerCard)

    switch {
    // 1. Pair of aces → Split
    case playerScore == 22:
        return "P"

    // 2. Blackjack (21) → Win or Stand
    case playerScore == 21 && dealerScore < 10:
        return "W"
    case playerScore == 21 && dealerScore >= 10:
        return "S"

    // 3. Stand on 17–20
    case playerScore >= 17 && playerScore <= 20:
        return "S"

    // 4. 12–16 → Hit if dealer has 7+, otherwise Stand
    case playerScore >= 12 && playerScore <= 16 && dealerScore >= 7:
        return "H"
    case playerScore >= 12 && playerScore <= 16 && dealerScore < 7:
        return "S"

    // 5. ≤ 11 → Hit
    case playerScore <= 11:
        return "H"
    }

    // Fallback (just in case)
    return "S"
}
