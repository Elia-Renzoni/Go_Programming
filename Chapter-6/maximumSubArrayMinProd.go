/**
*
*	@author Elia Renzoni
*	@date 13/09/2023
*	@brief array ex. Go
**/

package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

type numbers [10]int

func main() {
	var (
		nums *numbers = new(numbers)
		err := initArray(nums)
	)
	if err != nil {
		fmt.Printf("Errore!")
		os.Exit(1)
	}
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

func findMinProd(nums *numbers) (maxMinProd int) {
	// TODO
}