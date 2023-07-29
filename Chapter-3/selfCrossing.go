/**
 * 
 * 	@author Elia Renzoni
 * 	@date 29/07/2023
 * 	@brief self crossing distance
 * */

 package main

 import "fmt"

 const correctLenght = 5

 func main() {
 	var arrPoint []int16 = []int16 {1, 2, 3, 4, 1}
 	var isCrossing bool
 	checkCross(arrPoint, &isCrossing)
 	if isCrossing {
 		fmt.Printf("Crossing")
 	} else {
 		fmt.Printf("Not Crossing")
 	}
 }

 func checkCross(points []int16, isCross* bool) {
 	*isCross := false
 	for index := range points {
 		if len(points) >= correctLenght {
 			if index == 0 {
 				if points[index] >= points[index + 2] {
 					*isCross = true
 				}
 			} else if index == correctLenght - 1  {
 				if points[index] >= points[index - (correctLenght - 1)] {
 					*isCross = true
 				}
 			}
 		} else if len(points) < correctLenght {
 			if index == 0 {
 				if points[index] >= points[index + 2] {
 					*isCross = true
 				}
 			} else if index == 1  {
 				if points[index] <= points[index + 1] {
 					*isCross = true
 				}
 			}
 		}
 	}
}