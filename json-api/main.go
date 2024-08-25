package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nbktechworld/json-api/math3"
)

type Book struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func handleSingleBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var book Book = Book{
	// 	Id:    1,
	// 	Title: "Book Title 1",
	// }

	book2 := map[string]interface{}{
		"id":    1,
		"title": "Book TItle 2",
	}
	encodedBook, err := json.Marshal(book2)
	if err != nil {
		log.Fatal("Couldnt marshal the book:", err)
	}
	w.Write(encodedBook)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// This is a catch-all route if no other matches
		w.Header().Set("X-Whatever", "something")
	})

	http.HandleFunc("GET /books/1", handleSingleBook)
	Setup()

	fmt.Println(math3.Sum(1, 2))

	// Start up a server
	http.ListenAndServe(":8080", nil)
}
