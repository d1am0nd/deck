package hearths

import (
    "cards/deck"
)

type Board struct {
    players [4]Player
    deck deck.Deck
    mainPile deck.Deck
    discardPile deck.Deck
    phase int
}

const P1 = 1
const P2 = 2

func NewBoard(players [4]Player) Board {
    return Board{
        players: players,
        deck: deck.NewDefaultDeck(),
        mainPile: deck.NewDeck([]deck.Card{}),
        discardPile: deck.NewDeck([]deck.Card{}),
        phase: 1 }
}

func (b *Board) P1ShuffleDeck() error {
    if b.phase != P1 {
        newErr("Wrong phase")
    }
    b.deck.Shuffle()
    return nil
}

func (b *Board) P1DealAll() error {
    if b.phase != P1 {
        newErr("Wrong phase")
    }
    for i := 0; i < 52; i++ {
        hand := b.players[i % 4].Hand()
        hand.PutOnTop(b.deck.Draw())
    }
    return nil
}
