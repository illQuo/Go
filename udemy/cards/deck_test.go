package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 9 {
		t.Errorf("expected deck length of 9, but got %d", len(cards))
	}
}

func TestSaveToFileAndLoadFromFile(t *testing.T) {
	os.Remove("_decktesting")

	newDeck().saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 9 {
		t.Errorf("expected 9 cards, got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
