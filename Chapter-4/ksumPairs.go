/**
*	@author Elia Renzoni
*	@date 23/08/2023
*	@brief go functions problems
*
**/

package main

import (
	"fmt"
	"math/rand"
)

const (
	maxSize = 12
	minOp   = 0
)

func main() {
	var (
		kValue = func() int {
			return rand.Intn(20)
		}()
		numsArr [maxSize]int = [maxSize]int{12, 3, 4, 5, 89, 1}
	)
	if maxOp := maxOperation(numsArr, kValue); maxOp != minOp {
		fmt.Printf("Operations : %d", maxOp)
	} else {
		fmt.Printf("No operations")
	}
}

func maxOperation(nums [maxSize]int, k int) (operations int) {
	for _, value := range nums {
		for _, sValue := range nums {
			if value+sValue == k {
				operations++
			}
		}
	}
	return
}
