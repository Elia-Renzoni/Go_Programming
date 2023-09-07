/**
 * 
 * 	@author Elia Renzoni
 * 	@date 07/09/2023
 * 	@brief panic, signaling and Handling error. 
 * 
 **/

package main 

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	panicMessage string = "Errore nella creazione dell'Array ! funzione createArray"
	maxElement, maxRandValue int = 12, 100
)

func main() {
	var (
		nums []int = createArray()
		maxPos, maxNeg int
	)
	selectSort(nums)
	maximumCount(&maxPos, &maxNeg, nums)
	if maxPos < maxNeg {
		fmt.Printf("Max (pos): %d\n", maxPos)
	} else if maxPos == maxNeg {
		fmt.Printf("Equal Value : %d\n", maxNeg)
	} else {
		fmt.Printf("Max (neg): %d\n", maxNeg)
	}
}

func createArray() (array []int) {
	var (
		conrtolCounter int
		suppArray [maxElement]int
	)
	rand.Seed(time.Now().Unix())
	for index := range suppArray {
		suppArray[index] = rand.Intn(maxRandValue)
	}
	for _, value := range suppArray {
		if value == 0 {
			conrtolCounter++
		}
	}
	if conrtolCounter == len(suppArray) {
		panic(panicMessage)						// panic and crash
	} 
	return 
}


func selectSort(nums []int) {
	var (
		minValue, indexMinValue, j int 
	)
	for i := 0; i < len(nums) - 1; i++ {
		minValue = nums[i]
		indexMinValue = i
		for j = i + 1; j < len(nums); j++ {
			if nums[i] < minValue {
				minValue = nums[i]
				indexMinValue = j
			}
		}
		if indexMinValue != i {
			nums[indexMinValue] = nums[i]
			nums[i] = minValue
		}
	}
}

func maximumCount(pos *int, neg *int, nums []int) {
	for index := range nums {
		if math.Signbit(float64(nums[index])) {
			*neg++
		} else {
			*pos++
		}
	}
}