package main

import "fmt"

// The task: https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/

func longestSubarray(nums []int) int {
	zi1, zi2, m := 0, 0, 0 // zi1, zi2 indexes of pre-last an last zeroes
	for i, n := range nums {
		if n == 1 {
			continue
		}
		l := i - zi1 - 1 // series length = current index - pre-last index - 1(skipped zero)
		if l > m {
			m = l
		}
		zi1, zi2 = zi2, i+1
	}
	l := len(nums) - zi1 - 1
	if l > m {
		m = l
	}
	return m
}

func main() {
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
}
