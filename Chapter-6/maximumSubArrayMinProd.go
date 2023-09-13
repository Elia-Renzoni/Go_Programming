/**
*	@author Elia Renzoni
*	@date 13/09/2023
*	@brief array ex. Go
**/

package main

import (
	"fmt"
	_"math/rand"
	_"time"
	"os"
)

type numbers [10]int

const movePattern int = 3

func main() {
	var (
		nums *numbers = new(numbers)
		err := initArray(nums)
	)
	if err != nil {
		fmt.Printf("Errore!")
		os.Exit(1)
	}
	selectSort(nums)
	fmt.Printf("Min Prod : %d\n", findMinProd(nums))
}

func initArray(nums *numbers) error {
	var control bool = true
	for index := range nums {
		for control {
			fmt.Printf("Inserisci un numero \n")
			numbers, err := fmt.Scanf("%d", nums[index])
			if number != 1 || nums[index] <= 0 {
				control = true
			} else {
				control = false
			}
		}
	}
	if err != nil {
		return err
	} 
	return nil
}

func selectSort(nums *numbers) {
	var (
		minValue, indexMinValue int 
	)
	for i := 0; i < len(nums) - 1; i++ {
		minValue = nums[i]
		indexMinValue = i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < minValue {
				minValue = nums[j]
				indexMinValue = j
			}
		}
		if indexMinValue != i {
			nums[indexMinValue] = nums[i]
			nums[i] = minValue
		}
	}
}

func findMinProd(nums *numbers) int {
	var (
		maxMinP int
		sum int
	)
	for i := len(nums); i >= len(nums) - movePattern; i-- {
		sum += nums[i]
	}
	return maxMinP * sum
}