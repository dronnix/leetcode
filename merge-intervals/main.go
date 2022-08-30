package main

import (
	"fmt"
	"sort"
)

// The task: https://leetcode.com/problems/merge-intervals/

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0, len(intervals))
	b, e := intervals[0][0], intervals[0][1]
	for i := range intervals {
		if intervals[i][0] <= e {
			if intervals[i][1] > e {
				e = intervals[i][1]
			}
			continue
		}
		res = append(res, []int{b, e})
		b, e = intervals[i][0], intervals[i][1]
	}
	res = append(res, []int{b, e})
	return res
}

func main() {
	ints := [][]int{
		{5, 9},
		{0, 3},
		{0, 7},
	}
	fmt.Println(merge(ints))
}
