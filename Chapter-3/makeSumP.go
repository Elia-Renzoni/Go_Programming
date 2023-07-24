//	@author Elia Renzoni
//	@date 21/07/2023

package main 

import "fmt"

const notDiv int8 = -1

func main() {
	var arrNumbs []int = []int {12, 67, 4, 56, 21, 89}
	var targetDiv int = 5
	var nSubArray int8 = notDiv
	checkMakeSumDivbyP(arrNumbs, targetDiv, &nSubArray)
	if nSubArray != notDiv {
		fmt.Println("Divisibile !")
	} else {
		fmt.Println("Non divisibile !")
	}
}

func checkMakeSumDivbyP(array []int, div int, isDiv* int8) {
	var sum int 
	for index := range array {
		sum += array[index]
		if sum % div == 0 {
			*isDiv += 1
		} else {
			array[index] = -1
			if *isDiv > notDiv {
				*isDiv -= 1 
			}
		}
	}
}