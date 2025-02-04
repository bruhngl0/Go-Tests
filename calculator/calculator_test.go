package calculator_test

import (
	"calculator"
	"testing"
)

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

func TestDivide(t *testing.T) {
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: 10, b: 5, want: 2},
		{a: 6, b: 3, want: 2},
		{a: 10, b: 0, want: 0},
	}

	for _, value := range testCases {
		got, err := calculator.Divide(value.a, value.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if value.want != got {
			t.Errorf("DIVIDE: (%f ,%f), want: %f, got: %f", value.a, value.b, value.want, got)
		}
	}
}
