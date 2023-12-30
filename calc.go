package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		// Strip off any whitespace & combine all input to one string
		input := strings.Replace(strings.Join(os.Args[1:], ""), " ", "", -1)
		fmt.Printf("Thanks for your input of %s, but we are not yet able to process it", input)
	} else {
		fmt.Println("You must specify an expression")
	}
}
