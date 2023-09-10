package main
import "fmt"

//Contains an example of how slice works

func SampleFunc() {											//Function names in lib packages should start with Caps.
	slice := []string{"index_0","index_1","index_2"}		//init slice
	slice = append(slice, "index_3", "index_4")				//append items to slice
	fmt.Println(slice)										//Prints the slice as [index_0 index_1 index_2 index_3 index_4]
	
//Extraction syntax: slice[startOfIndex:upToButNotIncludingIndex]

	subSlice := slice[0:2]									//Extract index 0 to < 2, i.e. 0 and 1
	fmt.Println(subSlice)									//Prints: [index_0 index_1]
	
	leftSlice := slice[:2]									//Same as slice[0:2]
	fmt.Println(leftSlice)									//Prints:[index_0 index_1]

	subSlice = slice[2:5]									//Extract index 2 to < 5, i.e. 2,3 and 4
	fmt.Println(subSlice)									//Prints: [index_2 index_3 index_4]

	rightSlice := slice[2:]									//Extract index 2 to end of the slice, which is same as the above subSlice
	fmt.Println(rightSlice)									//Prints: [index_2 index_3 index_4]
}
