/*
*	@author Elia Renzoni
*	@date 15/08/2023
*	@brief Array algorithm
**/

package main

import "fmt"

const (
	incrementOrd, incrementOrd2 int = 4, 2
	basicInt         = 1
	elements         = 5
)

func main() {
	var (
		arrElement       = createEelemnts()
		numberOfIntegers = makeIntegers(arrElement)
	)
	if numberOfIntegers != basicInt {
		fmt.Printf("You can Make %d Integers", numberOfIntegers)
	} else {
		fmt.Printf("You can make only one Integers")
	}
}

func createEelemnts() []int {
	var (
		supportArray [elements]int
		value        int
	)
	for i := 0; i < elements; i++ {
		value += incrementOrd
		supportArray[i] = value
	}
	return supportArray[:]
}

func makeIntegers(array []int) int {
	var (
		nIntegers, indexSx int = 1, 0
		indexDx            int
	)
	for index := range array {
		if index == 0 {
			nIntegers++
		} else if index < len(array)-1 {
			indexDx = index + 1
			for j := indexSx; j < indexDx; j++ {
				if array[j] != array[j + 1] {
					nIntegers += incrementOrd2
				} else {
					nIntegers++
				}
			}
		}
	}
	return nIntegers
}
