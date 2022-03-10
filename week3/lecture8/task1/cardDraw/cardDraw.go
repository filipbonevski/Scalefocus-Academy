package cardDraw

import (
	"scf/week3/lecture8/task1/cardGame"
)

type dealer interface {
	Deal() *cardGame.Card
}

func DrawAllCards(dealer dealer) []cardGame.Card {
	var newDeck []cardGame.Card

	for true {
		card := dealer.Deal()
		if card != nil {
			newDeck = append(newDeck, *card)
		} else {
			break
		}
	}
	return newDeck
}
