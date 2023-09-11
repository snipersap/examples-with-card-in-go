package main
import "testing"

func TestNewDeck(t *testing.T){
	testDeck := newDeck()

	if len(testDeck) != 16 {
		t.Errorf("Size of deck not correct. Expected 16, but got %v",len(testDeck))
	}
	// if len(testDeck) != 20 {
	// 	t.Errorf("Size of deck not correct. Expected 20, but got %v",len(testDeck))
	// }
	if testDeck[0] != "Ace of Hearts" {
		t.Errorf("Expected Ace of Hearts as first card, but got %v",testDeck[0])
	}
	if(testDeck[len(testDeck)-1] != "Four of Club") {
		t.Errorf("Expected the last card to be Four of Club, but got %v",testDeck[len(testDeck)-1])
	}
}