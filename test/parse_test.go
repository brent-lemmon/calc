package test

import (
	"calc/pkg/rpn"
	"reflect"
	"testing"
)

func TestRpnArithmetic(t *testing.T) {
	got, _ := rpn.ToRpn("3 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3")

	want := []string{"3", "4", "2", "*", "1", "5", "-", "2", "3", "^", "^", "/", "+"}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("ToRpn Arithmetic failed\n wanted: %v\n got:    %v\n", want, *got)
	}
}

func TestRpnFunctions(t *testing.T) {
	got, _ := rpn.ToRpn("sin(max(2,3)/3*pi)")

	want := []string{"2", "3", "max", "3", "/", "pi", "*", "sin"}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("ToRpn functions failed\n wanted: %v\n got:    %v\n", want, *got)
	}
}
