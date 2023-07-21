package main

import "fmt"

const notFound = -1

func main() {
	var arrElements [5]int16 = [5]int16{12, 44, 3, 6, 8, 7}
	var target int16 = 34
	var firstIndex, secondIndex int16 = controlSum(arrElements, target)
	var result bool = setOutput(firstIndex, secondIndex) 

}

func controlSum(array []int16, target int16) int16 {
	var firstIndex, secondIndex int16
	for index := range array {
		if sum := array[index] + array[index + 1]; sum == target {
			firstIndex, secondIndex = index, index + 1
		} else {
			firstIndex, secondIndex = notFound, notFound
		}
	}
	return firstIndex, secondIndex
}

func setOutput(index1 int16, index2 int16) bool {
	if index1 != notFound && index2 != notFound {
		return false
	} else {
		return true
	}
}