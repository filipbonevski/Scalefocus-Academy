package main

import (
	"fmt"
	"scf/week3/lecture9/task1/cardDraw"
	"scf/week3/lecture9/task1/cardGame"
)

func main() {
	var newDeck cardGame.Deck = cardGame.NewDeck()
	newDeck.Shuffle()
	// fmt.Println(newDeck.Done())
	fmt.Println(cardDraw.DrawAllCards(&newDeck))

	// if err != nil {
	// 	log.Fatalf("there was an error: %v", err())
	// } else {
	// 	fmt.Println(cardDraw.DrawAllCards(&newDeck))
	// }

}
