package main

import "fmt"

func pivotIndex(nums []int) int {
	rs := 0
	for i := 1; i < len(nums); i++ {
		rs += nums[i]
	}
	if rs == 0 {
		return 0
	}

	ls := 0
	for i := 1; i < len(nums); i++ {
		ls += nums[i-1]
		rs -= nums[i]
		if ls == rs {
			return i
		}
	}
	if ls == 0 {
		return len(nums) - 1
	}
	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 2, 3, 2, 1}))
}
