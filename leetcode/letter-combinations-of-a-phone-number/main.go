package main

import "fmt"

// Follow link to see details
// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

var (
	digitMap = map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
)

func letterCombinations(digits string) []string {
	res := make([]string, 0)

	letterCombinationHelper(digits, &res)

	return res
}

func letterCombinationHelper(digits string, res *[]string) {
	if len(digits) == 0 {
		return
	}

	digit := string(digits[0])
	if values, ok := digitMap[digit]; ok {
		curLen := len(*res)

		for _, val := range values {
			if curLen == 0 {
				*res = append(*res, val)
			} else {
				for i := 0; i < curLen; i++ {
					*res = append(*res, (*res)[i]+val)
				}
			}
		}

		*res = (*res)[curLen:]
	}

	letterCombinationHelper(digits[1:], res)
}

func main() {
	fmt.Println(letterCombinations("23")) // ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	fmt.Println(letterCombinations(""))   // []
	fmt.Println(letterCombinations("2"))  // ["a","b","c"]
}
