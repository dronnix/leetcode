package main

import "fmt"

func findLength(nums1, nums2 []int) (max int) {
	for i := range nums1 {
		for j := range nums2 {
			m := countMatch(nums1[i:], nums2[j:])
			if m > max {
				max = m
			}
		}
	}
	return max
}

func countMatch(n1, n2 []int) (l int) {
	for i := 0; i < len(n1) && i < len(n2); i++ {
		if n1[i] == n2[i] {
			l++
			continue
		}
		break
	}
	return l
}

func main() {
	fmt.Println(findLength([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 6}))
}
