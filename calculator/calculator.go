package calculator

import "errors"

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
		return 0, errors.New("cannot divide by zero")
	}
	return (a / b), nil
}
