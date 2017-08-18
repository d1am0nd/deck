package deck

import (
    "time"
    "math/rand"
)

/**
 * Simple deck
 * Some rules:
 * * Cards are assumed to be unique (one of each max)
 * * Top card is card with index 0
 */

type Deck struct {
    cards []Card
}

func NewDeck(cards []Card) Deck {
    return Deck{cards: cards}
}

func (d *Deck) Cards() []Card {
    return d.cards
}

func (d *Deck) SetCards(c []Card) {
    d.cards = c
}

// Returns the card on index i
func (d *Deck) Card(i int) Card {
    return d.cards[i]
}

// Shuffles deck's cards
func (d *Deck) Shuffle() {
    cards := d.Cards()
    clen := len(cards)
    for i := clen - 1; i > 0; i-- {
        rand.Seed(time.Now().Unix())
        swapi := rand.Intn(i)
        cards[swapi], cards[i] = cards[i], cards[swapi]
    }
}

// Returns number of cards in deck
func (d *Deck) Count() int {
    return len(d.Cards())
}


/** Drawing from deck */

// Draws top card
func (d *Deck) Draw() Card {
    if d.Count() == 0 {
        return NewCard("")
    }
    c := d.Card(0)
    d.SetCards(d.cards[1:])
    return c
}

// Draws top n cards
func (d *Deck) DrawCards(n int) ([]Card, error) {
    if d.Count() < n {
        return []Card{}, newErr("Not enough cards in deck")
    }

    cards := d.Cards()
    drawn := cards[:n]
    d.SetCards(cards[n:])

    return drawn, nil
}

// Draws n cards starting from index i
func (d *Deck) DrawCardsFromN(i, n int) ([]Card, error) {
    clen := d.Count()
    if i + n > clen   {
        return []Card{}, newErr("Not enough cards in deck")
    }

    cards := d.Cards()
    drawn := cards[i:(i + n)]
    d.SetCards(append(cards[:i], cards[i + n:]...))

    return drawn, nil
}


/** Appending to deck */

// Puts given card on top of deck
func (d *Deck) PutOnTop(c Card) {
    d.SetCards(append([]Card{c}, d.Cards()...))
}

// Puts given card at the end of deck
func (d *Deck) PutOnBot(c Card) {
    d.SetCards(append(d.Cards(), c))
}

func (d *Deck) PutInN(i int, c Card) error {
    if i > d.Count() {
        return newErr("not enought cards in deck")
    }
    cards := d.Cards()
    d.SetCards(append(cards[:i], append([]Card{c}, cards[i:]...)...))
    return nil
}

// Puts given card in random place in deck
func (d *Deck) PutInRnd(c Card) {
    rand.Seed(time.Now().Unix())
    i := rand.Intn(d.Count())
    d.PutInN(i, c)
}
