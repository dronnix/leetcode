package main

import "fmt"

// The task: https://leetcode.com/problems/partition-equal-subset-sum/
func canPartition(nums []int) bool {
	const maxNum, maxNumCnt = 100, 200
	reachable := make([]bool, maxNum*maxNumCnt) // one bit for each possible number.
	s := 0
	for _, n := range nums {
		s += n
		for i := len(reachable) - n; i > 0; i-- {
			if reachable[i] {
				reachable[i+n] = true
			}
		}
		reachable[n] = true
	}
	if s%2 > 0 {
		return false
	}
	s /= 2

	return reachable[s]
}

func main() {
	fmt.Println(canPartition([]int{1, 2, 2, 1}))
}
