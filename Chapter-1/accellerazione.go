package main

import "fmt"

const tempo = 60
const spazio = 30
const accelerazione = 10
const divisore = 0.5

func main() {
	var (
		velocitaMissile int 
		distanzaMissile int
	)

	velocitaMissile = accelerazione * tempo
	distanzaMissile = divisore * accelerazione * (tempo * tempo)

	fmt.Println("Velocita' = ", velocitaMissile)
	fmt.Printf("Distanza = %d", distanzaMissile)

}