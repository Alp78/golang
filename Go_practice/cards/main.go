package main

import (
	"fmt"
)

func main() {
	// explicitly declaring the type will force the variable to only take this type
	// var card string = "Ace of Spades"

	// declaration and initialization with "var" can be refactored with ":"
	// card := "Ace of Spades"

	// card := newCard()
	// fmt.Println(card)

	// Array:	fixed length list
	// Slice:	list that can grow or shrink
	// -> both must be of the same type

	// cards := deck{newCard()}
	// cards = append(cards, "Ace of Spades")
	/* func newCard() string {
	return "Ace of Diamonds"
	}


	deck1 := newDeck()
	hand1, remain1 := deal(deck1, 5)

	hand1.print()
	fmt.Println(" ")
	remain1.print()
	*/
	deck1 := newDeck()
	deck1.saveDeckToFile("deck1.txt")
	deck2 := loadDeckFromFile("deck1.txt")

	deck2.shuffle()
	fmt.Println(deck2)
}
