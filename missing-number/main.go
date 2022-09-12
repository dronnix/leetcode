package main

import "fmt"

// The task: https://leetcode.com/problems/missing-number/
func missingNumber(nums []int) int {
	nsum, tsum := 0, len(nums)*(len(nums)+1)/2
	for _, n := range nums {
		nsum += n
	}
	return tsum - nsum
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
}
