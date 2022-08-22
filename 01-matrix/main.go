package main

import "fmt"

// task: https://leetcode.com/problems/01-matrix/

const unknown = 65535

func updateMatrix(mat [][]int) [][]int { //nolint:gocognit,gocyclo
	result := make([][]int, len(mat))
	rowLen := len(mat[0])

	for row := range result {
		result[row] = make([]int, rowLen)
		for col := 0; col < rowLen; col++ {
			if mat[row][col] == 0 {
				result[row][col] = 0
				continue
			}
			result[row][col] = unknown
			if row > 0 && result[row-1][col]+1 < result[row][col] {
				result[row][col] = result[row-1][col] + 1
			}
			if col > 0 && result[row][col-1]+1 < result[row][col] {
				result[row][col] = result[row][col-1] + 1
			}
		}
	}

	rMax, cMax := len(result)-1, len(result[0])-1
	for row := rMax; row >= 0; row-- {
		for col := cMax; col >= 0; col-- {
			if row < rMax && result[row+1][col]+1 < result[row][col] {
				result[row][col] = result[row+1][col] + 1
			}
			if col < cMax && result[row][col+1]+1 < result[row][col] {
				result[row][col] = result[row][col+1] + 1
			}

		}
	}

	return result
}

func main() {
	m := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	fmt.Println(updateMatrix(m))
}
