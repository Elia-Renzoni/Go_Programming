/**
*	@author Elia Renzoni
*	@date 19/11/2023
*	@brief HashTable go coding problems
 */

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var (
		words1          = make([]string, 10)
		word2           = make([]string, 10)
		parserAdtWords1 = make(map[string]int)
		parserAdtWords2 = make(map[string]int)
		firstString     string
		secondoString   string
	)
	if err := initWords(words1, word2); err != nil {
		os.Exit(1)
	} else if err := getStrings(&firstString, &secondoString); err != nil {
		os.Exit(1)
	}
	parseWords(parserAdtWords1, words1)
	parseWords(parserAdtWords2, word2)
	switch {
	case isSubset(parserAdtWords1, firstString):
		fmt.Printf("La Stringa  %s e' un sottoinsieme !", firstString)
		fallthrough
	case isSubset(parserAdtWords2, secondoString):
		fmt.Printf("La Stringa %s e' anche universale !", firstString)
	default:
		fmt.Printf("La String %s non e' un sottoinsieme !", firstString)
	}
}

func initWords(words1, words2 []string) (err error) {
	var (
		control     bool = true
		index       int
		values      int
		secondValue int
	)
	for control || index < 10 {
		fmt.Printf("Inserisci delle stringe\n")
		values, err = fmt.Scanf("%s", words1[index])
		secondValue, err = fmt.Scanf("%s", words2[index])
		switch {
		case err != nil:
			return err
		case values == 0, secondValue == 0:
			control = true
		default:
			control = false
			index++
		}
	}
	return
}

func getStrings(firstString, secondString *string) (err error) {
	control := true
	for control {
		fmt.Printf("Inserisci la prima stringa\n")
		_, err = fmt.Scanf("%s", firstString)
		fmt.Printf("Inserisci la seconda stringa\n")
		_, err = fmt.Scanf("%s", secondString)
		if err != nil {
			return errors.New("Errore nella funzione getStrings")
		} else {
			control = false
		}
	}
	return
}

func parseWords(parser map[string]int, words []string) {
	var firstSliceOccu int
	for index := range words {
		var tmp string = words[index]
		for _, values := range words[index] {
			for _, value := range tmp {
				if !(values != value) {
					firstSliceOccu += 1
					parser[string(value)] = firstSliceOccu
				}
			}
		}
	}
}

func isSubset(parser map[string]int, values string) bool {
	var tmp map[string]int
	var occurrencies int

	// TODO
	defer func() {
		for key := range parser {
			// main logic:
			// if the key exist
			// if the key(parser) exist in tmp and occurencies are the same
			// then result = true
			// else
			// thena result = false
			//
		}
	}()

	for _, char := range values {
		for index := range values {
			if char == rune(values[index]) {
				occurrencies++
				tmp[string(char)] = occurrencies
			}
		}
	}

	return true
}
