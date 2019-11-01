package calculator

import "testing"

type calculateTestCase struct {
	formula string
	result  string
}

func TestCalculate(t *testing.T) {
	cases := []calculateTestCase{
		// wrong formula
		calculateTestCase{
			formula: "this is a wrong formula",
			result:  "there is something wrong in the input",
		},
		// a number
		calculateTestCase{
			formula: "3007",
			result:  "3007",
		},
		// addition
		calculateTestCase{
			formula: "3327 + 612",
			result:  "3939",
		},
		// substraction
		calculateTestCase{
			formula: "2524 - 879",
			result:  "1645",
		},
		// multiplication
		calculateTestCase{
			formula: "1052 * 4437",
			result:  "4667724",
		},
		// division
		calculateTestCase{
			formula: "3733 / 1063",
			result:  "3.5117591721542802",
		},
		// mix
		calculateTestCase{
			formula: "4574 + 1492 - 1061 * 456 / 197",
			result:  "11585.177664974619",
		},
	}

	for i, c := range cases {
		result := Calculate(c.formula)
		if result != c.result {
			t.Errorf("#%d Calculate(%s) = %s; want %s", i, c.formula, result, c.result)
		}
	}
}
