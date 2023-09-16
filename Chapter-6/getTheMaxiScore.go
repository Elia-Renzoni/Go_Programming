/**
 * 
 * 	@author Elia Renzoni
 * 	@date 16/09/2023
 * 	@brief array ex. Go
 * */

package main 

import (
 	"fmt"
 	"math/rand"
 	"time"
)

type firstArr [10]int
type secondArr [10]int

const (
	max = 50
	msg = "Error, func initNums1"
)

func main() {
 	var (
 		nums1 *firstArr = new(firstArr)
 		nums2 *secondArr = new(secondArr)
 	)
 	initNums1(nums1)
 	initNums2(nums2)
 	fmt.Printf("Max Score : %d", maxSum(nums1, nums2))
 }

func initNums1(nums1 *firstArr) {
	control := 0 
	rand.Seed(time.Now().Unix())
	for index := range nums1 {
		nums1[index] = rand.Intn(max)
		if nums1[index] == 0 {
			control++
		}
	}
	if control == len(nums1) {
		panic(msg)
	}
}

func initNums2(nums2 *secondArr) {
	control := 0 
	rand.Seed(time.Now().Unix())
	for index := range nums2 {
		nums2[index] = rand.Intn(max)
		if nums2[index] == 0 {
			control++
		}
	}
	if control == len(nums2) {
		panic(msg)
	}
}

func maxSum(nums1 *firstArr, nums2 *secondArr) (max int) {
	for index := range nums1 {
		for _, value := range nums2 {
			if nums1[index] == value {
				max += value
				i := index
				for i < len(nums1) {
					if value == nums1[i] {
						for j := i; j < len(nums1); j++ {
							max += nums1[j]
						}
						return
					}
					i++
				}
			} else {
				max += nums1[index]
			}
		}
		return
	}
	return 
}