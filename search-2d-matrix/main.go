package main

import (
	"fmt"
	"sort"
)

func searchMatrix(matrix [][]int, target int) bool {
	sz := len(matrix) * len(matrix[0])
	l := len(matrix[0])
	to2Didx := func(i int) (r, c int) {
		return i / l, i % l
	}

	idx := sort.Search(sz, func(i int) bool {
		r, c := to2Didx(i)
		return matrix[r][c] >= target
	})
	if idx == sz {
		return false
	}
	r, c := to2Didx(idx)
	return matrix[r][c] == target

}

func main() {
	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(searchMatrix(m, 6))
}
