package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func main() {
	// Load environment variables
	dialect := "postgres"   // os.Getenv("DIALECT")
	dbHost := "localhost"   // os.Getenv("DB_HOST")
	dbPort := "5432"        // os.Getenv("DB_PORT")
	dbUser := "mufyaheer"   // os.Getenv("DB_USER")
	dbName := "book_keeper" // os.Getenv("DB_NAME")
	dbPassword := "secret"  // os.Getenv("DB_PASSWORD")

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, dbUser, dbName, dbPassword, dbPort)

	fmt.Println(dbURI)

	// Open database connection
	db, err = gorm.Open(dialect, dbURI)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Connected to database successfully")

	// Make migrations
	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Book{})

	/*----------- API routes ------------*/
	router := mux.NewRouter()

	// books
	router.HandleFunc("/books", CreateBook).Methods("POST")
	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")

	// people
	router.HandleFunc("/people", CreatePerson).Methods("POST")
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
