package main

import "fmt"

// Task: https://leetcode.com/problems/permutations/

func permute(nums []int) [][]int {
	n := 1
	for i := 1; i <= len(nums); i++ {
		n *= i
	}
	result := make([][]int, 0, n)
	swap(nums, 0, &result)
	return result
}

func swap(nums []int, pos int, result *[][]int) {
	if pos == len(nums)-1 {
		*result = append(*result, nums)
	}
	for i := pos; i < len(nums); i++ {
		ns := make([]int, len(nums))
		copy(ns, nums)
		ns[pos], ns[i] = ns[i], ns[pos]
		swap(ns, pos+1, result)
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
