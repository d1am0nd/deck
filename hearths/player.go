package hearths

import (
	"cards/deck"
)

type Player struct {
	hand    deck.Deck
	garbage deck.Deck
}

func NewPlayer() Player {
	return Player{hand: deck.NewDeck([]deck.Card{}), garbage: deck.NewDeck([]deck.Card{})}
}

func (p *Player) Hand() *deck.Deck {
	return &p.hand
}

func (p *Player) Garbage() *deck.Deck {
	return &p.garbage
}
