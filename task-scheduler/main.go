package main

import (
	"fmt"
	"sort"
)

// The task: https://leetcode.com/problems/task-scheduler/
func leastInterval(tasks []byte, n int) int {
	mp := make([]int, 26)
	for _, t := range tasks {
		mp[t-'A']++
	}

	res := 0
	last := 0
	for {
		sort.Slice(mp, func(i, j int) bool { return mp[i] > mp[j] })
		e := sort.Search(len(mp), func(i int) bool { return mp[i] == 0 })
		if e == 0 {
			return res - last
		}
		mp = mp[0:e]
		if e > n+1 {
			e = n + 1
		}

		s := mp[e-1]
		if len(mp) > e {
			s = mp[e]
		}
		for i := 0; i < e; i++ {
			mp[i] -= s
		}
		res += (n + 1) * s
		last = (n + 1) - e
	}
}

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'E'}, 2))
}
