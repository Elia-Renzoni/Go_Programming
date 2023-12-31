/**
*	@author Elia Renzoni
*	@date 12/09/2023
*	@brief array ex. go
***/

package main

import (
	"fmt"
	_ "math/rand"
	_ "time"
)

const infLimit byte = 0

func main() {
	firstArr := [5]int{1, 2, 3, 2, 1}
	secondArr := [5]int{3, 2, 1, 4, 7}
	if maxSub := findMaxSubArray(firstArr, secondArr); maxSub < infLimit {
		fmt.Printf("No subarray")
	} else {
		fmt.Printf("N subarray : %d", maxSub)
	}
}

func findMaxSubArray(nums1 [5]int, nums2 [5]int) (maxSub byte) {
	for i := range nums1 {
		for j := range nums2 {
			if nums1[i] == nums2[j] {
				if i == 0 && j == 0 {
					if nums2[j+1] == nums1[i+1] {
						maxSub++
					}
				} else if i == len(nums1) && j == len(nums2) {
					if nums2[j-1] == nums1[i-1] {
						maxSub++
					}
				} else {
					if nums2[j-1] == nums1[i-1] || nums2[j+1] == nums1[i+1] {
						maxSub++
					}
				}
			}
		}
	}
	return
}
