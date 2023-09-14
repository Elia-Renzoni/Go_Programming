/*
*	@author Elia Renzoni
*	@date 17/08/2023
*	@brief go function ex.
*
**/

package main

import "fmt"

const (
	arrElements        = 12
	incrementOrder     = 5
	maxStrongestValues = 3
)

func main() {
	var (
		arrValue [arrElements]int
		medianValue int   = func() int {
			return (arrElements - 1) / 2
		}()
		strongestValues [maxStrongestValues]int       = getStrongest(arrValue, medianValue)
		arrayVisit      func([]int) = visitStrongestValues
	)
	arrayVisit(strongestValues)
}

func createArr(arrValue []int) {
	value := 0
	for index := range suppArray {
		value += incrementOrder
		suppArray[index] = value
	}
	return
}

func getStrongest(values []int, median int) (array [maxStrongestValues]int) {
	counter := 0
	for index := range values {
		for j := range values {
			if (values[index] - values[median]) >= (values[j] - values[median]) {
				counter++
				if counter <= maxStrongestValues {
					array[index] = values[index]
				}
			}
		}
	}
	return
}

func visitStrongestValues(strongestValues []int) {
	for _, value := range strongestValues {
		fmt.Printf("N = %d\n", value)
	}
}
