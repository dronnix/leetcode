package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0, len(intervals)+1)
	pos := 0
	for ; pos < len(intervals) && intervals[pos][1] < newInterval[0]; pos++ {
		res = append(res, intervals[pos])
	}
	beg, end := newInterval[0], newInterval[1]
	if pos < len(intervals) && intervals[pos][0] < beg {
		beg = intervals[pos][0]
	}

	// TODO: Use binary search instead:
	for pos = len(intervals) - 1; pos >= 0 && intervals[pos][0] > end; pos-- {
	}
	if pos >= 0 && intervals[pos][1] > end {
		end = intervals[pos][1]
	}
	res = append(res, []int{beg, end})

	pos++
	for ; pos < len(intervals); pos++ {
		res = append(res, intervals[pos])
	}

	return res
}

func main() {
	fmt.Println(insert([][]int{{1, 3}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}
