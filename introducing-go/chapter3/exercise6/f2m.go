package main

import "fmt"

func main() {
	var feet float64
	fmt.Scanf("%f", &feet)
	meters := feet * 0.3048
	fmt.Println("meters =", meters)
}
