package main

import (
	"fmt"
	"github.com/brent-lemmon/calc/internal/ui"
	"github.com/brent-lemmon/calc/pkg/rpn"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		input := strings.Join(os.Args[1:], "")
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
	} else {
		ui.Start()
	}
}
