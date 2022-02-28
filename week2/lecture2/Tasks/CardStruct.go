package main

import "fmt"

type Card struct {
	value int
	suit  int
}

func main() {
	var cardOne Card
	var cardTwo Card

	fmt.Println("Enter card one value(2-14)")
	fmt.Scanf("%d\n", &cardOne.value)

	fmt.Println("Enter card one suit(1-4)")
	fmt.Scanf("%d\n", &cardOne.suit)

	fmt.Println("Enter card two value(2-14)")
	fmt.Scanf("%d\n", &cardTwo.value)

	fmt.Println("Enter card two suit(1-4)")
	fmt.Scanf("%d\n", &cardTwo.suit)

	if cardOne.value > 1 && cardOne.value < 15 && cardTwo.value > 1 && cardTwo.value < 15 && cardOne.suit > 0 && cardOne.suit < 5 && cardTwo.suit > 0 && cardTwo.suit < 5 {
		fmt.Println(compareCards(cardOne, cardTwo))
	} else {
		fmt.Println("Enter proper values")
	}

	var cards []Card

	cards = append(cards, Card{10, 1})
	cards = append(cards, Card{1, 1})
	cards = append(cards, Card{10, 3})
	cards = append(cards, Card{10, 2})
	cards = append(cards, Card{11, 2})
	cards = append(cards, Card{11, 4})
	cards = append(cards, Card{2, 2})
	cards = append(cards, Card{5, 2})
	cards = append(cards, Card{7, 2})
	cards = append(cards, Card{3, 2})

	fmt.Println(maxCard(cards))
}

func compareCards(cardOne Card, cardTwo Card) int {
	if (cardOne.value > cardTwo.value) || (cardOne.value == cardTwo.value && cardOne.suit > cardTwo.suit) {
		return 1
	} else if (cardTwo.value > cardOne.value) || (cardOne.value == cardTwo.value && cardTwo.suit > cardOne.suit) {
		return -1
	}

	return 0
}

func maxCard(cards []Card) Card {
	var maxCard Card
	maxCard = Card{1, 1}

	for _, card := range cards {
		if compareCards(card, maxCard) == 1 {
			maxCard = card
		}
	}
	return maxCard
}
