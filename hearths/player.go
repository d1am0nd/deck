package hearths

import (
    "cards/deck"
)

type Player struct {
    hand deck.Deck
}

func NewPlayer(d deck.Deck) Player {
    return Player{hand: d}
}

func (p *Player) Hand() *deck.Deck {
    return &p.hand
}
