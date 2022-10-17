package main

import (
	"fmt"
	"sort"
)

// The task: https://leetcode.com/problems/rectangle-area/

func computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 int) int {
	return (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1) - IntersectLen(ax1, ax2, bx1, bx2)*IntersectLen(ay1, ay2, by1, by2)
}

func IntersectLen(a1, a2, b1, b2 int) int {
	arr := []int{a1, a2, b1, b2}
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	if (arr[0] == a1 && arr[1] == a2) || (arr[0] == b1 && arr[1] == b2) {
		return 0
	}
	return arr[2] - arr[1]
}

func main() {
	fmt.Println(computeArea(0, 0, 2, 2, 1, 1, 3, 3))
}
