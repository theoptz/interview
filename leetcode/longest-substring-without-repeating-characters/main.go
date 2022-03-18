package main

import "fmt"

// Follow link to see details
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	strLength := len(s)
	if strLength <= 1 {
		return strLength
	}

	maxLength := 1
	visited := make(map[rune]int)
	start := 0

	for index, ch := range s {
		if visited[ch] > start {
			start = visited[ch]
		}

		tmp := index - start + 1
		if tmp > maxLength {
			maxLength = tmp
		}

		visited[ch] = index + 1
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
