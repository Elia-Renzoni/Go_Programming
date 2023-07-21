package main

import "fmt"

const notFound = -1

func main() {
	var arrElements [6]int16 = [6]int16{12, 44, 3, 6, 8, 7}
	var target int16 = 34
	var firstIndex, secondIndex int = controlSum(arrElements, target)
	var result bool = setOutput(firstIndex, secondIndex) 
	if result {
		fmt.Printf("Trovato !")
	} else {
		fmt.Printf("Non trovato !")
	}
}

func controlSum(array [6]int16, target int16) int {
	var firstIndex, secondIndex int
	for index := range array {
		if sum := array[index] + array[index + 1]; sum == target {
			firstIndex, secondIndex = index, index + 1
		} else {
			firstIndex, secondIndex = notFound, notFound
		}
	}
	return firstIndex, secondIndex
}

func setOutput(index1 int, index2 int) bool {
	if index1 != notFound && index2 != notFound {
		return false
	} else {
		return true
	}
}