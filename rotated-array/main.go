package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 3}
	fmt.Println(search(a, 2))
}

func search(nums []int, target int) int {
	maxIdx := sort.Search(len(nums), func(i int) bool { return nums[i] < nums[0] })

	sub, offset := nums[maxIdx:], maxIdx
	if target >= nums[0] {
		sub, offset = nums[0:maxIdx], 0
	}

	pos := sort.Search(len(sub), func(i int) bool { return target <= sub[i] })
	if pos == len(sub) || sub[pos] != target {
		return -1
	}

	return offset + pos
}
