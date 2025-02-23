package main

import (
	"books/books" // Importing the custom package
	"fmt"
)

func main() {
	// Calling the Info function from books package
	result := books.Info("for the love of go", "second edition", 200)
	fmt.Printf("%s\n", result)
}
