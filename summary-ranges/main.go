package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) (res []string) {
	start := 0
	for i := range nums {
		if (i < len(nums)-1) && (nums[i] == nums[i+1]-1) {
			continue
		}
		s := strconv.Itoa(nums[start])
		if i > start {
			s += "->" + strconv.Itoa(nums[i])
		}
		res = append(res, s)
		start = i + 1
	}
	return res
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}
