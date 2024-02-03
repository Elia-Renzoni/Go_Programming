/**
*	@author Elia Renzoni
*	@date 03/02/2024
*	@brief OOP in Go - interfaces
 */

package duolingo

import "fmt"

type UserImpl struct {
	username, password, languageToLearn string
	UserManager
}

func (u *UserImpl) setUserName() {
	var control bool = true
	for control {
		fmt.Printf("Inserisci il tuo nome utente\n")
		_, err := fmt.Scanf("%s\n", &u.username)
		switch {
		case err != nil, u.username == "":
			panic(err)
		default:
			control = false
		}
	}
}

func (u *UserImpl) setPassword() {
	var control bool = true
	for control {
		fmt.Printf("Inserisci la tua password \n")
		_, err := fmt.Scanf("%s", &u.password)
		switch {
		case err != nil, u.password == "":
			panic(err)
		default:
			control = false
		}
	}
}

func (u *UserImpl) setLanguageToLearn() {
	var control bool = true
	for control {
		fmt.Printf("Inserisci la lingua che vuoi imparare \n")
		_, err := fmt.Scanf("%s\n", &u.languageToLearn)
		switch {
		case err != nil, u.languageToLearn == "":
			panic(err)
		default:
			control = false
		}
	}
}
