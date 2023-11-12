/**
*	@author Elia Renzoni
*	@date 12/11/2023
*	@brief Map coding problems
**/

package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		charSequence string
		parser       = make(map[string]int)
	)
	if err := initSequence(&charSequence); err != nil {
		fmt.Printf("Errore ")
		os.Exit(1)
	}

	parsingOccurences(charSequence, parser)
	switch result := controlOccurences(parser, charSequence); result {
	case true:
		fmt.Printf("Numero di caratteri uguali !\n")
		func() {
			for _, values := range charSequence {
				fmt.Printf("value : %c \n", values)
			}
		}()
	case false:
		fmt.Printf("Sequenza invalida !")
	}

}

/**
*	@param pointer to the sequence
*	@return if the scanf check some errors
 */
func initSequence(sequence *string) (somethingWrong error) {
	var control bool = true
	for control {
		fmt.Printf("Inserisci una sequena di caratteri : \n")
		values, err := fmt.Scanf("%s", sequence)
		if err != nil {
			somethingWrong = err
			control = false
		} else if !(values != 1) {
			control = false
		}
	}
	return
}

/**
*	@param sequence of runes
*	@param Hash Map, <key = rune> <value = rune occurences>
**/
func parsingOccurences(sequence string, hashParser map[string]int) {
	var key int = 1
	for _, values := range sequence {
		hashParser[string(values)] = key
		for _, value := range sequence {
			if values == value {
				hashParser[string(values)] += 1
			}
		}
	}
	/*
	* optional compilation, only for debug
	 */
	func() {
		for key, value := range hashParser {
			fmt.Printf(" Key : %s - Value : %d ", key, value)
		}

		fmt.Println()
	}()
}

/**
*	@param Hash Map
*	@param sequence to control
*	@return if the sequence is good or not
**/
func controlOccurences(hashParser map[string]int, sequence string) (result bool) {
	var firstValue = hashParser[getFirstKey(hashParser)]
	for _, value := range hashParser {
		if firstValue == value {
			result = true
		} else {
			result = false
			return
		}
	}
	return
}

/**
*	@param Hash Map
*	@return first Key of the Hash Map, or Empty String (in case of some errors)
**/
func getFirstKey(hashParser map[string]int) string {
	var empty string = ""
	for key := range hashParser {
		return string(key)
	}
	return empty
}
