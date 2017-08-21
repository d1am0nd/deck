package main

import (
	"cards/deck"
	"fmt"
)

func main() {
	d := deck.NewDeck([]deck.Card{deck.NewCard("Th"),
		deck.NewCard("5h"),
		deck.NewCard("As"),
		deck.NewCard("3h"),
		deck.NewCard("7s"),
		deck.NewCard("2c")})

	fmt.Println(d.Cards())
	d.Shuffle()
	fmt.Println(d.Cards())

	fmt.Println("Hello world")
}
