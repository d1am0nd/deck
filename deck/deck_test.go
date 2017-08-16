package deck

import (
    "testing"
)

func TestNewDeck(t *testing.T) {
    cards := []Card{
        Card{face: "H", value: "2"},
        Card{face: "H", value: "3"} }

    deck := NewDeck(cards)
    if len(deck.cards) != 2 {
        t.Fatal("NewDeck(2 cards) len = ", len(deck.cards), " expected 2")
    }

    deck = NewDeck([]Card{})
    if len(deck.cards) != 0 {
        t.Fatal("NewDeck(0 cards) len = ", len(deck.cards), " expected 0")
    }
}
