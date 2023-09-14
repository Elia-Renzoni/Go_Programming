/**
 * 
 * 	@author Elia Renzoni
 * 	@date 26/07/2023
 * 	@brief 
 * */

 package main

 import "fmt"

 func main() {
 	var (
 		arrNumbers [6]int = [6]int {0, 1, 7, 4, 4, 5}
 		lower, upper int = 3, 6
 		rightCouples int 
 	)
 	if countRightCouples(arrNumbers, lower, upper, &rightCouples); rightCouples > 0 {
 		fmt.Printf("N right Couples = %d", rightCouples)
 	} else {
 		fmt.Printf("No right Couples")
 	}
 }

 func countRightCouples(arrNumbers [6]int, lower, upper int, couples* int) {
 	for i := 0; i < len(arrNumbers) - 1; i++ {
 		for j := i + 1; j < len(arrNumbers) - 1; j++ {
 			if sum := arrNumbers[i] + arrNumbers[j]; lower <= sum && sum <= upper {
 				*couples++
 			}
 		}
 	}
 }
