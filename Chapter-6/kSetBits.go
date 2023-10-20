/*
*	@author Elia Renzoni
*	@date 20/10/2023
*	@brief Slice
**/

package main 

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

const sigBit byte = 1

func main() {
	var (
		err error
		kValue int
		sliceLenght int
		sum int
	)

	if err, kValue, sliceLenght = setValues(); err != nil {
		fmt.Printf("Errore nell'acquisizione dei valori !")
		os.Exit(1)
	} else {
		if nums := make([]int, sliceLenght); nums == nil {
			os.Exit(1);
		} else {
			initNums(nums)
			nBits(nums, &sum, kValue)
			fmt.Printf("Somma : %d" sum)
		}
	}
}

func setValues() (err error, k int, lenght int) {
	control := true
	for control {
		fmt.Printf("Inserisci il numero K : \n")
		_, err = fmt.Scanf("%d", &k)
		fmt.Printf("Inserici la lunghezza massima dello slice : \n")
		_, err = fmt.Scanf("%d", &lenght)
		if err != nil {
			return errors.New("Errore nell'acquisizione !"), nil, nil
		} 
		if k < 0 || lenght < 0{
			fmt.Printf("Errore, valodi non validi ! \n")
		} else {
			control = false;
		} 
	}
	return 
}

func initNums(nums ...int) {
	rand.Seed(time.Now().Unix())
	for index := range nums {
		nums[index] = rand.Intn(100)
	}
}

func nBits(nums ...int, sum* int, k int) {
	for _, value := range nums {
		if result := encoding(value, k); result {
			*sum += value
		}
	} 
}

func encoding(value, k int) (result bool) {
	// TODO
}







