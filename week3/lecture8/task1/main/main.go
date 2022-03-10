package main

import (
	"fmt"
	"scf/week3/lecture8/task1/cardDraw"
	"scf/week3/lecture8/task1/cardGame"
)

func main() {
	var newDeck cardGame.Deck = cardGame.NewDeck()
	newDeck.Shuffle()
	fmt.Println(newDeck)
	fmt.Println(cardDraw.DrawAllCards(&newDeck))
}
