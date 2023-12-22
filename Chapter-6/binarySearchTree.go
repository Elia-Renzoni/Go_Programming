/**
*	@author Elia Renzoni
*	@date 22/12/2023
*	@brief Binary Search Tree Implementation in Golang
*
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	valueToStore int
	leftChild    *Tree
	rightChild   *Tree
	root         *Tree
}

const maxRandomValue int = 50

func main() {
	var tree *Tree = new(Tree)

	rand.Seed(time.Now().Unix())

	tree.addNewNode(rand.Intn(maxRandomValue))
	tree.addNewNode(rand.Intn(maxRandomValue))
	tree.addNewNode(rand.Intn(maxRandomValue))
	tree.addNewNode(rand.Intn(maxRandomValue))
	tree.addNewNode(rand.Intn(maxRandomValue))
	tree.addNewNode(rand.Intn(maxRandomValue))
	tree.addNewNode(rand.Intn(maxRandomValue))

	visitBST(tree.root)
}

func (tree *Tree) addNewNode(value int) {
	var newNode *Tree = new(Tree)
	newNode.valueToStore = value
	newNode.leftChild, newNode.rightChild = nil, nil

	var (
		node       *Tree = tree.root
		parentNode *Tree = tree.root
	)
	for node != nil {
		parentNode = node
		if value < node.valueToStore {
			node = node.leftChild
		} else {
			node = node.rightChild
		}
	}

	if node == nil {
		if node == tree.root {
			tree.root = newNode
		} else {
			if value < parentNode.leftChild.valueToStore {
				parentNode.leftChild = newNode
			} else {
				parentNode.rightChild = newNode
			}
		}
	}
}

func visitBST(root *Tree) {
	if root != nil {
		visitBST(root.leftChild)
		fmt.Println("%d", root.valueToStore)
		visitBST(root.rightChild)
	}
}
