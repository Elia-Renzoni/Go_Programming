/**
 * 
 * 	@author Elia Renzoni
 * 	@date 28/07/2023
 * 	@brief Levy, Marshall Hall and Mersenne Conjecture
 * */

 package main

 import "fmt"
 import "math"

 const (
 	levyConj = iota
 	marshallHallConj
 	mersenneConj
 )
 const maxPrimeNumb = 1000

 func main() {
 	var (
 		levyInput int = 13
 		marshallHallInput1, marshallHallInput2 int = -6, 12
 		mersenneInput int = 19
 		conjSelected byte = marshallHallConj 
 	)
 	switch conjSelected {
 	case levyConj:
 		if resultLevyConj := verLevyConjcture(levyInput); resultLevyConj {
 			fmt.Printf("Verified")
 		} else {
 			fmt.Printf("Not verified")
 		}
 	case marshallHallConj:
 		if resultMarshhallConj := verMarshallHallConjecture(marshallHallInput1, marshallHallInput2); resultMarshhallConj {
 			fmt.Printf("Verfied")
 		} else {
 			fmt.Printf("Not Verified")
 		}
 	case mersenneConj:
 		if resultMersenneConj := verMersenneConjecture(mersenneInput); resultMersenneConj {
 			fmt.Printf("Verified")
 		} else {
 			fmt.Printf("Not Verified")
 		}
 	}
 }

func checkPrimeNumbers(number int) bool {
	isPrime := true 
	for index := 2; index < number && isPrime != false; index++ {
		if index % number == 0 {
			isPrime = false
		}
	}
	return isPrime
}

 func verLevyConjcture(levyNumbInput int) bool {
    isVer := false
 	for firstNumber := 0; firstNumber < maxPrimeNumb; firstNumber++ {
 		if firstNumber % 2 != 0 && checkPrimeNumbers(firstNumber) {
 			for secondNumber := 0; secondNumber < maxPrimeNumb; secondNumber++ {
 				if checkPrimeNumbers(secondNumber) {
 					if levyNumbInput == (firstNumber + secondNumber * secondNumber) {
                        isVer = true
                    }
 				}
 			}
 		}
 	}
    return isVer
 }

 func verMarshallHallConjecture(firstNumb, secondNumb int) bool {
    cube := Pow(float64(firstNumb), 3)
    quad := Pow(float64(secondNumb), 2)
    isVer := false
    if cube != quad {
        isVer = true
    }
    return isVer
 }

 func verMersenneConjecture(mersenneIn int) bool {
    isVer := false
    if checkPrimeNumbers(mersenneIn) {
        isVer = true
    }
    return isVer
 }
