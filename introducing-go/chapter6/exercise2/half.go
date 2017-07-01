package main

import "fmt"

func main() {
	var input int
	fmt.Scanf("%d", &input)
	val, result := half(input)
	fmt.Println("val=", val, "result=", result)
}

func half(num int) (float64, bool) {
	return float64(num / 2), num%2 == 0
}
