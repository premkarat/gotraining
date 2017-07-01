package main

import "fmt"

func main() {
	var farh float64
	fmt.Scanf("%f", &farh)
	cels := (farh - 32) * 5 / 9
	fmt.Println("Celsius = ", cels)
}
