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
	p := NewPlayer(testDeck())
	if p.hand.Card(0) != testCards()[0] || p.hand.Card(1) != testCards()[1] {
		t.Fatal("NewPlayer(2 cards) doesnt have correct cards")
	}
}
