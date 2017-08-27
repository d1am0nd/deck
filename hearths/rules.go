package hearths

import (
    "cards/deck"
)

// Returns true if the card can be played on pile from hand, false otherwise
func canPlayOnPile(hand deck.Deck, card deck.Card, pile deck.Deck, hearthsBroken bool) error {
    if !hearthsBroken && isHearths(card) && hasOtherThanHearths(hand) {
        return newErr("hearths havent been broken yet")
    }
    plen := pile.Count()
    if plen == 0 {
        return nil
    }
    if card.Face() == pile.Cards()[plen - 1].Face() {
        return nil
    }
    if hand.HasFace(pile.Cards()[plen - 1].Face()) {
        return newErr("there wrong face")
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
