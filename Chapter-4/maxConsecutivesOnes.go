/*
*	@author Elia Renzoni
*	@date 16/08/2023
*	@brief go function ex.
**/

package main

import "fmt"

const arrElem = 12

func main() {
	var binaryArr []int = createArray()
	nConsecutives := func() int {
		var values int
		for _, value := range binaryArr {
			if value == 1 {
				values++
			}
		}
		return values
	}()
	switch {
	case nConsecutives == 0:
		fmt.Printf("No 1 Consecutives")
	default:
		fmt.Printf("N 1 consecutives : %d", nConsecutives)
	}
}

func createArray() (newArray []int) {
	var (
		array [arrElem]int
		value int = 0
	)
	for index := range array {
		array[index] = value
		if value == 0 {
			value = 1
		} else {
			value = 0
		}
	}
	return
}
