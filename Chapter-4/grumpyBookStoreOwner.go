/*
*	@author Elia Renzoni
*	@date 16/08/2023
*	@brief go function ex.
**/

package main

import "fmt"

const (
	arrElem    int = 8
	increments     = 3
)

func main() {
	var (
		clients                []int                       = createArrayClients()
		grumpy                 []int                       = createArrayGrumpy()
		minutes                int                         = 3
		grumpyBookStoreOwnCopy func([]int, []int, int) int = grumpyBookStoreOwner
	)
	fmt.Printf("N Customer Sutisfied %d", grumpyBookStoreOwnCopy(clients, grumpy, minutes))
}

func createArrayClients() (array []int) {
	var (
		emptyArray [arrElem]int
		value      int
	)
	for index := range emptyArray {
		value += increments
		emptyArray[index] = value
	}
	return
}

func createArrayGrumpy() (array []int) {
	var (
		emptyArray [arrElem]int
		value      int
	)
	for index := range emptyArray {
		emptyArray[index] = value
		if value == 0 {
			value = 1
		} else {
			value = 0
		}
	}
	return
}

func grumpyBookStoreOwner(clients []int, grumpy []int, minutes int) (sutisfaction int) {
	var clientSutified int
	for i := 0; i < arrElem; i++ {
		if grumpy[i] == 0 || arrElem-i == minutes {
			if arrElem-i == minutes {
				for j := arrElem - i; j < arrElem; j++ {
					clientSutified += clients[j]
				}
			}
			clientSutified += clients[i]
		}
	}
	return
}
