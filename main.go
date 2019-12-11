package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Book struct{
	ID int`json:id`
	Title string `json:title`
	Author string `json:author`
	Year string `json: year`
}

var books []Book

func main() {
	router := mux.NewRouter() 

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books/{id}", addBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")

}

