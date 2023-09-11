package main
import (
	"fmt"	
	"strings"		//For string operations
	"io/ioutil"		//Use ioutil for IO operations
	"os"
	"math/rand"
	"time"
)

//Part 2: Create a new type to represent a deck of cards as slice of string
type deck []string

//Part 2: Create a receiver function, that provides access to the variables of type deck to call it
//receivedDeck is the actual variable (reference) that called it (similar to self in Python)
func (receivedDeck deck) print(){
	fmt.Println("Printing Cards via print receiver function:>")
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
	if err != nil {
		fmt.Println("Error:", err)	//Print the error
		os.Exit(1)	//Exit with non-zero error code
	} else {
		fmt.Println("ReadFile Error:", err)
	}
	// Short cut to return the deck read from the file
	return deck(strings.Split(string(fileContentsAsByteSlice),","))

	// Long format for learning purpose
	// simpleText := string(fileContentsAsByteSlice)
	// Split the string into a string slice
	// cardsSlice := strings.Split(simpleText,",")
	// Convert slice of cards into the deck type
	// return deck(cardsSlice)
}

//Part 11: Create a deck shuffle function without seeding
func (d deck) shuffle() {
	// 1. Create a random number upto length of deck -1 (deck index starts from 0)
	// 2. Loop over the deck and swap each position with the randomly returned position
	for i := range d {
		newPosition := rand.Intn(len(d)-1)				//get random number
		d[i], d[newPosition] = d[newPosition], d[i]		//swap slices
	}
}

//Part 12: Shuffle the deck with seed
func (d deck) randomShuffle() {
	//1. Find a changing seed
	//2. Generate a random number in a range
	//3. Swap slice elements
	timeInInt64 := time.Now().UnixNano()	//time.Now returns a Time object, used to call UnixNano func returning a int64 number
	source := rand.NewSource(timeInInt64)
	randObj := rand.New(source)				//Creating a variable of type Rand, instead of using the rand package
// 	shortcut: randObj := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range d {
		newPosition := randObj.Intn(len(d)-1)			//Create a really random number with the unix nano time
		d[i], d[newPosition] = d[newPosition], d[i]		//swap slices
	}
}