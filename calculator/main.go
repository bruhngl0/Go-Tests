package main

import (
	"calculator/calculator"
	"fmt"
)

func main() {
	result, err := calculator.Divide(10, 5)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		fmt.Printf("answer to your question is %f", result)
	}
}
