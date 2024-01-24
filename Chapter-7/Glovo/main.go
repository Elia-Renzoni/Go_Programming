/**
*	@author Elia Renzoni
*	@date 22/01/2024
*	@brief OOP in Golang
**/

package main

import "glovomodule/glovofiles"

type arr [glovofiles.MAX_VALUE]int
type arrStr [glovofiles.MAX_VALUE]string

func main() {
	var (
		system                   = &glovofiles.GlovoSystem{}
		user1, user2, user3      = &glovofiles.User{}, &glovofiles.User{}, &glovofiles.User{}
		restaurant1, restaurant2 = &glovofiles.Restaurant{}, &glovofiles.Restaurant{}
		order1, order2, order3   = &glovofiles.Order{}, &glovofiles.Order{}, &glovofiles.Order{}
		prices                   = arr{12, 13, 5, 6, 12}
		menuDishesNames          = arrStr{"spaghetti alla carbonara",
			"filetto di Salmone alla Griglia",
			"Risotto ai Funghi Porcini",
			"Pollo al Curry con Riso Basmati",
			"Pizza Quattro Stagioni"}
	)

	user1.InsertNewName()
	user1.InsertNewAddress()
	user2.InsertNewName()
	user2.InsertNewAddress()
	user3.InsertNewName()
	user3.InsertNewAddress()

	restaurant1.StartRestaurant(prices, menuDishesNames)
	restaurant2.StartRestaurant(prices, menuDishesNames)

	order1.DishesOrder([]string{"spaghetti alla carbonara", "filetto di Salmone alla Griglia"})
	order2.DishesOrder([]string{"Pollo al Curry con Riso Basmati", "Pizza Quattro Stagioni"})
	order3.DishesOrder([]string{"Risotto ai Funghi Porcini", "filetto di Salmone alla Griglia"})

	system.AddOrderToStackOrders(user1, order1)
	system.AddOrderToStackOrders(user2, order2)
	system.AddOrderToStackOrders(user3, order3)

	system.ViewAllOrders()

}
