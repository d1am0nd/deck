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

func playPhase2player(b *Board, next int) error {
	if next != b.NextPlayerI() {
		return newErr("Wrong player")
	}
	b.Player(next).Hand().Sort(p2sort)
	drawn, err := b.Player(next).Hand().DrawCards(4)
	if err != nil {
		return err
	}
	b.Player((next + 1) % 4).Hand().PutManyOnTop(drawn)
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

func sortCards(bigger func(deck.Card, deck.Card) bool, cards []deck.Card) []deck.Card {
	if len(cards) < 2 {
		return cards
	}
	pivoti := len(cards) / 2
	left, right := 0, len(cards)-1
	pivot := cards[pivoti]

	cards[pivoti], cards[right] = cards[right], cards[pivoti]
	for i := 0; i < right; i++ {
		if bigger(cards[i], pivot) {
			cards[left], cards[i] = cards[i], cards[left]
			left++
		}
	}

	cards[left], cards[right] = cards[right], cards[left]
	sortCards(bigger, cards[:left])
	sortCards(bigger, cards[left+1:])

	return cards
}
