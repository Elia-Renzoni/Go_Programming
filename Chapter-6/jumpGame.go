/**
*	@Copyright
*	@author Elia Renzoni
*	@date 05/11/2023
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const panicMessage string = "Errore nell'acquisizione"

func main() {
	dim := arrayLenght()
	nums := make([]int, dim)
	initSlice(nums)
	if isJumpable := jumpGame(nums); isJumpable {
		fmt.Printf("Possibile Saltare all'ultimo elemento", isJumpable)
	} else {
		fmt.Printf("Non é possibile saltare all'ultimo elemento", isJumpable)
	}
}

func arrayLenght() (lenght int) {
	var control bool = true
	for control {
		fmt.Printf("Inserisci il numero degli elementi dell'array sottostante \n")
		_, err := fmt.Scanf("%d", &lenght)
		if err != nil {
			panic(panicMessage)
		} else if lenght <= 0 {
			control = true
		} else {
			control = false
		}
	}
	return
}

func initSlice(nums []int) {
	rand.Seed(time.Now().Unix())
	for i := range nums {
		nums[i] = rand.Intn(300)
	}
}

func jumpGame(nums []int) (canJumpToEnd bool) {
	var (
		index int
	)
	for index = 0; index < len(nums); {
		if index < len(nums) {
			index += nums[index]
		}
	}

	if index == len(nums)-1 {
		canJumpToEnd = true
	} else {
		canJumpToEnd = false
	}
	return
}
