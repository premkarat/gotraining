package main

import "fmt"

func main() {
	var n uint
	fmt.Scanf("%d", &n)
	fib(n)
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
