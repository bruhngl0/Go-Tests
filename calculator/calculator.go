package calculator

import (
	"errors"
	"math"
)

func Add(a, b float64) float64 {
	return (a + b)
}

func Sub(a, b float64) float64 {
	return (a - b)
}

func Multiply(a, b float64) float64 {
	return (a * b)
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("number not divided by zero")
	}

	return (a / b), nil
}

func Sqrt(a float64) (float64, error) {

	if a < 0 {
		return 0, errors.New("not a valid input, +ve integers allowed")
	}
	number := math.Sqrt(a)
	return number, nil
}
