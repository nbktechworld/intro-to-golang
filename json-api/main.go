package main

import (
	"fmt"
	"net/http"

	"github.com/nbktechworld/json-api/math3"
)

func handleFor(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler(w, r)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// This is a catch-all route if no other matches
		w.Header().Set("X-Whatever", "something")
	})

	http.HandleFunc("GET /books", handleFor(getBooks))
	http.HandleFunc("GET /books/{bookId}", handleFor(handleSingleBook))
	Setup()

	fmt.Println(math3.Sum(1, 2))

	// Start up a server
	http.ListenAndServe("localhost:8080", nil)
}
