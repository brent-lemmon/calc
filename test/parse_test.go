package test

import (
	"github.com/brent-lemmon/calc/pkg/rpn"
	"reflect"
	"testing"
)

func TestToRpnArithmetic(t *testing.T) {
	got, _ := rpn.ToRpn("3 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3")

	want := []string{"3", "4", "2", "*", "1", "5", "-", "2", "3", "^", "^", "/", "+"}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("ToRpn Arithmetic failed\n wanted: %v\n got:    %v\n", want, *got)
	}
}

func TestToRpnFunctions(t *testing.T) {
	got, _ := rpn.ToRpn("sin(max(2,3)/3*pi)")

	want := []string{"2", "3", "max", "3", "/", "pi", "*", "sin"}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("ToRpn functions failed\n wanted: %v\n got:    %v\n", want, *got)
	}
}

func TestToRpnMissingRightParen(t *testing.T) {
	_, err := rpn.ToRpn("max(sin(pi, 2)")

	if err == nil {
		t.Errorf("ToRpn failed to catch missing right parenthesis")
	}
}

func TestToRpnMissingLeftParen(t *testing.T) {
	_, err := rpn.ToRpn("max(sin pi),2)")

	if err == nil {
		t.Errorf("ToRpn failed to catch missing Left parenthesis")
	}
}

func TestToRpnDoubleOp(t *testing.T) {
	_, err := rpn.ToRpn("1+2**3")

	if err == nil {
		t.Errorf("ToRpn failed to catch back-to-back operators")
	}
}

func TestToRpnOpLeadingRightParen(t *testing.T) {
	_, err := rpn.ToRpn("sin(1+2*)")

	if err == nil {
		t.Errorf("ToRpn failed to catch operator leading right parenthesis")
	}
}
