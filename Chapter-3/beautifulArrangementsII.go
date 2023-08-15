/*
*
*	@author Elia Renzoni
*	@date 15/08/2023
*	@brief array algo.
*
**/

package main

import "fmt"

const (
	arrayElements = 12
	diffValue     = 5 // differece of value between adj. elements
)

func main() {
	var newArray = beautifulArrangements(arrayElements, diffValue)
	visitArray(newArray)
}

/*
*	@param numb of array elements
*	@param difference value between elements
 */
func beautifulArrangements(elements, kValue int) []int {
	var (
		arrayBuilt [arrayElements]int
		firstNumb  int
	)
	for i := 0; i < elements; i++ {
		firstNumb += kValue
		arrayBuilt[i] = firstNumb
	}
	return arrayBuilt[:]
}

/*
*	array visit algorithm
 */
func visitArray(array []int) {
	for index := range array {
		fmt.Println("Elem : %d", array[index])
	}
}
