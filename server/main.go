package main

import (
	"fmt"
	"net/http"
)

type Homehandler struct{}

func (h Homehandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "custom handler")
}

func main() {

	mux := http.NewServeMux()

	mux.Handle("/custom", Homehandler{})

	mux.HandleFunc("/life", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{status: ok}`)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "yo lets have fun at holi") // ✅ Fix: Writing response correctly
	})

	mux.HandleFunc("/holi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "event")
	})

	fmt.Println("starting server at port :8080")

	http.ListenAndServe(":8080", mux) // Start server
}
