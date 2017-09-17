package hearths

import (
	"cards/deck"
	"testing"
)

func emptyBoard() Board {
	return NewBoard([4]Player{NewPlayer(), NewPlayer(), NewPlayer(), NewPlayer()})
}

func testBoardP2ap() Board {
	b := emptyBoard()
	b.P1DealAll()
	return b
}

func testBoardP3ap() *Board {
	b := testBoardP2ap()
	for i := 0; i < 4; i++ {
		PlayPhase2player(&b, b.NextPlayerI())
	}
    return &b
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

func TestPlayPhase2player(t *testing.T) {
	b := testBoardP2ap()
	for i := 0; i < 4; i++ {
		err := PlayPhase2player(&b, b.NextPlayerI())
		if err != nil {
			t.Fatal("playPhase2player", i, ", returned err", err, "shouldnt")
		}
		if i == 3 {
			if b.Player(b.NextPlayerI()).Hand().Count() != b.Player((b.NextPlayerI()+3)%4).Hand().Count() {
				t.Fatal("playPhase2player", i, ", p1 should have same cards as next, doesnt. p1:", b.Player((b.NextPlayerI()+3)%4).Hand().Count(), "p2", b.Player(b.NextPlayerI()).Hand().Count())
			}
		} else if i == 0 {
			if b.Player(b.NextPlayerI()).Hand().Count() != b.Player((b.NextPlayerI()+3)%4).Hand().Count()+6 {
				t.Fatal("playPhase2player", i, ", p1 should have 6 less cards than p2, doesnt. p1:", b.Player((b.NextPlayerI()+3)%4).Hand().Count(), "p2", b.Player(b.NextPlayerI()).Hand().Count())
			}
		} else {
			if b.Player(b.NextPlayerI()).Hand().Count() != b.Player((b.NextPlayerI()+3)%4).Hand().Count()+3 {
				t.Fatal("playPhase2player", i, ", p1 should have 3 less cards than p2, doesnt. p1:", b.Player((b.NextPlayerI()+3)%4).Hand().Count(), "p2", b.Player(b.NextPlayerI()).Hand().Count())
			}
		}
	}
}

func TestPlayPhase3player(t *testing.T) {
    b := testBoardP3ap()
	for i := 0; i < 52; i++ {
		next := b.Player(b.NextPlayerI())
		scount := next.Hand().Count()
		err := PlayPhase3player(b, b.NextPlayerI())
        if err != nil {
			t.Fatal("PlayPhase3player", i, "returned err", err, "shouldnt")
		}
		if next.Hand().Count() != scount-1 {
			t.Fatal("PlayPhase3player", i, "player has", next.Hand().Count(), "cards, should have", scount-1)
		}
	}
    if b.Phase() != P4 {
        t.Fatal("PlayPhase3player P3 should have finished, didnt")
    }
    if b.results[0] + b.results[1] + b.results[2] + b.results[3] != 26 {
        t.Fatal("Wrong results", b.results[0] + b.results[1] + b.results[2] + b.results[3])
    }
}
