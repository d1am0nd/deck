package hearths

import (
	"cards/deck"
)

type Board struct {
	players     [4]Player
	deck        deck.Deck
	mainPile    deck.Deck
	discardPile deck.Deck
    nextPlayer int
	phase       int
}

const P1 = 1
const P2 = 2

func NewBoard(players [4]Player) Board {
	return Board{
		players:     players,
		deck:        deck.NewDefaultDeck(),
		mainPile:    deck.NewDeck([]deck.Card{}),
		discardPile: deck.NewDeck([]deck.Card{}),
        nextPlayer: 0,
		phase:       1}
}

func (b *Board) Phase() int {
    return b.phase
}

func (b *Board) NextPlayerI() int {
    return b.nextPlayer % 4
}

func (b *Board) Player(i int) *Player {
    return &b.players[i % 4]
}

func isNext (first, second int) bool {
    if first % 4 > second % 4 || (first % 4 == 3 && second % 4 == 0) {
        return true
    }
    return false
}

func (b *Board) P1ShuffleDeck() error {
	if b.phase != P1 {
		newErr("Wrong phase")
	}
	b.deck.Shuffle()
	return nil
}

func (b *Board) P1DealAll() error {
	if b.Phase() != P1 {
		newErr("Wrong phase")
	}
	for i := 0; i < 52; i++ {
		hand := b.players[i%4].Hand()
		hand.PutOnTop(b.deck.Draw())
	}
	b.phase = P2
	return nil
}

func (b *Board) P2Trade(fromi int, card deck.Card) error {
	if b.Phase() != P2 {
		return newErr("Wrong phase")
	}
    if b.NextPlayerI() != fromi {
        return newErr("Wrong player")
    }
    to := b.Player((fromi + 1) % 4)
    from := b.Player(fromi)
    var err error
    card, err = from.Hand().FindAndDraw(card)
    if err != nil {
        return err
    }
    to.Hand().PutOnTop(card)
    b.nextPlayer++

    return err
}
