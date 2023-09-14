/**
*
*	@author Elia Renzoni
*	@date 07/09/2023
*	@brief defer call. Go function ex.
***/

package main

import (
	"fmt"
)

const (
	maxValue, zeroValue, maxRand int = 12, 0, 30
)

func main() {
	var (
		values [maxValue]int = [maxValue]int{23, 2, 1, 4, 11, 89}
	)
	moveZeroes(values)
}

func moveZeroes(values [maxValue]int) {
	var (
		repeat  int = 0
		support int
	)
	defer visitArray(values)

	for repeat != len(values) {
		for index := range values {
			if values[index] == zeroValue {
				if index < len(values) {
					support = values[index]
					values[index] = values[index+1]
					values[index+1] = support
				}
			}
		}
		repeat++
	}
}

func visitArray(values [maxValue]int) {
	for index := range values {
		fmt.Printf("Value : %d\n", values[index])
	}
}
