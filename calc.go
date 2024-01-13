package main

import (
	"bufio"
	"fmt"
	"github.com/brent-lemmon/calc/pkg/rpn"
	"os"
	"strings"
)

func main() {
	var input string
	var err error
	if len(os.Args) > 1 {
		input = strings.Join(os.Args[1:], "")
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter an expression to evaluate:\n> ")
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	tokens, err := rpn.ToRpn(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	output, err := rpn.Evaluate(tokens)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Value: %f\n", output)
	}
}
