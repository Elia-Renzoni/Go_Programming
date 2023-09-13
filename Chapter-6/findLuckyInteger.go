/**
*	@author Elia Renzoni
*	@date 13/05/2023
*	@brief array ex. Go
**/

package main

import (
	"fmt"
)

type arrNumbers [10]int

const negResult int = -1

func main() {
	var (
		nums   = &arrNumbers{1, 2, 2, 3, 3, 3}
		result = findLucky(nums)
	)
	if result != negResult {
		fmt.Printf("N = %d", result)
	} else {
		fmt.Printf("No Lucky Integer")
	}
}

func findLucky(nums *arrNumbers) (nLuckyInt int) {
	var lucky int
	for _, value := range nums {
		for j := range nums {
			if value == nums[j] {
				lucky++
			}
		}
		if value != lucky {
			lucky = -1
		}
	}
	return
}
