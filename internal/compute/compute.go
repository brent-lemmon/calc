package compute

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func popArg(nums *[]float64) (float64, error) {
	if len(*nums) < 1 {
		return 0, errors.New("no arg")
	}
	arg := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-1]
	return arg, nil
}

func popTwoArgs(nums *[]float64) (float64, float64, error) {
	if len(*nums) < 2 {
		return 0, 0, errors.New("not enough args")
	}
	arg1 := (*nums)[len(*nums)-2]
	arg2 := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-2]
	return arg1, arg2, nil
}

func EvaluateRpn(rpn *[]string) (float64, error) {
	//fmt.Printf("EvaluateRpn received: %v\n", *rpn)
	nums := make([]float64, 0, len(*rpn))
	for _, tok := range *rpn {
		switch tok {
		case "pi":
			nums = append(nums, math.Pi)
		case "+":
			arg1, arg2, err := popTwoArgs(&nums)
			if err != nil {
				return 0, fmt.Errorf("%s for addition", err.Error())
			}
			nums = append(nums, arg1+arg2)
		case "-":
			arg1, arg2, err := popTwoArgs(&nums)
			if err != nil {
				return 0, fmt.Errorf("%s for subtraction", err.Error())
			}
			nums = append(nums, arg1-arg2)
		case "*":
			arg1, arg2, err := popTwoArgs(&nums)
			if err != nil {
				return 0, fmt.Errorf("%s for multiplication", err.Error())
			}
			nums = append(nums, arg1*arg2)
		case "/":
			arg1, arg2, err := popTwoArgs(&nums)
			if err != nil {
				return 0, fmt.Errorf("%s for division", err.Error())
			}
			nums = append(nums, arg1/arg2)
		case "^":
			arg1, arg2, err := popTwoArgs(&nums)
			if err != nil {
				return 0, fmt.Errorf("%s for power operation", err.Error())
			}
			nums = append(nums, math.Pow(arg1, arg2))
		case "sin":
			arg, err := popArg(&nums)
			if err != nil {
				return 0, fmt.Errorf("%s for sin operation", err.Error())
			}
			nums = append(nums, math.Sin(arg))
		case "max":
			arg1, arg2, err := popTwoArgs(&nums)
			if err != nil {
				return 0, fmt.Errorf("%s for power operation", err.Error())
			}
			nums = append(nums, math.Max(arg1, arg2))
		default:
			num, err := strconv.ParseFloat(tok, 64)
			if err != nil {
				return 0, fmt.Errorf("error evaluating expression at unsupported token %s", tok)
			}
			nums = append(nums, num)
		}
	}
	if len(nums) != 1 {
		return 0, fmt.Errorf("expected one number after evaluation but recieved %d: %v", len(nums), nums)
	}
	return nums[0], nil
}
