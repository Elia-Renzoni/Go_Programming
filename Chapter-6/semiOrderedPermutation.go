/**
*	@aurhor Elia Renzoni
*	@date 11/09/2023
*	@brief array ex.
*
**/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type arr [10]int

func main() {
	var (
		nums      *arr = new(arr)
		initArray      = func(nums *arr) {
			rand.Seed(time.Now().Unix())
			for i := 0; i < len(nums); i++ {
				nums[i] = rand.Intn(100)
			}
		}
		nOperation int
	)
	initArray(nums)
	semiOrderedPermutation(nums, &nOperation)
	fmt.Printf("N operazioni : %d", nOperation)
}

func semiOrderedPermutation(nums *arr, operation *int) {
	if nums[0] == 1 {
		*operation = 0
	} else {
		for index := range nums {
			if nums[index] == 1 {
				for i := index - 1; i > 0; i-- {
					support := nums[i]
					nums[i] = nums[index]
					nums[index] = support
					*operation++
				}
			}
		}
	}
}
