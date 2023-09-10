package main
import (
	"fmt"
	"os"
)

func isErr(err error) {
	if err != Nil {
		return true
	}
}

func handleReadFromFileErr(err error) {
	fmt.Println("Error:", err)	//Print the error
	os.Exit(1)	//Exit with non-zero error code
}