package main

import "fmt"

// The task: https://leetcode.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int) {
	w, r1, r2 := n+m-1, m-1, n-1
	for r1 >= 0 && r2 >= 0 {
		if nums1[r1] > nums2[r2] {
			nums1[w] = nums1[r1]
			r1--
		} else {
			nums1[w] = nums2[r2]
			r2--
		}
		w--
	}

	for ; r2 >= 0; r2-- {
		nums1[w] = nums2[r2]
		w--
	}
}

func main() {
	a1, a2 := []int{1, 5, 8, 0, 0, 0}, []int{2, 6, 9}
	merge(a1, 3, a2, 3)
	fmt.Println(a1)
}
