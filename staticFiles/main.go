package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static")) //serves files from static folder

	http.Handle("/static/", http.StripPrefix("/static/", fs)) //remove static prefix
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	log.Println("server starting at port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
