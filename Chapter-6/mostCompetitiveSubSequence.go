/*
*	@author Elia Renzoni
*	@date 29/10/2023
*	@brief Slice Ex.
**/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const maxLenght, maxCapacity, maxRandVal int = 10, 10, 100

func main() {
	var value int
	var err = func() error {
		var control bool
		var err error
		for control {
			fmt.Printf("Inserisci la dimensione K di ogni sotto-array ")
			_, err = fmt.Scanf("%d", &value)
			if value != 0 {
				control = false
			}
		}
		return err
	}()
	if err != nil {
		os.Exit(1)
	} else {
		var nums []int = make([]int, maxLenght, maxCapacity)
		initSlice(nums)
		result := competitiveSubSequence(nums, value)
		printSlice(result)
	}
}

func initSlice(nums []int) {
	rand.Seed(time.Now().Unix())
	for index := range nums {
		nums[index] = rand.Intn(maxRandVal)
	}
}

func competitiveSubSequence(nums []int, k int) (result []int) {
	var (
		sum          int
		sum2         int
		j            int = k
		tmp          int
		precSubSlide []int
	)
	for i := 0; ((i+k)-1) < len(nums) && (j < len(nums)); i++ {
		subSlide := nums[i:j]
		if tmp <= j {
			sum = sumSubSlide(subSlide)
		} else {
			precSubSlide = subSlide
			sum2 = sumSubSlide(subSlide)
		}
		if sum > sum2 {
			result = subSlide
		} else {
			result = precSubSlide
		}
		tmp = j
		j += k
	}
	return
}

func sumSubSlide(subNums []int) (sum int) {
	for index := range subNums {
		sum += subNums[index]
	}
	return
}

func printSlice(nums []int) {
	for _, value := range nums {
		fmt.Printf("%d\n", value)
	}
}
