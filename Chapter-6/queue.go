package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Queue struct {
	keyValue int
	nextNode *Queue
	tail     *Queue
	head     *Queue
}

func main() {
	var queue *Queue = new(Queue)

	rand.Seed(time.Now().Unix())
	queue.addNewElement(rand.Intn(50))
	queue.addNewElement(rand.Intn(50))
	queue.addNewElement(rand.Intn(50))
	queue.addNewElement(rand.Intn(50))

	queue.visit()

	fmt.Println("Bottom : %d - Top : %d", queue.head.keyValue, queue.tail.keyValue)

}

func (queue *Queue) visit() {
	var node *Queue
	for node = queue.head; node != nil; node = node.nextNode {
		fmt.Println(node.keyValue)
	}
}

func (queue *Queue) getTail() *Queue {
	return queue.tail
}

func (queue *Queue) getHead() *Queue {
	return queue.head
}

func (queue *Queue) addNewElement(value int) {
	var (
		node *Queue = new(Queue)
		top  *Queue = queue.getTail()
	)
	node.keyValue = value
	node.nextNode = nil
	if top != nil {
		queue.tail.nextNode = node
	} else {
		queue.head = node
	}
	queue.tail = node
}
