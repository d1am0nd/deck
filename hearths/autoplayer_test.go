package hearths

import (
	"cards/deck"
	"testing"
)

func emptyBoard() Board {
	return NewBoard([4]Player{NewPlayer(), NewPlayer(), NewPlayer(), NewPlayer()})
}

func boardP2() Board {
	b := emptyBoard()
	b.P1DealAll()
	return b
}

func TestP2Sort(t *testing.T) {
	if p2sort(deck.NewCard("Ah"), deck.NewCard("Qc")) {
		t.Fatal("Ah is higher than Qc, shouldnt be")
	}
	if p2sort(deck.NewCard("Kh"), deck.NewCard("Qc")) {
		t.Fatal("Kh is higher than Qc, shouldnt be")
	}
	if p2sort(deck.NewCard("Th"), deck.NewCard("Ah")) {
		t.Fatal("Th is higher than Ah, shouldnt be")
	}
}

func TestPlayNextMove(t *testing.T) {
    b := boardP2()
    playPhase2player(&b, b.NextPlayerI())
/*
    for i := 0; i < 4; i++ {
        err := b.PlayNextMove()
        if err != nil {
            t.Fatal("PlayNextMove p2 returned error", err, "shouldnt")
        }
        if b.NextPlayerI() != (i + 1) % 4 {
            t.Fatal("PlayNextMove p2 i", i, "board next player is ", b.NextPlayerI(), "should be ", (i + 1) % 4)
        }
    }
*/
}
