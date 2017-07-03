package main

import "fmt"

func main() {
	var n uint
	fmt.Scanf("%d", &n)
	fib(n)

	fgenerator := fibGenerator(n)
	for i := uint(2); i <= n; i++ {
		fmt.Printf("%d ", fgenerator())
	}
	fmt.Println()
}

func fib(n uint) {
	pre := 1
	cur := 1
	fmt.Printf("%d %d ", pre, cur)
	for i := uint(3); i <= n; i++ {
		seq := pre + cur
		fmt.Printf("%d ", seq)
		pre, cur = cur, seq
	}
	fmt.Println()
}

// implementation using closure
func fibGenerator(n uint) func() uint {
	pre := uint(0)
	cur := uint(1)
	fmt.Printf("%d ", cur)
	return func() uint {
		seq := uint(pre + cur)
		pre, cur = cur, seq
		return seq
	}
}
