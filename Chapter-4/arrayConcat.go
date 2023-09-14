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
	arrElem int = 12
	maxValue                                      int = 50
)

type arr [arrElem]int

func main() {
	var (
		firstArray  *arr = new(arr) 
		secondArray *arr = new(arr)
		err, error1 error
	)
	if err = createFirstArray(firstArray); err != nil {
		fmt.Printf("Errore durante l'esecuzione ! :", err)
		os.Exit(1)
	}
	if error1 = createSecondArray(secondArray, firstArray); error1 != nil {
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

func createFirstArray(array *arr) error {
	rand.Seed(time.Now().Local().Unix())
	for index := range array {
		array[index] = rand.Intn(maxValue)
	}
	if check := isEmpty(array); check == true {
		return errors.New("Impossibile creare il primo array !")
	}
	return nil
}

func createSecondArray(secondArray *arr, firstArray *arr) error {
	var j int
	for index := range secondArray {
		if index <= firstArrayNumsElement {
			secondArray[index] = firstArray[index]
		} else {
			secondArray[index] = firstArray[j]
			j++
		}
	}
	if check := isEmpty(secondArray); check == true {
		return errors.New("Impossibile creare il secondo array !")
	}
	return nil
}

func isEmpty(arrayToCheck *arr) bool {
	var (
		control int
		result  bool
	)
	for _, value := range arrayToCheck {
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
