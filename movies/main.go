package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json: "director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
	fmt.Fprintln(w, "get movies")
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = "23445"
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
	fmt.Fprintln(w, "add a movie")
}

func getAnyMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Tyoe", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	fmt.Fprintln(w, "get any movie")
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
		fmt.Fprintln(w, "movie deleted")
	}
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "update a movie")
}

func main() {
	mux := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "2345", Title: "three idiots", Director: &Director{FirstName: "lucky", LastName: "bhist"}})
	movies = append(movies, Movie{ID: "2", Isbn: "4567", Title: "hangover", Director: &Director{FirstName: "ashutosh", LastName: "gavarikar"}})

	mux.Handle("/movies", http.HandlerFunc(getMovies)).Methods("GET")
	mux.Handle("/movies/{id}", http.HandlerFunc(getAnyMovie)).Methods("GET")
	mux.Handle("/movies", http.HandlerFunc(addMovie)).Methods("POST")
	mux.Handle("/movies/{id}", http.HandlerFunc(deleteMovie)).Methods("DELETE")
	mux.Handle("/movies/{id}", http.HandlerFunc(updateMovie)).Methods("UPDATE")

	fmt.Println("server starting at :8080")
	http.ListenAndServe(":8080", mux)

}
