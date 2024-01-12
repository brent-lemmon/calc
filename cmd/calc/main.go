package main

import (
	"bufio"
	"calc/internal/compute"
	"calc/internal/parse"
	"fmt"
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

	rpn, err := parse.Rpn(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	output, err := compute.Evaluate(rpn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("RPN: %v\n", output)
	}
}
