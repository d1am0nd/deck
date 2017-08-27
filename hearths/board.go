package hearths

import (
	"cards/deck"
)

type Board struct {
	gamen         int
	players       [4]Player
	deck          deck.Deck
	mainPile      deck.Deck
	centerPile    deck.Deck
	hearthsBroken bool
	turn          int
	phase         int
}

const P1 = 1 // Suffle + deal
const P2 = 2 // Trade
const P3 = 3 // Play

func NewBoard(players [4]Player) Board {
	return Board{
		gamen:      0,
		players:    players,
		deck:       deck.NewDefaultDeck(),
		mainPile:   deck.NewDeck([]deck.Card{}),
		centerPile: deck.NewDeck([]deck.Card{}),
		hearthsBroken: false,
		turn:       0,
		phase:      1}
}

func (b *Board) Phase() int {
	return b.phase
}

func (b *Board) NextPlayerI() int {
	return b.turn % 4
}

func (b *Board) Player(i int) *Player {
	return &b.players[i%4]
}

func isNext(first, second int) bool {
	if first%4 > second%4 || (first%4 == 3 && second%4 == 0) {
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

// Expects 3 cards
func (b *Board) P2Trade(fromi int, cards []deck.Card) error {
	if b.Phase() != P2 {
		return newErr("Wrong phase")
	}
	if b.NextPlayerI() != fromi {
		return newErr("Wrong player")
	}
	if len(cards) != 3 {
		return newErr("Need exactly 3 cards")
	}
	to := b.Player((fromi + (b.gamen % 3) + 1) % 4)
	from := b.Player(fromi)

	var err error
	err = from.Hand().DrawSpecificCards(cards)
	if err != nil {
		return err
	}
	to.Hand().PutManyOnTop(cards)
	b.turn++

	if b.turn == 4 {
		b.phase = P3
	}

	return err
}
