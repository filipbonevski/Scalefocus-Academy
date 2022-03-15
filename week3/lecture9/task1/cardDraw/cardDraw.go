package cardDraw

import (
	"scf/week3/lecture9/task1/cardGame"
)

type dealer interface {
	Deal() (*cardGame.Card, error)
	Done() bool
}

func DrawAllCards(dealer dealer) ([]cardGame.Card, error) {
	var newDeck []cardGame.Card

	for true {
		card, err := dealer.Deal()
		// done := dealer.Done()
		// fmt.Println(done)
		if err != nil && dealer.Done() {
			// newDeck = append(newDeck, *card)
			return newDeck, nil
		} else {
			newDeck = append(newDeck, *card)
		}
	}
	return newDeck, nil
}
