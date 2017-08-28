package hearths

import (
	"cards/deck"
)

type Board struct {
	gamen         int
	players       [4]Player
	deck          deck.Deck
	mainPile      deck.Deck
	finished      bool
	results       [4]int
	hearthsBroken bool
	p3started     int
	turn          int
	phase         int
}

const P1 = 1 // Suffle + deal
const P2 = 2 // Trade
const P3 = 3 // Play
const P4 = 4

const P1lastTurn = 0
const P2lastTurn = 3
const P3lastTurn = 55

func NewBoard(players [4]Player) Board {
	return Board{
		gamen:         0,
		players:       players,
		deck:          deck.NewDefaultDeck(),
		mainPile:      deck.NewDeck([]deck.Card{}),
		results:       [4]int{0, 0, 0, 0},
		finished:      false,
		hearthsBroken: false,
		p3started:     0,
		turn:          0,
		phase:         1}
}

func (b *Board) Finished() bool {
	return b.finished
}

func (b *Board) Results() [4]int {
	return b.results
}

func (b *Board) Phase() int {
	return b.phase
}

func (b *Board) MainPile() *deck.Deck {
	return &b.mainPile
}

func (b *Board) NextPlayerI() int {
	if b.Phase() == P3 {
		return (b.turn + b.p3started) % 4
	}
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
	for i, p := range b.players {
		if p.Hand().Find(deck.NewCard("2c")) {
			b.p3started = i
			break
		}
	}
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

	if b.turn > P2lastTurn {
		b.phase = P3
	}

	return err
}

func (b *Board) P3PutOnPile(fromi int, card deck.Card) error {
	if b.turn == P2lastTurn+1 && (card.Face() != "c" || card.Value() != "2") {
		return newErr("Must put 2c")
	}
	if b.Phase() != P3 {
		return newErr("Wrong phase")
	}
	if b.NextPlayerI() != fromi {
		return newErr("Wrong player")
	}
	err := canPlayOnPile(*b.Player(fromi).Hand(), card, *b.MainPile(), b.hearthsBroken)
	if err != nil {
		return err
	}
	card, err = b.Player(fromi).Hand().FindAndDraw(card)
	if err != nil {
		return err
	}
	b.MainPile().PutOnTop(card)

	if isHearths(card) {
		b.hearthsBroken = true
	}

	if b.MainPile().Count() == 4 {
		// cards in pile have reverse ordfer to players indicies
		win := b.p3started + (4-winnerI(*b.MainPile()))%4
		cards := make([]deck.Card, 4)
		copy(cards, b.MainPile().Cards())
		err = b.MainPile().DrawSpecificCards(cards)
		if err != nil {
			return err
		}
		b.Player(win).Garbage().PutManyOnTop(cards)
	}

	b.turn++
	if b.turn > P3lastTurn {
		b.phase = P4
		for i, p := range b.players {
			b.results[i] = sumResult(*p.Garbage())
		}
		b.finished = true
	}
	return nil
}
