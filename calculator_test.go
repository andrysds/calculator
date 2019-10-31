package calculator

import "testing"

type calculateTestCase struct {
	formula string
	result  string
	err     error
}

func TestCalculate(t *testing.T) {
	cases := []calculateTestCase{
		// a number
		calculateTestCase{
			formula: "3007",
			result:  "3007",
			err:     nil,
		},
		// addition
		calculateTestCase{
			formula: "3327 + 612",
			result:  "3939",
			err:     nil,
		},
		// substraction
		calculateTestCase{
			formula: "2524 - 879",
			result:  "1645",
			err:     nil,
		},
		// multiplication
		calculateTestCase{
			formula: "1052 * 4437",
			result:  "4667724",
			err:     nil,
		},
		// division
		calculateTestCase{
			formula: "3733 / 1063",
			result:  "3.5117591721542802",
			err:     nil,
		},
		// mix
		calculateTestCase{
			formula: "4574 + 1492 - 1061 * 456 / 197",
			result:  "11585.177664974619",
			err:     nil,
		},
	}

	for i, c := range cases {
		result, err := Calculate(c.formula)
		if result != c.result {
			t.Errorf("#%d Calculate(%s) = %s; want %s", i, c.formula, result, c.result)
		}
		if err != c.err {
			t.Errorf("#%d error = %s, want %s", i, err, c.err)
		}
	}
}
