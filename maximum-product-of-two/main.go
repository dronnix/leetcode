package main

import "fmt"

func main() {
	nums := []int{3, 4, 5, 2}

	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	m1, m2 := 0, 1
	if nums[0] < nums[1] {
		m1, m2 = 1, 0
	}
	l := len(nums)
	for i := 2; i < l; i++ {
		if nums[i] > nums[m1] {
			m1, m2 = i, m1
			continue
		}
		if nums[i] > nums[m2] {
			m2 = i
		}
	}
	return (nums[m1] - 1) * (nums[m2] - 1)
}
