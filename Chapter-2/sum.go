package main

import (
	"fmt"
)

const primo_numero int = 10
const secondo_numero int = 20

func main() {
	fmt.Println(somma())
}

func somma() int {
	var sum int = primo_numero + secondo_numero
	return sum
}
