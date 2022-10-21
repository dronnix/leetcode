package main

import (
	"fmt"
	"math"
	"time"
)

// The task: https://leetcode.com/problems/coin-change/

func coinChange(coins []int, amount int) int {
	res := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		min := math.MaxInt
		for _, c := range coins {
			d := i - c
			if d >= 0 && res[d] >= 0 {
				if res[d]+1 < min {
					min = res[d] + 1
				}
			}
		}
		res[i] = -1
		if min < math.MaxInt {
			res[i] = min
		}
	}
	return res[amount]
}

func main() {
	t := time.Now()
	fmt.Println(coinChange([]int{1}, 10000))
	fmt.Println(time.Since(t))
}
