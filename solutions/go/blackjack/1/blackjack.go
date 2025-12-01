package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "ten", "jack", "queen", "king":
		return 10
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	}
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	pc := ParseCard(card1) + ParseCard(card2)
	dc := ParseCard(dealerCard)
	switch {
	case pc == 22:
		return "P"
	case pc == 21:
		if dc == 11 || dc == 10 {
			return "S"
		}
		return "W"
	case 17 <= pc && pc <= 20:
		return "S"
	case 12 <= pc && pc <= 16:
		if dc >= 7 {
			return "H"
		}
		return "S"
	}
	return "H"
}
