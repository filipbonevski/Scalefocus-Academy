package cardGame

type Card struct {
	Value int
	Suit  int
}

func compareCards(cardOne Card, cardTwo Card) int {
	if (cardOne.Value > cardTwo.Value) || (cardOne.Value == cardTwo.Value && cardOne.Suit > cardTwo.Suit) {
		return 1
	} else if (cardTwo.Value > cardOne.Value) || (cardOne.Value == cardTwo.Value && cardTwo.Suit > cardOne.Suit) {
		return -1
	}

	return 0
}

func MaxCard(cards []Card) Card {
	var maxCard Card
	maxCard = Card{1, 1}

	for _, card := range cards {
		if compareCards(card, maxCard) == 1 {
			maxCard = card
		}
	}
	return maxCard
}
