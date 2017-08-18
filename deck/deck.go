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

func (d *Deck) Card(i int) Card {
    return d.cards[i]
}

func (d *Deck) Shuffle() {
    cards := d.Cards()
    clen := len(cards)
    for i := clen - 1; i > 0; i-- {
        rand.Seed(time.Now().Unix())
        min := 0
        max := i
        swapi := rand.Intn(max - min) + min
        cards[swapi], cards[i] = cards[i], cards[swapi]
    }
}

func (d *Deck) Count() int {
    return len(d.Cards())
}

func (d *Deck) Draw() Card {
    if d.Count() == 0 {
        return NewCard("")
    }
    c := d.Card(0)
    d.SetCards(d.cards[1:])
    return c
}

func (d *Deck) DrawCards(n int) ([]Card, error) {
    if d.Count() < n {
        return []Card{}, newErr("Not enough cards in deck")
    }

    cards := d.Cards()
    drawn := cards[:n]
    d.SetCards(cards[n:])

    return drawn, nil
}

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
