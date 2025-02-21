package main

import (
	"fmt"
	"books/books" // Importing the custom package
)

func main() {
	// Calling the Info function from books package
	result := books.Info("for the love of go", "second edition" , 99)
	fmt.Printf("%s\n", result)
}