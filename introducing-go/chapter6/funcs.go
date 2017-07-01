package main

import "fmt"

func main() {
	xs := []float64{12, 13, 14, 15, 16, 17}
	fmt.Println("average=", average(xs))
	fmt.Println("total=", add(13, 24, 56, 73, 55))
	xt := []int{234, 234, 4235, 232}
	fmt.Println("total slice elements=", add(xt...))
	fmt.Println("elements in array xt=", xt)
	sub := func(x, y float64) float64 {
		return x - y
	}
	fmt.Println("Substract=", sub(5, 6))
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println("x=", increment())
	fmt.Println("x=", increment())
	nextEven := makeEvenGenerator(10)
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println("Factorial=", factorial(6))
	defer second()
	//first()
	/*defer func() {
		str := recover()
		fmt.Println(str)
	}()*/
	//panic("PANIC")
	// pointer reference
	var px *int
	xvar := 10
	px = &xvar
	ptrfunc(px)
	fmt.Println("px=", *px)
	newpx := new(int)
	*newpx = 10
	fmt.Println("newpx=", *newpx)
}

func average(xs []float64) float64 {
	total := 0.0
	for _, val := range xs {
		total += val
	}
	return total / float64(len(xs))
}

func add(avgs ...int) int {
	total := 0
	for _, val := range avgs {
		total += val
	}
	return total
}

func makeEvenGenerator(x int) func() uint {
	i := uint(x)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x int) int {
	if x == 1 {
		return 1
	}
	return x * factorial(x-1)
}

func first() {
	fmt.Println("first")
	panic("just panicking")
}

func second() {
	fmt.Println("second")
}

func ptrfunc(xptr *int) {
	*xptr = 10100
}
