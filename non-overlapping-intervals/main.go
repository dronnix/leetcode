package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	prev := intervals[0]
	res := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < prev[1] {
			res++
			continue
		}
		prev = intervals[i]
	}
	return res
}

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
}
