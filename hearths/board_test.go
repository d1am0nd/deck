package hearths

import (
    "cards/deck"
    "testing"
)

func testBoard() Board {
    return NewBoard([4]Player{NewPlayer(deck.Deck{}), NewPlayer(deck.Deck{}), NewPlayer(deck.Deck{}), NewPlayer(deck.Deck{})})
}

func TestNewBoard(t *testing.T) {
    board := NewBoard([4]Player{})

    if board.deck.Count() != 52 {
        t.Fatal("NewBoard didnt set correct default deck")
    }

    if board.mainPile.Count() != 0 {
        t.Fatal("NewBoard didnt set correct mainPile")
    }

    if board.discardPile.Count() != 0 {
        t.Fatal("NewBoard didnt set correct discardPile")
    }
    if board.phase != P1 {
        t.Fatal("NewBoard didnt set phase to 1")
    }
}

func TestDealAll(t *testing.T) {
    b := testBoard()
    b.P1DealAll()
    for i := 0; i < 4; i++ {
        if b.players[i].hand.Count() != 52/4 {
            t.Fatal("P1DealAll: player", i, "has", b.players[i].hand.Count(), "cards, should have", 52/4)
        }
    }
    if b.deck.Count() != 0 {
        t.Fatal("P1DealAll: deck has", b.deck.Count(), "cards, should have 0")
    }
}
