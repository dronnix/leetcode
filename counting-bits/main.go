package main

import "fmt"

func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	r := make([]int, n+1)
	r[1] = 1
	d := 1
	for i := 1; i <= n; i++ {
		if i>>1 == d {
			d *= 2
		}
		r[i] = r[i-d] + 1
	}
	return r
}

func main() {
	fmt.Println(countBits(7))
}
