/**
*	@author Elia Renzoni
*	@date 13/09/2023
*	@brief array ex. Go
**/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type arrWork [5]int
type arrProf [5]int
type arrDiff [5]int

func main() {
	var (
		workers    *arrWork = new(arrWork)
		difficulty *arrDiff = new(arrDiff)
		profit     *arrProf = new(arrProf)
		errW                = initWorkers(workers)
		errD                = initDifficulty(difficulty)
		errP                = initProfit(profit)
	)
	switch {
	case errW != nil, errD != nil, errP != nil:
		fmt.Printf("Errore nella generazione dei valori !")
		os.Exit(1)
	}
	fmt.Printf("Max Profit: %d", maxProfit(workers, difficulty, profit))
}

func initWorkers(workers *arrWork) error {
	var (
		max int
		err error
	)

	if max, err = setMax(); err != nil {
		return err
	}

	defer func() {
		for _, value := range workers {
			fmt.Println("%d", value)
		}
	}()
	rand.Seed(time.Now().Unix())
	for index := range workers {
		workers[index] = rand.Intn(max)
	}
	return nil
}

func initDifficulty(difficulty *arrDiff) error {
	var (
		max int
		err error
	)

	if max, err = setMax(); err != nil {
		return err
	}

	defer func() {
		for _, value := range difficulty {
			fmt.Println("%d", value)
		}
	}()
	rand.Seed(time.Now().Unix())
	for index := range difficulty {
		difficulty[index] = rand.Intn(max)
	}
	return nil
}

func initProfit(profit *arrProf) error {
	var (
		max int
		err error
	)
	if max, err = setMax(); err != nil {
		return err
	}
	defer func() {
		for _, value := range profit {
			fmt.Println("%d", value)
		}
	}()
	rand.Seed(time.Now().Unix())
	for index := range profit {
		profit[index] = rand.Intn(max)
	}
	return nil
}

func setMax() (int, error) {
	var (
		maxValue int
		err      error
		control  bool = true
	)
	for control {
		fmt.Printf("Inserisci il limite Max\n")
		_, err = fmt.Scanf("%d", &maxValue)
		if maxValue > 0 {
			control = false
		} else {
			fmt.Printf("Errore, ripeti l'azione !\n")
		}
	}
	return maxValue, err
}

func maxProfit(workers *arrWork, difficulty *arrDiff, profit *arrProf) (maxProf int) {
	// TODO
}
