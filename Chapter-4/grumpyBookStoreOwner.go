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
		clients                [arrElem]int                              = [arrElem]int{12, 34, 5, 12, 4, 6, 7}
		grumpy                 [arrElem]int                              = [arrElem]int{4, 5, 1, 5, 6, 7, 2, 1}
		minutes                int                                       = 3
		grumpyBookStoreOwnCopy func([arrElem]int, [arrElem]int, int) int = grumpyBookStoreOwner
	)
	fmt.Printf("N Customer Sutisfied %d", grumpyBookStoreOwnCopy(clients, grumpy, minutes))
}

func grumpyBookStoreOwner(clients [arrElem]int, grumpy [arrElem]int, minutes int) (sutisfaction int) {
	for i := 0; i < arrElem; i++ {
		if grumpy[i] == 0 || arrElem-i == minutes {
			if arrElem-i == minutes {
				for j := arrElem - i; j < arrElem; j++ {
					sutisfaction += clients[j]
				}
			}
			sutisfaction += clients[i]
		}
	}
	return
}
