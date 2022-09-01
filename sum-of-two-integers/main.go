package main

import "fmt"

// Task: https://leetcode.com/problems/sum-of-two-integers/

func getSum(a, b int) int {
	xor := a ^ b
	and := a & b

	for and != 0 {
		and <<= 1
		xor, and = xor^and, xor&and
	}

	return xor
}

func main() {
	fmt.Println(getSum(1, 1))
}
