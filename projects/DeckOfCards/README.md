# Deck of Cards

A flexible Go package for creating and manipulating decks of playing cards with customizable options.

## Features

- Generate standard 52-card decks
- Add jokers to your deck
- Sort cards with custom sorting functions
- Shuffle decks randomly
- Filter out specific cards
- Combine multiple decks
- Chainable options for flexible deck creation

## Installation

```bash
go get github.com/yourusername/DeckOfCards
```

## Usage

### Basic Usage

```go
package main

import (
    "DeckOfCards/deck"
    "fmt"
)

func main() {
    // Generate a standard deck
    cards := deck.GenerateNew()
    fmt.Println(len(cards)) // 52
}
```

### Using Options

The package uses a functional options pattern to customize deck generation:

```go
// Create a shuffled deck with 2 jokers
cards := deck.GenerateNew(
    deck.OptionAddJokers(2),
    deck.OptionShuffle(),
)

// Create a sorted deck excluding all diamonds
cards := deck.GenerateNew(
    deck.OptionSort(deck.DefaultOptionSortFunction),
    deck.OptionExclude(func(c deck.Card) bool {
        return c.Suit == deck.SuitDiamonds
    }),
)
```

## API Reference

### Types

#### Card

```go
type Card struct {
    Suit  Suit
    Value Value
}
```

#### Suit

```go
type Suit int

const (
    SuitSpades   Suit = iota // ♠
    SuitHearts               // ♥
    SuitDiamonds             // ♦
    SuitClubs                // ♣
    SuitJoker                // J
)
```

#### Value

```go
type Value int

const (
    ValueTwo   Value = 2
    ValueThree Value = 3
    // ... up to ...
    ValueAce   Value = 14
)
```

Face cards are represented as:

- Jack: `ValueJack` (displays as "J")
- Queen: `ValueQueen` (displays as "Q")
- King: `ValueKing` (displays as "K")
- Ace: `ValueAce` (displays as "A")

### Functions

#### GenerateNew

```go
func GenerateNew(options ...Option) []Card
```

Creates a new deck of cards. Without options, generates a standard 52-card deck.

### Options

#### OptionSort

```go
func OptionSort(sortFunction func(i, j Card) bool) Option
```

Sorts the deck using the provided comparison function.

**Default Sort Function:**

```go
var DefaultOptionSortFunction = func(i, j Card) bool {
    return i.Suit < j.Suit || (i.Suit == j.Suit && i.Value < j.Value)
}
```

#### OptionShuffle

```go
func OptionShuffle() Option
```

Randomly shuffles the deck.

#### OptionAddJokers

```go
func OptionAddJokers(n int) Option
```

Adds `n` jokers to the deck.

#### OptionExclude

```go
func OptionExclude(filterFunc func(Card) bool) Option
```

Excludes cards that match the filter function. Cards are excluded if the function returns `true`.

#### OptionCompose

```go
func OptionCompose(decks ...[]Card) Option
```

Combines multiple decks into one.

## Examples

### Custom Sorting

```go
// Sort by value only
cards := deck.GenerateNew(
    deck.OptionSort(func(i, j deck.Card) bool {
        return i.Value < j.Value
    }),
)
```

### Filtering Cards

```go
// Create a deck with only red cards
cards := deck.GenerateNew(
    deck.OptionExclude(func(c deck.Card) bool {
        return c.Suit == deck.SuitSpades || c.Suit == deck.SuitClubs
    }),
)

// Create a deck with only face cards
cards := deck.GenerateNew(
    deck.OptionExclude(func(c deck.Card) bool {
        return c.Value < deck.ValueJack
    }),
)
```

### Combining Operations

```go
// Complex deck: sorted, shuffled, with jokers, excluding diamonds,
// and combined with another deck
cards := deck.GenerateNew(
    deck.OptionSort(deck.DefaultOptionSortFunction),
    deck.OptionShuffle(),
    deck.OptionAddJokers(2),
    deck.OptionExclude(func(c deck.Card) bool {
        return c.Suit == deck.SuitDiamonds
    }),
    deck.OptionCompose(deck.GenerateNew()),
)
```

## String Representation

Cards have built-in string representations:

- Suits: ♠ ♥ ♦ ♣ J (for Joker)
- Values: 2-10, J, Q, K, A
