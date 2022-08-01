package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab" //nolint:lll

	fmt.Println(wordBreakBFS(s, wordDict))

}

//nolint:unused,deadcode
func wordBreak(s string, wordDict []string) bool {
	sort.Slice(wordDict, func(i, j int) bool {
		return len(wordDict[i]) > len(wordDict[j])
	})
	tried := make([]bool, len(s)+1)
	return wordBreakWithMemoization(s, wordDict, tried)
}

//nolint:unused
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

//nolint:gocognit
func wordBreakBFS(s string, wordDict []string) bool {
	wordMap := make(map[string]struct{}, len(wordDict))
	wordsLegth := make([]bool, len(s)+1000)
	maxWordLen := 0
	for _, w := range wordDict {
		wordMap[w] = struct{}{}
		wordsLegth[len(w)] = true
		if len(w) > maxWordLen {
			maxWordLen = len(w)
		}
	}

	visited := make([]bool, len(s))

	stack := make([]int, 0, len(s))
	stack = append(stack, 0)
	for len(stack) > 0 {
		curLength := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[curLength] {
			continue
		}
		for i := curLength + 1; i <= len(s); i++ {
			wLen := i - curLength
			if wLen > maxWordLen {
				break
			}
			if !wordsLegth[wLen] {
				continue
			}
			curStr := s[curLength:i]
			if _, ok := wordMap[curStr]; ok {
				if len(curStr)+curLength == len(s) {
					return true
				}
				stack = append(stack, len(curStr)+curLength)
			}
		}
		visited[curLength] = true
	}
	return false
}
