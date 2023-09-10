package main

//Create a new type to represent a deck of cards as slice of string
type deck []string

func (receivedDeck deck) print(){
	for i,card := range receivedDeck {
		fmt.Println(i,">",card)
	}
}