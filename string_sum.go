package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has (-|\+){0,1}\d+ of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	oper, count := Oper(input)
	if oper == "" && count > 1 {
		err = fmt.Errorf("%w", errorNotTwoOperands)
		return "", err
	}
	if oper == "" && count == 1 {
		err = fmt.Errorf("%w", errorEmptyInput)
		return "", err
	}
	num := strings.Split(input, oper)
	if len(num) != 2 {
		return "", errors.New("nums less or more then 2")
	}
	num1, err := strconv.Atoi(num[0])
	if err != nil {
		return "", err
	}
	num2, err := strconv.Atoi(num[1])
	if err != nil {
		return "", err
	}
	result := Math(num1, num2, oper)
	output = strconv.Itoa(result)
	return output, nil
}

func Oper(input string) (operad string, count int) {
	for _, w := range input {
		if w == '+' {
			operad = "+"
			count++
		} else if w == '-' {
			operad = "-"
			count++
		} else if w == '*' {
			operad = "*"
			count++
		} else if w == '%' {
			operad = "%"
			count++
		} else if w == '/' {
			operad = "/"
			count++
		}
	}

	if count > 1 {
		operad = ""
		return operad, count
	}
	return operad, count
}

func Math(num1, num2 int, oper string) (result int) {
	if oper == "+" {
		result = num1 + num2
	} else if oper == "-" {
		result = num1 - num2
	} else if oper == "/" {
		result = num1 / num2
	} else if oper == "%" {
		result = num1 % num2
	} else if oper == "*" {
		result = num1 * num2
	}
	return
}
