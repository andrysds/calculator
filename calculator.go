package calculator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Calculate receives a calculator/math formula string and returns the result
func Calculate(formula string) (string, error) {
	formula = strings.ReplaceAll(formula, " ", "")
	if len(formula) > 0 && formula[0] != '-' {
		formula = "+" + formula
	}

	result := float64(0)
	units := regexp.MustCompile(`(\+|-|\*|/)?\d+`).FindAllString(formula, -1)
	for _, u := range units {
		operator, number := u[0:1], u[1:len(u)]

		i, err := strconv.ParseFloat(number, 10)
		if err != nil {
			return "0", err
		}

		switch operator {
		case "+":
			result += i
		case "-":
			result -= i
		case "*":
			result *= i
		case "/":
			result /= i
		default:
			return "0", fmt.Errorf("unknown symbol: %s", operator)
		}
	}

	return strconv.FormatFloat(result, 'f', -1, 64), nil
}
