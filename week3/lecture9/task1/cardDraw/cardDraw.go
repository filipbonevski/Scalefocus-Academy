package cardDraw

import (
	"scf/week3/lecture9/task1/cardGame"
)

type dealer interface {
	Deal() *cardGame.Card, error
	Done() bool
}

func DrawAllCards(dealer dealer) []cardGame.Card {
	var newDeck []cardGame.Card

	for true {
		card, err := dealer.Deal()
		// done := dealer.Done()
		// fmt.Println(done)
		if card != nil {
			newDeck = append(newDeck, *card)
		} else {
			break
		}
	}
	return newDeck
}
