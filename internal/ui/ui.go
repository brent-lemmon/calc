package ui

import (
	"bufio"
	"fmt"
	"github.com/brent-lemmon/calc/pkg/rpn"
	"math"
	"os"
)

var red = "\033[31m"
var green = "\033[32m"
var resetColor = "\033[0m"

func displayPrompt() {
	fmt.Print(resetColor, "> ")
}

func DisplayResult(res float64) {
	var output string
	if math.Trunc(res) == res {
		output = fmt.Sprintf("= %d\n", int(res))
	} else {
		output = fmt.Sprintf("= %f\n", res)
	}
	fmt.Print(green, output)
}

func DisplayError(err error) {
	output := fmt.Sprintf("X %s\n", err.Error())
	fmt.Print(red, output)
}

func Start() {
	reader := bufio.NewReader(os.Stdin)
	displayPrompt()
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
		displayPrompt()
		input, err = reader.ReadString('\n')
	}
}
