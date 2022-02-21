package main

import "fmt"

func main() {
	var cardOneVal int
	var cardOneSuit int
	var cardTwoVal int
	var cardTwoSuit int

	fmt.Println("Enter card one value")
	fmt.Scanf("%d\n", &cardOneVal)

	fmt.Println("Enter card one suit")
	fmt.Scanf("%d\n", &cardOneSuit)

	fmt.Println("Enter card two value")
	fmt.Scanf("%d\n", &cardTwoVal)

	fmt.Println("Enter card two suit")
	fmt.Scanf("%d\n", &cardTwoSuit)

	fmt.Println(compareCards(cardOneVal, cardOneSuit, cardTwoVal, cardTwoSuit))
}

func compareCards(cardOneVal int, cardOneSuit int, cardTwoVal int, cardTwoSuit int) int {
	if cardOneVal > cardTwoVal {
		return 1
	} else if cardTwoVal > cardOneVal {
		return -1
	} else if cardTwoVal == cardOneVal {
		if cardOneSuit > cardTwoSuit {
			return 1
		} else if cardTwoSuit > cardOneSuit {
			return -1
		}
	}
	return 0
}
