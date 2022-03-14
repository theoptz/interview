package main

import "fmt"

// Follow link to see details
// https://leetcode.com/problems/longest-palindromic-substring/

func longestPalindrome(s string) string {
	strLen := len(s)
	if len(s) < 2 {
		return s
	}

	var res string

	getPalindrome := func(left, right int) string {
		var begin, end int
		for left >= 0 && right < strLen && s[left] == s[right] {
			begin = left
			end = right

			left--
			right++
		}

		return s[begin : end+1]
	}

	for i := 0; i < strLen; i++ {
		word := getPalindrome(i, i)
		if len(word) > len(res) {
			res = word
		}

		word = getPalindrome(i, i+1)
		if len(word) > len(res) {
			res = word
		}
	}

	return res
}

func main() {
	fmt.Println("babad", longestPalindrome("babad"))
	fmt.Println("cbbd", longestPalindrome("cbbd"))
}
