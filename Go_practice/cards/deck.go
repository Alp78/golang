package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new custom type 'deck' (slice of string)
// []string is its underlying type
// declared after 'import': package scope (global)
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spade", "Diamond", "Heart", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}
	return cards
}

// create a function that belongs to the 'deck' type
// 'deck' is the receiver
// -> any variable of type 'deck' has access to the function
// 'd' is the actual copy of the deck (variable) in the function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// range of a slice: deck[0:2]
// -> from 0 included up to 2 not included
// -> can ommit 0 if from start, can ommit 2 if it is the slice's end

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// type conversion 'deck' -> []byte
// 1 - helper function to convert in string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// 2 - svae to file (conversion in the function)
func (d deck) saveDeckToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func loadDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option 1: log the error and return a call to newDeck()
		// Option 2: log the error and quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	ss := strings.Split(string(bs), ",")

	return deck(ss)
}

func (d deck) shuffle() {
	// set up the random generator with a new seed each run
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		// generate a random integer between 0 and deck length (index)
		newPosition := r.Intn(len(d) - 1)
		// swap the elements
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
