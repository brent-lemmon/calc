package ui

import (
	"bufio"
	"fmt"
	"github.com/brent-lemmon/calc/pkg/rpn"
	"os"
)

func Start() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	input, err := reader.ReadString('\n')
	for input != "q\r\n" && input != "quit\r\n" && input != "exit\r\n" {
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		tokens, err := rpn.ToRpn(input)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			output, err := rpn.Evaluate(tokens)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("= %f\n", output)
			}
		}
		fmt.Print("> ")
		input, err = reader.ReadString('\n')
	}
}
