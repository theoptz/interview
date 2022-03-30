package main

import "fmt"

// Follow link to see details
// https://leetcode.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	count := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][0:count]
			}
		}

		count++
	}

	return strs[0][0:count]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flower", "flower", "flower"})) // "flower"
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))             // "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))                // ""
}
