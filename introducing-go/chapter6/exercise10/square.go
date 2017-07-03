package main

import "fmt"

func main() {
	x := 1.5
	square(&x)
}

func square(xptr *float64) {
	fmt.Println(*xptr * *xptr)
}
