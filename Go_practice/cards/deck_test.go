package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}
}

func TestCardsName(t *testing.T) {
	d := newDeck()

	if d[0] != "Ace of Spade" {
		t.Errorf("Expected first card's name '%v', but got '%v'", "Ace of Spade", d[0])
	}

	if d[len(d)-1] != "Four of Club" {
		t.Errorf("Expected last card's name '%v', but got '%v'", "Four of Club", d[len(d)-1])
	}
}

func TestSaveDeckToFileAndLoadDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveDeckToFile("_decktesting")
	loadedDeck := loadDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
