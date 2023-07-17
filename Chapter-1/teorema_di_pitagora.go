package main

import "fmt"

const (
	primoNumero int = 12
	secondoNumero int = 8
	mulOperando int8 = 2
)


func main() {
	var lato1, lato2 int64
	var ipotenusa int64

	lato1 = (primoNumero * primoNumero) - (secondoNumero * secondoNumero)
	lato2 = mulOperando * primoNumero * secondoNumero
	ipotenusa = (primoNumero * primoNumero) + (secondoNumero * secondoNumero)

	fmt.Println("Lato 1 = %d", lato1)
	fmt.Println("Lato 2 = %d", lato2)
	fmt.Printf("Ipotenusa = %d", ipotenusa)

}