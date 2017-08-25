package hearths

import (
	"cards/deck"
	"testing"
)

func testBoard() Board {
	return NewBoard([4]Player{NewPlayer(deck.Deck{}), NewPlayer(deck.Deck{}), NewPlayer(deck.Deck{}), NewPlayer(deck.Deck{})})
}

func testBoardP2() Board {
	b := testBoard()
	b.P1DealAll()
	return b
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
	if b.phase != P2 {
		t.Fatal("P1DealAll should leave deck in p2, didnt")
	}
}

func TestP2Trade(t *testing.T) {
	b := testBoardP2()
	err := b.P2Trade(0, b.Player(0).Hand().Card(0))
	if err != nil {
		t.Fatal("b.p2trade returned error,", err, " shouldnt")
	}
	err = b.P2Trade(1, b.Player(1).Hand().Card(0))
	if err != nil {
		t.Fatal("b.p2trade 2 returned error,", err, " shouldnt")
	}
	err = b.P2Trade(3, b.Player(3).Hand().Card(0))
	if err == nil {
		t.Fatal("b.P2Trade(wrong index) didnt return an error, should")
	}
	err = b.P2Trade(2, b.Player(3).Hand().Card(0))
	if err == nil {
		t.Fatal("b.P2Trade(wrong card) didnt return an error, should")
	}
	if b.Player(0).Hand().Count() != 12 {
		t.Fatal("b.p2trade p1 has", b.Player(0).Hand().Count(), "cards, should have 12")
	}
	if b.Player(1).Hand().Count() != 13 {
		t.Fatal("b.p2trade p2 has", b.Player(1).Hand().Count(), "cards, should have 13")
	}
	if b.Player(2).Hand().Count() != 14 {
		t.Fatal("b.p2trade p3 has", b.Player(2).Hand().Count(), "cards, should have 14")
	}
}
