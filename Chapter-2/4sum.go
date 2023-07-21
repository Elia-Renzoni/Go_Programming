package main 

import "fmt"

const maxIter = 5

func main() {
	var numbers []int = []int {12, 45, 6, 11, 8, 34, 5, 32, 7, 55}
	var target int = 73
	if quadSum := checkFourSum(numbers, target); quadSum > 0 {
		fmt.Printf("%d", quadSum)
	} else {
		fmt.Printf("0 sums")
	}
}

func checkFourSum(numbers []int, target int) int {
	var nSum, sum int
	counter := 0 
	for _, val := range numbers {
		if counter < maxIter {
			sum += val
			if counter == maxIter - 1 {
				if sum == target {
					nSum++
				}
			}
		}
		counter++
	}
	return nSum 
}