package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	testDeck := newDeck()

	if len(testDeck) != 16 {
		t.Errorf("Size of deck not correct. Expected 16, but got %v", len(testDeck))
	}
	// if len(testDeck) != 20 {
	// 	t.Errorf("Size of deck not correct. Expected 20, but got %v",len(testDeck))
	// }
	if testDeck[0] != "Ace of Hearts" {
		t.Errorf("Expected Ace of Hearts as first card, but got %v", testDeck[0])
	}
	if testDeck[len(testDeck)-1] != "Four of Clubs" {
		t.Errorf("Expected the last card to be Four of Clubs, but got %v", testDeck[len(testDeck)-1])
	}
}

func TestSaveToFile(t *testing.T) {
	// Setup
	fileName := "_deckTesting"
	os.Remove(fileName)
	deck := newDeck()

	// Execute
	err := deck.saveToFile(fileName)
	if err != nil {
		t.Errorf("Saving deck to File %s failed. It resulted in the error:%v ", fileName, err)
	}

	// Teardown
	os.Remove(fileName)
}

func TestNewDeckFromFile(t *testing.T) {
	// Setup
	fileName := "_deckTesting"
	os.Remove(fileName)
	deck := newDeck()
	err := deck.saveToFile(fileName)
	if err != nil {
		t.Errorf("Saving deck to File %s failed. It resulted in the error:%v ", fileName, err)
	}

	// Execute
	deckFromFile := newDeckFromFile(fileName)

	if len(deckFromFile) != 16 {
		t.Errorf("Length of deck in %s expected to be 16, found %v instead", fileName, len(deckFromFile))
	}
	if deckFromFile[0] != "Ace of Hearts" {
		t.Errorf("Expected Ace of Hearts as first card, but got %v", deckFromFile[0])
	}
	if deckFromFile[len(deckFromFile)-1] != "Four of Clubs" {
		t.Errorf("Expected the last card to be Four of Clubs, but got %v", deckFromFile[len(deckFromFile)-1])
	}

	// Teardown
	os.Remove(fileName)
}
