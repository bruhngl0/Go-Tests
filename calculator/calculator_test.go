package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

// test for ADDITION
func TestAdd(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 1, b: -1, want: 0},
		{a: 2, b: 3, want: 5},
		{a: 5, b: -1, want: 4},
	}

	for _, value := range testCases {
		got := calculator.Add(value.a, value.b)
		if value.want != got {
			t.Errorf("ADD: (%f, %f), want: %f, got: %f", value.a, value.b, value.want, got)
		}
	}
}

// test for SUBSTRACTION
func TestSub(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 1, b: 2, want: -1},
		{a: 3, b: 2, want: 1},
		{a: 4, b: 2, want: 2},
	}

	for _, value := range testCases {
		got := calculator.Sub(value.a, value.b)
		if value.want != got {
			t.Errorf("SUB: (%f %f) , want: %f, got: %f ", value.a, value.b, value.want, got)
		}
	}
}

// TEST FOR MULTIPLY
func TestMultiply(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 1},
		{a: 3, b: 2, want: 6},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add: (%f, %f), want: %f, got: %f", tc.a, tc.b, tc.want, got)
		}
	}

}

//TEST FOR DIVIDE

// one behaviour one test---------- only for +integers!!
func TestDivide(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 3, b: 1, want: 3},
		{a: 10, b: 5, want: 2},
		{a: 49, b: 7, want: 7},
		{a: 56, b: 8, want: 7},
		{a: 63, b: 7, want: 9},
		{a: 24, b: 4, want: 6},
	}

	for _, value := range testCases {
		got, err := calculator.Divide(value.a, value.b)

		if err != nil {
			t.Fatalf("invalid input")
		}

		if got != value.want {
			t.Errorf("DIVIDE (%f, %f), want: %f , got : %f", value.a, value.b, value.want, got)
		}

	}
}

// one behaviour one test--- Divide fn for b=0, (invalid input)
func TestDivideInvalidInput(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
	}
	testCases := []testCase{
		{a: 3, b: 0},
		{a: 5, b: 0},
	}
	for _, value := range testCases {
		_, err := calculator.Divide(value.a, value.b)

		if err == nil {
			t.Errorf("no error detected")
		}
	}
}

//one behaviour one test-- cases where a<b--tolerence define

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestDivideTolerance(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64

		want float64
	}

	testCases := []testCase{
		{a: 1, b: 3, want: 0.333},
		{a: 1, b: 3, want: 0.333},
		{a: 1, b: 3, want: 0.333},
	}

	for _, value := range testCases {

		got, _ := calculator.Divide(value.a, value.b)
		result := closeEnough(value.want, got, 0.01)
		if !result {
			t.Errorf("(DIVIDE A/B: %f, %f), want: %f, got: %f", value.a, value.b, value.want, got)
		}

	}

}

// TEST FOR SQRT FN

// helper fn for tolrence
func closeEnoughSqrt(a, b, tolerance float64) bool {
	result := math.Abs(a-b) <= tolerance
	return result

}

// one behaviour -one test--+ve integers
func TestSqrt(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a    float64
		want float64
	}
	testCases := []testCase{
		{a: 49, want: 7},
		{a: 100, want: 10},
		{a: 10, want: 3.163},
	}

	for _, value := range testCases {
		got, _ := calculator.Sqrt(value.a)

		//input the closeEnough function here--compare value
		result := closeEnoughSqrt(value.want, got, 0.01)
		if !result {
			t.Errorf("SQRT: %f, want: %f, got: %f", value.a, value.want, got)
		}

	}
}

//one behaviour-one test--for invalidInput(-ve integers)

func TestInvalidInputSqrt(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a float64
	}

	testCases := []testCase{
		{a: -4},
		{a: -8},
		{a: -3},
	}

	for _, value := range testCases {

		_, err := calculator.Sqrt(value.a)
		if err == nil {
			t.Errorf("no error detected+error is there")
		}
	}

}
