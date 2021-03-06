package deck

import (
	"testing"
)

func testCards() []Card {
	return []Card{
		Card{face: "h", value: "2", valid: true},
		Card{face: "h", value: "3", valid: true}}
}
func testCards6() []Card {
	return []Card{
		Card{face: "h", value: "2", valid: true},
		Card{face: "h", value: "3", valid: true},
		Card{face: "s", value: "A", valid: true},
		Card{face: "c", value: "T", valid: true},
		Card{face: "d", value: "7", valid: true},
		Card{face: "d", value: "4", valid: true}}
}

func testDeck() Deck {
	return NewDeck(testCards())
}

func testDeck3() Deck {
	d := testDeck()
	d.cards = append(d.cards, Card{face: "s", value: "A"})
	return d
}

func testDeck6() Deck {
	return NewDeck(testCards6())
}

func testEqCard(a, b []Card) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestNewDefaultDeck(t *testing.T) {
	deck := NewDefaultDeck()
	for f := range faces {
	Vals:
		for v := range vals {
			card := Card{face: f, value: v, valid: true}
			for _, c := range deck.Cards() {
				if c == card {
					continue Vals
				}
			}
			t.Fatal("NewDefaultDeck() card", v, f, "not found")
		}
	}
}

func TestNewDeck(t *testing.T) {
	cards := []Card{
		Card{face: "H", value: "2"},
		Card{face: "H", value: "3"}}

	deck := NewDeck(cards)
	if len(deck.cards) != 2 {
		t.Fatal("NewDeck(2 cards) len = ", len(deck.cards), " expected 2")
	}

	deck = NewDeck([]Card{})
	if len(deck.cards) != 0 {
		t.Fatal("NewDeck(0 cards) len = ", len(deck.cards), " expected 0")
	}
}

func TestCards(t *testing.T) {
	deck := testDeck()
	if !testEqCard(testCards(), deck.Cards()) {
		t.Fatal("testDeck.Cards() does not equal input cards")
	}
}

func TestSetCards(t *testing.T) {
	deck := NewDeck([]Card{})
	deck.SetCards(testCards())
	if !testEqCard(testCards(), deck.cards) {
		t.Fatal("testDeck.Cards() does not equal input cards")
	}
}

func TestFind(t *testing.T) {
	deck := testDeck()
	if deck.Find(testCards()[0]) == false {
		t.Fatal("testDeck.Find didnt find a card that should exist")
	}
	if deck.Find(Card{face: "c", value: "7s", valid: true}) == true {
		t.Fatal("testDeck.Find found a card that shouldnt be in")
	}
	deck = NewDeck([]Card{})
	if deck.Find(Card{face: "c", value: "7s", valid: true}) == true {
		t.Fatal("testDeck.Find 2 found a card that shouldnt be in")
	}
}

func TestCard(t *testing.T) {
	deck := testDeck()
	if deck.Card(0) != testCards()[0] {
		t.Fatal("testDeck.Card(0) does not equal input card 0")
	}
	if deck.Card(1) != testCards()[1] {
		t.Fatal("testDeck.Card(1) does not equal input card 1")
	}
}

func TestSortByComp(t *testing.T) {
	deck := testDeck6()
	deck.SortByComp(false)
	for i := 0; i < len(deck.cards)-1; i++ {
		if deck.cards[i].Comparable() > deck.cards[i+1].Comparable() {
			t.Fatal("deck.SortByComp deck[", i, "] ", deck.cards[i], " is bigger than deck[", i+1, "]", deck.cards[i+1])
		}
	}
	deck.SortByComp(true)
	for i := 0; i < len(deck.cards)-1; i++ {
		if deck.cards[i].Comparable() < deck.cards[i+1].Comparable() {
			t.Fatal("deck.SortByComp deck[", i, "] is smaller than deck[", i+1, "]")
		}
	}
}

func TestHasFace(t *testing.T) {
	deck := testDeck()
	if deck.HasFace("c") {
		t.Fatal("deck.HasFace(c) returned true, shouldnt")
	}
	if !deck.HasFace("h") {
		t.Fatal("deck.HasFace(h) returned false, shouldnt")
	}
}

func TestHasValue(t *testing.T) {
	deck := testDeck()
	if deck.HasValue("5") {
		t.Fatal("deck.HasValue(5) returned true, shouldnt")
	}
	if !deck.HasValue("3") {
		t.Fatal("deck.HasValue(3) returned false, shouldnt")
	}
}

func TestDraw(t *testing.T) {
	deck := testDeck()
	c := deck.Draw()
	if c != testCards()[0] {
		t.Fatal("testDeck.Draw() does not equal input card 0")
	}
	if len(deck.cards) != 1 {
		t.Fatal("testing.Draw() count is ", len(deck.cards), " should be 1")
	}
	c = deck.Draw()
	if c != testCards()[1] {
		t.Fatal("testDeck.Draw() 2 does not equal input card 1")
	}
	if len(deck.cards) != 0 {
		t.Fatal("testing.Draw() 2 count is ", len(deck.cards), " should be 0")
	}
	c = deck.Draw()
	if c.valid {
		t.Fatal("testing.Draw() 3 returned a valid card, shouldnt")
	}
}

func TestDrawCardFromN(t *testing.T) {
	deck := testDeck()
	c, err := deck.DrawCardFromN(2)
	if err == nil {
		t.Fatal("testDeck.DrawCardFromN(2) didnt return an error, should")
	}
	if c.valid {
		t.Fatal("testDeck.DrawCardFromN(2) returned valid card, shouldnt")
	}
	c, err = deck.DrawCardFromN(1)
	if err != nil {
		t.Fatal("testDeck.DrawCardFromN(1) returned error, shouldnt")
	}
	if c != testCards()[1] {
		t.Fatal("testDeck.DrawCardFromN(1) returned wrong card, should be 3h")
	}
	if len(deck.cards) != 1 {
		t.Fatal("testDeck.DrawCardFromN(1) deck has ", len(deck.cards), "cards, should have 1")
	}
}

func TestDrawCards(t *testing.T) {
	deck := testDeck3()
	c, err := deck.DrawCards(2)
	if err != nil {
		t.Fatal("testDeck.DrawCards(2) returned an error, shouldnt")
	}
	if len(c) != 2 {
		t.Fatal("testDeck.DrawCards(2) does not contain 2 cards")
	}
	c, err = deck.DrawCards(2)
	if err == nil {
		t.Fatal("testDeck.DrawCards(2) 2 doesnt return an error, should")
	}
	if len(deck.cards) != 1 {
		t.Fatal("testing.DrawCards(2) 2 deck count is ", len(deck.cards), " should be 1")
	}
	c, err = deck.DrawCards(1)
	if err != nil {
		t.Fatal("testDeck.DrawCards(1) returned an error, shouldnt")
	}
	if len(c) != 1 {
		t.Fatal("testDeck.DrawCards(1) does not contain 1 cards")
	}
	if len(deck.cards) != 0 {
		t.Fatal("testDeck.DrawCards(1) deck afterwards does not contain 0 cards")
	}
}

func TestDrawCardsFromN(t *testing.T) {
	deck := testDeck6()
	c, err := deck.DrawCardsFromN(1, 3)
	if err != nil {
		t.Fatal("testDeck.DrawCardsFromN(1, 3) returned error, shouldnt")
	}
	if len(deck.cards) != 3 {
		t.Fatal("testDeck.DrawCardsFromN(1, 3) deck got ", len(deck.cards), " should have 3")
	}
	if deck.cards[1] != testCards6()[4] {
		t.Fatal("testDeck.DrawCardsFromN(1, 3) deck's 2nd card is wrong")
	}
	c, err = deck.DrawCardsFromN(5, 0)
	if err == nil {
		t.Fatal("testDeck.DrawCardsFromN(5, 0) returned no error, should (index too big)")
	}
	if len(c) != 0 {
		t.Fatal("testing.DrawCardsFromN(5, 0) returned non-empty cards array, shouldnt")
	}
	if len(deck.cards) != 3 {
		t.Fatal("testDeck.DrawCardsFromN(5, 0) returned ", len(deck.cards), " should have 3")
	}
	c, err = deck.DrawCardsFromN(0, 4)
	if err == nil {
		t.Fatal("testDeck.DrawCardsFromN(0, 4) returned no error, should (not enough cards in deck)")
	}
	if len(c) != 0 {
		t.Fatal("testing.DrawCardsFromN(0, 4) returned non-empty cards array, shouldnt")
	}
	if len(deck.cards) != 3 {
		t.Fatal("testDeck.DrawCardsFromN(0, 4) returned ", len(deck.cards), " should have 3")
	}
	c, err = deck.DrawCardsFromN(0, 3)
	if err != nil {
		t.Fatal("testDeck.DrawCardsFromN(0, 3) returned error, shouldnt")
	}
	if len(deck.cards) != 0 {
		t.Fatal("testDeck.DrawCardsFromN(0, 3) deck got ", len(deck.cards), " should have 0")
	}
	if c[1] != testCards6()[4] {
		t.Fatal("testDeck.DrawCardsFromN(0, 3) returned wrong 2nd card")
	}
}

func TestDrawSpecificCards(t *testing.T) {
	deck := testDeck6()
	toDraw := testCards()
	err := deck.DrawSpecificCards(toDraw)
	if err != nil {
		t.Fatal("testDeck6.DrawSpecificCards returned err,", err, " shouldnt")
	}
	err = deck.DrawSpecificCards([]Card{Card{face: "c", value: "T", valid: true}, Card{face: "h", value: "3", valid: true}})
	if len(deck.cards) != 4 {
		t.Fatal("testDeck6.DrawSpecificCards has", len(deck.cards), "cards, should have 4")
	}
}

func TestFindAndDraw(t *testing.T) {
	deck := testDeck()
	c, err := deck.FindAndDraw(testCards()[1])
	if err != nil {
		t.Fatal("testDeck.FindAndDraw(3h) returned error, shouldnt")
	}
	if c != testCards()[1] {
		t.Fatal("testDeck.FindAndDraw(3h) returned wrong card, should be 3h")
	}
	c, err = deck.FindAndDraw(testCards()[1])
	if err == nil {
		t.Fatal("testDeck.FindAndDraw(3h) didnt return an error, should")
	}
	if c.valid {
		t.Fatal("testDeck.FindAndDraw(3h) returned valid card, shouldnt")
	}
	if len(deck.cards) != 1 {
		t.Fatal("testDeck.FindAndDraw(3h) deck got ", len(deck.cards), " should have 1")
	}
}

func TestPutOnTop(t *testing.T) {
	deck := testDeck()
	c := Card{face: "d", value: "7", valid: true}
	deck.PutOnTop(c)
	if deck.cards[0] != c {
		t.Fatal("testDeck.PutOnTop(7d), deck's first card isnt correct")
	}
	if len(deck.cards) != 3 {
		t.Fatal("testDeck.PutOnTop(7d) deck has ", len(deck.cards), " cards, should have 3")
	}
	c = Card{face: "h", value: "7", valid: true}
	deck.PutOnTop(c)
	if deck.cards[0] != c {
		t.Fatal("testDeck.PutOnTop(7d), deck's first card isnt correct")
	}
	if len(deck.cards) != 4 {
		t.Fatal("testDeck.PutOnTop(7h) deck has ", len(deck.cards), " cards, should have 4")
	}
}

func TestPutManyOnTop(t *testing.T) {
	deck := testDeck()
	c := []Card{Card{face: "d", value: "7", valid: true}, Card{face: "d", value: "8", valid: true}}
	deck.PutManyOnTop(c)
	if deck.Card(0) != c[0] || deck.Card(1) != c[1] {
		t.Fatal("deck.PutManyOnTop(7d, 8d): card with index 0 or 1 is not correct")
	}
}

func TestPutOnBot(t *testing.T) {
	deck := testDeck()
	c := Card{face: "d", value: "7", valid: true}
	deck.PutOnBot(c)
	if deck.cards[2] != c {
		t.Fatal("testDeck.PutOnBot(7d), deck's last card isnt correct")
	}
	if len(deck.cards) != 3 {
		t.Fatal("testDeck.PutOnBot(7d) deck has ", len(deck.cards), " cards, should have 3")
	}
}

func TestPutManyOnBot(t *testing.T) {
	deck := testDeck()
	c := []Card{Card{face: "d", value: "7", valid: true}, Card{face: "d", value: "8", valid: true}}
	deck.PutManyOnBot(c)
	if deck.Card(deck.Count() - 2) != c[0] || deck.Card(deck.Count() - 1) != c[1] {
		t.Fatal("deck.PutManyOnTop(7d, 8d): card with index 0 or 1 is not correct")
	}
}

func TestPutInN(t *testing.T) {
	deck := testDeck()
	c := Card{face: "d", value: "7", valid: true}
	err := deck.PutInN(2, c)
	if err != nil {
		t.Fatal("testDeck.PutInN(2, 7d) returned error ", err, ", shouldnt")
	}
	if len(deck.cards) != 3 {
		t.Fatal("testDeck.PutInN(2, 7d) returned deck with ", len(deck.cards), "cards, should have 3")
	}
	if deck.cards[2] != c {
		t.Fatal("testDeck.PutInN(2, 7d) deck's card[2] is wrong")
	}
	err = deck.PutInN(4, c)
	if err == nil {
		t.Fatal("testDeck.PutInN(3, 7d) didnt return an error, should (n too high)")
	}
}

func TestSort(t *testing.T) {
    deck := testDeck6()
    comp := func (c1 Card, c2 Card) bool { return c1.Comparable() > c2.Comparable() }
    deck.Sort(comp)
    for i := 0; i < deck.Count() - 2; i++ {
        if comp(deck.Card(i + 1), deck.Card(i)) {
            t.Fatal("cards", i + 1, "bigger than card", i, "shouldnt be")
        }
    }
    deck = testDeck6()
    comp = func (c1 Card, c2 Card) bool { return c1.Comparable() < c2.Comparable() }
    deck.Sort(comp)
    for i := 0; i < deck.Count() - 2; i++ {
        if comp(deck.Card(i + 1), deck.Card(i)) {
            t.Fatal("cards", i + 1, "smaller than card", i, "shouldnt be")
        }
    }
}
