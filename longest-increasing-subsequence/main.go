package main

import (
	"fmt"
	"sort"
)

func main() {
	seq := []int{4, 10, 4, 3, 8, 9}
	fmt.Println(lengthOfLIS(seq))
}

func lengthOfLIS(nums []int) int {
	maxs := make([]int, 1, len(nums))
	maxs[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxs[len(maxs)-1] {
			maxs = append(maxs, nums[i])
			continue
		}
		pos := sort.Search(len(maxs), func(j int) bool {
			return nums[i] <= maxs[j]
		})
		maxs[pos] = nums[i]
	}
	return len(maxs)
}
