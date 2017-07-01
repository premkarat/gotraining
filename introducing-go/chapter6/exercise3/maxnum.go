package main

func max(args ...int) int {
	max := args[0]
	for _, val := range args {
		if val > max {
			max = val
		}
	}
	return max
}
