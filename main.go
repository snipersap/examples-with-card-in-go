package main	//This file belongs to the package "main", which makes it executable
import "fmt"	//Import all functions from a reusable or library package

//OR import (	//alternative syntax to import multiple packages at once
//	"fmt"
//  "String"
//)

func newCard(cardName string)string  {			//New function, in the same package and file.
	//Does not need to be imported 
return cardName
}

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

//Part 4: Create a new Deck with the help of a function - 4 numbers of all 4 suites
	myNewDeck := newDeck()
	myNewDeck.print()

//Part 5: Create a deck of cards, create a hand and also print the remaining cards in the deck
	part5Deck := newDeck()
	hand, leftDeck := deal(part5Deck, 3)		//Use := because we are initializing the variables before storing the return value
	fmt.Println("The hand dealt in the deal:")
	hand.print()								//Reuse the print function to print the values of the deck type
	fmt.Println("The remaining deck:")
	leftDeck.print()

//Part 6: Slice extraction example: 
	SampleFunc()

// Part 7: Convert deck to String
	// Create a new Deck, to avoid any overlap
	deckOfStrings := newDeck()
	fmt.Println(deckOfStrings.toString())	//Convert to string by calling the func
	// Prints: Ace of Hearts,Two of Hearts,Three of Hearts,Four of Hearts,Ace of Spades,Two of Spades,Three of Spades,Four of Spades,
	// Ace of Diamonds,Two of Diamonds,Three of Diamonds,Four of Diamonds,Ace of Clubs,Two of Clubs,Three of Clubs,Four of Clubs

//Part 8: Convert deck to string and then to byte slice
	//Use new Deck
	sampleDeck2 := newDeck()
	fmt.Println(sampleDeck2.toByteSlice())
// Prints: [65 99 101 32 111 102 32 72 101 97 114 116 115 44 84 119 111 32 111 102 32 72 101 97 114 116 115 44 84 104 114 101 101 
// 32 111 102 32 72 101 97 114 116 115 44 70 111 117 114 32 111 102 32 72 101 97 114 116 115 44 65 99 101 32 111 102 32 83 112 97
// 100 101 115 44 84 119 111 32 111 102 32 83 112 97 100 101 115 44 84 104 114 101 101 32 111 102 32 83 112 97 100 101 115 44 70 111 
// 117 114 32 111 102 32 83 112 97 100 101 115 44 65 99 101 32 111 102 32 68 105 97 109 111 110 100 115 44 84 119 111 32 111 102 32 
// 68 105 97 109 111 110 100 115 44 84 104 114 101 101 32 111 102 32 68 105 97 109 111 110 100 115 44 70 111 117 114 32 111 102 32 
// 68 105 97 109 111 110 100 115 44 65 99 101 32 111 102 32 67 108 117 98 115 44 84 119 111 32 111 102 32 67 108 117 98 115 44 84 104 
// 114 101 101 32 111 102 32 67 108 117 98 115 44 70 111 117 114 32 111 102 32 67 108 117 98 115]

//Part 9: Write the byte slice to file
	cardsToFile := newDeck()
	fmt.Println(cardsToFile.saveToFile("my_cards"))
	// Prints: <Nil> or the error
	cardsToFile.saveToFile("my_cards_2")		 
	// Prints nothing as return was not used. Therefore, collecting a returned value is optional

//Part 10: Reading from Hard Drive and Handle error
	cardsFromFile := newDeckFromFile("my_cards")
	cardsFromFile.print()	//Should print the deck read from the file


}