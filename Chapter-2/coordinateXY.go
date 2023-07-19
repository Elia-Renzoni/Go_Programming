package main 

import "fmt"

const (
	primoQuadrante = 1
	secondoQuadrante = 2
	terzoQuadrante = 3
	quartoQuadrante = 4
	controlloQuadrante float32 = 0.0
)

func main() {
	var xCoord, yCoord float32 = -1.0, -2.5
	if quadrante := controllaGrafico(xCoord, yCoord); quadrante == primoQuadrante {
		fmt.Println("le Coordinate di trovano nel primo quadrante")
	} else if quadrante == secondoQuadrante {
		fmt.Println("le Coordinate si trovano nel secondo quadrante")
	} else if quadrante == terzoQuadrante {
		fmt.Println("le Coordinate si trovano nel terzo quadrante")
	} else {
		fmt.Println("le Coordinate si trovano nel quarto quadrante")
	}

}

func controllaGrafico(x float32, y float32) int8 {
	var quadrante int8 
	if x < controlloQuadrante && y < controlloQuadrante {
		quadrante = terzoQuadrante
	} else if x > controlloQuadrante && y < controlloQuadrante {
		quadrante = quartoQuadrante
	} else if x > controlloQuadrante && y > controlloQuadrante {
		quadrante = primoQuadrante
	} else if x < controlloQuadrante && y > controlloQuadrante {
		quadrante = secondoQuadrante
	}
	return quadrante
}

