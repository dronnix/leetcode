package main

// The task: https://leetcode.com/problems/find-subarrays-with-equal-sum/

func findSubarrays(nums []int) bool {
	m := make(map[int]bool, len(nums))
	for i := 1; i < len(nums); i++ {
		s := nums[i-1] + nums[i]
		if m[s] {
			return true
		}
		m[s] = true
	}
	return false
}

func main() {
	println(findSubarrays([]int{4, 0, 4}))
}
