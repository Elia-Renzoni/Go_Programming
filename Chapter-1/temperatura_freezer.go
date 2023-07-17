package main

import "fmt"

const (
	_primoValore = 2.0 * iota
	secondoValore
	terzoValore
)

const differenza int = 20
const tempoTrascorso int = 14

func main() {
	var temperatura int

	temperatura = (terzoValore * (tempoTrascorso * tempoTrascorso)) / (tempoTrascorso + secondoValore) - differenza
	fmt.Printf("Temperatura = %d", temperatura)

}