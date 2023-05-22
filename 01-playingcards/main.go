package main

import (
	"fmt"
	"game/deck"
)

func main(){
	card := deck.Card{deck.Spades, deck.Seven}
	deck := deck.Deck{[]deck.Card{card}}
	fmt.Println(deck)
}