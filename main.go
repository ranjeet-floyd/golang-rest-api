package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Ranjeet

// Book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func main() {
	r := mux.NewRouter()

	// Mock data
	books = append(books, Book{ID: "1", Isbn: "4242424",
		Title: "title Book 1", Author: &Author{
			FirstName: "Ranjeet", LastName: "Kumar"}})
	books = append(books, Book{ID: "2", Isbn: "4242425",
		Title: "title Book 2", Author: &Author{
			FirstName: "Ranjeet2", LastName: "Kumar2"}})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	fmt.Println("Starting the application....")
	// http server
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Init books var as a slice Book struct
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r) // get rest param

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})

}

func createBook(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Create book....")
	w.Header().Set("Content-type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	log.Fatal(" book.... {}", book)
	book.ID = strconv.Itoa(rand.Intn(100000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	params := mux.Vars(r)

	book.ID = params["id"]

	for index, item := range books {
		if item.ID == book.ID {
			books = append(books[:index], books[index+1:]...)
			books = append(books, book)
			break
		}
	}
	json.NewEncoder(w).Encode(book)

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	id := params["id"]

	for index, item := range books {
		if item.ID == id {
			books = append(books[:index], books[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(books)
}
