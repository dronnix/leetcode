package main

import "fmt"

// The task: https://leetcode.com/problems/letter-combinations-of-a-phone-number/

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	res := make([]string, 0, 1<<(len(digits)*4))
	buildCombination(digits, "", &res)
	return res
}

func buildCombination(digits, prefix string, result *[]string) {
	if digits == "" {
		*result = append(*result, prefix)
		return
	}
	for _, l := range digit2Letters(digits[0]) {
		buildCombination(digits[1:], prefix+string(l), result)
	}
}

func digit2Letters(l byte) string {
	switch l {
	case '2':
		return "abc"
	case '3':
		return "def"
	case '4':
		return "ghi"
	case '5':
		return "jkl"
	case '6':
		return "mno"
	case '7':
		return "pqrs"
	case '8':
		return "tuv"
	case '9':
		return "wxyz"
	default:
		panic("unknown digit:" + string(l))
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
