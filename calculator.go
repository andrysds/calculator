package calculator

import (
	"regexp"
	"strconv"
	"strings"
)

// Calculate receives a calculator/math formula string and returns the result
func Calculate(formula string) string {
	formula = strings.ReplaceAll(formula, " ", "")
	if len(formula) > 0 && formula[0] != '-' {
		formula = "+" + formula
	}

	units := regexp.MustCompile(`(\+|-|\*|/)?\d+`).FindAllString(formula, -1)
	if len(units) == 0 {
		return "there is something wrong in the input"
	}

	result := float64(0)
	for _, u := range units {
		operator, number := u[0:1], u[1:len(u)]

		i, err := strconv.ParseFloat(number, 10)
		if err != nil {
			return err.Error()
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
			return "unknown symbol: %s " + operator
		}
	}

	return strconv.FormatFloat(result, 'f', -1, 64)
}
