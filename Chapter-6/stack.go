/**
*	@author Elia Renzoni
*	@date 22/12/2023
*	@brief Stack Data Structure implemented in Golang
**/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Stack struct {
	keyValue int
	nextNode *Stack
	top      *Stack
}

func main() {
	var stack *Stack = new(Stack)
	rand.Seed(time.Now().Unix())
	stack.addValue(rand.Intn(50))

	stack.addValue(rand.Intn(50))
	stack.addValue(rand.Intn(50))
	stack.addValue(rand.Intn(50))
	stack.addValue(rand.Intn(50))

	stack.visitStack()

	fmt.Printf("Top of the Stack : %d", stack.top.keyValue)

}

func (stack *Stack) addValue(valueToStore int) {
	var newNode *Stack = &Stack{}
	newNode.keyValue = valueToStore

	if stack.top != nil {
		newNode.nextNode = stack.top
		stack.top = newNode
	} else {
		stack.top = newNode
	}
}

func (stack *Stack) visitStack() {
	for node := stack.top; node != nil; node = node.nextNode {
		fmt.Println(node.keyValue)
	}
}
