package main

import "fmt"

const (
	acqua, mercurio, mercurio1, rame, argento, oro int = 100, 350, 357, 1187, 2193, 2660
)

func main() {
	var eboSostanza = 2666
	verificaSostanza(eboSostanza)
}

func verificaSostanza(eboSostanza int) {
	switch eboSostanza {
	case acqua:
		fmt.Printf("Acqua")
	case mercurio, mercurio1:		// logical or operator 
		fmt.Printf("mercurio")
	case rame:
		fmt.Printf("Rame")
	case argento:
		fmt.Println("Argento")
	case oro:
		fmt.Printf("Oro")
	default:
		fmt.Printf("Sostanza non riconosciuta !")
	}
}