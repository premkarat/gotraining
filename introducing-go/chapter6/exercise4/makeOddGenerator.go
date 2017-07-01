package main

import "fmt"

func main() {
	nextOdd := makeOddGenerator()
	for i := 0; i < 5; i++ {
		fmt.Println(nextOdd())
	}
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() uint {
		ret := i
		i += 2
		return ret
	}
}
