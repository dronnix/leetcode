package main

import "fmt"

// https://leetcode.com/problems/minimum-height-trees/
func findMinHeightTrees(n int, edges [][]int) (res []int) {
	if n == 1 {
		return []int{0}
	}
	conns := make([]map[int]struct{}, n)
	for i := 0; i < n; i++ {
		conns[i] = make(map[int]struct{})
	}

	for _, e := range edges {
		conns[e[0]][e[1]] = struct{}{}
		conns[e[1]][e[0]] = struct{}{}
	}

	toDel := make([]int, 0, n)
	for k, v := range conns {
		if len(v) == 1 {
			toDel = append(toDel, k)
		}
	}

	remain := n
	for remain > 2 {
		newDel := make([]int, 0, n)
		for _, k := range toDel {
			remain--
			for d := range conns[k] {
				delete(conns[d], k)
				if len(conns[d]) == 1 {
					newDel = append(newDel, d)
				}
			}
		}
		toDel = newDel
	}

	return toDel
}

func main() {
	e := [][]int{
		{3, 0},
		{3, 1},
		{3, 2},
		{3, 4},
		{5, 4},
	}
	fmt.Println(findMinHeightTrees(6, e))
}
