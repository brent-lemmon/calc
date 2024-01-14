package main

import (
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
			ui.DisplayError(err)
		} else {
			res, err := rpn.Evaluate(tokens)
			if err != nil {
				ui.DisplayError(err)
			} else {
				ui.DisplayResult(res)
			}
		}
	} else {
		ui.Start()
	}
}
