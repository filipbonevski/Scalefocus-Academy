package main

import (
	"fmt"
	"scf/week6/lecture17/cardGame"
)

func main() {

	var cards []cardGame.Card

	cards = append(cards, cardGame.Card{10, 1})
	cards = append(cards, cardGame.Card{1, 1})
	cards = append(cards, cardGame.Card{10, 3})
	cards = append(cards, cardGame.Card{10, 2})
	cards = append(cards, cardGame.Card{11, 2})
	cards = append(cards, cardGame.Card{11, 1})
	cards = append(cards, cardGame.Card{2, 2})
	cards = append(cards, cardGame.Card{5, 2})
	cards = append(cards, cardGame.Card{7, 2})
	cards = append(cards, cardGame.Card{3, 2})

	fmt.Println(cardGame.MaxCard(cards))
}
