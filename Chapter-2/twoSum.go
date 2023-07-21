package main 

import "fmt"

const notFound = -1

func main() {
	var numbers []int = []int {12, 45, 67, 12, 5, 78, 45, 6, 7}
	var target int = 34
	var index1, index2 int = controlSum(numbers, target)
	if result := resultSum(index1, index2); result {
		fmt.Println("trovato ! %d %d", index1, index2)
	} else {
		fmt.Println("Non trovato !")
	}
}

func controlSum(numbers []int, target int) (int, int) {
	var firstIndex, secondIndex int
	for i := 0; i < len(numbers) - 1; i++ {
		if sum := numbers[i] + numbers[i + 1]; sum == target {
			firstIndex, secondIndex = i, i + 1
		} else {
			firstIndex, secondIndex = notFound, notFound
		}

	}
	return firstIndex, secondIndex
}

func resultSum(index1 int, index2 int) bool {
	if index1 != notFound && index2 != notFound {
		return true
	} else {
		return false
	}
}
