package cardGame

import (
	"errors"
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

func NewDeck() Deck {
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

	return Deck{deck, 52}
}

func (d *Deck) Shuffle() *Deck {
	for i := 0; i < d.length; i++ {
		r := rand.Intn(i + 1)
		if i != r {
			d.deck[r], d.deck[i] = d.deck[i], d.deck[r]
		}
	}

	return d
}

func (d *Deck) Done() bool {
	return len(d.deck) == 0
	// return false
}

func (d *Deck) Deal() (*Card, error) {
	if len(d.deck) != 0 {
		// if !d.Done() {
		firstCard := d.deck[0]
		d.deck = d.deck[1:]

		return &firstCard, nil
	} else {
		return nil, errors.New("The deck is empty")
	}
}
