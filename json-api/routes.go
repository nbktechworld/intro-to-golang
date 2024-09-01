package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func getBooks(w http.ResponseWriter, r *http.Request) {
	fileContent, fileContentError := os.ReadFile("db.json")
	if fileContentError != nil {
		log.Println(fileContentError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(fileContent)
}

func handleSingleBook(w http.ResponseWriter, r *http.Request) {
	var bookId string = r.PathValue("bookId")

	fileContent, fileContentError := os.ReadFile("db.json")
	if fileContentError != nil {
		log.Println(fileContentError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var books []Book
	booksError := json.Unmarshal(fileContent, &books)
	if booksError != nil {
		log.Println(booksError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var foundBook *Book
	for i := 0; i < len(books); i++ {
		book := books[i]
		if book.Id == bookId {
			foundBook = &book
			break
		}
	}
	if foundBook == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// var book Book = Book{
	// 	Id:    1,
	// 	Title: "Book Title 1",
	// }

	// book2 := map[string]interface{}{
	// 	"id":    1,
	// 	"title": "Book TItle 2",
	// }
	encodedBook, err := json.Marshal(foundBook)
	if err != nil {
		log.Println("Couldnt marshal the book:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(encodedBook)
}
