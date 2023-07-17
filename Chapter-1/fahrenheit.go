package main

import "fmt"

const (
	conv1 float32 = 0.55
	conv2 int = 32
)

func main() {
	var fahrenheit int = 23
	celsius := conv1 * (float32(fahrenheit) - float32(conv2)) 

	fmt.Printf("Gradi Celsius = %f", celsius)
}