package hearths

import (
    "cards/deck"
)

// Returns true if the card can be played on pile from hand, false otherwise
func canPlayOnPile(hand deck.Deck, card deck.Card, pile deck.Deck, hearthsBroken bool) bool {
    plen := pile.Count()
    if plen == 0 {
        return hearthsBroken || !isHearths(card)
    }
    if card.Face() == pile.Cards()[plen - 1].Face() {
        return true
    }
    if isHearths(card) && !hearthsBroken {
        return false
    }
    return !hand.HasFace(card.Face())
}

func isHearths(c deck.Card) bool {
    return c.Face() == "h"
}
