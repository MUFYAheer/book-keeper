package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Book struct
type Book struct {
	gorm.Model

	Title      string
	Author     string
	CallNumber int `gorm:"unique_index"`
	PersonID   int
}

// GetBook returns book by id
func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book Book

	db.First(&book, params["id"])

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&book)
}

// GetBooks returns all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book

	db.Find(&books)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&books)
}

// CreateBook creates new book record in db
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	createdBook := db.Create(&book)
	err = createdBook.Error
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&createdBook)
}

// DeleteBook deletes book by id
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var book Book

	db.First(&book, params["id"])
	db.Delete(&book)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&book)
}
