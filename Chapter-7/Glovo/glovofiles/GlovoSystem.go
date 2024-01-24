/**
*	@author Elia Renzoni
*	@date 22/01/2024
*	@brief OOP in Golang
*
**/

package glovofiles

import "fmt"

type GlovoSystem struct {
	user      User
	order     Order
	nextOrder *GlovoSystem
	head      *GlovoSystem
}

func (g *GlovoSystem) AddOrderToStackOrders(user *User, order *Order) {
	var node *GlovoSystem = new(GlovoSystem)
	node.user = *user
	node.order = *order

	if g.head == nil {
		g.head = node
	} else {
		node.nextOrder = g.head
		g.head = node
	}
}

func (g GlovoSystem) ViewAllOrders() {
	for node := g.head; node != nil; node = node.nextOrder {
		func() {
			for _, value := range node.order.order {
				fmt.Printf("%s\t", value)
			}
		}()
		fmt.Printf("Price : %d\n", node.order.price)
		fmt.Printf("User name : %s\n", node.user.name)
		fmt.Printf("Restaurant Name : %s\n", node.order.Restaurant.name)
		fmt.Printf("User address : %s\n", node.user.address)
	}
}
