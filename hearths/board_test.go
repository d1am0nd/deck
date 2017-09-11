package hearths

import (
	"cards/deck"
	"fmt"
	"testing"
)

func testBoard() Board {
	return NewBoard([4]Player{NewPlayer(), NewPlayer(), NewPlayer(), NewPlayer()})
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
	if board.finished != false {
		t.Fatal("NewBoard didnt set correct finished")
	}
	if board.phase != P1 {
		t.Fatal("NewBoard didnt set phase to 1")
	}
	if board.gamen != 0 {
		t.Fatal("NewBoard didnt set phase to 1")
	}
	if board.hearthsBroken {
		t.Fatal("NewBoard didnt set hearthsBroken to false")
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
	if !b.Player(b.p3started).Hand().Find(deck.NewCard("2c")) {
		t.Fatal("P1DealAll didnt set correct p3started")
	}
}

func TestP2Trade(t *testing.T) {
	b := testBoardP2()
	err := b.P2Trade(0, b.Player(0).Hand().Cards()[0:3])
	if err != nil {
		t.Fatal("b.p2trade returned error,", err, " shouldnt")
	}
	err = b.P2Trade(1, b.Player(1).Hand().Cards()[0:3])
	if err != nil {
		t.Fatal("b.p2trade 2 returned error,", err, " shouldnt")
	}
	err = b.P2Trade(3, b.Player(3).Hand().Cards()[0:3])
	if err == nil {
		t.Fatal("b.P2Trade(wrong index) didnt return an error, should")
	}
	err = b.P2Trade(2, b.Player(3).Hand().Cards()[0:3])
	if err == nil {
		t.Fatal("b.P2Trade(wrong card) didnt return an error, should")
	}
	if b.Player(0).Hand().Count() != 10 {
		t.Fatal("b.p2trade p1 has", b.Player(0).Hand().Count(), "cards, should have 10")
	}
	if b.Player(1).Hand().Count() != 13 {
		t.Fatal("b.p2trade p2 has", b.Player(1).Hand().Count(), "cards, should have 13")
	}
	if b.Player(2).Hand().Count() != 16 {
		t.Fatal("b.p2trade p3 has", b.Player(2).Hand().Count(), "cards, should have 16")
	}

}

func TestP3PutOnPile(t *testing.T) {
	var err error
	err = nil
	b := testBoardP2()
	b.turn = P2lastTurn + 1
	b.phase = P3
	startTurn := b.turn
	nexti := b.NextPlayerI()
	var found = false
	for i, c := range b.Player(nexti).Hand().Cards() {
		err = b.P3PutOnPile(nexti, c)
		if c.Face() == "c" && c.Value() == "2" {
			if err != nil {
				t.Fatal("b.P3lastTurn(2c) returned err, shouldnt")
			}
			found = true
			break
		}
		if err == nil && (c.Face() != "c" || c.Value() != "2") {
			t.Fatal("b.P3PutOnPile 1: wrong card", c, i, "no error")
		}
	}
	if !found {
		t.Fatal("b.P3lastTurn 1 NextPlayerI didnt have 2c", b.Player(nexti).Hand().Cards())
	}
	if b.turn != startTurn+1 {
		t.Fatal("b.P3lastTurn 1 didnt increase turn", b.turn, startTurn)
	}
	for b.turn <= P3lastTurn {
		n := b.NextPlayerI()
		for _, c := range b.Player(n).Hand().Cards() {
			err := b.P3PutOnPile(n, c)
			if err == nil {
				break
			}
		}
	}
	if b.turn != P3lastTurn+1 {
		t.Fatal("b.P3 finish: finished on turn", b.turn, ", should have turn ", P3lastTurn+1)
	}
	if b.MainPile().Count() != 0 {
		t.Fatal("b.P3 finish: main pile has", b.MainPile().Count(), "cards, should have 0")
	}
	if b.Player(0).Hand().Count() != 0 {
		t.Fatal("b.P3 finish: player 0 has", b.Player(0).Hand().Count(), "cards, should have 0")
	}
	if b.Player(1).Hand().Count() != 0 {
		t.Fatal("b.P3 finish: player 1 has", b.Player(1).Hand().Count(), "cards, should have 0")
	}
	if b.Player(2).Hand().Count() != 0 {
		t.Fatal("b.P3 finish: player 2 has", b.Player(2).Hand().Count(), "cards, should have 0")
	}
	if b.Player(3).Hand().Count() != 0 {
		t.Fatal("b.P3 finish: player 3 has", b.Player(3).Hand().Count(), "cards, should have 0")
	}
	if b.Phase() != P4 {
		t.Fatal("b.P3 finish: left phase in ", b.Phase(), "should have", P4)
	}

	gcount := 0
	for _, p := range b.players {
		gcount += p.Garbage().Count()
	}
	if gcount != 52 {
		t.Fatal("b.P3 finish: left garbage count at", gcount, "should be 52")
	}
	if !b.Finished() {
		t.Fatal("b.P3 finish: finished false, should be true")
	}
	fmt.Println(b.Results())
}
