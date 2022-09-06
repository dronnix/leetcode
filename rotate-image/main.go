package main

import "fmt"

// The task:https://leetcode.com/problems/rotate-image/
func rotate(m [][]int) {
	n := len(m)/2 + len(m)%2
	e := len(m) - 1
	for i := 0; i < n; i++ {
		for j := i; j < e; j++ {
			swap(m, i, j)
		}
		e--
	}
}

func swap(m [][]int, i, j int) {
	n := len(m) - 1
	m[i][j], m[n-j][i], m[n-i][n-j], m[j][n-i] = m[n-j][i], m[n-i][n-j], m[j][n-i], m[i][j]
}

func main() {
	m := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	rotate(m)
	for i := range m {
		fmt.Println(m[i])
	}
}
