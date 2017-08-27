package hearths

import (
    "testing"
    "cards/deck"
)

func testHand() deck.Deck {
    return deck.NewDeck([]deck.Card{
        deck.NewCard("Ac"), deck.NewCard("Qh"), deck.NewCard("5c") })
}

func testPile() deck.Deck {
    return deck.NewDeck([]deck.Card{
        deck.NewCard("8c")})
}

func testPile2() deck.Deck {
    return deck.NewDeck([]deck.Card{})
}

func TestCanPlayOnPil(t *testing.T) {
    p := testPile()
    h := testHand()
    if !canPlayOnPile(h, deck.NewCard("5c"), p, true) {
        t.Fatal("canPlayOnPile 1: returned false, shouldnt")
    }
    if canPlayOnPile(h, deck.NewCard("Qh"), p, true) {
        t.Fatal("canPlayOnPile 2: returned true, shouldnt")
    }
    h = testPile2()
    if !canPlayOnPile(h, deck.NewCard("Qh"), p, true) {
        t.Fatal("canPlayOnPile 3: returend false, shouldnt")
    }
    if canPlayOnPile(h, deck.NewCard("Qh"), p, false) {
        t.Fatal("canPlayOnPile 4: returend false, shouldnt")
    }
}
