package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// The task

func largestNumber(nums []int) string {
	strs := make([]string, 0, len(nums))
	for _, n := range nums {
		strs = append(strs, strconv.Itoa(n))
	}

	comp := func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	}
	sort.Slice(strs, comp)

	if strs[0] == "0" {
		return "0"
	}
	return strings.Join(strs, "")
}

func main() {
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
	fmt.Println(largestNumber([]int{8308, 8308, 830}))
	fmt.Println(largestNumber([]int{111311, 1113}))
}
