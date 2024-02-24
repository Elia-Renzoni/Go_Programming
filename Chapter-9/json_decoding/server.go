/**
*	@author Elia Renzoni
*	@date 2402/2024
*	@brief json decoder
**/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	name   string `json:"name"`
	author string `json:"book_author"`
	price  int    `json:"price"`
}

func (b *Book) writeInStruct(w http.ResponseWriter, req *http.Request) {
	var book Book
	decoder := json.NewDecoder(req.Body)          // io.Reader - prende dal body della richiesta il json
	if err := decoder.Decode(&book); err != nil { // io.Writer - lo scrive nella variabile rispettanto le json key
		fmt.Printf("Error !")
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("OK"))
	printStruct(book)
}

func printStruct(book Book) {
	fmt.Printf("Name : %s - Author : %s - Price : %d", book.name, book.author, book.price)
}

func main() {
	var book Book = Book{}
	http.HandleFunc("/books", book.writeInStruct)
	http.ListenAndServe(":8080", nil)
}
