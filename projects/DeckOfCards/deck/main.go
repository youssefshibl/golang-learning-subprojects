package deck

import (
	"math/rand"
	"sort"
	"strconv"
)

type Suit int

const (
	SuitSpades Suit = iota
	SuitHearts
	SuitDiamonds
	SuitClubs
	SuitJoker
)

func (s Suit) String() string {
	switch s {
	case SuitSpades:
		return "♠"
	case SuitHearts:
		return "♥"
	case SuitDiamonds:
		return "♦"
	case SuitClubs:
		return "♣"
	case SuitJoker:
		return "J"
	}
	return ""
}

type Value int

const (
	valueTwo Value = iota + 2
	ValueThree
	ValueFour
	ValueFive
	ValueSix
	ValueSeven
	ValueEight
	ValueNine
	ValueTen
	ValueJack
	ValueQueen
	ValueKing
	ValueAce
)

func (v Value) String() string {
	switch v {
	case ValueJack:
		return "J"
	case ValueQueen:
		return "Q"
	case ValueKing:
		return "K"
	case ValueAce:
		return "A"
	default:
		return strconv.Itoa(int(v))
	}
}

type Card struct {
	Suit  Suit
	Value Value
}

type Option func([]Card) []Card

func GenerateNew(options ...Option) []Card {

	var res []Card

	for suit := SuitSpades; suit <= SuitClubs; suit++ {
		for value := valueTwo; value <= ValueAce; value++ {

			res = append(res, Card{Suit: suit, Value: value})
		}
	}

	for _, option := range options {
		res = option(res)
	}
	return res
}

var DefaultOptionSortFunction func(i, j Card) bool = func(i, j Card) bool {
	return i.Suit < j.Suit || (i.Suit == j.Suit && i.Value < j.Value)
}

func OptionSort(sorFunction func(i, j Card) bool) Option {
	return func(c []Card) []Card {
		sort.Slice(c, func(i, j int) bool {
			return sorFunction(c[i], c[j])
		})
		return c
	}
}

func OptionShuffle() Option {
	return func(c []Card) []Card {
		rand.Shuffle(len(c), func(i, j int) {
			c[i], c[j] = c[j], c[i]
		})
		return c
	}
}

func OptionAddJokers(n int) Option {
	return func(deck []Card) []Card {
		for i := 1; i <= n; i++ {
			deck = append(deck, Card{Suit: SuitJoker})
		}
		return deck
	}
}

func OptionExclude(fun func(Card) bool) Option {
	return func(cards []Card) []Card {
		var filteredCards []Card
		for _, card := range cards {
			if fun(card) {
				continue
			}
			filteredCards = append(filteredCards, card)
		}
		return filteredCards
	}
}

func OptionCompose(decks ...[]Card) Option {
	return func(deck []Card) []Card {
		for _, d := range decks {
			deck = append(deck, d...)
		}
		return deck
	}
}
