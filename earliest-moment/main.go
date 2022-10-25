package main

import (
	"fmt"
	"sort"
)

func earliestAcq(logs [][]int, n int) int { //nolint:gocognit
	sort.Slice(logs, func(i, j int) bool { return logs[i][0] < logs[j][0] })

	grps := make([]map[int]struct{}, 0, n)
	for i := 0; i < n; i++ {
		g := make(map[int]struct{})
		g[i] = struct{}{}
		grps = append(grps, g)
	}

	for _, l := range logs {
		t, a, b := l[0], l[1], l[2]
		aIdx, bIdx := -1, -1
		for i := range grps {
			if _, ok := grps[i][a]; ok {
				aIdx = i
			}
			if _, ok := grps[i][b]; ok {
				bIdx = i
			}
			if aIdx != -1 && bIdx != -1 {
				break
			}
		}

		if aIdx == bIdx {
			continue
		}

		for e := range grps[bIdx] {
			grps[aIdx][e] = struct{}{}
		}
		if len(grps) == 2 {
			return t
		}

		grps[bIdx], grps[len(grps)-1] = grps[len(grps)-1], grps[bIdx]
		grps = grps[:len(grps)-1]
	}
	return -1
}

func main() {
	fmt.Println(earliestAcq([][]int{}, 2))
}
