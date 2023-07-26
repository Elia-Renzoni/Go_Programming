/**
 * 	@author Elia Renzoni
 * 	@date 26/07/2023
 * 	@brief
 * */

 package main

 import "fmt"

 func main() {
 	var (
 		array []int = []int {1, 2, 3, 6, 2, 3, 4, 7, 8}
 		groupSize int = 3
 		isDiv bool
 	)
 	if checkPokerHands(array, groupSize, &isDiv); isDiv {
 		fmt.Printf("Div - Ord");
 	} else {
 		fmt.Printf("Non Div - Not Ord")
 	}
 }

 func checkPokerHands(array []int, groupSize int, isDiv* bool) {
 	if (len(array) % groupSize) == 0 {
        if div := len(array) / groupSize; (div * groupSize) == len(array) {
            if isOrd := checkOrd(array); isOrd {
                *isDiv = true
            }
        }
    } else {
        *isDiv = false
    }
 }

// bubblesort and check
func checkOrd(array []int) bool {
    var ( 
        tmp, i, j int
        result bool 
    )
    for i = 1; i < len(array); i++ {
        for j = len(array) - 1; j >= i; j-- {
            tmp = array[j - 1]
            array[j - 1] = array[j]
            array[j] = tmp;
        }
    }
    for index := 0; index < len(array) - 1; index++ {
        if array[index] < array[index + 1] {
            result = true 
        } else {
            result = false
        }
    }
    return result 
}
