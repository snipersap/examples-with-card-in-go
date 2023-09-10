package main	//This file belongs to the package "main", which makes it executable
import "fmt"	//Import all functions from a reusable or library package

//OR import (	//alternative syntax to import multiple packages at once
//	"fmt"
//  "String"
//)

func  main()  {											//Execution starts here (mandatory for executable file)
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


}
func newCard(cardName string)string  {			//New function, in the same package and file.
												//Does not need to be imported 
		return cardName
}