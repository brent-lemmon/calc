package ui

import (
	"bufio"
	"fmt"
	"github.com/brent-lemmon/calc/pkg/rpn"
	"math"
	"os"
)

func DisplayResult(res float64) {
	if math.Trunc(res) == res {
		fmt.Printf("= %d\n", int(res))
	} else {
		fmt.Printf("= %f\n", res)
	}
}

func DisplayError(err error) {
	fmt.Printf("X %s\n", err.Error())
}

func Start() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	input, err := reader.ReadString('\n')
	for input != "q\r\n" && input != "quit\r\n" && input != "exit\r\n" {
		if err != nil {
			DisplayError(err)
			break
		}
		tokens, err := rpn.ToRpn(input)
		if err != nil {
			DisplayError(err)
		} else {
			res, err := rpn.Evaluate(tokens)
			if err != nil {
				DisplayError(err)
			} else {
				DisplayResult(res)
			}
		}
		fmt.Print("> ")
		input, err = reader.ReadString('\n')
	}
}
