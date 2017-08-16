package deck

import (
)

type Deck struct {
    cards []Card
}

func NewDeck(cards []Card) Deck {
    return Deck{cards: cards}
}
