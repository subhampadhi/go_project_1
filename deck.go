package main

import "fmt"

// pretty self explanatory
// just replace unused variables with underscore _
func newDeck() deck {
	cards := deck {}
	cardsSuite := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string {"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King"}
	for _, suite := range cardsSuite {
		for _, value := range cardValues {
			cards = append(cards, value+ " of " +suite)
		}
	}
	return cards
}

type deck []string

// they call it receiver function . lol....
// I think I got it .
// So this is apparently convention to call it a single letter variable as the receiver argument

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// so go can return more than one value . cools
func deal (d deck, handSize int ) (deck, deck) {
	return d[:handSize], d[handSize:]
}
