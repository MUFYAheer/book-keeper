package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Person struct
type Person struct {
	gorm.Model

	Name  string
	Email string `gorm:"type=varchar(100);unique_index"`
	Books []Book
}

// GetPerson returns single person
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	var books []Book

	db.First(&person, params["id"])
	db.Model(&person).Related(&books)

	person.Books = books

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&person)
}

// GetPeople returns all people in db
func GetPeople(w http.ResponseWriter, r *http.Request) {
	var people []Person

	db.Find(&people)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&people)
}

// CreatePerson create a new record in db
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	json.NewDecoder(r.Body).Decode(&person)

	createdPerson := db.Create(&person)
	err = createdPerson.Error
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&createdPerson)
}

// DeletePerson deletes person from db
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var person Person

	db.First(&person, params["id"])
	db.Delete(&person)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&person)
}
