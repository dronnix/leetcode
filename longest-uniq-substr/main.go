package main

import "fmt"

// See task at: https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	max, start := 0, 0
	ltrs := make([]uint16, 128)
	for i, c := range s {
		if start < int(ltrs[c]) {
			start = int(ltrs[c])
		}
		if i-start+1 > max {
			max = i - start + 1
		}
		ltrs[c] = uint16(i + 1)
	}
	return max
}

func main() {
	s := "abba"
	fmt.Println(lengthOfLongestSubstring(s))
}
