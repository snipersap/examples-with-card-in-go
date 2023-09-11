package main
import (
	"fmt"	
	"strings"		//For string operations
	"io/ioutil"		//Use ioutil for IO operations
	"os"
)

//Part 2: Create a new type to represent a deck of cards as slice of string
type deck []string

//Part 2: Create a receiver function, that provides access to the variables of type deck to call it
//receivedDeck is the actual variable (reference) that called it (similar to self in Python)
func (receivedDeck deck) print(){
	for i,card := range receivedDeck {
		fmt.Println(i,">",card)
	}
}

//Part 3: Declare few more types with primary types in go
type deckSize int
type deckName string

func (receivedDeckSize deckSize) showDeckSize() {
	fmt.Println("size of the deck is", receivedDeckSize)
}

func (name deckName) getDeckName() deckName{
	return name
}


//Part 4: Function to create a new Deck of cards of a combination of 4 numbers and 4 suites
func newDeck() deck {
	cards := deck{}				//Init deck
	cardSuites := []string{"Hearts","Spades","Diamonds","Clubs"} //init the slice of strings as the 4 suits
	cardValues := []string{"Ace","Two","Three","Four"}

	//Create deck as a combo of suit and values
	for _,suit := range cardSuites {			//Use _ to tell the compiler that you don't want to use the required variable
		for _,value := range cardValues {
			cards = append(cards, value+" of "+suit)	//Combine strings using +
		}
	}
	return cards
}

//Part 5: Create a deal: a hand of few cards and the remaining cards from the deck
func deal(receivedDeck deck, sizeOfHand int) (deck, deck) {
	return receivedDeck[:sizeOfHand], receivedDeck[sizeOfHand:]
}

//Part 7: Deck to String. Used String instead of toString as per Effective Go.
func (d deck) String() string {
	return strings.Join([]string(d),",")
}

//Part 8: Convert Deck to byte slice
func (d deck) ByteSlice() []byte {
	return []byte(d.String())
}

// Part 9: Save the byte slice to file
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, d.ByteSlice(), 0644)
}

//Part 10: Reading from Hard Drive and Error Handling
func newDeckFromFile(fileName string) deck{
	fileContentsAsByteSlice, err := os.ReadFile(fileName)
	fmt.Println(err)
	// if isErr(err) {
	// 	handleReadFromFileErr(err)	//Handle error
	// }
	// Short cut to return the deck read from the file
	// return deck(strings.Split(string(fileContentsAsByteSlice),","))

	//Long format for learning purpose
	simpleText := string(fileContentsAsByteSlice)
	// Split the string into a string slice
	cardsSlice := strings.Split(simpleText,",")
	// Convert slice of cards into the deck type
	return deck(cardsSlice)
}