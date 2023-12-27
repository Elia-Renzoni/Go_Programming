/**
*
*	@author Elia Renzoni
*	@date 27/12/2023
*	@brief Library Management in Golang
 */

package main

import (
	"fmt"
	"math/rand"
)

const loan, refund string = "Loan", "Refund"

type book struct {
	title            string
	author           string
	pubblicationYear int
	availableCopies  int
	loanedCopies     int
}

type user struct {
	name       string
	surname    string
	idNumber   int
	bookLoaded *bookLoaded
}

type bookLoaded struct {
	bookInfo    book
	nextNode    *bookLoaded
	previusNode *bookLoaded
	head        *bookLoaded
}

type Transactions struct {
	userInfo          user
	bookInfo          book
	date, hour        int
	typeOfTransaction string
	nextTransaction   *Transactions
	head              *Transactions
	previusNode       *Transactions
}

/**
* books implementation
 */
var book1 = book{
	title:            "Mastering API Architecture",
	author:           "Yan Gordon",
	pubblicationYear: 2022,
	availableCopies:  500,
	loanedCopies:     30,
}

var book2 = book{
	title:            "Distribuited Services in Go",
	author:           "Gordon Rayn",
	pubblicationYear: 2021,
	availableCopies:  60,
	loanedCopies:     0,
}

var book3 = book{
	title:            "Learning Machine Learning with Rust",
	author:           "Sarah Huston",
	pubblicationYear: 2018,
	availableCopies:  500,
	loanedCopies:     400,
}

var book4 = book{
	title:            "Learning Java From 0 to Hero",
	author:           "Elia Renzoni",
	pubblicationYear: 2019,
	availableCopies:  20,
	loanedCopies:     4,
}

/**
*	user implementation
 */
var user1 = user{
	name:       "Selene",
	surname:    "Venozni",
	idNumber:   rand.Intn(50),
	bookLoaded: nil,
}

var user2 = user{
	name:       "Matteo",
	surname:    "Gramellini",
	idNumber:   rand.Intn(50),
	bookLoaded: nil,
}
var user3 = user{
	name:       "Piero",
	surname:    "Vecchioni",
	idNumber:   rand.Intn(50),
	bookLoaded: nil,
}

func main() {

}

func (book *bookLoaded) addBookLoaned(newBook book) {
	var newNode *bookLoaded = new(bookLoaded)
	newNode.bookInfo = newBook

	if newNode == book.head {
		book.head = newNode
	} else {
		book.previusNode.nextNode = newNode
	}

	book.previusNode = newNode
}

func (book *bookLoaded) visitBookLoaned() {
	for node := book.head; node != nil; node = node.nextNode {
		fmt.Println(book.bookInfo)
	}
}

func (transation *Transactions) addNewTransation(myUs user, myBook book, date, hour int, transitionType string) {
	var newNode *Transactions = new(Transactions)
	newNode.userInfo = myUs
	newNode.bookInfo = myBook
	newNode.date = date
	newNode.hour = hour
	newNode.typeOfTransaction = transitionType

	if newNode == transation.head {
		transation.head = newNode
	} else {
		transation.previusNode = newNode
		fmt.Printf("New Trasation Occur ! type %s", newNode.typeOfTransaction)
		fmt.Printf("Hour and Date : %d %d", newNode.hour, newNode.date)
	}

	transation.previusNode = newNode
}
