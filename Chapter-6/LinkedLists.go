package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	keyValue int
	nextNode *Node
}

type LinkedList struct {
	head *Node
}

func main() {
	var linkedlist = &LinkedList{}

	rand.Seed(time.Now().Unix())
	linkedlist.addNodeToLinkedList(rand.Intn(50))
	linkedlist.addNodeToLinkedList(rand.Intn(50))
	linkedlist.addNodeToLinkedList(rand.Intn(50))
	linkedlist.addNodeToLinkedList(rand.Intn(50))

	fmt.Println(linkedlist.head.keyValue)
	linkedlist.visitLinkedList()

}

func (linkedList *LinkedList) getLastNode() *Node {
	var lastNode *Node
	for node := linkedList.head; node != nil; node = node.nextNode {
		lastNode = node
	}
	return lastNode
}

func (linkedList *LinkedList) addNodeToLinkedList(valueToAdd int) {
	var newNode *Node = new(Node)
	newNode.keyValue = valueToAdd
	newNode.nextNode = nil
	if previusNode := linkedList.getLastNode(); previusNode == nil {
		linkedList.head = newNode
	} else {
		previusNode.nextNode = newNode
	}
}

func (linkedList *LinkedList) visitLinkedList() {
	var node *Node
	for node = linkedList.head; node != nil; node = node.nextNode {
		fmt.Println(node.keyValue)
	}
}
