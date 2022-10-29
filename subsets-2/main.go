package main

import "fmt"

func subsetsWithDup(nums []int) [][]int {
	numCounts := make(map[int]int, len(nums))
	for _, n := range nums {
		numCounts[n]++
	}

	res := make([][]int, 0, 1<<len(nums))
	res = append(res, []int{})
	for n, cnt := range numCounts {
		for i := range res {
			e := make([]int, 0, cnt)
			for c := 0; c < cnt; c++ {
				e = append(e, n)
				nr := make([]int, len(res[i]), len(res[i])+len(e))
				copy(nr, res[i])
				res = append(res, append(nr, e...))
			}
		}
	}
	return res
}

func main() {
	fmt.Println(subsetsWithDup([]int{0, 1, 2, 3, 4}))
}
