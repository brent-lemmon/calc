package rpn

import (
	"errors"
	"fmt"
	"go/scanner"
	"go/token"
	"math"
	"strconv"
)

// operators map go/Token values to their corresponding strings for all supported operators
var operators = map[token.Token]string{
	token.ADD: "+",
	token.SUB: "-",
	token.MUL: "*",
	token.QUO: "/",
	token.XOR: "^",
}

// operatorPrecedences map all supported operators to their corresponding precedence
var operatorPrecedences = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
	"^": 3,
}

// rightAssociativeOperators lists all supported left associative operators
var rightAssociativeOperators = []string{"^"}

// functions lists of all supported functions
var functions = []string{"sin", "cos", "tan", "max", "min"}

// constants maps all supported constants to their respective values
var constants = map[string]float64{
	"pi": math.Pi,
}

// comparePrecedence evaluates the precedence difference between operators
func comparePrecedence(op1 string, op2 string) (int, error) {
	prec1, found := operatorPrecedences[op1]
	if !found {
		return 0, fmt.Errorf("no operator precedence found for '%s'", op1)
	}
	prec2, found := operatorPrecedences[op2]
	if !found {
		return 0, fmt.Errorf("no operator precedence found for '%s'", op1)
	}
	return prec1 - prec2, nil
}

// isNumber evaluates whether the supplied string is a number or constant
func isNumber(num string) bool {
	_, parseErr := strconv.ParseFloat(num, 64)
	_, isConst := constants[num]
	return parseErr == nil || isConst
}

// isFunction evaluates whether the supplied string is a supported function
func isFunction(fun string) bool {
	for _, function := range functions {
		if function == fun {
			return true
		}
	}
	return false
}

// isLeftAssociative evaluates whether the supplied string is a left associative operator
func isLeftAssociative(op string) bool {
	for _, operator := range rightAssociativeOperators {
		if operator == op {
			return false
		}
	}
	return true
}

// processOperator processes an operator according to the Shunting Yard Algorithm
func processOperator(rpn *[]string, ops *[]string, op string) error {
	for len(*ops) > 0 && (*ops)[len(*ops)-1] != "(" {
		diff, err := comparePrecedence((*ops)[len(*ops)-1], op)
		if err != nil {
			return err
		}
		if diff > 0 || (diff == 0 && isLeftAssociative(op)) {
			*rpn = append(*rpn, (*ops)[len(*ops)-1])
			*ops = (*ops)[:len(*ops)-1]
		} else {
			break
		}
	}
	*ops = append(*ops, op)
	return nil
}

// processComma processes a comma according to the Shunting Yard Algorithm
func processComma(rpn *[]string, ops *[]string) {
	for len(*ops) > 0 && (*ops)[len(*ops)-1] != "(" {
		*rpn = append(*rpn, (*ops)[len(*ops)-1])
		*ops = (*ops)[:len(*ops)-1]
	}
}

// processRightParen processes a right parenthesis according to the Shunting Yard Algorithm
func processRightParen(rpn *[]string, ops *[]string) error {
	for i := len(*ops) - 1; i >= 0 && (*ops)[i] != "("; i-- {
		*rpn = append(*rpn, (*ops)[i])
		*ops = (*ops)[:i]
	}
	if len(*ops) > 0 && (*ops)[len(*ops)-1] == "(" {
		*ops = (*ops)[:len(*ops)-1]
		if len(*ops) > 0 && isFunction((*ops)[len(*ops)-1]) {
			*rpn = append(*rpn, (*ops)[len(*ops)-1])
			*ops = (*ops)[:len(*ops)-1]
		}
		return nil
	}
	return errors.New("mismatched parenthesis")
}

// validate checks whether the previous and current token are a valid input sequence
func validate(tok token.Token, lit string, op string, isOp bool, prev string, wasPrevOp bool) error {
	switch {
	case isNumber(prev) || prev == ")":
		if isNumber(lit) || isFunction(lit) {
			return fmt.Errorf("invalid input sequence '%s' '%s'", prev, lit)
		}
		if tok == token.LPAREN {
			return fmt.Errorf("invalid input sequence '%s' '('", prev)
		}
	case isFunction(prev):
		if isNumber(lit) || isFunction(lit) {
			return fmt.Errorf("invalid input sequence '%s' '%s'", prev, lit)
		}
	case wasPrevOp || prev == "," || prev == "(":
		if isOp {
			return fmt.Errorf("invalid input sequence '%s' '%s'", prev, op)
		}
		if tok == token.RPAREN {
			return fmt.Errorf("invalid input sequence '%s' ')'", prev)
		}
		if tok == token.COMMA {
			return fmt.Errorf("invalid input sequence '%s' ','", prev)
		}
	}
	return nil
}

// initScanner creates a scanner to read the provided string
func initScanner(in string) scanner.Scanner {
	var scnr scanner.Scanner
	src := []byte(in)
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	scnr.Init(file, src, nil, 0)
	return scnr
}

// ToRpn takes an input string and returns the inputs tokenized and in Reverse Polish Notation.
// See https://en.wikipedia.org/wiki/Shunting_yard_algorithm#The_algorithm_in_detail
func ToRpn(in string) (*[]string, error) {
	scnr := initScanner(in)
	rpn := make([]string, 0, len(in))
	ops := make([]string, 0, len(in))
	_, tok, lit := scnr.Scan()
	prev := ""
	wasPrevOp := false
	for tok != token.EOF {
		op, isOp := operators[tok]
		err := validate(tok, lit, op, isOp, prev, wasPrevOp)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("tok: %s    lit: %s\n", tok.String(), lit)
		switch {
		case isNumber(lit):
			rpn = append(rpn, lit)
			prev = lit
		case isFunction(lit):
			ops = append(ops, lit)
			prev = lit
		case isOp:
			err := processOperator(&rpn, &ops, op)
			if err != nil {
				return nil, err
			}
			prev = op
		case tok == token.COMMA:
			processComma(&rpn, &ops)
			prev = ","
		case tok == token.LPAREN:
			ops = append(ops, "(")
			prev = "("
		case tok == token.RPAREN:
			err := processRightParen(&rpn, &ops)
			if err != nil {
				return nil, err
			}
			prev = ")"
		case tok == token.SEMICOLON: //Ignore
		default:
			return nil, fmt.Errorf("error parsing input at token %s with value %s", tok.String(), lit)
		}
		wasPrevOp = isOp
		//fmt.Printf("RPN: %v, OPS: %v\n", rpn, ops)
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
