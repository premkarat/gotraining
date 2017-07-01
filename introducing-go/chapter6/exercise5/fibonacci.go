package main

import "fmt"

func main() {
	var n uint
	fmt.Scanf("%d", &n)
	fib(n)
}

func fib(n uint) uint {
	if n != 0 {
		fmt.Printf(n + fib(n-1))
	} else {
		return 0
	}
	return 0
}

fmt.Printf(format, a)
