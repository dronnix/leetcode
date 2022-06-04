package main

import "fmt"

func findCenter(edges [][]int) int {
	v := make(map[int]int)
	for _, e := range edges {
		v[e[0]]++
		if v[e[0]] > 1 {
			return e[0]
		}
		v[e[1]]++
		if v[e[1]] > 1 {
			return e[1]
		}
	}
	panic("no center")
}

func main() {
	fmt.Println(findCenter([][]int{{0, 1}, {0, 2}, {0, 2}}))
}
