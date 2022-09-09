package main

import "fmt"

// The task: https://leetcode.com/problems/daily-temperatures/
func dailyTemperatures(temp []int) []int {
	res := make([]int, len(temp))
	stk := make([]int, len(temp))
	top := -1
	for i, t := range temp {
		for top >= 0 && temp[stk[top]] < t {
			res[stk[top]] = i - stk[top]
			top--
		}
		top++
		stk[top] = i
	}
	return res
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
