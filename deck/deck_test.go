package deck

import (
    "testing"
)

func testCards() []Card {
    return []Card{
        Card{face: "H", value: "2"},
        Card{face: "H", value: "3"} }
}

func testDeck() Deck {
    return NewDeck(testCards())
}

func testEqCard(a, b []Card) bool {

    if a == nil && b == nil {
        return true;
    }

    if a == nil || b == nil {
        return false;
    }

    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}

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

func TestCards(t *testing.T) {
    deck := testDeck()
    if !testEqCard(testCards(), deck.Cards()) {
        t.Fatal("testDeck.Cards() does not equal input cards")
    }
}

func TestCard(t *testing.T) {
    deck := testDeck()
    if deck.Card(0) != testCards()[0] {
        t.Fatal("testDeck.Card(0) does not equal input card 0")
    }
    if deck.Card(1) != testCards()[1] {
        t.Fatal("testDeck.Card(1) does not equal input card 1")
    }
}

func TestDraw(t *testing.T) {
    deck := testDeck()
    c := deck.Draw()
    if c != testCards()[0] {
        t.Fatal("testDeck.Draw() does not equal input card 0")
    }
    if len(deck.cards) != 1 {
        t.Fatal("testing.Draw() count is ", len(deck.cards), " should be 1")
    }
    c = deck.Draw()
    if c != testCards()[1] {
        t.Fatal("testDeck.Draw() 2 does not equal input card 1")
    }
    if len(deck.cards) != 0 {
        t.Fatal("testing.Draw() 2 count is ", len(deck.cards), " should be 0")
    }
}
