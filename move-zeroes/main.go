package main

import "fmt"

func moveZeroes(nums []int) {
	w := 0
	for r := range nums {
		if nums[r] == 0 {
			continue
		}
		nums[w], nums[r] = nums[r], nums[w]
		w++
	}
}

func main() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	fmt.Println(arr)
}
