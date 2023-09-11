/**
*	@author Elia Renzoni
*	@date 11/09/2023
*	@brief array ex.
**/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type myArr [10]int

func main() {
	var (
		nums        *myArr = new(myArr)
		init               = initArray
		valueUt     int
		control     bool = true
		searchValue      = func() int {
			for control {
				fmt.Printf("Inserisci il valore da ricercare\n")
				number, _ := fmt.Scanf("%d", &valueUt)
				if number != 1 && valueUt <= 0 {
					control = true
				} else {
					control = false
				}
			}
			return valueUt
		}()
	)
	init(nums)
	evenOdd(nums, searchValue)
}

func initArray(nums *myArr) {
	rand.Seed(time.Now().Unix())
	for index := range nums {
		nums[index] = rand.Intn(100)
	}
}

func evenOdd(nums *myArr, value int) {
	for index, value := range nums {
		if value%2 == 0 {
			fmt.Print("Pari\n")
			value *= 2
			nums[index] = value
		} else {
			fmt.Printf("Disapri\n")
			value *= 3
			nums[index] = value
		}
	}
}
