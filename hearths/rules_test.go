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

func TestCanPlayOnPile(t *testing.T) {
    p := testPile()
    h := testHand()
    err := canPlayOnPile(h, deck.NewCard("5c"), p, true)
    if err != nil {
        t.Fatal("canPlayOnPile 1: returned err,", err, "shouldnt")
    }
    err = canPlayOnPile(h, deck.NewCard("Qh"), p, true)
    if err == nil {
        t.Fatal("canPlayOnPile 2: returned no error, shouldn")
    }
    p = testPile2()
    err = canPlayOnPile(h, deck.NewCard("Qh"), p, true)
    if err != nil {
        t.Fatal("canPlayOnPile 3: returned err,", err, "shouldnt")
    }
    err = canPlayOnPile(h, deck.NewCard("Qh"), p, false)
    if err == nil {
        t.Fatal("canPlayOnPile 4: returned no error, should")
    }

    h = testPile2()
    h.PutOnTop(deck.NewCard("Th"))
    err = canPlayOnPile(h, deck.NewCard("Th"), p, false)
    if err != nil {
        t.Fatal("canPlayOnPile 5: returned error", err, "shouldnt")
    }
}
