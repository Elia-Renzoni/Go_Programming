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
	bookInfo book
	nextNode *bookLoaded
	head     *bookLoaded
}

type Transactions struct {
	userInfo          user
	bookInfo          book
	date, hour        int
	typeOfTransaction string
	nextTransaction   *Transactions
	head              *Transactions
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
	name:     "Selene",
	surname:  "Venozni",
	idNumber: rand.Intn(50),
}

var user2 = user{
	name:     "Matteo",
	surname:  "Gramellini",
	idNumber: rand.Intn(50),
}
var user3 = user{
	name:     "Piero",
	surname:  "Vecchioni",
	idNumber: rand.Intn(50),
}

func main() {

	user1.bookLoaded = new(bookLoaded)
	user2.bookLoaded = new(bookLoaded)
	user3.bookLoaded = new(bookLoaded)

	user1.bookLoaded.addBookLoaned(book1)
	user2.bookLoaded.addBookLoaned(book2)
	user3.bookLoaded.addBookLoaned(book2)
	user3.bookLoaded.addBookLoaned(book3)

	user3.bookLoaded.visitBookLoaned()

	var transaction *Transactions = new(Transactions)

	transaction.addNewTransation(user1, book3, 23, 2, loan)

}

func (book *bookLoaded) addBookLoaned(newBook book) {
	var newNode *bookLoaded = new(bookLoaded)
	newNode.bookInfo = newBook

	if book.getLastNode() == nil {
		book.head = newNode
	} else {
		book.getLastNode().nextNode = newNode
	}
}

func (book *bookLoaded) visitBookLoaned() {
	for node := book.head; node != nil; node = node.nextNode {
		fmt.Println("Title : ", node.bookInfo.title+"  - Author : ", node.bookInfo.author)
	}
}

func (transation *Transactions) addNewTransation(myUs user, myBook book, date, hour int, transitionType string) {
	var newNode *Transactions = new(Transactions)
	newNode.userInfo = myUs
	newNode.bookInfo = myBook
	newNode.date = date
	newNode.hour = hour
	newNode.typeOfTransaction = transitionType

	if transation.getLastTransaction() == nil {
		transation.head = newNode
	} else {
		transation.getLastTransaction().nextTransaction = newNode
	}

	fmt.Printf("New Trasation Occur ! type %s\n", newNode.typeOfTransaction)
	fmt.Printf("Hour and Date : %d %d", newNode.hour, newNode.date)
}

func (book *bookLoaded) getLastNode() *bookLoaded {
	var toReturn *bookLoaded
	for node := book.head; node != nil; node = node.nextNode {
		toReturn = node
	}

	return toReturn
}

func (transactions *Transactions) getLastTransaction() *Transactions {
	var toReturn *Transactions
	for node := transactions.head; node != nil; node = node.nextTransaction {
		toReturn = node
	}

	return toReturn
}
