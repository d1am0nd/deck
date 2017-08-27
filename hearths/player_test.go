package hearths

import (
	"testing"

	"cards/deck"
)

func testCards() []deck.Card {
	return []deck.Card{deck.NewCard("2h"), deck.NewCard("3h")}
}
func testCards6() []deck.Card {
	return []deck.Card{deck.NewCard("2h"), deck.NewCard("3h"), deck.NewCard("As"),
		deck.NewCard("Ac"), deck.NewCard("Td"), deck.NewCard("9d")}
}

func testDeck() deck.Deck {
	return deck.NewDeck(testCards())
}

func TestNewPlayer(t *testing.T) {
	p := NewPlayer()
	if p.hand.Count() != 0 || p.garbage.Count() != 0 {
		t.Fatal("NewPlayer() doesnt have 0 cards, should have")
	}
}
