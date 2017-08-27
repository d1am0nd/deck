package hearths

import (
	"cards/deck"
	"testing"
)

func testHand() deck.Deck {
	return deck.NewDeck([]deck.Card{
		deck.NewCard("Ac"), deck.NewCard("Qh"), deck.NewCard("5c")})
}

func testPile() deck.Deck {
	return deck.NewDeck([]deck.Card{
		deck.NewCard("8c")})
}

func testPile2() deck.Deck {
	return deck.NewDeck([]deck.Card{})
}

func TestCanPlayOnPile(t *testing.T) {
	p := testPile()
	h := testHand()
	err := canPlayOnPile(h, deck.NewCard("5c"), p, true)
	if err != nil {
		t.Fatal("canPlayOnPile 1: returned err,", err, "shouldnt")
	}
	err = canPlayOnPile(h, deck.NewCard("Qh"), p, true)
	if err == nil {
		t.Fatal("canPlayOnPile 2: returned no error, shouldn")
	}
	p = testPile2()
	err = canPlayOnPile(h, deck.NewCard("Qh"), p, true)
	if err != nil {
		t.Fatal("canPlayOnPile 3: returned err,", err, "shouldnt")
	}
	err = canPlayOnPile(h, deck.NewCard("Qh"), p, false)
	if err == nil {
		t.Fatal("canPlayOnPile 4: returned no error, should")
	}

	h = testPile2()
	h.PutOnTop(deck.NewCard("Th"))
	err = canPlayOnPile(h, deck.NewCard("Th"), p, false)
	if err != nil {
		t.Fatal("canPlayOnPile 5: returned error", err, "shouldnt")
	}
	p = testHand()
	p.PutOnTop(deck.NewCard("Ts"))
	err = canPlayOnPile(h, deck.NewCard("Th"), p, false)
	if err == nil {
		t.Fatal("canPlayOnPile 6: returned no error, should")
	}
}

func TestWinnerI(t *testing.T) {
	d := deck.NewDeck([]deck.Card{
		deck.NewCard("Ac"), deck.NewCard("Qh"), deck.NewCard("5c"), deck.NewCard("4c")})
	if winnerI(d) != 0 {
		t.Fatal("winnerI 1: should return 0, returned", winnerI(d))
	}
	d = deck.NewDeck([]deck.Card{
		deck.NewCard("2c"), deck.NewCard("Qh"), deck.NewCard("5c"), deck.NewCard("Tc")})
	if winnerI(d) != 3 {
		t.Fatal("winnerI 2: should return 3, returned", winnerI(d))
	}
	d = deck.NewDeck([]deck.Card{
		deck.NewCard("As"), deck.NewCard("Qh"), deck.NewCard("Ah"), deck.NewCard("Qs")})
	if winnerI(d) != 0 {
		t.Fatal("winnerI 3: should return 0, returned", winnerI(d))
	}
}
