package main

import (
	"fmt"
)

// Deck class
type Deck []string

func newDeck() Deck {
	cards := Deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearths", "Clubs"}
	cardValues := []string{"Ace", "Tow", "Three", "Four", "Five"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d Deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d Deck) deal(handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}
