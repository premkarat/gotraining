package main

import "fmt"

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(xptr *int, yptr *int) {
	tmp := *xptr
	*xptr = *yptr
	*yptr = tmp
}
