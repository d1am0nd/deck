package deck

import (
	"fmt"
	"math/rand"
	"time"
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

func NewDefaultDeck() Deck {
	cards := []Card{}
	for v := range vals {
		for f := range faces {
			cards = append(cards, NewCard(v+f))
		}
	}
	return NewDeck(cards)
}

func (d *Deck) Cards() []Card {
	return d.cards
}

func (d *Deck) SetCards(c []Card) {
	d.cards = c
}

func (d *Deck) Find(c Card) bool {
	for i := range d.cards {
		if d.cards[i] == c {
			return true
		}
	}
	return false
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
		rng := rand.New(rand.NewSource(time.Now().UnixNano()))
		swapi := rng.Intn(i)
		cards[swapi], cards[i] = cards[i], cards[swapi]
	}
}

// Returns number of cards in deck
func (d *Deck) Count() int {
	return len(d.Cards())
}

// Sorting by comparable()
func (d *Deck) SortByComp(desc bool) {
	if desc {
		sort(func (c1 Card, c2 Card) bool {return c1.Comparable() > c2.Comparable()},d.Cards())
	} else {
		sort(func (c1 Card, c2 Card) bool {return c1.Comparable() < c2.Comparable()},d.Cards())
	}
}

func (d *Deck) Sort(comp func(Card, Card) bool) {
	sort(comp, d.Cards())
}

func sort(comp func(Card, Card) bool, cards []Card) []Card {
	clen := len(cards)
	if clen <= 1 {
		return cards
	}
	pivoti := clen / 2
	pivot := cards[pivoti]
	left, right := 0, clen-1
	cards[pivoti], cards[right] = cards[right], cards[pivoti]
	for i := 0; i < right; i++ {
		if comp(cards[i], pivot) {
			cards[left], cards[i] = cards[i], cards[left]
			left++
		}
	}
	cards[left], cards[right] = cards[right], cards[left]
	sort(comp, cards[:left])
	sort(comp, cards[left+1:])

	return cards
}

/** Finding cards */
func (d *Deck) HasFace(f string) bool {
	for _, c := range d.Cards() {
		if c.Face() == f {
			return true
		}
	}
	return false
}

func (d *Deck) HasValue(v string) bool {
	for _, c := range d.Cards() {
		if c.Value() == v {
			return true
		}
	}
	return false
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

// Draws card from index n if exists
func (d *Deck) DrawCardFromN(i int) (Card, error) {
	cards := d.Cards()
	if i >= len(cards) {
		return Card{}, newErr("Not enough cards in deck")
	}
	card := cards[i]
	d.SetCards(append(cards[:i], cards[i+1:]...))
	return card, nil
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
	if i+n > clen {
		return []Card{}, newErr("Not enough cards in deck")
	}

	cards := d.Cards()
	drawn := cards[i:(i + n)]
	d.SetCards(append(cards[:i], cards[i+n:]...))

	return drawn, nil
}

// Draws all passed cards or returns an error, if at least one of them is not found
func (d *Deck) DrawSpecificCards(cards []Card) error {
	for i := range cards {
		if !d.Find(cards[i]) {
			return newErr("Card not found")
		}
	}
	for i := range cards {
		d.FindAndDraw(cards[i])
	}
	return nil
}

// Draws specific card if found
func (d *Deck) FindAndDraw(c Card) (Card, error) {
	for i := range d.Cards() {
		if d.Card(i) == c {
			return d.DrawCardFromN(i)
		}
	}
	return Card{}, newErr("Card not found")
}

/** Appending to deck */

// Puts given card on top of deck
func (d *Deck) PutOnTop(c Card) {
	d.SetCards(append([]Card{c}, d.Cards()...))
}

// Puts given slice of cards on top of deck
func (d *Deck) PutManyOnTop(c []Card) {
	d.SetCards(append(c, d.Cards()...))
}

// Puts given card at the end of deck
func (d *Deck) PutOnBot(c Card) {
	d.SetCards(append(d.Cards(), c))
}

// Puts given slice of cards on bot of deck
func (d *Deck) PutManyOnBot(c []Card) {
	d.SetCards(append(d.Cards(), c...))
}

// Append card to deck to nth index
func (d *Deck) PutInN(i int, c Card) error {
	if i > d.Count() {
		return newErr("not enough cards in deck")
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

func (d *Deck) Print() {
	for _, c := range d.Cards() {
		fmt.Print(c.Value(), c.Face(), ",")
	}
}
