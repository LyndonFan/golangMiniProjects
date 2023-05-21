package main

import (
	"fmt"

	"example.com/deck"
)

func main(){
	card := deck.Card{deck.Spades, deck.Seven}
	deck := deck.Deck{[]deck.Card{card}}
	fmt.Println(deck)
}