/**
*	@author Elia Renzoni
*	@date 15/09/2023
*	@brief array ex. Go
*
**/
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const minLimit int = 0
const maxElement = 20

type arr [maxElement]int

func main() {
	var (
		nums  *arr = new(arr)
		value      = 0
		err   error
	)

	if value, err = setMaxValueToRand(); err != nil {
		os.Exit(1)
	}
	initArray(nums, value)
	if minOperation := minIntegerForUnique(nums); minOperation != minLimit {
		fmt.Printf("Unique Array, %d operation needed\n", minLimit)
	} else {
		fmt.Printf("Unique Array, 0 operation needed\n")
	}
}

func setMaxValueToRand() (value int, err error) {
	control := true
	for control {
		fmt.Printf("Inserisci il numero massimo da generare !\n")
		_, err = fmt.Scanf("%d", &value)
		if value > minLimit {
			control = false
		} else {
			fmt.Printf("Errore! Ripeti\n")
		}
	}
	if err != nil {
		return
	}
	return
}

func initArray(nums *arr, maxRand int) {
	defer visitArray(nums)
	rand.Seed(time.Now().Unix())
	for index := range nums {
		nums[index] = rand.Intn(maxRand)
	}
}

func visitArray(nums *arr) {
	for _, value := range nums {
		fmt.Printf("%d\n", value)
	}
}

func minIntegerForUnique(nums *arr) (minInt int) {
	defer visitArray(nums)
	defer fmt.Printf("*********** New Values *************\n")
	var maxValue int
	for k := range nums {
		if k < len(nums)-1 {
			if nums[k] > nums[k+1] {
				maxValue = nums[k]
			} else {
				maxValue = nums[k+1]
			}
		}

	}
	for index := 0; index < len(nums); index++ {
		for _, value := range nums {
			if nums[index] == value {
				minInt += 1
				maxValue += 1
				nums[index] = maxValue
			}
		}
	}
	return
}
