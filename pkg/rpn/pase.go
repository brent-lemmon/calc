package rpn

import (
	"errors"
	"fmt"
	"go/scanner"
	"go/token"
	"math"
)

func initScanner(in string) scanner.Scanner {
	var scnr scanner.Scanner
	src := []byte(in)
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	scnr.Init(file, src, nil, 0)
	return scnr
}

func getSupportedOperators() *map[token.Token]string {
	return &map[token.Token]string{
		token.ADD: "+",
		token.SUB: "-",
		token.MUL: "*",
		token.QUO: "/",
		token.XOR: "^",
	}
}

func getSupportedConstants() *map[string]float64 {
	return &map[string]float64{
		"pi": math.Pi,
	}
}

func getPrecedence(op string) int {
	switch op {
	case "+":
		return 1
	case "-":
		return 1
	case "*":
		return 2
	case "/":
		return 2
	case "^":
		return 3
	default:
		return 0
	}
}

func isSupportedFunction(fun string) bool {
	return fun == "sin" || fun == "max"
}

func isLeftAssociative(op string) bool {
	return op != "^"
}

func processOperators(rpn *[]string, ops *[]string, op string) {
	for len(*ops) > 0 && (*ops)[len(*ops)-1] != "(" &&
		(getPrecedence((*ops)[len(*ops)-1]) > getPrecedence(op) ||
			getPrecedence((*ops)[len(*ops)-1]) == getPrecedence(op) && isLeftAssociative(op)) { //
		*rpn = append(*rpn, (*ops)[len(*ops)-1])
		*ops = (*ops)[:len(*ops)-1]
	}
	*ops = append(*ops, op)
}

func processComma(rpn *[]string, ops *[]string) {
	for len(*ops) > 0 && (*ops)[len(*ops)-1] != "(" {
		*rpn = append(*rpn, (*ops)[len(*ops)-1])
		*ops = (*ops)[:len(*ops)-1]
	}
}

func processRightParen(rpn *[]string, ops *[]string) error {
	for i := len(*ops) - 1; i >= 0 && (*ops)[i] != "("; i-- {
		*rpn = append(*rpn, (*ops)[i])
		*ops = (*ops)[:i]
	}
	if len(*ops) > 0 && (*ops)[len(*ops)-1] == "(" {
		*ops = (*ops)[:len(*ops)-1]
		if len(*ops) > 0 && isSupportedFunction((*ops)[len(*ops)-1]) {
			*rpn = append(*rpn, (*ops)[len(*ops)-1])
			*ops = (*ops)[:len(*ops)-1]
		}
		return nil
	}
	return errors.New("mismatched parenthesis")
}

// ToRpn takes an input string and returns the inputs tokenized and in Reverse Polish Notation.
// See https://en.wikipedia.org/wiki/Shunting_yard_algorithm#The_algorithm_in_detail
func ToRpn(in string) (*[]string, error) {
	scnr := initScanner(in)
	supOps := getSupportedOperators()
	supConsts := getSupportedConstants()
	rpn := make([]string, 0, len(in))
	ops := make([]string, 0, len(in))
	_, tok, lit := scnr.Scan()
	prevOp := ""
	wasPrevOp := false
	for tok != token.EOF {
		op, isOp := (*supOps)[tok]
		//fmt.Printf("tok: %s    lit: %s\n", tok.String(), lit)
		switch {
		case tok == token.INT || tok == token.FLOAT:
			rpn = append(rpn, lit)
		case tok == token.STRING || tok == token.IDENT:
			_, isConst := (*supConsts)[lit]
			if isConst {
				rpn = append(rpn, lit)
			} else if isSupportedFunction(lit) {
				ops = append(ops, lit)
			} else {
				return nil, fmt.Errorf("error parsing invalid token %s", lit)
			}
		case isOp:
			if wasPrevOp {
				return nil, fmt.Errorf("error parsing back to back operators: %s%s\n", prevOp, op)
			}
			processOperators(&rpn, &ops, op)
		case tok == token.COMMA:
			if wasPrevOp {
				return nil, fmt.Errorf("error parsing '%s' followed by ','\n", prevOp)
			}
			processComma(&rpn, &ops)
		case tok == token.LPAREN:
			ops = append(ops, "(")
		case tok == token.RPAREN:
			err := processRightParen(&rpn, &ops)
			if err != nil {
				return nil, err
			}
		case tok == token.SEMICOLON: //Ignore
		default:
			return nil, fmt.Errorf("error parsing input at token %s with value %s\n", tok.String(), lit)
		}
		//fmt.Printf("RPN: %v, OPS: %v\n", rpn, ops)
		prevOp = op
		wasPrevOp = isOp
		_, tok, lit = scnr.Scan()
	}
	for len(ops) > 0 {
		if ops[len(ops)-1] == "(" {
			return nil, errors.New("mismatched parenthesis")
		}
		rpn = append(rpn, ops[len(ops)-1])
		ops = ops[:len(ops)-1]
	}
	return &rpn, nil
}
