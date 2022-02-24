package main

import "fmt"

type Suite = int

const (
	clubs Suite = iota
	diamonds
	hearts
	spades
)

func main() {
	var cardOneVal int
	var cardOneSuit int
	var cardTwoVal int
	var cardTwoSuit int

	fmt.Println(clubs)

	fmt.Println("Enter card one value(2-13)")
	fmt.Scanf("%d\n", &cardOneVal)

	fmt.Println("Enter card one suit(0-3)")
	fmt.Scanf("%d\n", &cardOneSuit)

	fmt.Println("Enter card two value(2-13)")
	fmt.Scanf("%d\n", &cardTwoVal)

	fmt.Println("Enter card two suit(0-3)")
	fmt.Scanf("%d\n", &cardTwoSuit)

	if cardOneVal > 1 && cardOneVal < 15 && cardTwoVal > 1 && cardTwoVal < 15 && cardOneSuit > 0 && cardOneSuit < 5 && cardTwoSuit > 0 && cardTwoSuit < 5 {
		fmt.Println(compareCards(cardOneVal, cardOneSuit, cardTwoVal, cardTwoSuit))
	} else {
		fmt.Println("Enter proper values")
	}
}

func compareCards(cardOneVal int, cardOneSuit int, cardTwoVal int, cardTwoSuit int) int {
	if (cardOneVal > cardTwoVal) || (cardOneVal == cardTwoVal && cardOneSuit > cardTwoSuit) {
		return 1
	} else if (cardTwoVal > cardOneVal) || (cardOneVal == cardTwoVal && cardTwoSuit > cardOneSuit) {
		return -1
	}

	return 0
}
