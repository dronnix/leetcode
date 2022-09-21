package main

import "fmt"

// https://leetcode.com/problems/permutations-ii/
func permuteUnique(nums []int) [][]int {
	cur := make([]int, len(nums))
	sz := 1
	for i := 2; i < len(nums); i++ {
		sz *= i
	}
	res := make([][]int, 0, sz)
	fillNum(&res, nums, cur, 0)
	return res
}

func fillNum(res *[][]int, nums, cur []int, pos int) {
	if pos == len(nums) {
		r := make([]int, len(cur))
		copy(r, cur)
		*res = append(*res, r)
		return
	}
	const used = 42
	uniqNums := make(map[int]int, len(nums))
	for p, n := range nums {
		if n == used {
			continue
		}
		uniqNums[n] = p
	}

	for n, p := range uniqNums {
		cur[pos] = n
		nums[p] = used
		fillNum(res, nums, cur, pos+1)
		nums[p] = n
	}
}

func main() {
	for _, r := range permuteUnique([]int{1, 2, 3}) {
		fmt.Println(r)
	}
}
