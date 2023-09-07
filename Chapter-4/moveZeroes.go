/**
*
*	@author Elia Renzoni
*	@date 07/09/2023
*	@brief defer call. Go function ex.
***/

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	maxValue, zeroValue, maxRand int = 12, 0, 30
)

func main() {
	var (
		values []int
		err    error
	)
	if values, err = createArray(); err != nil {
		fmt.Printf("Errore ! Impossibile creare l'array ")
		os.Exit(1)
	} else {
		moveZeroes(values)
	}
}

func createArray() ([]int, error) {
	var (
		suppArray      [maxValue]int
		controlCounter int
	)
	rand.Seed(time.Now().Unix())
	for index := range suppArray {
		suppArray[index] = rand.Intn(maxRand)
	}
	for _, value := range suppArray {
		if value == zeroValue {
			controlCounter++
		}
	}
	if controlCounter != len(suppArray) {
		return nil, errors.New("Errore nella generazione dell'Array !")
	}
	return suppArray[:], nil
}

func moveZeroes(values []int) {
	var (
		repeat  int = 0
		support int
	)
	defer visitArray(values)

	for repeat != len(values) {
		for index := range values {
			if values[index] == zeroValue {
				if index < len(values) {
					support = values[index]
					values[index] = values[index+1]
					values[index+1] = support
				}
			}
		}
		repeat++
	}
}

func visitArray(values []int) {
	for index := range values {
		fmt.Printf("Value : %d\n", values[index])
	}
}
