/**
*	@author Elia Renzoni
*	@date 22/01/2024
*	@brief OOP in Golang
 */

package glovofiles

import "fmt"

type User struct {
	name, address string
}

func (u *User) InsertNewName() {
	var control bool = true
	for control {
		fmt.Printf("Inserisci un nuovo nome : \n")
		_, err := fmt.Scanf("%s\n", &u.name)
		if err != nil {
			panic(err)
		} else {
			control = false
		}
	}
}

func (u *User) InsertNewAddress() {
	var control bool = true
	for control {
		fmt.Printf("Inserisci un nuovo indirizzo : \n")
		_, err := fmt.Scanf("%s\n", &u.address)
		if err != nil {
			panic(err)
		} else {
			control = false
		}
	}
}

func (u User) getUserNameAndAddress() (string, string) {
	return u.name, u.address
}
