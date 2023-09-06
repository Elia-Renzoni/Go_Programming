/*
*	@author Elia Renzoni
*	@date 06/09/2023
*	@brief concat of two arrays. Go function ex
*
**/

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	firstArrayNumsElement, secondArrayNumsElement int = 12, 12 * 2
)

func main() {
	var (
		firstArray, secondArray []int
		err, error1             error
	)
	if firstArray, err = createFirstArray(); err != nil {
		fmt.Printf("Errore durante l'esecuzione ! :", err)
		os.Exit(1)
	}
	if secondArray, error1 = createSecondArray(firstArray); error1 != nil {
		fmt.Printf("Errore durante l'esecuzione ! : ", error1)
		os.Exit(1)
	}
	visitConcatArray := func(concat []int) {
		for _, value := range concat {
			fmt.Printf("Value : %d \n", value)
		}
	}
	visitConcatArray(secondArray)
}

func createFirstArray() ([]int, error) {
	var supportArray [firstArrayNumsElement]int
	rand.Seed(time.Now())
	for index := range supportArray {
		supportArray[index] = rand.Intn()
	}
	if check := isEmpty(supportArray); check == true {
		return nil, errors.New("Impossibile creare il primo array !")
	}
	return supportArray[:], nil
}

func createSecondArray(firstArray []int) ([]int, error) {
	var (
		supportArray [secondArrayNumsElement]int
		j            int
	)

	for index := range supportArray {
		if index <= firstArrayNumsElement {
			supportArray[i] = firstArray[i]
		} else {
			supportArray[i] = firstArray[j]
			j++
		}
	}
	if check := isEmpty(supportArray); check == true {
		return nil, errors.New("Impossibile creare il secondo array !")
	}
	return supportArray[:], nil
}

func isEmpty(arrayToCheck []int) bool {
	var (
		control int
		result  bool
	)
	for value := range arrayToCheck {
		if value == 0 {
			control++
		}
	}
	if control == len(arrayToCheck) {
		result = true
	} else {
		result = false
	}
	return result
}
