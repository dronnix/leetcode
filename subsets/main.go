package main

import "fmt"

// The task: https://leetcode.com/problems/subsets/

func subsets(nums []int) [][]int {
	n := 1<<len(nums) - 1
	res := make([][]int, 0, n)
	for i := 0; i <= n; i++ {
		r := make([]int, 0, len(nums))
		for j := 0; j < len(nums); j++ {
			if i&(1<<j) > 0 {
				r = append(r, nums[j])
			}
		}
		res = append(res, r)
	}
	return res
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
