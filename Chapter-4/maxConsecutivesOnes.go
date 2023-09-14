/*
*	@author Elia Renzoni
*	@date 16/08/2023
*	@brief go function ex.
**/

package main

import "fmt"

const arrElem = 12

func main() {
	var binaryArr [arrElem]int = [arrElem]int{1, 44, 6, 2, 78, 9, 12}
	nConsecutives := func() int {
		var values int
		for _, value := range binaryArr {
			if value == 1 {
				values++
			}
		}
		return values
	}()
	switch {
	case nConsecutives == 0:
		fmt.Printf("No 1 Consecutives")
	default:
		fmt.Printf("N 1 consecutives : %d", nConsecutives)
	}
}
