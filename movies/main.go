package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get movies")
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "add a movie")
}

func getAnyMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get any movie")
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "delete a movie")
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "update a movie")
}

func main() {
	mux := mux.NewRouter()

	mux.Handle("/movies", http.HandlerFunc(getMovies)).Methods("GET")
	mux.Handle("/movies/{id}", http.HandlerFunc(getAnyMovie)).Methods("GET")
	mux.Handle("/movies", http.HandlerFunc(addMovie)).Methods("POST")
	mux.Handle("/movies/{id}", http.HandlerFunc(deleteMovie)).Methods("DELETE")
	mux.Handle("/movies/{id}", http.HandlerFunc(updateMovie)).Methods("UPDATE")

	fmt.Println("server starting at :8080")
	http.ListenAndServe(":8080", mux)

}
