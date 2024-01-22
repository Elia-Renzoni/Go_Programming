/**
*	@author Elia Renzoni
*	@date 22/01/2024
*	@brief OOP in Golang
*
**/

package glovofiles

type Order struct {
	User
	Restaurant
	order []string
	price int
}

// l'ordinazione puó avere massimo 10 richieste
func (o *Order) DishesOrder(order []string) {
	o.order = make([]string, 10)

	defer func() {
		for _, value := range o.order {
			for key, mapValue := range o.Restaurant.menu {
				if value == mapValue {
					o.price += key
				}
			}
		}
	}()

	for index := range o.order {
		o.order[index] = order[index] // NOTE : possibile crash del programma
	}
}
