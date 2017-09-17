package hearths

import (
	"cards/deck"
)

/*
func PlayNextMove(b *Board) error {
    next := b.NextPlayerI()
    phase := b.Phase()
}
*/

func PlayPhase2player(b *Board, next int) error {
	if next != b.NextPlayerI() {
		return newErr("Wrong player")
	}
	b.Player(next).Hand().Sort(p2sort)
	drawn := b.Player(next).Hand().Cards()[:3]
	err := b.P2Trade(next, drawn)
	if err != nil {
		return err
	}
	return nil
}

func p2sort(c1 deck.Card, c2 deck.Card) bool {
	if c1 == deck.NewCard("Qc") {
		return true
	} else if c2 == deck.NewCard("Qc") {
		return false
	}
	if c1.Value() == "A" {
		return true
	} else if c2.Value() == "A" {
		return false
	}
	return c1.Comparable() >= c2.Comparable()
}

func PlayPhase3player(b *Board, next int) error {
	if next != b.NextPlayerI() {
		return newErr("Wrong player")
	}
	if b.Turn() == P2lastTurn + 1 {
		return b.P3PutOnPile(next, deck.NewCard("2c"))
	}
	if b.Player(next).Hand().Count() == 0 {
		b.Player(next).Hand().Sort(p3sortEmpty)
	} else {
		b.Player(next).Hand().Sort(p3sort)
	}
	var err error
	err = nil
	for _, c := range b.Player(next).Hand().Cards() {
		err = b.P3PutOnPile(next, c)
		if err == nil {
			return err
		}
	}
	return err
}

func p3sort(c1 deck.Card, c2 deck.Card) bool {
	if c1.Value() == "A" {
		return false
	} else if c2.Value() == "A" {
		return true
	}
	return c1.Comparable() <= c2.Comparable()
}

func p3sortEmpty(c1 deck.Card, c2 deck.Card) bool {
	if c1.Value() == "A" {
		return true
	} else if c2.Value() == "A" {
		return false
	}
	return c1.Comparable() >= c2.Comparable()
}
