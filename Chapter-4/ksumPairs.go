/**
*	@author Elia Renzoni
*	@date 23/08/2023
*	@brief go functions problems
*
**/

package main

import "fmt"
import "math/rand"

const (
	maxSize = 12
	minOp = 0
)

func main() {
	var (
		createArray = func() (array []int) {
			var newArray[maxSize] int
			for index := range newArray {
				newArray[index] = rand.Intn(20)
			}
			return
		}
		kValue = func() int {
			return rand.Intn(20)
		} ()
		numsArr = createArray()
	)
	if maxOp := maxOperation(numsArr, kValue); maxOp != minOp {
		fmt.Printf("Operations : %d", maxOp)
	} else {
		fmt.Printf("No operations")
	}
}

func maxOperation(nums []int, k int) (operations int) {
	var ope int
	for _, value := range nums {
		for _, sValue := range nums {
			if value + sValue == k {
				ope++
			}
		}
	}
	return 
}
