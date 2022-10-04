package main

import (
	"fmt"
)

// The task: https://leetcode.com/problems/fair-distribution-of-cookies/
func distributeCookies(cookies []int, k int) int {
	children := make([]int, k)
	return distribute(cookies, children, 0)
}

func distribute(cookies, children []int, cookiesLeft int) int {
	if len(cookies) == cookiesLeft {
		max := 0
		for _, c := range children {
			if max < c {
				max = c
			}
		}
		return max
	}

	min := 1000000
	for i := range children {
		children[i] += cookies[cookiesLeft]
		m := distribute(cookies, children, cookiesLeft+1)
		if min > m {
			min = m
		}
		children[i] -= cookies[cookiesLeft]
	}

	return min
}

func main() {
	fmt.Println(distributeCookies([]int{8, 15, 10, 20, 8}, 2))
}
