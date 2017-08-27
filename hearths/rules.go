package hearths

import (
	"cards/deck"
)

// Returns true if the card can be played on pile from hand, false otherwise
func canPlayOnPile(hand deck.Deck, card deck.Card, pile deck.Deck, hearthsBroken bool) error {
	plen := pile.Count()
	if plen == 4 {
		return newErr("pile full")
	}
	if !hearthsBroken && isHearths(card) && hasOtherThanHearths(hand) {
		return newErr("hearths havent been broken yet")
	}
	if plen == 0 {
		return nil
	}
	if card.Face() == pile.Cards()[plen-1].Face() {
		return nil
	}
	if hand.HasFace(pile.Cards()[plen-1].Face()) {
		return newErr("there wrong face")
	}
	if !hand.Find(card) {
		return newErr("hand doesnt have the card")
	}
	return nil
}

func hasOtherThanHearths(d deck.Deck) bool {
	for _, c := range d.Cards() {
		if !isHearths(c) {
			return true
		}
	}
	return false
}

func isHearths(c deck.Card) bool {
	return c.Face() == "h"
}

func winnerI(d deck.Deck) int {
	if d.Count() != 4 {
		return -1
	}
	r := 3
	face := d.Cards()[d.Count() - 1].Face()
	for i, c := range d.Cards() {
		c1 := c.Comparable()
		c2 := d.Cards()[r].Comparable()
		if c1 == 1 {
			c1 = 15
		}
		if c2 == 1 {
			c2 = 15
		}
		if c.Face() == face && c1 > c2 {
			r = i
		}
	}
	return r
}
