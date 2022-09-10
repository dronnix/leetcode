package main

import "fmt"

const visited = 1
const pacific = 2
const atlantic = 4

// The task: https://leetcode.com/problems/pacific-atlantic-water-flow/
func pacificAtlantic(heights [][]int) [][]int {
	sinks := make([][]byte, len(heights))
	for i := range sinks {
		sinks[i] = make([]byte, len(heights[0]))
	}

	res := make([][]int, 0, len(heights)+len(heights[0]))
	for i := range heights {
		checkSink(heights, sinks, i, 0, 0, pacific, &res)
		checkSink(heights, sinks, i, len(heights[0])-1, 0, atlantic, &res)
	}
	for j := range heights[0] {
		checkSink(heights, sinks, 0, j, 0, pacific, &res)
		checkSink(heights, sinks, len(heights)-1, j, 0, atlantic, &res)
	}
	return res
}

func checkSink(heights [][]int, sinks [][]byte, r, c, prevHeight int, sink byte, res *[][]int) {
	if r < 0 || c < 0 || r >= len(heights) || c >= len(heights[0]) {
		return
	}
	if (sinks[r][c]&visited > 0) || (sinks[r][c]&sink > 0) {
		return
	}
	if heights[r][c] < prevHeight {
		return
	}
	sinks[r][c] |= visited | sink
	checkSink(heights, sinks, r-1, c, heights[r][c], sink, res)
	checkSink(heights, sinks, r, c-1, heights[r][c], sink, res)
	checkSink(heights, sinks, r+1, c, heights[r][c], sink, res)
	checkSink(heights, sinks, r, c+1, heights[r][c], sink, res)
	sinks[r][c] -= visited
	if sinks[r][c] >= pacific+atlantic {
		*res = append(*res, []int{r, c})
	}
}

func main() {
	h := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	fmt.Println(pacificAtlantic(h))
}
