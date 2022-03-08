package main

import (
	"fmt"
	"math/rand"
)

type Card struct {
	value int
	suit  int
}

type Deck struct {
	deck   []Card
	length int
}

func newDeck() *Deck {
	var deck []Card

	for i := 2; i < 15; i++ {
		for y := 1; y < 5; y++ {
			card := Card{
				value: i,
				suit:  y,
			}
			deck = append(deck, card)
		}
	}

	return &Deck{deck, 52}
}

func (d *Deck) shuffle() *Deck {
	for i := 0; i < d.length; i++ {
		r := rand.Intn(i + 1)
		if i != r {
			d.deck[r], d.deck[i] = d.deck[i], d.deck[r]
		}
	}

	return d
}

func (d *Deck) deal() *Card {
	if len(d.deck) != 0 {
		firstCard := d.deck[0]
		d.deck = d.deck[1:]

		return &firstCard
	} else {
		return nil
	}
}

func main() {
	newDeck := newDeck()
	// fmt.Println(newDeck)
	newDeck.shuffle()
	// fmt.Println(newDeck)
	fmt.Printf("Card dealt: %d\n", newDeck.deal())
	fmt.Printf("Card dealt: %d\n", newDeck.deal())
	fmt.Printf("Card dealt: %d\n", newDeck.deal())
	fmt.Printf("Card dealt: %d\n", newDeck.deal())

	fmt.Printf("Card left in the deck: %d\n", newDeck)
}
