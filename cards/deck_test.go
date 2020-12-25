package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected deck length of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected deck length of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestFileIO(t *testing.T) {
	os.Remove("_deckTest")

	deck := newDeck()
	deck.saveToFile("_deckTest")

	loadDeck := newDeckFromFile("_deckTest")

	if len(loadDeck) != 16 {
		t.Errorf("Expected loadDeck length of 16, but got %v", len(loadDeck))
	}

	os.Remove("_deckTest")
}
