/**
*	@author Elia Renzoni
*	@date 22/01/2024
*	@brief OOP in Golang
*
**/

package glovofiles

import "fmt"

type Restaurant struct {
	name string
	menu map[int]string
}

const MAX_VALUE int = 5

func setNewRestaurantName() (name string) {
	var control bool = true
	for control {
		fmt.Printf("Inserisci un nome di un ristorante : ")
		_, err := fmt.Scanf("%s", &name)
		if err != nil {
			panic("Errore nella funzione SetNewRestaurant")
		} else {
			control = false
		}
	}
	return
}

// ogni menu, di ogni ristorante, ha solo 5 piatti
func (r *Restaurant) StartRestaurant(prices [MAX_VALUE]int, menu [MAX_VALUE]string) {
	var index int
	r.name = setNewRestaurantName()
	r.menu = map[int]string{
		prices[index]:   menu[index],
		prices[index+1]: menu[index+1],
		prices[index+2]: menu[index+2],
		prices[index+3]: menu[index+3],
		prices[index+4]: menu[index+4],
	}
}
