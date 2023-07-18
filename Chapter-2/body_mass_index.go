package main

import "fmt"

const (
	sottopeso float32 = 18.5
	normopeso float32 = 24.9
	sovrappesoPrimoRange, sovrappresoSecondoRange float32 = 25.0, 29.9
	_obeso float32 = 30.0
)

func main() {
	var peso, altezza float32 = 67.9, 1.80
	var bmiUtente float32 = impostaBmi(peso, altezza)
	controllaBmi(bmiUtente)
}

func impostaBmi(peso float32, altezza float32) float32 {
	var bmi float32 = peso / (altezza * altezza)
	return bmi
}

func controllaBmi(bmiUtente float32) {
	if bmiUtente >= sottopeso && bmiUtente < normopeso {
		fmt.Println("Sei Sottopeso")
	} else if bmiUtente >= normopeso && bmiUtente < sovrappesoPrimoRange {
		fmt.Println("Sei Normopeso")
	} else if bmiUtente >= sovrappesoPrimoRange && bmiUtente <= sovrappresoSecondoRange {
		fmt.Println("Sei Sovrappeso")
	} else {
		fmt.Println("Sei Obeso")
	}
}