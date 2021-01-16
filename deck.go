package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) toString() string {
	// conversion of slice to string to upload stuff . Okay
	return strings.Join([]string(d), ",")
}

// so go can return more than one value . cools
func deal (d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func(d deck) shuffle() {
	// since we have to make it different everytime. Well always user time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		//swapping values
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func newDeckFromFile(filename string) deck {
	byteString, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("Error:",error)
		os.Exit(1)
	}
	s := strings.Split(string(byteString), ",")
	return deck(s)
}

