package cardGame

import (
	"testing"
)

func TestCompareCards(t *testing.T) {
	firstCard := Card{11, 2}
	secondCard := Card{12, 3}

	ans := compareCards(firstCard, secondCard)
	if ans != -1 {
		t.Errorf("expected -1, got %d", ans)
	}

	ans = compareCards(firstCard, firstCard)
	if ans != 0 {
		t.Errorf("expected 0, got %d", ans)
	}

	ans = compareCards(secondCard, firstCard)
	if ans != 1 {
		t.Errorf("expected 1, got %d", ans)
	}
}

func TestMaxCard(t *testing.T) {
	var cards []Card

	cards = append(cards, Card{10, 1})
	cards = append(cards, Card{1, 1})
	cards = append(cards, Card{10, 3})
	cards = append(cards, Card{14, 2})
	cards = append(cards, Card{11, 2})

	var tests = []struct {
		input    []Card
		expected Card
	}{
		{cards, Card{14, 2}},
	}
	for _, test := range tests {
		ans := MaxCard(cards)
		if ans != test.expected {
			t.Errorf("input %d, expected %d, got %d", test.input, test.expected, ans)
		}
	}
}
