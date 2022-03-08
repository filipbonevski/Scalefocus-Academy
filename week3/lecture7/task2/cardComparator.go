package main

import "fmt"

type Card struct {
	value int
	suit  int
}

type CardComparator func(cardOne Card, cardTwo Card) int

func (comparatorFunc CardComparator) compareCards(cardOne Card, cardTwo Card) int {
	if (cardOne.value > cardTwo.value) || (cardOne.value == cardTwo.value && cardOne.suit > cardTwo.suit) {
		return 1
	} else if (cardTwo.value > cardOne.value) || (cardOne.value == cardTwo.value && cardTwo.suit > cardOne.suit) {
		return -1
	}
	return 0
}

func compareCards(cardOne Card, cardTwo Card) int {
	if (cardOne.value > cardTwo.value) || (cardOne.value == cardTwo.value && cardOne.suit > cardTwo.suit) {
		return 1
	} else if (cardTwo.value > cardOne.value) || (cardOne.value == cardTwo.value && cardTwo.suit > cardOne.suit) {
		return -1
	}
	return 0
}

func maxCard(cards []Card, comparatorFunc CardComparator) Card {
	var maxCard Card
	maxCard = Card{1, 1}

	for _, card := range cards {
		if comparatorFunc.compareCards(card, maxCard) == 1 {
			maxCard = card
		}
	}
	return maxCard
}

func main() {
	var cards []Card

	cards = append(cards, Card{10, 1})
	cards = append(cards, Card{6, 1})
	cards = append(cards, Card{12, 3})
	cards = append(cards, Card{12, 2})
	cards = append(cards, Card{11, 2})

	fmt.Println(maxCard(cards, compareCards))

	fmt.Println(maxCard(cards, func(cardOne Card, cardTwo Card) int {
		return 0
	}))
}
