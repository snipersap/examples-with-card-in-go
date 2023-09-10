package main

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
	cardSuites := []string{"Hearts","Spades","Diamonds","Clubs"}
	cardValues := []string{"Ace","Two","Three","Four"}

	//Create deck as a combo of suit and values
	for _,suit := range cardSuites {
		for _,value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}