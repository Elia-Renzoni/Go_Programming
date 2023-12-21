package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Queue struct {
	keyValue int
	nextNode *Queue
}

type QueueImportantElements struct {
	top    *Queue
	bottom *Queue
}

func main() {
	var queue *Queue = new(Queue)
	var utilityMethods *QueueImportantElements = new(QueueImportantElements)

	rand.Seed(time.Now().Unix())
	queue.addNewElement(rand.Intn(50))
	queue.addNewElement(rand.Intn(50))
	queue.addNewElement(rand.Intn(50))
	queue.addNewElement(rand.Intn(50))

	queue.visit()

	fmt.Println("Bottom : %d - Top : %d", utilityMethods.getBottomOfQueue().keyValue, utilityMethods.getTopOfQueue().keyValue)

}

func (queue *Queue) visit() {
	var node *Queue
	var utilityMethods = &QueueImportantElements{}
	for node = utilityMethods.getBottomOfQueue(); node != nil; node = node.nextNode {
		fmt.Println(node.keyValue)
	}
}

func (utiliyMethods *QueueImportantElements) getTopOfQueue() *Queue {
	return utiliyMethods.top
}

func (utilityMethods *QueueImportantElements) getBottomOfQueue() *Queue {
	return utilityMethods.bottom
}

func (queue *Queue) addNewElement(value int) {
	var (
		node           *Queue                  = new(Queue)
		utilityMethods *QueueImportantElements = &QueueImportantElements{}
		top            *Queue                  = utilityMethods.getTopOfQueue()
	)
	node.keyValue = value
	node.nextNode = nil
	if top != nil {
		top.nextNode = node
	} else {
		utilityMethods.bottom = node
	}
	utilityMethods.top = node
}
