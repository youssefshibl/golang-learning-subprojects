package main

import (
	"DeckOfCards/deck"
	"fmt"
)

func main() {

	cards := deck.GenerateNew(deck.OptionSort(deck.DefaultOptionSortFunction), deck.OptionSort(func(i, j deck.Card) bool {
		return i.Value < j.Value
	}), deck.OptionShuffle(), deck.OptionAddJokers(20), deck.OptionExclude(func(c deck.Card) bool {
		return c.Suit == deck.SuitDiamonds
	}), deck.OptionCompose(deck.GenerateNew()))
	fmt.Println(cards)
}
