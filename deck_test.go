package main
import "testing"

func TestNewDeck(t *testing.T){
	testDeck := newDeck()

	if len(testDeck) != 16 {
		t.Errorf("Size of deck not correct. Expected 16, but got %v",len(testDeck))
	}
}