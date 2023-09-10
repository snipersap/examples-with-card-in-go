package main	//This file belongs to the package "main", which makes it executable
import "fmt"	//Import all functions from a reusable or library package

//OR import (	//alternative syntax to import multiple packages at once
//	"fmt"
//  "String"
//)

func  main()  {											//Execution starts here (mandatory for executable file)

//Part 1: Print our cards, which is a slice of string, one by one
	//Create new Slice of cards and initialize them
	cards := []string {"2 of Spades", newCard("3 of Diamonds")}
	
	//Append new cards to the cards Slice
	cards = append(cards, newCard("4 of Hearts"))

	//print the cards in the slice
	fmt.Println("The Slice of cards:",cards)

	//Loop over and print each card
	for i,card := range cards {
		fmt.Println(i,">",card)
	}

//Part 2: Declare a collection of cards from the type deck and print them out one by one
	deckOfCards := deck{"2 of Spades", newCard("3 of diamonds")}
	deckOfCards = append(deckOfCards, newCard("4 of Hearts"))
	fmt.Println("The full deck of cards is:",deckOfCards)
	deckOfCards.print()

//Part 3: demonstrate the use of type with single data types in go
	var theSizeOfDeck deckSize = 3	//Full format of initialization
	sizeOfDeck := deckSize(2)		//Shortcut for initialization with custom type
	sizeOfDeck.showDeckSize()		//Call receiver function
	theSizeOfDeck.showDeckSize()

	nameOfDeck := deckName("My Deck")
	name := nameOfDeck.getDeckName()
	fmt.Println("Name of the deck is:", name)

}

func newCard(cardName string)string  {			//New function, in the same package and file.
												//Does not need to be imported 
		return cardName
}-