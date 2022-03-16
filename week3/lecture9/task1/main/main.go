package main

import (
	"fmt"
	"log"
	"scf/week3/lecture9/task1/cardDraw"
	"scf/week3/lecture9/task1/cardGame"
)

func main() {
	var newDeck cardGame.Deck = cardGame.NewDeck()
	newDeck.Shuffle()

	deck, err := cardDraw.DrawAllCards(&newDeck)
	if err != nil {
		log.Fatal("there was an error")
	} else {
		fmt.Println(deck)
	}
}
