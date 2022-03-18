package main

import "fmt"

// Follow link to see details
// https://leetcode.com/problems/remove-duplicate-letters/

func removeDuplicateLetters(s string) string {
	if len(s) < 2 {
		return s
	}

	counters := make(map[rune]int)
	for _, ch := range s {
		counters[ch]++
	}

	used := make(map[rune]bool)
	var res []rune

	for _, ch := range s {
		if used[ch] {
			counters[ch]--
			continue
		}

		for len(res) > 0 && res[len(res)-1] > ch && counters[res[len(res)-1]] > 1 {
			lastCh := res[len(res)-1]
			counters[lastCh]--
			used[lastCh] = false

			res = res[:len(res)-1]
		}

		res = append(res, ch)
		used[ch] = true
	}

	return string(res)
}

func main() {
	fmt.Println(removeDuplicateLetters("bcabc"))    // abc
	fmt.Println(removeDuplicateLetters("cbacdcbc")) // acdb
}
