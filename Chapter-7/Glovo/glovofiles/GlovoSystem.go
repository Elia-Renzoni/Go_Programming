/**
*	@author Elia Renzoni
*	@date 22/01/2024
*	@brief OOP in Golang
*
**/

package glovofiles

type GlovoSystem struct {
	user      User
	order     Order
	nextOrder *GlovoSystem
	head      *GlovoSystem
}

func (g *GlovoSystem) AddOrderToStackOrders(user User, order Order) {
	var node *GlovoSystem = new(GlovoSystem)
	node.user = user
	node.order = order

	if g.head == node {
		g.head = node
	} else {
		g.head.nextOrder = node
		g.head = node
	}
}

func (g *GlovoSystem) ViewAllOrders() {
	for node := g.head; node != nil; node = node.nextOrder {
		// TODO
	}
}
