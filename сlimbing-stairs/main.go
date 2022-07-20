package main

import "fmt"

func main() {
	fmt.Println(climbStairs(3))
}

func climbStairs(n int) int {
	w := [3]int{0, 1, 2}
	for i := 3; i <= n; i++ {
		w[0], w[1] = w[1], w[2]
		w[2] = w[1] + w[0]
	}
	if n < 3 {
		return w[n]
	}
	return w[2]
}
