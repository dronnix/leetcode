package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab" //nolint:lll

	fmt.Println(wordBreak(s, wordDict))

}

func wordBreak(s string, wordDict []string) bool {
	sort.Slice(wordDict, func(i, j int) bool {
		return len(wordDict[i]) > len(wordDict[j])
	})
	tried := make([]bool, len(s)+1)
	return wordBreakWithMemoization(s, wordDict, tried)
}

func wordBreakWithMemoization(s string, wordDict []string, tried []bool) bool {
	for i := range wordDict {
		if tried[len(s)] {
			return false
		}
		if strings.HasPrefix(s, wordDict[i]) {
			if len(s) == len(wordDict[i]) {
				return true
			}
			if wordBreakWithMemoization(s[len(wordDict[i]):], wordDict, tried) {
				return true
			}
		}
	}
	tried[len(s)] = true
	return false
}
