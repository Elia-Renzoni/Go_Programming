// @author Elia Renzoni
// @date 24/07/2023

package main

import "fmt"

const (
	multiply1, multiply2, multiplt3 int = 7, 11, 13
)

func main() {
	var numberList []int = []int {104, 373, 13, 121, 77, 30751}
	var isMult, isPair, isPrime bool 
	checkMult(&isMult, numberList)
	checkIfPair(&isPair, numberList)
	checkIfPrime(&isPrime, numberList)
	checkResults(isMult, isPair, isPrime)
}

func checkMult(isMult* bool, numberList []int) {
	for index := range numberList {
		switch {
		case numberList[index] % multiply1 == 0, 
		     numberList[index] % multiply2 == 0,
		     numberList[index] % multiplt3 == 0:
			 *isMult = true
		default:
			 *isMult = false
		}
	}
}

func checkIfPair(isPair* bool, numberList []int) {
	for _, val := range numberList {
		if val % 2 == 0 {
			*isPair = true 
		} else {
			*isPair = false
		}
	}
}

func checkIfPrime(isPrime* bool, numberList []int) {
	*isPrime = true
	for _, value := range numberList {
		for index := 2; index < value &&  *isPrime != false; index++ {
			if value % index == 0 {
				*isPrime = true
			}
		}
	}
}

func checkResults(isMult bool, isPair bool, isPrime bool) {
	if isMult {
		fmt.Println("Ha multipli !")
	} else {
		fmt.Println("Non ha Multipli !")
	}
	if isPair {
		fmt.Println("pari !")
	} else {
		fmt.Println("dispari !")
	}
	if isPrime {
		fmt.Println("Primi !")
	} else {
		fmt.Println("Non primi !")
	}
}
